package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

func Test1(t *testing.T) {
	//input, err := antlr.NewFileStream("../conf/ok.conf")
	//if err != nil {
	//	panic(err)
	//}
	input := antlr.NewInputStream(` worker_processes 1 ;
events {
    worker_connections 1024 ;
}
http {
    include mime.types ;
    default_type application/octet-stream ;
    sendfile on ;
    keepalive_timeout 65 ;
    server {
        listen 80 ;
        server_name localhost ;
        error_page 500 502 503 504 /50x.html ;
        rewrite_by_lua_block {
            local uri = ngx.var.uri;
local path = ""
local regex = "/t.*"

if (path ~= "" and uri == path) or (regex ~= "" and string.match(uri, regex))
then
ngx.header["Server"] = "mock"
ngx.header["Content-Type"] = "application/json"

ngx.say('{"a":1}')
ngx.exit(200)
end
  ;
        }
        location / {
            root html ;
            index index.html index.htm ;
        }
        location = /50x.html {
            root html ;
        }
    }
}`)
	lexer := NewNginxLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewNginxParser(stream)

	//p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Config()
	fmt.Println(tree.ToStringTree(nil, p))
	tree.GetChildren()
	//visitor := newMappingVisitor()
	//config := tree.Accept(visitor).(*Config)
	//fmt.Println(config)
	//
	//file, err := os.OpenFile("out.conf", os.O_CREATE, 0666)
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()

	//config.EasyDumpToFile(file)
}
