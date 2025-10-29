package parser

import (
	"bufio"
	"fmt"
	"strings"
	"yas80/errorstore"
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
	fileNmae   string
	ErrorStore *errorstore.ErrorStore
}

// yyLexer インターフェースメソッド
func (l *Lexer) Lex(lval *yySymType) int {
	tok := l.NextToken()
	lval.token = tok
	// fmt.Println("[token]", tok.String())
	if tok.TokenType == INVALID {
		return int([]rune(tok.Literal)[0])
	} else {
		return int(tok.TokenType)
	}
}

// yyLexer インターフェースメソッド
func (l *Lexer) Error(s string, args ...any) {
	var line int

	switch len(args) {
	case 0:
		line = l.lineNumber
	case 1:
		n, ok := args[0].(int)
		if !ok {
			panic(fmt.Sprintf("invalid argument for Lexer.Error(string, %T)", args[0]))
		}
		line = n
	default:
		panic(fmt.Sprintf("too much args for Lexer.Error() %#v", args))
	}

	if strings.HasPrefix(s, "[W]") {
		l.ErrorStore.AddWarning(l.fileNmae, line, s[3:])
	} else {
		l.ErrorStore.AddError(l.fileNmae, line, s)
	}
}

func NewLexer(r *bufio.Reader, filename string, es *errorstore.ErrorStore) *Lexer {
	l := &Lexer{scanner: bufio.NewScanner(r)}
	l.nextChar()
	l.fileNmae = filename
	l.ErrorStore = es
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
		return Token{TokenType: 0, Literal: "[EOF]", LineNumber: l.lineNumber}
	case l.curChar == ';':
		// コメント
		for l.curChar != '\n' && l.curChar != EOF {
			l.nextChar()
		}
		l.nextChar()
		return Token{TokenType: EOL, Literal: "\\n", LineNumber: l.lineNumber}
	case l.curChar == '\\':
		// マルチステートメント
		tok := Token{TokenType: EOL, Literal: "\\", LineNumber: l.lineNumber}
		l.nextChar()
		return tok
	case l.curChar == '\n':
		// EOL
		l.nextChar()
		return Token{TokenType: EOL, Literal: "\\n", LineNumber: l.lineNumber}
	case l.curChar == '-': // 単項、2項の両方あるので個別トークンとする
		ch := l.curChar
		l.nextChar()
		return Token{TokenType: TokenType(ch), Literal: string(ch), LineNumber: l.lineNumber}
	case l.curChar == '+' || l.curChar == '^':
		ch := l.curChar
		l.nextChar()
		return Token{TokenType: ADD, TokenSubType: TokenSubType(ch), Literal: string(ch), LineNumber: l.lineNumber}
	case l.curChar == '*' || l.curChar == '/':
		ch := l.curChar
		l.nextChar()
		return Token{TokenType: MULDIV, TokenSubType: TokenSubType(ch), Literal: string(ch), LineNumber: l.lineNumber}
	case l.isTowCharTokenStart(l.curChar):
		tok := l.checkTwoCharToken(l.curChar)
		tok.LineNumber = l.lineNumber
		return tok
	case l.curChar == '(' || l.curChar == ')':
		// 1文字トークン
		ch = l.curChar
		l.nextChar()
		return Token{TokenType: TokenType(ch), Literal: string(ch), LineNumber: l.lineNumber}
	case l.curChar == '&':
		ch = l.curChar
		l.nextChar()
		return Token{TokenType: MULDIV, TokenSubType: TokenSubType(ch), Literal: string(ch), LineNumber: l.lineNumber}
	case l.curChar == '~':
		ch = l.curChar
		l.nextChar()
		return Token{TokenType: UNARY, TokenSubType: TokenSubType(ch), Literal: string(ch), LineNumber: l.lineNumber}
	case l.curChar == '0' && (l.peekChar() == 'x' || l.peekChar() == 'X'):
		// 16進数リテラル(0x)
		l.nextChar() // '0'をスキップ
		literal = "0" + string(l.curChar)
		l.nextChar() // 'x'または'X'をスキップ
		literal += l.readWord()
		l.nextChar()
		return Token{TokenType: NUMBER, Literal: literal, LineNumber: l.lineNumber}
	case l.curChar == '$' || l.curChar == '%':
		// 16進数リテラル($) or 2進数リテラル
		literal = string(l.curChar)
		l.nextChar() // '$'をスキップ
		literal += l.readWord()
		l.nextChar()
		return Token{TokenType: NUMBER, Literal: literal, LineNumber: l.lineNumber}
	case l.isDigit(l.curChar):
		// 10進、16進(0x)、2進(0b)、8進（0o)リテラル
		literal = l.readWord()
		l.nextChar()
		return Token{TokenType: NUMBER, Literal: literal, LineNumber: l.lineNumber}
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
		// z80 予約語
		tok, ok := z80ReservedWords[strings.ToUpper(literal)]
		if ok {
			tok.LineNumber = l.lineNumber
			return tok
		}
		// これ以外は識別子
		return Token{TokenType: IDENT, Literal: literal}
	case l.curChar == '@' || l.curChar == '.':
		literal = string(l.curChar)
		l.nextChar()
		literal += l.readWord()
		l.nextChar()
		return Token{TokenType: IDENT, Literal: literal, LineNumber: l.lineNumber}
	default:
		literal = string(l.curChar)
		l.nextChar()
		return Token{TokenType: INVALID, Literal: literal, LineNumber: l.lineNumber}
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
		tok = Token{TokenType: COMP, TokenSubType: LE, Literal: "<="}
	case ch1 == '<' && ch2 == '<':
		tok = Token{TokenType: SHIFT, TokenSubType: SL, Literal: "<<"}
	case ch1 == '>' && ch2 == '=':
		tok = Token{TokenType: COMP, TokenSubType: GE, Literal: ">="}
	case ch1 == '>' && ch2 == '>':
		tok = Token{TokenType: SHIFT, TokenSubType: SR, Literal: ">>"}
	case ch1 == '<' || ch1 == '>':
		tok = Token{TokenType: COMP, TokenSubType: TokenSubType(ch1), Literal: string(ch1)}
	case ch1 == '=' && ch2 == '=':
		tok = Token{TokenType: COMP, TokenSubType: EQ, Literal: "=="}
	case ch1 == '!' && ch2 == '=':
		tok = Token{TokenType: COMP, TokenSubType: NEQ, Literal: "!="}
	case ch1 == '!':
		return Token{TokenType: UNARY, TokenSubType: TokenSubType(ch1), Literal: string(ch1)}
	case ch1 == '&' && ch2 == '&':
		tok = Token{TokenType: AND, TokenSubType: 0, Literal: "&&"}
	case ch1 == '&':
		return Token{TokenType: MULDIV, TokenSubType: TokenSubType(ch1), Literal: string(ch1)}
	case ch1 == '|' && ch2 == '|':
		tok = Token{TokenType: OR, TokenSubType: 0, Literal: "||"}
	case ch1 == '|':
		return Token{TokenType: ADD, TokenSubType: TokenSubType(ch1), Literal: string(ch1)}
	default:
		// 1文字トークンを返す
		return Token{TokenType: TokenType(ch1), Literal: string(rune(ch1))}
	}
	l.nextChar()
	return tok
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
