// Package parser
// mapping nginx.conf to struct; mapping struct to nginx.conf
package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"strings"
)

//java org.antlr.v4.Tool -Dlanguage=Go -visitor -no-listener .\Nginx.g4
const (
	Server   = "server"
	Http     = "http"
	Upstream = "upstream"
	Events   = "events"
	Location = "location"
)

type Config struct {
	Blocks     []Block
	Statements map[string]Statement
	blockList  []ListResult //所有的block
	idCounter  int          //blockId
}
type Statement struct {
	Key   string
	Value string
}
type IfStatement struct {
	Condition  string //contains '(' ')'
	Statements map[string]Statement
}
type Block struct {
	Header       string //contains Type
	Type         string //e.g. location, server
	Blocks       []Block
	Statements   map[string]Statement
	IfStatements []IfStatement
}
type mappingVisitor struct {
	NginxVisitor
	Config   *Config
	context  interface{}
	parentId int
}

func newConfig() *Config {
	return &Config{Statements: make(map[string]Statement)}
}

//func newBlock(config *Config, parentId int) *Block {
//	block := &Block{Statements: make(map[string]Statement)}
//	if config != nil {
//		config.idCounter = config.idCounter + 1
//		config.blockList = append(config.blockList, *newListResult(block, config.idCounter, parentId))
//	}
//	return block
//}
func newBlock() *Block {
	block := &Block{Statements: make(map[string]Statement)}
	return block
}
func newStatement() *Statement {
	return &Statement{}
}
func newIfStatement() *IfStatement {
	return &IfStatement{Statements: make(map[string]Statement)}
}
func newMappingVisitor() NginxVisitor {
	return &mappingVisitor{Config: newConfig(), context: nil}
}
func (v *mappingVisitor) VisitConfig(ctx *ConfigContext) interface{} {
	for _, s := range ctx.AllStatement() {
		v.context = v.Config.Statements
		s.Accept(v)
	}
	for _, s := range ctx.AllBlock() {
		v.parentId = 0
		child := s.Accept(v).(Block)
		v.Config.Blocks = append(v.Config.Blocks, child)
	}
	return v.Config
}

func (v *mappingVisitor) VisitStatement(ctx *StatementContext) interface{} {
	m := v.context.(map[string]Statement)
	if ctx.RewriteStatement() != nil {
		s := ctx.RewriteStatement().Accept(v).(Statement)
		m[s.Key] = s
	}
	if ctx.GenericStatement() != nil {
		s := ctx.GenericStatement().Accept(v).(Statement)
		m[s.Key] = s
	}
	if ctx.RegexHeaderStatement() != nil {
		s := ctx.RegexHeaderStatement().Accept(v).(Statement)
		m[s.Key] = s
	}
	return nil
}

func (v *mappingVisitor) VisitGenericStatement(ctx *GenericStatementContext) interface{} {
	s := newStatement()
	children := ctx.GetChildren()
	s.Key = children[0].GetPayload().(antlr.Token).GetText()
	s.Value = concatChildrenString(children[1:], " ")
	return *s
}

func (v *mappingVisitor) VisitRegexHeaderStatement(ctx *RegexHeaderStatementContext) interface{} {
	s := newStatement()
	s.Key = ctx.REGEXP_PREFIXED().GetText()
	s.Value = ctx.Value().GetText()
	return *s
}

func (v *mappingVisitor) VisitRewriteStatement(ctx *RewriteStatementContext) interface{} {
	s := newStatement()
	children := ctx.GetChildren()
	s.Key = "rewrite"
	s.Value = concatChildrenString(children[1:], " ")
	return *s
}

func (v *mappingVisitor) VisitBlock(ctx *BlockContext) interface{} {
	block := newBlock()
	if ctx.GenericBlockHeader() != nil {
		block.Header = ctx.GenericBlockHeader().Accept(v).(string)
	}
	if ctx.LocationBlockHeader() != nil {
		block.Header = ctx.LocationBlockHeader().Accept(v).(string)
	}
	if strings.Contains(block.Header, " ") {
		block.Type = block.Header[:strings.Index(block.Header, " ")]
	} else {
		block.Type = block.Header
	}

	for _, s := range ctx.AllStatement() {
		v.context = block.Statements
		s.Accept(v)
	}
	for _, s := range ctx.AllBlock() {
		child := s.Accept(v).(Block)
		block.Blocks = append(block.Blocks, child)
	}
	for _, s := range ctx.AllIfStatement() {
		block.IfStatements = append(block.IfStatements, s.Accept(v).(IfStatement))
	}

	return *block
}

func (v *mappingVisitor) VisitGenericBlockHeader(ctx *GenericBlockHeaderContext) interface{} {
	return concatChildrenString(ctx.GetChildren(), " ")
}

func (v *mappingVisitor) VisitLocationBlockHeader(ctx *LocationBlockHeaderContext) interface{} {
	return concatChildrenString(ctx.GetChildren(), " ")
}

func (v *mappingVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	ifStatement := newIfStatement()
	ifStatement.Condition = ctx.IfBody().Accept(v).(string)
	for _, s := range ctx.AllStatement() {
		v.context = ifStatement.Statements
		s.Accept(v)
	}
	return *ifStatement
}

func (v *mappingVisitor) VisitIfBody(ctx *IfBodyContext) interface{} {
	return concatChildrenString(ctx.GetChildren(), " ")
}

// VisitRegexp unused
func (v *mappingVisitor) VisitRegexp(ctx *RegexpContext) interface{} {
	return ctx.GetText()
}

