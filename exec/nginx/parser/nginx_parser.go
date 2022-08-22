// Code generated from Nginx.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Nginx

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type NginxParser struct {
	*antlr.BaseParser
}

var nginxParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func nginxParserInit() {
	staticData := &nginxParserStaticData
	staticData.literalNames = []string{
		"", "';'", "'{'", "'}'", "'if'", "')'", "'('", "'\\.'", "'^'", "'location'",
		"'rewrite'", "'last'", "'break'", "'redirect'", "'permanent'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "LUA_HEADER",
		"Value", "STR_EXT", "LINE_COMMENT", "REGEXP_PREFIXED", "QUOTED_STRING",
		"SINGLE_QUOTED", "WS", "NEWLINE",
	}
	staticData.ruleNames = []string{
		"config", "statement", "genericStatement", "regexHeaderStatement", "luaBlock",
		"luaStatement", "block", "genericBlockHeader", "ifStatement", "ifBody",
		"regexp", "locationBlockHeader", "rewriteStatement",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 23, 197, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 4, 0, 30, 8, 0, 11, 0,
		12, 0, 31, 1, 1, 1, 1, 1, 1, 3, 1, 37, 8, 1, 1, 1, 1, 1, 3, 1, 41, 8, 1,
		1, 2, 1, 2, 1, 2, 5, 2, 46, 8, 2, 10, 2, 12, 2, 49, 9, 2, 1, 3, 1, 3, 1,
		3, 1, 4, 1, 4, 3, 4, 56, 8, 4, 1, 4, 1, 4, 3, 4, 60, 8, 4, 1, 4, 5, 4,
		63, 8, 4, 10, 4, 12, 4, 66, 9, 4, 1, 4, 1, 4, 3, 4, 70, 8, 4, 1, 5, 4,
		5, 73, 8, 5, 11, 5, 12, 5, 74, 1, 5, 3, 5, 78, 8, 5, 1, 5, 3, 5, 81, 8,
		5, 1, 5, 3, 5, 84, 8, 5, 1, 6, 1, 6, 3, 6, 88, 8, 6, 1, 6, 3, 6, 91, 8,
		6, 1, 6, 1, 6, 3, 6, 95, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 101, 8, 6,
		1, 6, 3, 6, 104, 8, 6, 5, 6, 106, 8, 6, 10, 6, 12, 6, 109, 9, 6, 1, 6,
		1, 6, 3, 6, 113, 8, 6, 1, 7, 1, 7, 1, 7, 5, 7, 118, 8, 7, 10, 7, 12, 7,
		121, 9, 7, 1, 8, 1, 8, 3, 8, 125, 8, 8, 1, 8, 1, 8, 3, 8, 129, 8, 8, 1,
		8, 1, 8, 3, 8, 133, 8, 8, 1, 8, 1, 8, 3, 8, 137, 8, 8, 5, 8, 139, 8, 8,
		10, 8, 12, 8, 142, 9, 8, 1, 8, 1, 8, 3, 8, 146, 8, 8, 1, 9, 1, 9, 3, 9,
		150, 8, 9, 1, 9, 4, 9, 153, 8, 9, 11, 9, 12, 9, 154, 1, 9, 3, 9, 158, 8,
		9, 1, 9, 1, 9, 3, 9, 162, 8, 9, 1, 9, 3, 9, 165, 8, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 4, 10, 176, 8, 10, 11, 10, 12,
		10, 177, 1, 11, 1, 11, 3, 11, 182, 8, 11, 1, 11, 1, 11, 3, 11, 186, 8,
		11, 1, 12, 1, 12, 1, 12, 3, 12, 191, 8, 12, 1, 12, 1, 12, 3, 12, 195, 8,
		12, 1, 12, 0, 0, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 0,
		2, 2, 0, 4, 6, 16, 16, 1, 0, 11, 14, 230, 0, 29, 1, 0, 0, 0, 2, 36, 1,
		0, 0, 0, 4, 42, 1, 0, 0, 0, 6, 50, 1, 0, 0, 0, 8, 53, 1, 0, 0, 0, 10, 80,
		1, 0, 0, 0, 12, 87, 1, 0, 0, 0, 14, 114, 1, 0, 0, 0, 16, 122, 1, 0, 0,
		0, 18, 147, 1, 0, 0, 0, 20, 175, 1, 0, 0, 0, 22, 179, 1, 0, 0, 0, 24, 187,
		1, 0, 0, 0, 26, 30, 3, 2, 1, 0, 27, 30, 3, 12, 6, 0, 28, 30, 3, 8, 4, 0,
		29, 26, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 28, 1, 0, 0, 0, 30, 31, 1,
		0, 0, 0, 31, 29, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 1, 1, 0, 0, 0, 33,
		37, 3, 24, 12, 0, 34, 37, 3, 4, 2, 0, 35, 37, 3, 6, 3, 0, 36, 33, 1, 0,
		0, 0, 36, 34, 1, 0, 0, 0, 36, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 40,
		5, 1, 0, 0, 39, 41, 5, 23, 0, 0, 40, 39, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0,
		41, 3, 1, 0, 0, 0, 42, 47, 5, 16, 0, 0, 43, 46, 5, 16, 0, 0, 44, 46, 3,
		20, 10, 0, 45, 43, 1, 0, 0, 0, 45, 44, 1, 0, 0, 0, 46, 49, 1, 0, 0, 0,
		47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 5, 1, 0, 0, 0, 49, 47, 1, 0,
		0, 0, 50, 51, 5, 19, 0, 0, 51, 52, 5, 16, 0, 0, 52, 7, 1, 0, 0, 0, 53,
		55, 5, 15, 0, 0, 54, 56, 5, 23, 0, 0, 55, 54, 1, 0, 0, 0, 55, 56, 1, 0,
		0, 0, 56, 57, 1, 0, 0, 0, 57, 59, 5, 2, 0, 0, 58, 60, 5, 23, 0, 0, 59,
		58, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 64, 1, 0, 0, 0, 61, 63, 3, 10,
		5, 0, 62, 61, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 65,
		1, 0, 0, 0, 65, 67, 1, 0, 0, 0, 66, 64, 1, 0, 0, 0, 67, 69, 5, 3, 0, 0,
		68, 70, 5, 23, 0, 0, 69, 68, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 9, 1,
		0, 0, 0, 71, 73, 7, 0, 0, 0, 72, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74,
		72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 77, 1, 0, 0, 0, 76, 78, 5, 1, 0,
		0, 77, 76, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0, 79, 81,
		5, 1, 0, 0, 80, 72, 1, 0, 0, 0, 80, 79, 1, 0, 0, 0, 81, 83, 1, 0, 0, 0,
		82, 84, 5, 23, 0, 0, 83, 82, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 11, 1,
		0, 0, 0, 85, 88, 3, 22, 11, 0, 86, 88, 3, 14, 7, 0, 87, 85, 1, 0, 0, 0,
		87, 86, 1, 0, 0, 0, 88, 90, 1, 0, 0, 0, 89, 91, 5, 23, 0, 0, 90, 89, 1,
		0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 94, 5, 2, 0, 0, 93,
		95, 5, 23, 0, 0, 94, 93, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 107, 1, 0,
		0, 0, 96, 101, 3, 2, 1, 0, 97, 101, 3, 12, 6, 0, 98, 101, 3, 16, 8, 0,
		99, 101, 3, 8, 4, 0, 100, 96, 1, 0, 0, 0, 100, 97, 1, 0, 0, 0, 100, 98,
		1, 0, 0, 0, 100, 99, 1, 0, 0, 0, 101, 103, 1, 0, 0, 0, 102, 104, 5, 23,
		0, 0, 103, 102, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 106, 1, 0, 0, 0,
		105, 100, 1, 0, 0, 0, 106, 109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107,
		108, 1, 0, 0, 0, 108, 110, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 110, 112,
		5, 3, 0, 0, 111, 113, 5, 23, 0, 0, 112, 111, 1, 0, 0, 0, 112, 113, 1, 0,
		0, 0, 113, 13, 1, 0, 0, 0, 114, 119, 5, 16, 0, 0, 115, 118, 5, 16, 0, 0,
		116, 118, 3, 20, 10, 0, 117, 115, 1, 0, 0, 0, 117, 116, 1, 0, 0, 0, 118,
		121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 15, 1,
		0, 0, 0, 121, 119, 1, 0, 0, 0, 122, 124, 5, 4, 0, 0, 123, 125, 5, 23, 0,
		0, 124, 123, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126,
		128, 3, 18, 9, 0, 127, 129, 5, 23, 0, 0, 128, 127, 1, 0, 0, 0, 128, 129,
		1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 132, 5, 2, 0, 0, 131, 133, 5, 23,
		0, 0, 132, 131, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 140, 1, 0, 0, 0,
		134, 136, 3, 2, 1, 0, 135, 137, 5, 23, 0, 0, 136, 135, 1, 0, 0, 0, 136,
		137, 1, 0, 0, 0, 137, 139, 1, 0, 0, 0, 138, 134, 1, 0, 0, 0, 139, 142,
		1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141, 143, 1, 0,
		0, 0, 142, 140, 1, 0, 0, 0, 143, 145, 5, 3, 0, 0, 144, 146, 5, 23, 0, 0,
		145, 144, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 17, 1, 0, 0, 0, 147, 149,
		5, 6, 0, 0, 148, 150, 5, 23, 0, 0, 149, 148, 1, 0, 0, 0, 149, 150, 1, 0,
		0, 0, 150, 152, 1, 0, 0, 0, 151, 153, 5, 16, 0, 0, 152, 151, 1, 0, 0, 0,
		153, 154, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155,
		157, 1, 0, 0, 0, 156, 158, 5, 23, 0, 0, 157, 156, 1, 0, 0, 0, 157, 158,
		1, 0, 0, 0, 158, 161, 1, 0, 0, 0, 159, 162, 5, 16, 0, 0, 160, 162, 3, 20,
		10, 0, 161, 159, 1, 0, 0, 0, 161, 160, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0,
		162, 164, 1, 0, 0, 0, 163, 165, 5, 23, 0, 0, 164, 163, 1, 0, 0, 0, 164,
		165, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 5, 5, 0, 0, 167, 19, 1,
		0, 0, 0, 168, 176, 5, 7, 0, 0, 169, 176, 5, 8, 0, 0, 170, 176, 5, 16, 0,
		0, 171, 172, 5, 6, 0, 0, 172, 173, 3, 20, 10, 0, 173, 174, 5, 5, 0, 0,
		174, 176, 1, 0, 0, 0, 175, 168, 1, 0, 0, 0, 175, 169, 1, 0, 0, 0, 175,
		170, 1, 0, 0, 0, 175, 171, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 175,
		1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 21, 1, 0, 0, 0, 179, 181, 5, 9,
		0, 0, 180, 182, 5, 16, 0, 0, 181, 180, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0,
		182, 185, 1, 0, 0, 0, 183, 186, 5, 16, 0, 0, 184, 186, 3, 20, 10, 0, 185,
		183, 1, 0, 0, 0, 185, 184, 1, 0, 0, 0, 186, 23, 1, 0, 0, 0, 187, 190, 5,
		10, 0, 0, 188, 191, 5, 16, 0, 0, 189, 191, 3, 20, 10, 0, 190, 188, 1, 0,
		0, 0, 190, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 194, 5, 16, 0, 0,
		193, 195, 7, 1, 0, 0, 194, 193, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195,
		25, 1, 0, 0, 0, 40, 29, 31, 36, 40, 45, 47, 55, 59, 64, 69, 74, 77, 80,
		83, 87, 90, 94, 100, 103, 107, 112, 117, 119, 124, 128, 132, 136, 140,
		145, 149, 154, 157, 161, 164, 175, 177, 181, 185, 190, 194,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// NginxParserInit initializes any static state used to implement NginxParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNginxParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NginxParserInit() {
	staticData := &nginxParserStaticData
	staticData.once.Do(nginxParserInit)
}

// NewNginxParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNginxParser(input antlr.TokenStream) *NginxParser {
	NginxParserInit()
	this := new(NginxParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &nginxParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Nginx.g4"

	return this
}

// NginxParser tokens.
const (
	NginxParserEOF             = antlr.TokenEOF
	NginxParserT__0            = 1
	NginxParserT__1            = 2
	NginxParserT__2            = 3
	NginxParserT__3            = 4
	NginxParserT__4            = 5
	NginxParserT__5            = 6
	NginxParserT__6            = 7
	NginxParserT__7            = 8
	NginxParserT__8            = 9
	NginxParserT__9            = 10
	NginxParserT__10           = 11
	NginxParserT__11           = 12
	NginxParserT__12           = 13
	NginxParserT__13           = 14
	NginxParserLUA_HEADER      = 15
	NginxParserValue           = 16
	NginxParserSTR_EXT         = 17
	NginxParserLINE_COMMENT    = 18
	NginxParserREGEXP_PREFIXED = 19
	NginxParserQUOTED_STRING   = 20
	NginxParserSINGLE_QUOTED   = 21
	NginxParserWS              = 22
	NginxParserNEWLINE         = 23
)

// NginxParser rules.
const (
	NginxParserRULE_config               = 0
	NginxParserRULE_statement            = 1
	NginxParserRULE_genericStatement     = 2
	NginxParserRULE_regexHeaderStatement = 3
	NginxParserRULE_luaBlock             = 4
	NginxParserRULE_luaStatement         = 5
	NginxParserRULE_block                = 6
	NginxParserRULE_genericBlockHeader   = 7
	NginxParserRULE_ifStatement          = 8
	NginxParserRULE_ifBody               = 9
	NginxParserRULE_regexp               = 10
	NginxParserRULE_locationBlockHeader  = 11
	NginxParserRULE_rewriteStatement     = 12
)

// IConfigContext is an interface to support dynamic dispatch.
type IConfigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConfigContext differentiates from other interfaces.
	IsConfigContext()
}

type ConfigContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfigContext() *ConfigContext {
	var p = new(ConfigContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_config
	return p
}

func (*ConfigContext) IsConfigContext() {}

func NewConfigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfigContext {
	var p = new(ConfigContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_config

	return p
}

func (s *ConfigContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfigContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ConfigContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ConfigContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *ConfigContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ConfigContext) AllLuaBlock() []ILuaBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILuaBlockContext); ok {
			len++
		}
	}

	tst := make([]ILuaBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILuaBlockContext); ok {
			tst[i] = t.(ILuaBlockContext)
			i++
		}
	}

	return tst
}

func (s *ConfigContext) LuaBlock(i int) ILuaBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILuaBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILuaBlockContext)
}

func (s *ConfigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConfigContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitConfig(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) Config() (localctx IConfigContext) {
	this := p
	_ = this

	localctx = NewConfigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NginxParserRULE_config)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__8)|(1<<NginxParserT__9)|(1<<NginxParserLUA_HEADER)|(1<<NginxParserValue)|(1<<NginxParserREGEXP_PREFIXED))) != 0) {
		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(26)
				p.Statement()
			}

		case 2:
			{
				p.SetState(27)
				p.Block()
			}

		case 3:
			{
				p.SetState(28)
				p.LuaBlock()
			}

		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) RewriteStatement() IRewriteStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRewriteStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRewriteStatementContext)
}

