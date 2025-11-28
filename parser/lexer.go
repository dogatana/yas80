package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"yas80/fileblock"
	"yas80/logger"
)

const EOF = 0

type FileBlockProvider func() *fileblock.FileBlock

type LexerContext struct {
	filename   string
	lineNumber int
	index      int
	fileBlock  *fileblock.FileBlock
	curChar    rune
}

// 最低限必要な構造体を定義
type Lexer struct {
	scanner *bufio.Scanner
	isEOF   bool
	text    []byte

	logger   *logger.Logger
	program  *Program
	callback FileBlockProvider
	ctx      *LexerContext
}

func NewLexer(filename string, fb *fileblock.FileBlock, logger *logger.Logger) *Lexer {
	sb := bytes.NewReader(fb.Content)
	ctx := &LexerContext{filename: fb.Filename, fileBlock: fb}
	l := &Lexer{scanner: bufio.NewScanner(sb), program: &Program{}, ctx: ctx, logger: logger}
	l.nextChar()
	return l
}

// yyLexer インターフェースメソッド
func (l *Lexer) Lex(lval *yySymType) int {
	// if l.ctx == nil {
	// 	l.ctx = &LexerContext{
	// 		fileBlock:  l.callback(),
	// 		index:      0,
	// 		curChar:    0,
	// 		lineNumber: 0,
	// 	}
	// }
	// if l.ctx.fileBlock == nil {
	// 	// これ以上 FileBlock が得られない場合 EOF
	// 	tok := Token{TokenType: 0, Literal: "[EOF]", LineNumber: l.ctx.lineNumber}
	// 	lval.token = tok
	// 	return int(tok.TokenSubType)
	// }
	tok := l.NextToken()
	lval.token = tok
	return int(tok.TokenType)
}

