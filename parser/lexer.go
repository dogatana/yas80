package parser

import (
	"bufio"
	"fmt"
	"strings"
)

const EOF = 0

// 最低限必要な構造体を定義
type Lexer struct {
	scanner    *bufio.Scanner
	isEOF      bool
	text       []rune
	index      int
	curChar    rune
	lineNumber int
}

// yyLexer インターフェースメソッド
func (l *Lexer) Lex(lval *yySymType) int {
	tok := l.NextToken()
	lval.token = tok
	// fmt.Println("[token]", tok.String())
	if tok.Type == INVALID {
		return int([]rune(tok.Literal)[0])
	} else {
		return tok.Type
	}
}

// yyLexer インターフェースメソッド
func (l *Lexer) Error(s string) {
	msg := strings.Replace(s, "unexpected", "ここでは使用不可", 1)
	fmt.Println("[error]", msg)
}

func NewLexer(r *bufio.Reader) *Lexer {
	l := &Lexer{scanner: bufio.NewScanner(r)}
	l.nextChar()
	return l
}

func (l *Lexer) NextToken() Token {
	// 空白をスキップ
	l.skipWhitespace()

	var literal string
	var ch rune
	// var tokType int
	switch {
	case l.curChar == EOF:
		// EOF
		return Token{Type: 0, Literal: "[EOF]"}
	case l.curChar == '\n':
		// EOL
		l.nextChar()
		return Token{Type: int(EOL), Literal: "\\n"}
	case l.curChar == '-': // 単項、2項の両方あるので個別トークンとする
		ch := l.curChar
		l.nextChar()
		return Token{Type: int(ch), Literal: string(ch)}
	case l.curChar == '+' || l.curChar == '^':
		ch := l.curChar
		l.nextChar()
		return Token{Type: ADD, SubType: int(ch), Literal: string(ch)}
	case l.curChar == '*' || l.curChar == '/':
		ch := l.curChar
		l.nextChar()
		return Token{Type: MULDIV, SubType: int(ch), Literal: string(ch)}
	case l.isTowCharTokenStart(l.curChar):
		return l.checkTwoCharToken(l.curChar)
	case l.curChar == '(' || l.curChar == ')':
		// 1文字トークン
		ch = l.curChar
		l.nextChar()
		return Token{Type: int(ch), Literal: string(ch)}
	case l.curChar == '&':
		ch = l.curChar
		l.nextChar()
		return Token{Type: MULDIV, SubType: int(ch), Literal: string(ch)}
	case l.curChar == '~':
		ch = l.curChar
		l.nextChar()
		return Token{Type: UNARY, SubType: int(ch), Literal: string(ch)}
	case l.curChar == '0' && (l.peekChar() == 'x' || l.peekChar() == 'X'):
		// 16進数リテラル(0x)
		l.nextChar() // '0'をスキップ
		literal = "0" + string(l.curChar)
		l.nextChar() // 'x'または'X'をスキップ
		literal += l.readWord()
		l.nextChar()
		return Token{Type: NUMBER, Literal: literal}
	case l.curChar == '$' || l.curChar == '%':
		// 16進数リテラル($) or 2進数リテラル
		l.nextChar() // '$'をスキップ
		literal = "$" + l.readWord()
		l.nextChar()
		return Token{Type: NUMBER, Literal: literal}
	case l.isDigit(l.curChar):
		// 10進、16進(0x)、2進(0b)、8進（0o)リテラル
		literal = l.readWord()
		l.nextChar()
		return Token{Type: NUMBER, Literal: literal}
	case l.isAlpha(l.curChar):
		// IDENT、予約語
		literal = l.readWord()
		peek := l.peekChar()
		if peek == '@' || peek == '.' {
			l.nextChar()
			l.nextChar()
			literal += string(peek) + l.readWord()
		}
		l.nextChar()
		// AF’の対処
		if len(literal) == 2 && l.curChar == '\'' {
			literal += "'"
			l.nextChar()
		}
		tok, ok := l.checkZ80ReservedWord(literal)
		// fmt.Println("search", literal, tok, ok)
		if ok {
			return tok
		}
		return Token{Type: IDENT, Literal: literal}
	case l.curChar == '@' || l.curChar == '.':
		literal = string(l.curChar)
		l.nextChar()
		literal += l.readWord()
		l.nextChar()
		return Token{Type: IDENT, Literal: literal}
	default:
		literal = string(l.curChar)
		l.nextChar()
		return Token{Type: INVALID, Literal: literal}
	}
}