func (s *StatementContext) GenericStatement() IGenericStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericStatementContext)
}

func (s *StatementContext) RegexHeaderStatement() IRegexHeaderStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexHeaderStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexHeaderStatementContext)
}

func (s *StatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(NginxParserNEWLINE, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NginxParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(36)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NginxParserT__9:
		{
			p.SetState(33)
			p.RewriteStatement()
		}

	case NginxParserValue:
		{
			p.SetState(34)
			p.GenericStatement()
		}

	case NginxParserREGEXP_PREFIXED:
		{
			p.SetState(35)
			p.RegexHeaderStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(38)
		p.Match(NginxParserT__0)
	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(39)
			p.Match(NginxParserNEWLINE)
		}

	}

	return localctx
}

// IGenericStatementContext is an interface to support dynamic dispatch.
type IGenericStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGenericStatementContext differentiates from other interfaces.
	IsGenericStatementContext()
}

type GenericStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenericStatementContext() *GenericStatementContext {
	var p = new(GenericStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_genericStatement
	return p
}

func (*GenericStatementContext) IsGenericStatementContext() {}

func NewGenericStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericStatementContext {
	var p = new(GenericStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_genericStatement

	return p
}

func (s *GenericStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericStatementContext) AllValue() []antlr.TerminalNode {
	return s.GetTokens(NginxParserValue)
}

func (s *GenericStatementContext) Value(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserValue, i)
}

func (s *GenericStatementContext) AllRegexp() []IRegexpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRegexpContext); ok {
			len++
		}
	}

	tst := make([]IRegexpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRegexpContext); ok {
			tst[i] = t.(IRegexpContext)
			i++
		}
	}

	return tst
}

