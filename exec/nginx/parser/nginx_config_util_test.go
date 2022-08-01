package parser

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		isFile bool
	}{
		{
			name:   "testConfigParser",
			input:  "test.conf",
			isFile: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var lexer *NginxLexer
			if tt.isFile {
				input, err := antlr.NewFileStream(tt.input)
				if err != nil {
					t.Errorf(fmt.Sprintf("parser test err: %s", err))
				}
				lexer = NewNginxLexer(input)
			} else {
				input := antlr.NewInputStream(tt.input)
				lexer = NewNginxLexer(input)
			}

			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := NewNginxParser(stream)
			//p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
			p.BuildParseTrees = true
			tree := p.Config()
			//fmt.Println(tree.ToStringTree(nil, p))
			visitor := newMappingVisitor()
			config := tree.Accept(visitor).(*Config)
			if config == nil {
				t.Errorf("LoadConfig() err")
			}
		})
	}
}
