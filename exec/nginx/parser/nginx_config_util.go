package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"strings"
)

//go:generate java org.antlr.v4.Tool -Dlanguage=Go -visitor -no-listener Nginx.g4

const (
	Server   = "server"
	HTTP     = "http"
	Upstream = "upstream"
	Events   = "events"
	Location = "location"
	Lua      = "lua"
)

type Config struct {
	Blocks     []Block
	Statements []Statement
	blockList  []ListResult //所有的block
	idCounter  int          //blockId
}
type Statement struct {
	Key   string
	Value string
}
type IfStatement struct {
	Condition  string //contains '(' ')'
	Statements []Statement
}
type Block struct {
	Header       string //contains Type
	Type         string //e.g. location, server
	Blocks       []Block
	Statements   []Statement
	IfStatements []IfStatement
}
type mappingVisitor struct {
	NginxVisitor
	Config   *Config
	context  interface{}
	parentId int
}

func newConfig() *Config {
	return &Config{}
}

func NewBlock() *Block {
	return &Block{}
}
func NewStatement() *Statement {
	return &Statement{}
}
func newIfStatement() *IfStatement {
	return &IfStatement{}
}
func newMappingVisitor() NginxVisitor {
	return &mappingVisitor{Config: newConfig(), context: nil}
}
func (v *mappingVisitor) VisitConfig(ctx *ConfigContext) interface{} {
	for _, s := range ctx.AllStatement() {
		v.Config.Statements = append(v.Config.Statements, s.Accept(v).(Statement))
	}
	for _, s := range ctx.AllBlock() {
		v.parentId = 0
		child := s.Accept(v).(Block)
		v.Config.Blocks = append(v.Config.Blocks, child)
	}
	return v.Config
}

func (v *mappingVisitor) VisitStatement(ctx *StatementContext) interface{} {
	if ctx.RewriteStatement() != nil {
		return ctx.RewriteStatement().Accept(v).(Statement)
	}
	if ctx.GenericStatement() != nil {
		return ctx.GenericStatement().Accept(v).(Statement)
	}
	if ctx.RegexHeaderStatement() != nil {
		return ctx.RegexHeaderStatement().Accept(v).(Statement)
	}
	return nil
}

func (v *mappingVisitor) VisitGenericStatement(ctx *GenericStatementContext) interface{} {
	s := NewStatement()
	children := ctx.GetChildren()
	s.Key = children[0].GetPayload().(antlr.Token).GetText()
	s.Value = concatChildrenString(children[1:], " ") // value = "" when lua statement
	return *s
}

func (v *mappingVisitor) VisitRegexHeaderStatement(ctx *RegexHeaderStatementContext) interface{} {
	s := NewStatement()
	s.Key = ctx.REGEXP_PREFIXED().GetText()
	s.Value = ctx.Value().GetText()
	return *s
}

func (v *mappingVisitor) VisitRewriteStatement(ctx *RewriteStatementContext) interface{} {
	s := NewStatement()
	children := ctx.GetChildren()
	s.Key = "rewrite"
	s.Value = concatChildrenString(children[1:], " ")
	return *s
}

func (v *mappingVisitor) VisitBlock(ctx *BlockContext) interface{} {
	block := NewBlock()
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
		block.Statements = append(block.Statements, s.Accept(v).(Statement))
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
		ifStatement.Statements = append(ifStatement.Statements, s.Accept(v).(Statement))
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
	//fmt.Println(tree.ToStringTree(nil, p))
	visitor := newMappingVisitor()
	config := tree.Accept(visitor).(*Config)
	return config, nil
}

func (c *Config) EasyDumpToFile(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
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

func dumpAllStatements(file *os.File, space string, indent int, statements []Statement) {
	if statements == nil {
		return
	}
	for _, s := range statements {
		writeWithIndent(file, space, indent, fmt.Sprintf("%s%s%s ;\n", s.Key, space, s.Value))
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
		//TODO
		panic(err)
	}
}

type ListResult struct {
	Id     int
	Type   string
	Header string
	Block  *Block
}

func newListResult(block *Block, id int) *ListResult {
	return &ListResult{
		Id:     id,
		Block:  block,
		Header: block.Header,
		Type:   block.Type,
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
	for i := 0; i < len(c.Blocks); i++ {
		id = c.printBlocks(&c.Blocks[i], 1, 4, id, false)
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

	startId++
	for i := 0; i < len(block.Blocks); i++ {
		startId = c.printBlocks(&block.Blocks[i], level+1, 4, startId, print)
	}
	if print && block.Type != Location && block.Type != HTTP {
		fmt.Println("│")
	}

	return startId
}

func fmtBlockName(block *Block, id int) string {
	desc := block.Header
	if block.Type == Server {
		desc = Server + "{"
		serverName := findStatement(block.Statements, "server_name")
		port := findStatement(block.Statements, "listen")
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

func findStatement(arr []Statement, key string) string {
	for _, s := range arr {
		if s.Key == key {
			return s.Value
		}
	}
	return ""
}

func SetStatement(arr []Statement, k, v string, addNew bool) []Statement {
	if addNew {
		return append(arr, Statement{Key: k, Value: v})
	} else {
		for i := 0; i < len(arr); i++ {
			if arr[i].Key == k {
				arr[i].Value = v
				return arr
			}
		}
		return append(arr, Statement{Key: k, Value: v})
	}
}