func (s *GenericStatementContext) Regexp(i int) IRegexpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexpContext)
}

func (s *GenericStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenericStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitGenericStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) GenericStatement() (localctx IGenericStatementContext) {
	this := p
	_ = this

	localctx = NewGenericStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NginxParserRULE_genericStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(NginxParserValue)
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__5)|(1<<NginxParserT__6)|(1<<NginxParserT__7)|(1<<NginxParserValue))) != 0 {
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(43)
				p.Match(NginxParserValue)
			}

		case 2:
			{
				p.SetState(44)
				p.Regexp()
			}

		}

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRegexHeaderStatementContext is an interface to support dynamic dispatch.
type IRegexHeaderStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegexHeaderStatementContext differentiates from other interfaces.
	IsRegexHeaderStatementContext()
}

type RegexHeaderStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegexHeaderStatementContext() *RegexHeaderStatementContext {
	var p = new(RegexHeaderStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_regexHeaderStatement
	return p
}

func (*RegexHeaderStatementContext) IsRegexHeaderStatementContext() {}

func NewRegexHeaderStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegexHeaderStatementContext {
	var p = new(RegexHeaderStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_regexHeaderStatement

	return p
}

func (s *RegexHeaderStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *RegexHeaderStatementContext) REGEXP_PREFIXED() antlr.TerminalNode {
	return s.GetToken(NginxParserREGEXP_PREFIXED, 0)
}

func (s *RegexHeaderStatementContext) Value() antlr.TerminalNode {
	return s.GetToken(NginxParserValue, 0)
}

func (s *RegexHeaderStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexHeaderStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegexHeaderStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitRegexHeaderStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) RegexHeaderStatement() (localctx IRegexHeaderStatementContext) {
	this := p
	_ = this

	localctx = NewRegexHeaderStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NginxParserRULE_regexHeaderStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(NginxParserREGEXP_PREFIXED)
	}
	{
		p.SetState(51)
		p.Match(NginxParserValue)
	}

	return localctx
}

// ILuaBlockContext is an interface to support dynamic dispatch.
type ILuaBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLuaBlockContext differentiates from other interfaces.
	IsLuaBlockContext()
}

type LuaBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLuaBlockContext() *LuaBlockContext {
	var p = new(LuaBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_luaBlock
	return p
}

func (*LuaBlockContext) IsLuaBlockContext() {}

func NewLuaBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LuaBlockContext {
	var p = new(LuaBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_luaBlock

	return p
}

func (s *LuaBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *LuaBlockContext) LUA_HEADER() antlr.TerminalNode {
	return s.GetToken(NginxParserLUA_HEADER, 0)
}

func (s *LuaBlockContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NginxParserNEWLINE)
}

func (s *LuaBlockContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserNEWLINE, i)
}

func (s *LuaBlockContext) AllLuaStatement() []ILuaStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILuaStatementContext); ok {
			len++
		}
	}

	tst := make([]ILuaStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILuaStatementContext); ok {
			tst[i] = t.(ILuaStatementContext)
			i++
		}
	}

	return tst
}

func (s *LuaBlockContext) LuaStatement(i int) ILuaStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILuaStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILuaStatementContext)
}

func (s *LuaBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LuaBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LuaBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitLuaBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) LuaBlock() (localctx ILuaBlockContext) {
	this := p
	_ = this

	localctx = NewLuaBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NginxParserRULE_luaBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(NginxParserLUA_HEADER)
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(54)
			p.Match(NginxParserNEWLINE)
		}

	}
	{
		p.SetState(57)
		p.Match(NginxParserT__1)
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(58)
			p.Match(NginxParserNEWLINE)
		}

	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__0)|(1<<NginxParserT__3)|(1<<NginxParserT__4)|(1<<NginxParserT__5)|(1<<NginxParserValue))) != 0 {
		{
			p.SetState(61)
			p.LuaStatement()
		}

		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(67)
		p.Match(NginxParserT__2)
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(68)
			p.Match(NginxParserNEWLINE)
		}

	}

	return localctx
}

// ILuaStatementContext is an interface to support dynamic dispatch.
type ILuaStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLuaStatementContext differentiates from other interfaces.
	IsLuaStatementContext()
}

type LuaStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLuaStatementContext() *LuaStatementContext {
	var p = new(LuaStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_luaStatement
	return p
}

func (*LuaStatementContext) IsLuaStatementContext() {}

func NewLuaStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LuaStatementContext {
	var p = new(LuaStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_luaStatement

	return p
}

func (s *LuaStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LuaStatementContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(NginxParserNEWLINE, 0)
}

func (s *LuaStatementContext) AllValue() []antlr.TerminalNode {
	return s.GetTokens(NginxParserValue)
}

func (s *LuaStatementContext) Value(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserValue, i)
}

func (s *LuaStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LuaStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LuaStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitLuaStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) LuaStatement() (localctx ILuaStatementContext) {
	this := p
	_ = this

	localctx = NewLuaStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NginxParserRULE_luaStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(80)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NginxParserT__3, NginxParserT__4, NginxParserT__5, NginxParserValue:
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(71)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__3)|(1<<NginxParserT__4)|(1<<NginxParserT__5)|(1<<NginxParserValue))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(74)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(76)
				p.Match(NginxParserT__0)
			}

		}

	case NginxParserT__0:
		{
			p.SetState(79)
			p.Match(NginxParserT__0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(82)
			p.Match(NginxParserNEWLINE)
		}

	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LocationBlockHeader() ILocationBlockHeaderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocationBlockHeaderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocationBlockHeaderContext)
}

func (s *BlockContext) GenericBlockHeader() IGenericBlockHeaderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenericBlockHeaderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenericBlockHeaderContext)
}

func (s *BlockContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NginxParserNEWLINE)
}

func (s *BlockContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserNEWLINE, i)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *BlockContext) AllIfStatement() []IIfStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIfStatementContext); ok {
			len++
		}
	}

	tst := make([]IIfStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIfStatementContext); ok {
			tst[i] = t.(IIfStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) IfStatement(i int) IIfStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *BlockContext) AllLuaBlock() []ILuaBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILuaBlockContext); ok {
			len++
		}
	}

	tst := make([]ILuaBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILuaBlockContext); ok {
			tst[i] = t.(ILuaBlockContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) LuaBlock(i int) ILuaBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILuaBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILuaBlockContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NginxParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(87)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NginxParserT__8:
		{
			p.SetState(85)
			p.LocationBlockHeader()
		}

	case NginxParserValue:
		{
			p.SetState(86)
			p.GenericBlockHeader()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(89)
			p.Match(NginxParserNEWLINE)
		}

	}
	{
		p.SetState(92)
		p.Match(NginxParserT__1)
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(93)
			p.Match(NginxParserNEWLINE)
		}

	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__3)|(1<<NginxParserT__8)|(1<<NginxParserT__9)|(1<<NginxParserLUA_HEADER)|(1<<NginxParserValue)|(1<<NginxParserREGEXP_PREFIXED))) != 0 {
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(96)
				p.Statement()
			}

		case 2:
			{
				p.SetState(97)
				p.Block()
			}

		case 3:
			{
				p.SetState(98)
				p.IfStatement()
			}

		case 4:
			{
				p.SetState(99)
				p.LuaBlock()
			}

		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NginxParserNEWLINE {
			{
				p.SetState(102)
				p.Match(NginxParserNEWLINE)
			}

		}

		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(110)
		p.Match(NginxParserT__2)
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(111)
			p.Match(NginxParserNEWLINE)
		}

	}

	return localctx
}

// IGenericBlockHeaderContext is an interface to support dynamic dispatch.
type IGenericBlockHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGenericBlockHeaderContext differentiates from other interfaces.
	IsGenericBlockHeaderContext()
}

type GenericBlockHeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenericBlockHeaderContext() *GenericBlockHeaderContext {
	var p = new(GenericBlockHeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_genericBlockHeader
	return p
}

func (*GenericBlockHeaderContext) IsGenericBlockHeaderContext() {}

func NewGenericBlockHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenericBlockHeaderContext {
	var p = new(GenericBlockHeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_genericBlockHeader

	return p
}

func (s *GenericBlockHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *GenericBlockHeaderContext) AllValue() []antlr.TerminalNode {
	return s.GetTokens(NginxParserValue)
}

func (s *GenericBlockHeaderContext) Value(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserValue, i)
}

func (s *GenericBlockHeaderContext) AllRegexp() []IRegexpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRegexpContext); ok {
			len++
		}
	}

	tst := make([]IRegexpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRegexpContext); ok {
			tst[i] = t.(IRegexpContext)
			i++
		}
	}

	return tst
}