//only for Value, Token
func concatChildrenString(tree []antlr.Tree, sep string) string {
	if len(tree) == 0 {
		return ""
	}
	s := ""
	for _, c := range tree {
		payload := c.GetPayload()
		switch payload.(type) {
		case antlr.Token:
			s += payload.(antlr.Token).GetText()
		case *antlr.BaseParserRuleContext:
			s += payload.(*antlr.BaseParserRuleContext).GetText()
		default:
			panic("unexpected child type")
		}
		s += sep
	}
	return s[:len(s)-len(sep)]
}

// LoadConfig Parse nginx.conf
func LoadConfig(file string) (*Config, error) {
	input, err := antlr.NewFileStream(file)
	if err != nil {
		return nil, err
	}
	lexer := NewNginxLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewNginxParser(stream)
	//p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Config()
	visitor := newMappingVisitor()
	//global block id=0 TODO
	config := tree.Accept(visitor).(*Config)
	return config, nil
}

func (c *Config) EasyDumpToFile(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	c.DumpToFile(file, " ", 0, 4)
	return nil
}

// DumpToFile Generate new nginx.conf
func (c *Config) DumpToFile(file *os.File, space string, indent, delta int) {
	dumpAllStatements(file, space, indent, c.Statements)
	dumpAllBlocks(file, space, indent, delta, c.Blocks)
}

func (c *Block) dumpToFile(file *os.File, space string, indent, delta int) {
	writeWithIndent(file, space, indent, fmt.Sprintf("%s {\n", c.Header))
	dumpAllStatements(file, space, indent+delta, c.Statements)
	dumpAllBlocks(file, space, indent+delta, delta, c.Blocks)
	for _, s := range c.IfStatements {
		s.dumpToFile(file, space, indent+delta, delta)
	}
	writeWithIndent(file, space, indent, "}\n")
}

func (c *IfStatement) dumpToFile(file *os.File, space string, indent, delta int) {
	writeWithIndent(file, space, indent, fmt.Sprintf(" if %s {\n", c.Condition))
	dumpAllStatements(file, space, indent+delta, c.Statements)
	writeWithIndent(file, space, indent, "}\n")
}

func dumpAllStatements(file *os.File, space string, indent int, m map[string]Statement) {
	if m == nil {
		return
	}
	for k, v := range m {
		writeWithIndent(file, space, indent, fmt.Sprintf("%s%s%s ;\n", k, space, v.Value))
	}
}

func dumpAllBlocks(file *os.File, space string, indent, delta int, blocks []Block) {
	if blocks == nil {
		return
	}
	for _, block := range blocks {
		block.dumpToFile(file, space, indent, delta)
	}
}

func writeWithIndent(file *os.File, space string, indent int, s string) {
	_, err := file.WriteString(strings.Repeat(space, indent) + s)
	if err != nil {
		panic(err)
	}
}

//func IsValidBlockHeader(header string) bool {
//	allBlockHeader := []string{server, http, upstream, events, location}
//	for _, i := range allBlockHeader {
//		if i == header {
//			return true
//		}
//	}
//	return false
//}

type ListResult struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	Header string `json:"blockHeader"`
	//ParentId int    `json:"parentId"` //parent block id
	Block *Block `json:"-"` //don't export
}

func newListResult(block *Block, id int) *ListResult {
	return &ListResult{
		Id:     id,
		Block:  block,
		Header: block.Header,
		//ParentId: parentId,
		Type: block.Type,
	}
}

// ListAllBlocks Print Block tree
func (c *Config) ListAllBlocks() {
	fmt.Println("Global[nginx.conf](id=0)")
	id := 1
	for _, block := range c.Blocks {
		id = c.printBlocks(&block, 1, 4, id, true)
	}
}

// GetBlocksList Return ListResult array corresponding to Block tree
func (c *Config) GetBlocksList() []ListResult {
	if len(c.blockList) > 0 {
		return c.blockList
	}
	id := 1
	for _, block := range c.Blocks {
		id = c.printBlocks(&block, 1, 4, id, false)
	}
	return c.blockList
}

func (c *Config) printBlocks(block *Block, level, indent, startId int, print bool) int {
	if print {
		if level > 1 {
			fmt.Printf("│%s", strings.Repeat(" ", (level-1)*(indent+1)))
		}
		fmt.Printf("│%s %s\n", strings.Repeat("─", indent), fmtBlockName(block, startId))
	} else {
		c.blockList = append(c.blockList, *newListResult(block, startId))
	}

	startId = startId + 1
	for _, b := range block.Blocks {
		startId = c.printBlocks(&b, level+1, 4, startId, print)
	}
	if print && block.Type != Location && block.Type != Http {
		fmt.Println("│")
	}

	return startId
}

func fmtBlockName(block *Block, id int) string {
	desc := block.Header
	if block.Type == Server {
		desc = Server + "{"
		serverName := block.Statements["server_name"].Value
		port := block.Statements["listen"].Value
		if serverName == "" {
			serverName = "unknown_host"
		}
		if port == "" {
			serverName = "unknown_port"
		}
		if strings.Count(serverName, " ") > 0 {
			for _, s := range strings.Split(serverName, " ") {
				desc += fmt.Sprintf("%s:%s,", s, port)
			}
		} else {
			desc += fmt.Sprintf("%s:%s", serverName, port)
		}
		desc += "}"
	}
	return fmt.Sprintf("[%s] (id=%d)", desc, id)
}
