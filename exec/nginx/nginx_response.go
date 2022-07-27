package nginx

import (
	"context"
	"fmt"
	"strconv"

	"github.com/chaosblade-io/chaosblade-exec-os/exec/category"
	"github.com/chaosblade-io/chaosblade-exec-os/exec/nginx/parser"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/chaosblade-io/chaosblade-spec-go/util"
)

const (
	NginxResponseBin           = "chaos_nginxresponse"
	defaultContentType         = "text/plain;charset=utf-8"
	contentTypeHeaderNameUpper = "Content-Type"
	contentTypeHeaderNameLower = "content-type"
)

var contentTypeMap = map[string]string{
	"json": "application/json",
	"txt":  "text/plain;charset=utf-8",
	// "html": "text/html;charset=utf-8",
}

type ResponseActionSpec struct {
	spec.BaseExpActionCommandSpec
}

//TODO regex支持
//TODO 支持匹配冲突问题，即regex情况下的多个匹配都满足的情况
//TODO 支持路由类型的响应替换？
//TODO 支持html文件类型，解决方案是自己启动一个web server
//目前暂定只支持固定url响应替换，而不支持regex
func NewResponseActionSpec() spec.ExpActionCommandSpec {
	return &ResponseActionSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name: "body",
					Desc: "change response body",
				},
				// &spec.ExpFlag{
				// 	Name: "body-file",
				// 	Desc: "change response body",
				// },
				&spec.ExpFlag{
					//为了使body有效，会自动设置content type
					Name: "header",
					Desc: "change response header",
				},
				&spec.ExpFlag{
					//为了使body有效，会自动设置content type
					Name:    "code",
					Desc:    "change response code, default 200",
					Default: "200",
				},
				&spec.ExpFlag{
					Name:     "path",
					Desc:     "change response path",
					Required: true,
				},
				&spec.ExpFlag{
					Name:    "type",
					Desc:    "new response body type, such as html, text/plain, json ,etc. this argument is same as setting --header='content-type=html'",
					Default: "json",
				},
			},
			ActionFlags:    []spec.ExpFlagSpec{},
			ActionExecutor: &NginxResponseExecutor{},
			ActionExample: `
# Set /test returns body='ok',code=200,type=json
blade create nginx response --path /test --body ok

# Set /test returns body='',code=500,type=json
blade create nginx response --path /test --code 500

# Set /test returns body='',code=500,type=json
blade create nginx response --path /test --code 200 --body '{"a":1}' --type json

# Revert config change to the oldest config file
blade destroy nginx response
			`,
			ActionPrograms:   []string{NginxResponseBin},
			ActionCategories: []string{category.Middleware},
		},
	}
}

func (*ResponseActionSpec) Name() string {
	return "response"
}

func (*ResponseActionSpec) Aliases() []string {
	return []string{}
}

func (*ResponseActionSpec) ShortDesc() string {
	return "Response experiment"
}

func (d *ResponseActionSpec) LongDesc() string {
	if d.ActionLongDesc != "" {
		return d.ActionLongDesc
	}
	return "Nginx response experiment"
}

type NginxResponseExecutor struct {
	channel spec.Channel
}

func (*NginxResponseExecutor) Name() string {
	return "response"
}

func (ng *NginxResponseExecutor) Exec(suid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	_, response := getNginxPid(ng.channel, ctx) // nginx process
	if response != nil {
		return response
	}

	dir, activeFile, response := getNginxConfigLocation(ng.channel, ctx)
	if response != nil {
		return response
	}
	backup := dir + configBackupName

	if _, ok := spec.IsDestroy(ctx); ok {
		return ng.stop(ctx, dir, activeFile, backup, model)
	}
	return ng.start(ctx, dir, activeFile, backup, model)
}