func (s *GenericBlockHeaderContext) Regexp(i int) IRegexpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexpContext)
}

func (s *GenericBlockHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenericBlockHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenericBlockHeaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitGenericBlockHeader(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) GenericBlockHeader() (localctx IGenericBlockHeaderContext) {
	this := p
	_ = this

	localctx = NewGenericBlockHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NginxParserRULE_genericBlockHeader)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(NginxParserValue)
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__5)|(1<<NginxParserT__6)|(1<<NginxParserT__7)|(1<<NginxParserValue))) != 0 {
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(115)
				p.Match(NginxParserValue)
			}

		case 2:
			{
				p.SetState(116)
				p.Regexp()
			}

		}

		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IfBody() IIfBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBodyContext)
}

func (s *IfStatementContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NginxParserNEWLINE)
}

func (s *IfStatementContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserNEWLINE, i)
}

func (s *IfStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) IfStatement() (localctx IIfStatementContext) {
	this := p
	_ = this

	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NginxParserRULE_ifStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(NginxParserT__3)
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(123)
			p.Match(NginxParserNEWLINE)
		}

	}
	{
		p.SetState(126)
		p.IfBody()
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(127)
			p.Match(NginxParserNEWLINE)
		}

	}
	{
		p.SetState(130)
		p.Match(NginxParserT__1)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(131)
			p.Match(NginxParserNEWLINE)
		}

	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__9)|(1<<NginxParserValue)|(1<<NginxParserREGEXP_PREFIXED))) != 0 {
		{
			p.SetState(134)
			p.Statement()
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NginxParserNEWLINE {
			{
				p.SetState(135)
				p.Match(NginxParserNEWLINE)
			}

		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(143)
		p.Match(NginxParserT__2)
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(144)
			p.Match(NginxParserNEWLINE)
		}

	}

	return localctx
}

// IIfBodyContext is an interface to support dynamic dispatch.
type IIfBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfBodyContext differentiates from other interfaces.
	IsIfBodyContext()
}

type IfBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfBodyContext() *IfBodyContext {
	var p = new(IfBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_ifBody
	return p
}

func (*IfBodyContext) IsIfBodyContext() {}

func NewIfBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfBodyContext {
	var p = new(IfBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_ifBody

	return p
}

func (s *IfBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *IfBodyContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(NginxParserNEWLINE)
}

func (s *IfBodyContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserNEWLINE, i)
}

func (s *IfBodyContext) AllValue() []antlr.TerminalNode {
	return s.GetTokens(NginxParserValue)
}

func (s *IfBodyContext) Value(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserValue, i)
}

func (s *IfBodyContext) Regexp() IRegexpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexpContext)
}

func (s *IfBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitIfBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) IfBody() (localctx IIfBodyContext) {
	this := p
	_ = this

	localctx = NewIfBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, NginxParserRULE_ifBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(NginxParserT__5)
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(148)
			p.Match(NginxParserNEWLINE)
		}

	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(151)
				p.Match(NginxParserValue)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(156)
			p.Match(NginxParserNEWLINE)
		}

	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(159)
			p.Match(NginxParserValue)
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(160)
			p.Regexp()
		}

	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NginxParserNEWLINE {
		{
			p.SetState(163)
			p.Match(NginxParserNEWLINE)
		}

	}
	{
		p.SetState(166)
		p.Match(NginxParserT__4)
	}

	return localctx
}

// IRegexpContext is an interface to support dynamic dispatch.
type IRegexpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegexpContext differentiates from other interfaces.
	IsRegexpContext()
}

type RegexpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegexpContext() *RegexpContext {
	var p = new(RegexpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_regexp
	return p
}

func (*RegexpContext) IsRegexpContext() {}

func NewRegexpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegexpContext {
	var p = new(RegexpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_regexp

	return p
}

func (s *RegexpContext) GetParser() antlr.Parser { return s.parser }

func (s *RegexpContext) AllValue() []antlr.TerminalNode {
	return s.GetTokens(NginxParserValue)
}

func (s *RegexpContext) Value(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserValue, i)
}

func (s *RegexpContext) AllRegexp() []IRegexpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRegexpContext); ok {
			len++
		}
	}

	tst := make([]IRegexpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRegexpContext); ok {
			tst[i] = t.(IRegexpContext)
			i++
		}
	}

	return tst
}

