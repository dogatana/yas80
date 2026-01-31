package parser

import (
	"strings"
	"testing"
	"yas80/filecontent"
	"yas80/logging"
)

func newLexerForTest(input string) *Lexer {
	file := "<string>"
	logger := logging.New(file)
	fc := filecontent.New(file, []byte(input))
	lex := NewLexer(logger, func() *filecontent.FileContent {
		ret := fc
		fc = nil
		return ret
	})

	// ./parser/*_test.go の後方互換性のため、最初の FILE トークンを読み飛ばす
	var lval yySymType
	lex.Lex(&lval)
	return lex
}

func ParseForTest(t *testing.T, lexer *Lexer, tn int) *BlockStatement {
	prog := Parse(lexer)
	return PreProrocess(lexer.logger, prog)
}

func nextToken(l *Lexer) Token {
	var lval yySymType
	l.Lex(&lval)
	return lval.token
}

func testInputEnd(t *testing.T, tn int, lexer *Lexer) {
	// EOL
	for {
		tok := lexer.NextToken()
		if tok.TokenType == EOL {
			break
		}
		if tok.TokenType == EOF {
			t.Errorf("[%d] no EOL before EOF", tn)
			return
		}
	}
	for {
		tok := lexer.NextToken()
		if tok.TokenType == EOL {
			continue
		}
		if tok.TokenType == EOF {
			return
		}
		t.Errorf("[%d] not EOF. got %s", tn, tok.String())
	}
}

// prog が 単一 Statement かどうか
func progHasOnlyOneStatement(t *testing.T, tn int, prog *BlockStatement) Statement {
	if len(prog.Block) != 1 {
		t.Fatalf("[%d] returns %d statements. not 1", tn, len(prog.Block))
	}
	return prog.Block[0]
}

func splitTrim(input string) string {
	strs := strings.Split(strings.ReplaceAll(input, "\\", "\n"), "\n")
	ret := []string{}
	for _, s := range strs {
		str := strings.Trim(s, " \n\t")
		if str != "" {
			ret = append(ret, str)
		}
	}
	return strings.Join(ret, "\n")
}

func testAssignStatement(t *testing.T, tn int, node Node) *AssignStatement {
	stmt, ok := node.(*AssignStatement)
	if !ok {
		t.Fatalf("[%d] not *AssignStatement. got %T", tn, node)
	}
	return stmt
}

func testNumberLiteral(t *testing.T, tn int, node Node, expected int) {
	literal, ok := node.(*NumberLiteral)
	if !ok {
		t.Errorf("[%d] not *NumberLiteral. got %T", tn, literal)
	}
	if literal.Value != expected {
		t.Errorf("[%d] not %d. got %d", tn, expected, literal.Value)
	}
}

func testStringLiteral(t *testing.T, tn int, node Node, expected string) {
	literal, ok := node.(*StringLiteral)
	if !ok {
		t.Errorf("[%d] not *StringLiteral. got %T", tn, literal)
	}
	if literal.Value != expected {
		t.Errorf("[%d] not %q. got %q", tn, expected, literal.Value)
	}
}