// yyLexer インターフェースメソッド
func (l *Lexer) Error(msg string, args ...any) {
	var line int
	var yyVAL, yylval *Token

	switch len(args) {
	case 0: // Error(msg)
		line = l.ctx.lineNumber
	case 1: // Error(msg, lineNumber)
		n, ok := args[0].(int)
		if !ok {
			panic(fmt.Sprintf("[SYSTEM] invalid argument for Lexer.Error(string, %T(#%v))", args[0], args[0]))
		}
		line = n
	case 2:
		line = l.ctx.lineNumber
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

func NewLexerProvider(callback FileBlockProvider, logger *logger.Logger) *Lexer {
	return &Lexer{logger: logger}
}

func (l *Lexer) Logger() *logger.Logger { return l.logger }

func (l *Lexer) NextToken() Token {
	var literal string
	var ch rune

LINE_CONT:
	// 空白をスキップ
	l.skipWhitespace()
	// fmt.Printf("curChar %q, peekChar %q\n", string(l.ctx.curChar), string(l.peekChar()))

	// var tokType int
	switch {
	case l.ctx.curChar == EOF:
		// EOF
		return Token{TokenType: 0, Literal: "[EOF]", LineNumber: l.ctx.lineNumber}
	case l.ctx.curChar == ';':
		// コメント
		for l.ctx.curChar != '\n' && l.ctx.curChar != EOF {
			l.nextChar()
		}
		l.nextChar()
		return Token{TokenType: EOL, Literal: "\\n", LineNumber: l.ctx.lineNumber}
	case l.ctx.curChar == '\\' && l.peekChar() == '\n':
		// 行継続
		l.nextChar()
		l.nextChar()
		goto LINE_CONT
	case l.ctx.curChar == '\\':
		// マルチステートメント
		tok := Token{TokenType: EOL, Literal: "\\", LineNumber: l.ctx.lineNumber}
		l.nextChar()
		return tok
	case l.ctx.curChar == '\n':
		// EOL
		l.nextChar()
		return Token{TokenType: EOL, Literal: "\\n", LineNumber: l.ctx.lineNumber}
	case l.ctx.curChar == '"':
		// 文字列リテラル
		s := l.readString()
		l.nextChar()
		l.nextChar()
		return Token{TokenType: STRING, Literal: s, LineNumber: l.ctx.lineNumber}
	case l.ctx.curChar == '+' || l.ctx.curChar == '-' || l.ctx.curChar == '^': // ADDSUB
		ch := l.ctx.curChar
		l.nextChar()
		return Token{TokenType: ADDSUB, TokenSubType: TokenSubType(ch), Literal: string(ch), LineNumber: l.ctx.lineNumber}
	case l.ctx.curChar == '*' || l.ctx.curChar == '/': // MULDIV
		ch := l.ctx.curChar
		l.nextChar()
		return Token{TokenType: MULDIV, TokenSubType: TokenSubType(ch), Literal: string(ch), LineNumber: l.ctx.lineNumber}
	case l.isTowCharTokenStart(l.ctx.curChar):
		tok := l.checkTwoCharToken(l.ctx.curChar)
		tok.LineNumber = l.ctx.lineNumber
		return tok
	case l.isOneCharToken(l.ctx.curChar):
		// 1文字トークン
		ch = l.ctx.curChar
		l.nextChar()
		return Token{TokenType: TokenType(ch), Literal: string(ch), LineNumber: l.ctx.lineNumber}
	case l.ctx.curChar == '~':
		ch = l.ctx.curChar
		l.nextChar()
		return Token{TokenType: UNARY, TokenSubType: TokenSubType(ch), Literal: string(ch), LineNumber: l.ctx.lineNumber}
	case l.ctx.curChar == '0' && (l.peekChar() == 'x' || l.peekChar() == 'X'):
		// 16進数リテラル(0x)
		l.nextChar() // '0'をスキップ
		literal = "0" + string(l.ctx.curChar)
		l.nextChar() // 'x'または'X'をスキップ
		literal += l.readWord()
		l.nextChar()
		return Token{TokenType: NUMBER, Literal: literal, LineNumber: l.ctx.lineNumber}
	case l.ctx.curChar == '$' && !l.isXDigit(l.peekChar()):
		// $ ローケーションカウンタ
		tok := Token{TokenType: IDENT, Literal: "$", LineNumber: l.ctx.lineNumber}
		l.nextChar()
		return tok
	case l.ctx.curChar == '$' || l.ctx.curChar == '%':
		// 16進数リテラル($) or 2進数リテラル
		literal = string(l.ctx.curChar)
		l.nextChar() // '$'をスキップ
		literal += l.readWord()
		l.nextChar()
		return Token{TokenType: NUMBER, Literal: literal, LineNumber: l.ctx.lineNumber}
	case l.isDigit(l.ctx.curChar):
		// 10進、16進(0x)、2進(0b)、8進（0o)リテラル
		literal = l.readWord()
		l.nextChar()
		return Token{TokenType: NUMBER, Literal: literal, LineNumber: l.ctx.lineNumber}
	case l.isAlpha(l.ctx.curChar):
		// IDENT、予約語
		literal = l.readWord()
		if l.peekChar() == '.' {
			// LABEL abc.def
			l.nextChar()
			literal += l.readWord()
			l.nextChar()
			return Token{TokenType: DOT_IDENT, Literal: literal, LineNumber: l.ctx.lineNumber}
		}
		l.nextChar()
		// AF’の対処
		if len(literal) == 2 && l.ctx.curChar == '\'' {
			literal += "'"
			l.nextChar()
		}
		// z80 予約語
		tok, ok := z80ReservedWords[strings.ToUpper(literal)]
		if ok {
			tok.LineNumber = l.ctx.lineNumber
			return tok
		}
		// yas80 予約語
		tok, ok = reservedWords[strings.ToUpper(literal)]
		if ok {
			tok.LineNumber = l.ctx.lineNumber
			return tok
		}
		// これ以外は識別子
		return Token{TokenType: IDENT, Literal: literal, LineNumber: l.ctx.lineNumber}
	case l.ctx.curChar == '@' || l.ctx.curChar == '.':
		prefix := l.ctx.curChar
		literal = string(l.ctx.curChar)
		l.nextChar()
		literal += l.readWord()
		l.nextChar()
		if prefix == '@' {
			return Token{TokenType: AT_IDENT, Literal: literal, LineNumber: l.ctx.lineNumber}
		} else {
			return Token{TokenType: LOCAL_IDENT, Literal: literal, LineNumber: l.ctx.lineNumber}
		}

	default:
		literal = string(l.ctx.curChar)
		l.nextChar()
		return Token{TokenType: INVALID, Literal: literal, LineNumber: l.ctx.lineNumber}
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
		tok = Token{TokenType: COMP, TokenSubType: LE, Literal: "<=", LineNumber: l.ctx.lineNumber}
	case ch1 == '<' && ch2 == '<':
		tok = Token{TokenType: SHIFT, TokenSubType: SL, Literal: "<<", LineNumber: l.ctx.lineNumber}
	case ch1 == '>' && ch2 == '=':
		tok = Token{TokenType: COMP, TokenSubType: GE, Literal: ">=", LineNumber: l.ctx.lineNumber}
	case ch1 == '>' && ch2 == '>':
		tok = Token{TokenType: SHIFT, TokenSubType: SR, Literal: ">>", LineNumber: l.ctx.lineNumber}
	case ch1 == '<' || ch1 == '>':
		tok = Token{TokenType: COMP, TokenSubType: TokenSubType(ch1), Literal: string(ch1), LineNumber: l.ctx.lineNumber}
	case ch1 == '=' && ch2 == '=':
		tok = Token{TokenType: COMP, TokenSubType: EQ, Literal: "==", LineNumber: l.ctx.lineNumber}
	case ch1 == '!' && ch2 == '=':
		tok = Token{TokenType: COMP, TokenSubType: NEQ, Literal: "!=", LineNumber: l.ctx.lineNumber}
	case ch1 == '!':
		return Token{TokenType: UNARY, TokenSubType: TokenSubType(ch1), Literal: string(ch1), LineNumber: l.ctx.lineNumber}
	case ch1 == '&' && ch2 == '&':
		tok = Token{TokenType: AND, TokenSubType: 0, Literal: "&&", LineNumber: l.ctx.lineNumber}
	case ch1 == '&':
		return Token{TokenType: MULDIV, TokenSubType: TokenSubType(ch1), Literal: string(ch1), LineNumber: l.ctx.lineNumber}
	case ch1 == '|' && ch2 == '|':
		tok = Token{TokenType: OR, TokenSubType: 0, Literal: "||", LineNumber: l.ctx.lineNumber}
	case ch1 == '|':
		return Token{TokenType: ADDSUB, TokenSubType: TokenSubType(ch1), Literal: string(ch1), LineNumber: l.ctx.lineNumber}
	default:
		// 1文字トークンを返す
		return Token{TokenType: TokenType(ch1), Literal: string(rune(ch1)), LineNumber: l.ctx.lineNumber}
	}
	l.nextChar()
	return tok
}

func (l *Lexer) nextChar() {
	if l.isEOF {
		l.ctx.curChar = EOF
		return
	}
	if l.ctx.index == 0 {
		scanned := l.scanner.Scan()
		if !scanned {
			if l.ctx.curChar != '\n' {
				l.ctx.curChar = '\n'
				return
			}
			l.ctx.curChar = EOF
			l.isEOF = true
			return
		}
		l.ctx.lineNumber++
		l.text = []byte(l.scanner.Text())
	}
	if l.ctx.index >= len(l.text) {
		l.ctx.curChar = '\n'
		l.ctx.index = 0
		return
	}
	start := l.ctx.index
	l.ctx.index += l.charSize(l.text[l.ctx.index])
	l.ctx.curChar = []rune(string(l.text[start:l.ctx.index]))[0]
}

func (l *Lexer) peekChar() rune {
	if l.ctx.index >= len(l.text) {
		return '\n'
	}
	return rune(l.text[l.ctx.index])
}

func (l *Lexer) skipWhitespace() {
	for !l.isEOF && (l.ctx.curChar == ' ' || l.ctx.curChar == '\t' || l.ctx.curChar == '\r') {
		l.nextChar()
	}
}

func (l *Lexer) readString() string {
	startIndex := l.ctx.index
	for l.ctx.index < len(l.text) && l.text[l.ctx.index] != '"' {
		l.ctx.index++
	}
	return string(l.text[startIndex:l.ctx.index])
}

func (l *Lexer) readWord() string {
	startIndex := l.ctx.index - 1
	for l.ctx.index < len(l.text) && l.isWordChar(rune(l.text[l.ctx.index])) {
		l.ctx.index++
	}
	return string(l.text[startIndex:l.ctx.index])
}

func (l *Lexer) isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) isXDigit(ch rune) bool {
	return l.isDigit(ch) || 'a' <= ch && ch <= 'f' || 'A' <= ch && ch <= 'F'
}

func (l *Lexer) isAlpha(ch rune) bool {
	return 'A' <= ch && ch <= 'Z' || 'a' <= ch && ch <= 'z' || ch == '_'
}

func (l *Lexer) isWordChar(ch rune) bool {
	return l.isDigit(ch) || l.isAlpha(ch)
}

func (l *Lexer) isOneCharToken(ch rune) bool {
	return ch == '(' || ch == ')' || ch == ',' || ch == ':' || ch == '[' || ch == ']'
}

func (l *Lexer) charSize(ch byte) int {
	switch {
	case ch < 0x80:
		return 1
	case ch < 0xc2: // invalid utf-8
		return 1
	case ch < 0xe0:
		return 2
	case ch < 0xf0:
		return 3
	case ch < 0xf5:
		return 4
	default:
		return 1 // invalid utf-8
	}
}