// 2文字トークンの最初の文字か
func (l *Lexer) isTowCharTokenStart(c rune) bool {
	return c == '=' || c == '<' || c == '>' || c == '!' || c == '|' || c == '&'
}

// 2文字トークンのチェック
func (l *Lexer) checkTwoCharToken(ch1 rune) Token {
	ch2 := l.peekChar()
	l.nextChar()

	var tok Token
	switch {
	// COMP
	case ch1 == '<' && ch2 == '=':
		tok = Token{Type: COMP, SubType: LE, Literal: "<="}
	case ch1 == '<' && ch2 == '<':
		tok = Token{Type: SHIFT, SubType: SL, Literal: "<<"}
	case ch1 == '>' && ch2 == '=':
		tok = Token{Type: COMP, SubType: GE, Literal: ">="}
	case ch1 == '>' && ch2 == '>':
		tok = Token{Type: SHIFT, SubType: SR, Literal: ">>"}
	case ch1 == '<' || ch1 == '>':
		tok = Token{Type: COMP, SubType: int(ch1), Literal: string(ch1)}
	case ch1 == '=' && ch2 == '=':
		tok = Token{Type: COMP, SubType: EQ, Literal: "=="}
	case ch1 == '!' && ch2 == '=':
		tok = Token{Type: COMP, SubType: NEQ, Literal: "!="}
	case ch1 == '!':
		return Token{Type: UNARY, SubType: int(ch1), Literal: string(ch1)}
	case ch1 == '&' && ch2 == '&':
		tok = Token{Type: AND, SubType: 0, Literal: "&&"}
	case ch1 == '&':
		return Token{Type: MULDIV, SubType: int(ch1), Literal: string(ch1)}
	case ch1 == '|' && ch2 == '|':
		tok = Token{Type: OR, SubType: 0, Literal: "||"}
	case ch1 == '|':
		return Token{Type: ADD, SubType: int(ch1), Literal: string(ch1)}
	default:
		// 1文字トークンを返す
		return Token{Type: int(ch1), Literal: string(rune(ch1))}
	}
	l.nextChar()
	return tok
}

func (l *Lexer) checkZ80ReservedWord(literal string) (Token, bool) {
	var (
		tok Token
		ok  bool
	)

	// 大文字化して検索
	word := strings.ToUpper(literal)

	// Z80 レジスタ、フラグ
	tok, ok = Z80Registers[word]
	if ok {
		tok.Literal = literal
		return tok, ok
	}
	// Z80 命令
	tok, ok = Z80OpCodes[word]
	if ok {
		tok.Literal = literal
		return tok, ok
	}
	return Token{}, false
}

func (l *Lexer) nextChar() {
	if l.isEOF {
		l.curChar = EOF
		return
	}
	if l.index == 0 {
		scanned := l.scanner.Scan()
		if !scanned {
			if l.curChar != '\n' {
				l.curChar = '\n'
				return
			}
			l.curChar = EOF
			l.isEOF = true
			return
		}
		l.lineNumber++
		l.text = []rune(l.scanner.Text())
	}
	if l.index >= len(l.text) {
		l.curChar = '\n'
		l.index = 0
		return
	}
	l.curChar = l.text[l.index]
	l.index++
}

func (l *Lexer) peekChar() rune {
	if l.index >= len(l.text) {
		return 0
	}
	return l.text[l.index]
}

func (l *Lexer) skipWhitespace() {
	for !l.isEOF && (l.curChar == ' ' || l.curChar == '\t') {
		l.nextChar()
	}
}

func (l *Lexer) readNumber() string {
	startIndex := l.index - 1
	for l.index < len(l.text) && l.isWordChar(l.text[l.index]) {
		l.index++
	}
	return string(l.text[startIndex:l.index])
}

func (l *Lexer) readWord() string {
	startIndex := l.index - 1
	for l.index < len(l.text) && l.isWordChar(l.text[l.index]) {
		l.index++
	}
	return string(l.text[startIndex:l.index])
}

func (l *Lexer) isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) isAlpha(ch rune) bool {
	return 'A' <= ch && ch <= 'Z' || 'a' <= ch && ch <= 'z' || ch == '_'
}

func (l *Lexer) isWordChar(ch rune) bool {
	return l.isDigit(ch) || l.isAlpha(ch)
}

func (l *Lexer) isOneCharToken(ch rune) bool {
	return ch == '(' || ch == ')' || ch == ','
}