func (s *RegexpContext) Regexp(i int) IRegexpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexpContext)
}

func (s *RegexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegexpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitRegexp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) Regexp() (localctx IRegexpContext) {
	this := p
	_ = this

	localctx = NewRegexpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, NginxParserRULE_regexp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(175)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case NginxParserT__6:
				{
					p.SetState(168)
					p.Match(NginxParserT__6)
				}

			case NginxParserT__7:
				{
					p.SetState(169)
					p.Match(NginxParserT__7)
				}

			case NginxParserValue:
				{
					p.SetState(170)
					p.Match(NginxParserValue)
				}

			case NginxParserT__5:
				{
					p.SetState(171)
					p.Match(NginxParserT__5)
				}
				{
					p.SetState(172)
					p.Regexp()
				}
				{
					p.SetState(173)
					p.Match(NginxParserT__4)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())
	}

	return localctx
}

// ILocationBlockHeaderContext is an interface to support dynamic dispatch.
type ILocationBlockHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocationBlockHeaderContext differentiates from other interfaces.
	IsLocationBlockHeaderContext()
}

type LocationBlockHeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocationBlockHeaderContext() *LocationBlockHeaderContext {
	var p = new(LocationBlockHeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_locationBlockHeader
	return p
}

func (*LocationBlockHeaderContext) IsLocationBlockHeaderContext() {}

func NewLocationBlockHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocationBlockHeaderContext {
	var p = new(LocationBlockHeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_locationBlockHeader

	return p
}

func (s *LocationBlockHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *LocationBlockHeaderContext) AllValue() []antlr.TerminalNode {
	return s.GetTokens(NginxParserValue)
}

func (s *LocationBlockHeaderContext) Value(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserValue, i)
}

func (s *LocationBlockHeaderContext) Regexp() IRegexpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexpContext)
}

func (s *LocationBlockHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocationBlockHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocationBlockHeaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitLocationBlockHeader(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) LocationBlockHeader() (localctx ILocationBlockHeaderContext) {
	this := p
	_ = this

	localctx = NewLocationBlockHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, NginxParserRULE_locationBlockHeader)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(NginxParserT__8)
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(180)
			p.Match(NginxParserValue)
		}

	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(183)
			p.Match(NginxParserValue)
		}

	case 2:
		{
			p.SetState(184)
			p.Regexp()
		}

	}

	return localctx
}

// IRewriteStatementContext is an interface to support dynamic dispatch.
type IRewriteStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRewriteStatementContext differentiates from other interfaces.
	IsRewriteStatementContext()
}

type RewriteStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRewriteStatementContext() *RewriteStatementContext {
	var p = new(RewriteStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NginxParserRULE_rewriteStatement
	return p
}

func (*RewriteStatementContext) IsRewriteStatementContext() {}

func NewRewriteStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RewriteStatementContext {
	var p = new(RewriteStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NginxParserRULE_rewriteStatement

	return p
}

func (s *RewriteStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *RewriteStatementContext) AllValue() []antlr.TerminalNode {
	return s.GetTokens(NginxParserValue)
}

func (s *RewriteStatementContext) Value(i int) antlr.TerminalNode {
	return s.GetToken(NginxParserValue, i)
}

func (s *RewriteStatementContext) Regexp() IRegexpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexpContext)
}

func (s *RewriteStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RewriteStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RewriteStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NginxVisitor:
		return t.VisitRewriteStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NginxParser) RewriteStatement() (localctx IRewriteStatementContext) {
	this := p
	_ = this

	localctx = NewRewriteStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, NginxParserRULE_rewriteStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Match(NginxParserT__9)
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(188)
			p.Match(NginxParserValue)
		}

	case 2:
		{
			p.SetState(189)
			p.Regexp()
		}

	}
	{
		p.SetState(192)
		p.Match(NginxParserValue)
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__10)|(1<<NginxParserT__11)|(1<<NginxParserT__12)|(1<<NginxParserT__13))) != 0 {
		{
			p.SetState(193)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NginxParserT__10)|(1<<NginxParserT__11)|(1<<NginxParserT__12)|(1<<NginxParserT__13))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}