func (ng *NginxResponseExecutor) start(ctx context.Context, dir, activeFile, backup string, model *spec.ExpModel) *spec.Response {
	contentType, response := getContentType(model.ActionFlags["type"])
	if response != nil {
		return response
	}
	config, _ := parser.LoadConfig(activeFile)
	path := model.ActionFlags["path"]
	code := model.ActionFlags["code"]
	body := model.ActionFlags["body"]
	header := model.ActionFlags["header"]
	//暂时不指定id，默认是第一个server内增加location
	server, response := findServerBlock(config)
	if response != nil {
		return response
	}
	newBlock, response := createNewLocationBlock(path, code, body, header, contentType)
	if response != nil {
		return response
	}
	newBlockList := []parser.Block{*newBlock}
	for _, b := range server.Blocks {
		if b.Type != newBlock.Type || b.Header != newBlock.Header{
			newBlockList = append(newBlockList, b)
		}
	}
	server.Blocks = newBlockList

	name := "nginx.chaosblade.tmp.conf"
	err := config.EasyDumpToFile(name)
	if err != nil {
		return spec.ReturnFail(spec.OsCmdExecFailed, err.Error())
	}
	if response := testNginxConfig(ng.channel, ctx, name, dir); response != nil {
		return response
	}

	cmd := ""
	if util.IsExist(backup) {
		//don't create backup
		//TODO version chain
		cmd = fmt.Sprintf("cp -f %s %s", name, activeFile)
	} else {
		cmd = fmt.Sprintf("cp %s %s && cp -f %s %s", activeFile, backup, name, activeFile)
	}
	cmd += fmt.Sprintf(" && rm %s && nginx -s reload", name)
	response = ng.channel.Run(ctx, cmd, "")
	if !response.Success {
		return response
	}

	return spec.ReturnSuccess("set response successfully")
}

func (ng *NginxResponseExecutor) stop(ctx context.Context, dir, activeFile, backup string, model *spec.ExpModel) *spec.Response {
	if !util.IsExist(backup) || util.IsDir(backup) {
		return spec.ReturnFail(spec.OsCmdExecFailed, fmt.Sprintf("backup file %s not exists", backup))
	}

	response := ng.channel.Run(ctx, fmt.Sprintf("mv -f %s %s && nginx -s reload", backup, activeFile), "")
	if !response.Success {
		return response
	}
	return spec.ReturnSuccess("nginx config restored")
}

func (ng *NginxResponseExecutor) SetChannel(channel spec.Channel) {
	ng.channel = channel
}

func getContentType(contentTypeKey string) (string, *spec.Response) {
	if contentTypeKey == "" {
		return defaultContentType, nil
	}
	if v, ok := contentTypeMap[contentTypeKey]; ok {
		return v, nil
	}
	support := ""
	for k := range contentTypeMap {
		support += k + ", "
	}
	return "", spec.ReturnFail(spec.OsCmdExecFailed, fmt.Sprintf("--type %s is not supported, only supports ( %s )", contentTypeKey, support))
}

//TODO not only first
func findServerBlock(config *parser.Config) (*parser.Block, *spec.Response) {
	var http *parser.Block = nil
	for i := 0; i < len(config.Blocks); i++ {
		b := &config.Blocks[i]
		if b.Type == parser.Http {
			http = b
			break
		}
	}
	for i := 0; i < len(http.Blocks); i++ {
		b := &http.Blocks[i]
		if b.Type == parser.Server {
			return b, nil
		}
	}
	return nil, spec.ReturnFail(spec.OsCmdExecFailed, "Server config not found in nginx.conf")
}

func createNewLocationBlock(path, code, body, header, contentType string) (*parser.Block, *spec.Response) {
	block := parser.NewBlock()
	block.Type = parser.Location
	block.Header = fmt.Sprintf("%s = %s", block.Type, path) //highest priority
	pairs := parseMultipleKvPairs(header)
	if pairs == nil && header != "" {
		return nil, spec.ReturnFail(spec.ParameterInvalid, fmt.Sprintf("--header=%s is not valid", header))
	}
	hasContentType := false
	for _, pair := range pairs {
		statement := parser.Statement{Key: pair[0], Value: pair[1]}
		block.Statements[statement.Key] = statement
		if statement.Key == contentTypeHeaderNameLower ||
			statement.Key == contentTypeHeaderNameUpper {
			hasContentType = true
		}
	}
	if !hasContentType {
		block.Statements["add_header"] = parser.Statement{Key: "add_header", Value: fmt.Sprintf("Content-Type '%s'", contentType)}
	}

	if _, err := strconv.Atoi(code); err != nil {
		return nil, spec.ReturnFail(spec.ParameterInvalid, fmt.Sprintf("--code=%s is not valid, %s", code, err))
	}
	statement := parser.Statement{Key: "return", Value: fmt.Sprintf("%s '%s'", code, body)}
	block.Statements[statement.Key] = statement
	return block, nil
}
