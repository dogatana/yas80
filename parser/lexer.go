package parser

import (
	"bufio"
	"fmt"
	"strings"
	"yas80/logger"
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
	logger     *logger.Logger
	program    *Program
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
func (l *Lexer) Error(msg string, args ...any) {
	var line int
	var yyVAL, yylval *Token

	switch len(args) {
	case 0: // Error(msg)
		line = l.lineNumber
	case 1: // Error(msg, lineNumber)
		n, ok := args[0].(int)
		if !ok {
			panic(fmt.Sprintf("[SYSTEM] invalid argument for Lexer.Error(string, %T)", args[0]))
		}
		line = n
	case 2:
		line = l.lineNumber
		yyVAL = args[0].(*Token)
		yylval = args[1].(*Token)
		msg = modifyYaccError(msg, yyVAL, yylval)
	default:
		panic(fmt.Sprintf("[SYSTEM] too much args for Lexer.Error() %#v", args))
	}

	if strings.HasPrefix(msg, "[I]") {
		l.logger.Info(msg[3:], line)
	} else if strings.HasPrefix(msg, "[W]") {
		l.logger.Warning(msg[3:], line)
	} else if strings.HasPrefix(msg, "[E]") {
		l.logger.Error(msg[3:], line)
	} else if msg[0] == '[' {
		l.logger.Error(msg, line)
	} else {
		l.logger.Error(msg, line)
	}
}

func NewLexer(r *bufio.Reader, filename string, logger *logger.Logger) *Lexer {
	l := &Lexer{scanner: bufio.NewScanner(r), program: &Program{}}
	l.nextChar()
	l.fileNmae = filename
	l.logger = logger
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
	case l.curChar == '"':
		// 文字列リテラル
		s := l.readString()
		l.nextChar()
		l.nextChar()
		return Token{TokenType: STRING, Literal: s, LineNumber: l.lineNumber}
	case l.curChar == '-': // 単項・2項両方あるので1文字トークンとする
		ch := l.curChar
		l.nextChar()
		return Token{TokenType: TokenType(ch), Literal: string(ch), LineNumber: l.lineNumber}
	case l.curChar == '+' || l.curChar == '^': // ADDSUB
		ch := l.curChar
		l.nextChar()
		return Token{TokenType: ADDSUB, TokenSubType: TokenSubType(ch), Literal: string(ch), LineNumber: l.lineNumber}
	case l.curChar == '*' || l.curChar == '/': // MULDIV
		ch := l.curChar
		l.nextChar()
		return Token{TokenType: MULDIV, TokenSubType: TokenSubType(ch), Literal: string(ch), LineNumber: l.lineNumber}
	case l.isTowCharTokenStart(l.curChar):
		tok := l.checkTwoCharToken(l.curChar)
		tok.LineNumber = l.lineNumber
		return tok
	case l.isOneCharToken(l.curChar):
		// 1文字トークン
		ch = l.curChar
		l.nextChar()
		return Token{TokenType: TokenType(ch), Literal: string(ch), LineNumber: l.lineNumber}
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
		if l.peekChar() == '.' {
			// LABEL abc.def
			l.nextChar()
			literal += l.readWord()
			l.nextChar()
			return Token{TokenType: DOT_IDENT, Literal: literal, LineNumber: l.lineNumber}
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
		// yas80 予約語
		tok, ok = reservedWords[strings.ToUpper(literal)]
		if ok {
			tok.LineNumber = l.lineNumber
			return tok
		}
		// これ以外は識別子
		return Token{TokenType: IDENT, Literal: literal, LineNumber: l.lineNumber}
	case l.curChar == '@' || l.curChar == '.':
		prefix := l.curChar
		literal = string(l.curChar)
		l.nextChar()
		literal += l.readWord()
		l.nextChar()
		if prefix == '@' {
			return Token{TokenType: AT_IDENT, Literal: literal, LineNumber: l.lineNumber}
		} else {
			return Token{TokenType: LOCAL_IDENT, Literal: literal, LineNumber: l.lineNumber}
		}

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
		tok = Token{TokenType: COMP, TokenSubType: LE, Literal: "<=", LineNumber: l.lineNumber}
	case ch1 == '<' && ch2 == '<':
		tok = Token{TokenType: SHIFT, TokenSubType: SL, Literal: "<<", LineNumber: l.lineNumber}
	case ch1 == '>' && ch2 == '=':
		tok = Token{TokenType: COMP, TokenSubType: GE, Literal: ">=", LineNumber: l.lineNumber}
	case ch1 == '>' && ch2 == '>':
		tok = Token{TokenType: SHIFT, TokenSubType: SR, Literal: ">>", LineNumber: l.lineNumber}
	case ch1 == '<' || ch1 == '>':
		tok = Token{TokenType: COMP, TokenSubType: TokenSubType(ch1), Literal: string(ch1), LineNumber: l.lineNumber}
	case ch1 == '=' && ch2 == '=':
		tok = Token{TokenType: COMP, TokenSubType: EQ, Literal: "==", LineNumber: l.lineNumber}
	case ch1 == '!' && ch2 == '=':
		tok = Token{TokenType: COMP, TokenSubType: NEQ, Literal: "!=", LineNumber: l.lineNumber}
	case ch1 == '!':
		return Token{TokenType: UNARY, TokenSubType: TokenSubType(ch1), Literal: string(ch1), LineNumber: l.lineNumber}
	case ch1 == '&' && ch2 == '&':
		tok = Token{TokenType: AND, TokenSubType: 0, Literal: "&&", LineNumber: l.lineNumber}
	case ch1 == '&':
		return Token{TokenType: MULDIV, TokenSubType: TokenSubType(ch1), Literal: string(ch1), LineNumber: l.lineNumber}
	case ch1 == '|' && ch2 == '|':
		tok = Token{TokenType: OR, TokenSubType: 0, Literal: "||", LineNumber: l.lineNumber}
	case ch1 == '|':
		return Token{TokenType: ADDSUB, TokenSubType: TokenSubType(ch1), Literal: string(ch1), LineNumber: l.lineNumber}
	default:
		// 1文字トークンを返す
		return Token{TokenType: TokenType(ch1), Literal: string(rune(ch1)), LineNumber: l.lineNumber}
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

func (l *Lexer) readString() string {
	startIndex := l.index
	for l.index < len(l.text) && l.text[l.index] != '"' {
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
	return ch == '(' || ch == ')' || ch == ',' || ch == ':'
}
