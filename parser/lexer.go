package parser

import (
	"fmt"
	"strings"
	"yas80/errcode"
	"yas80/filecontent"
	"yas80/logging"
)

const EOF = 0

type filecontentProvider func() *filecontent.FileContent

type LexerContext struct {
	filename    string
	lineNumber  int
	index       int
	fileContent *filecontent.FileContent
	curChar     rune
}

func (ctx *LexerContext) toContext(start int) *filecontent.Context {
	return &filecontent.Context{FileContent: ctx.fileContent, Line: ctx.lineNumber, Index: start}
}

// 最低限必要な構造体を定義
type Lexer struct {
	callback func() *filecontent.FileContent
	isEOF    bool
	start    int // token 開始 index
	logger   *logging.Logger
	program  *BlockStatement
	lctx     *LexerContext
	lexState int
	//	callback filecontentProvider
}

func NewLexerOrg(fc *filecontent.FileContent, logger *logging.Logger) *Lexer {
	lctx := &LexerContext{filename: fc.Filename, lineNumber: 1, fileContent: fc}
	l := &Lexer{logger: logger, program: &BlockStatement{}, lctx: lctx}
	l.nextChar()
	return l
}

func NewLexer(logger *logging.Logger, callback func() *filecontent.FileContent) *Lexer {
	l := &Lexer{logger: logger, callback: callback, program: &BlockStatement{}}
	return l
}

// yyLexer インターフェースメソッド
func (l *Lexer) Lex(lval *yySymType) int {
	var fc *filecontent.FileContent

	for {
		switch l.lexState {
		case 0:
			// l.ctx == nil
			fc = l.callback()
			if fc == nil {
				// これ以上 filecontent が得られない場合 EOF を返すステートへ移行
				l.lexState = 9
				break
			}
			l.lexState++

		case 1:
			// setup LexerContext
			lctx := &LexerContext{filename: fc.Filename, lineNumber: 1, fileContent: fc}
			l.isEOF = false
			l.lctx = lctx
			l.nextChar()

			// return FILE
			tok := Token{TokenType: FILE, Literal: fc.Filename}
			lval.token = tok

			l.lexState++
			return int(tok.TokenType)

		case 2:
			tok := l.NextToken()
			if tok.TokenType == EOF {
				l.lexState = 0
				break
			}
			lval.token = tok
			return int(tok.TokenType)
		case 9:
			tok := Token{TokenType: 0, Literal: "[EOF]"}
			lval.token = tok
			return int(tok.TokenType)
		}
	}
}

// yyLexer インターフェースメソッド
func (l *Lexer) Error(msg string, ctx *filecontent.Context) {
	if strings.HasPrefix(msg, "[I]") {
		l.logger.Info(msg[3:], ctx)
	} else if strings.HasPrefix(msg, "[W]") {
		l.logger.Warning(msg[3:], ctx)
	} else if strings.HasPrefix(msg, "[E]") {
		l.logger.Error(msg[3:], ctx)
	} else if msg[0] == '[' {
		l.logger.Error(msg, ctx)
	} else {
		l.logger.Error(msg, ctx)
	}
}

func NewLexerProvider(callback filecontentProvider, logger *logging.Logger) *Lexer {
	return &Lexer{logger: logger}
}

func (l *Lexer) Logger() *logging.Logger { return l.logger }

func (l *Lexer) NextToken() Token {
	var literal string
	var ch rune

LINE_CONT:
	// 空白をスキップ
	l.skipWhitespace()
	l.start = l.lctx.index - 1

	// var tokType int
	switch {
	case l.lctx.curChar == EOF:
		// EOF
		return Token{TokenType: 0, Literal: "[EOF]", Context: l.lctx.toContext(l.start)}

	case l.lctx.curChar == ';':
		// コメント
		for l.lctx.curChar != '\n' && l.lctx.curChar != EOF {
			l.nextChar()
		}
		l.nextChar()
		return Token{TokenType: EOL, Literal: "\\n", Context: l.lctx.toContext(l.start)}

	case l.lctx.curChar == '\\' && l.peekChar() == '\n':
		// 行継続
		l.nextChar()
		l.nextChar()
		goto LINE_CONT

	case l.lctx.curChar == '\\':
		// マルチステートメント
		tok := Token{TokenType: EOL, Literal: "\\", Context: l.lctx.toContext(l.start)}
		l.nextChar()
		return tok

	case l.lctx.curChar == '\n':
		// EOL
		l.nextChar()
		return Token{TokenType: EOL, Literal: "\\n", Context: l.lctx.toContext(l.start)}

	case l.lctx.curChar == '"' || l.lctx.curChar == '\'':
		// 文字列リテラル
		s := l.readString(l.lctx.curChar)
		if l.lctx.curChar != '\n' {
			// ESTR_END_QUOTE の場合、nextChar すると EOL トークンを出力しない
			l.nextChar()
		}
		return Token{TokenType: STRING, Literal: s, Context: l.lctx.toContext(l.start)}

	case l.lctx.curChar == '+' || l.lctx.curChar == '-' || l.lctx.curChar == '^':
		// ADDSUB
		ch := l.lctx.curChar
		l.nextChar()
		return Token{TokenType: ADDSUB, TokenSubType: TokenSubType(ch), Literal: string(ch), Context: l.lctx.toContext(l.start)}

	case l.lctx.curChar == '*' || l.lctx.curChar == '/':
		// MULDIV（%は数値リテラルとの識別のため別処理）
		ch := l.lctx.curChar
		l.nextChar()
		return Token{TokenType: MULDIV, TokenSubType: TokenSubType(ch), Literal: string(ch), Context: l.lctx.toContext(l.start)}

	case twoCharTokenChars[l.lctx.curChar]:
		// 2 文字トークン
		tok := l.checkTwoCharToken(l.lctx.curChar)
		tok.Context = l.lctx.toContext(l.start)
		return tok

	case oneCharTokenChars[l.lctx.curChar]:
		// 1文字トークン
		ch = l.lctx.curChar
		l.nextChar()
		return Token{TokenType: TokenType(ch), Literal: string(ch), Context: l.lctx.toContext(l.start)}

	case l.lctx.curChar == '~':
		// 単項演算子（ここでは ~ のみ。'!', '-' はパーサ側で判定
		ch = l.lctx.curChar
		l.nextChar()
		return Token{TokenType: UNARY, TokenSubType: TokenSubType(ch), Literal: string(ch), Context: l.lctx.toContext(l.start)}

	// case l.lctx.curChar == '$' && l.isXDigit(l.peekChar()):
	// 	// 16進数リテラル($)
	// 	l.nextChar()
	// 	literal = "$" + l.readHexString()
	// 	l.nextChar()
	// 	return Token{TokenType: NUMBER, Literal: literal, Context: l.lctx.toContext(l.start)}

	case l.lctx.curChar == '$' && l.isWordChar(l.peekChar()):
		// システム識別子($)
		l.nextChar()
		literal = "$" + l.readWord()
		l.nextChar()
		if l.isHexString(string(literal[1:])) && literal[1] != '_' {
			return Token{TokenType: NUMBER, Literal: literal, Context: l.lctx.toContext(l.start)}
		} else {
			return Token{TokenType: IDENT, Literal: literal, Context: l.lctx.toContext(l.start)}
		}

	case l.lctx.curChar == '$':
		// $ ローケーションカウンタ
		tok := Token{TokenType: IDENT, Literal: "$", Context: l.lctx.toContext(l.start)}
		l.nextChar()
		return tok

	case l.lctx.curChar == '%' && l.isBinDigit(l.peekChar()):
		// 2進数リテラル(%)
		l.nextChar()
		literal = "%" + l.readBinString()
		l.nextChar()
		return Token{TokenType: NUMBER, Literal: literal, Context: l.lctx.toContext(l.start)}

	case l.lctx.curChar == '%':
		// % 演算子
		ch := l.lctx.curChar
		l.nextChar()
		return Token{TokenType: MULDIV, TokenSubType: TokenSubType(ch), Literal: string(ch), Context: l.lctx.toContext(l.start)}

	case l.lctx.curChar == '0' && (l.peekChar() == 'x' || l.peekChar() == 'X'):
		// 16進数リテラル(0x)
		l.nextChar() // '0'をスキップ
		literal = "0" + string(l.lctx.curChar)
		l.nextChar() // 'x'または'X'をスキップ
		if !l.isHexChar(l.lctx.curChar) {
			l.logger.Error(fmt.Sprintf(errcode.ENUMBER, literal), l.lctx.toContext(l.start))
			return Token{TokenType: NUMBER, Literal: "0", Context: l.lctx.toContext(l.start)}
		}
		literal += l.readHexString()
		l.nextChar()
		return Token{TokenType: NUMBER, Literal: literal, Context: l.lctx.toContext(l.start)}

	case l.lctx.curChar == '0' && (l.peekChar() == 'b' || l.peekChar() == 'B'):
		// 2数リテラル(0b)
		l.nextChar() // '0'をスキップ
		literal = "0" + string(l.lctx.curChar)
		l.nextChar() // 'b'または'B'をスキップ
		if !l.isBinChar(l.lctx.curChar) {
			l.logger.Error(fmt.Sprintf(errcode.ENUMBER, literal), l.lctx.toContext(l.start))
			return Token{TokenType: NUMBER, Literal: "0", Context: l.lctx.toContext(l.start)}
		}
		literal += l.readBinString()
		l.nextChar()
		return Token{TokenType: NUMBER, Literal: literal, Context: l.lctx.toContext(l.start)}

	case l.lctx.curChar == '0' && (l.peekChar() == 'o' || l.peekChar() == 'O'):
		// 8数リテラル(0o)
		l.nextChar() // '0'をスキップ
		literal = "0" + string(l.lctx.curChar)
		l.nextChar() // 'o'または'O'をスキップ
		if !l.isOctChar(l.lctx.curChar) {
			l.logger.Error(fmt.Sprintf(errcode.ENUMBER, literal), l.lctx.toContext(l.start))
			return Token{TokenType: NUMBER, Literal: "0", Context: l.lctx.toContext(l.start)}
		}
		literal += l.readOctString()
		l.nextChar()
		return Token{TokenType: NUMBER, Literal: literal, Context: l.lctx.toContext(l.start)}

	case l.isDigit(l.lctx.curChar):
		// 10進リテラル, 16進リテラル（末尾 h）
		literal = l.readHexString()
		l.nextChar()
		ch := l.lctx.curChar
		if l.isHexString(literal) && (ch == 'h' || ch == 'H') {
			literal += string(ch)
			l.nextChar()
			return Token{TokenType: NUMBER, Literal: literal, Context: l.lctx.toContext(l.start)}
		}
		return Token{TokenType: NUMBER, Literal: literal, Context: l.lctx.toContext(l.start)}

	case l.isAlpha(l.lctx.curChar):
		// IDENT, DOT_IDENT, 予約語
		literal = l.readWord()
		if l.peekChar() == '.' {
			// LABEL abc.def
			l.nextChar()
			literal += l.readWord()
			l.nextChar()
			return Token{TokenType: DOT_IDENT, Literal: literal, Context: l.lctx.toContext(l.start)}
		}
		l.nextChar()
		// AF’の対処
		if strings.ToUpper(literal) == "AF" && l.lctx.curChar == '\'' {
			literal += "'"
			l.nextChar()
		}
		// z80 予約語
		tok, ok := z80ReservedWords[strings.ToUpper(literal)]
		if ok {
			tok.Context = l.lctx.toContext(l.start)
			return tok
		}
		// yas80 予約語
		tok, ok = reservedWords[strings.ToUpper(literal)]
		if ok {
			tok.Context = l.lctx.toContext(l.start)
			return tok
		}
		// これ以外は IDENT
		return Token{TokenType: IDENT, Literal: literal, Context: l.lctx.toContext(l.start)}

	case l.lctx.curChar == '@' || l.lctx.curChar == '.':
		// AT_IDENT, LOCAL_IDENT
		prefix := l.lctx.curChar
		literal = string(l.lctx.curChar)
		l.nextChar()
		literal += l.readWord()
		l.nextChar()
		if prefix == '@' {
			return Token{TokenType: AT_IDENT, Literal: literal, Context: l.lctx.toContext(l.start)}
		} else {
			return Token{TokenType: LOCAL_IDENT, Literal: literal, Context: l.lctx.toContext(l.start)}
		}

	default:
		literal = string(l.lctx.curChar)
		l.nextChar()
		return Token{TokenType: INVALID, Literal: literal, Context: l.lctx.toContext(l.start)}
	}
}

// 1 文字トークン
var oneCharTokenChars = map[rune]bool{
	'(': true,
	')': true,
	',': true,
	':': true,
	'[': true,
	']': true,
}

// 2文字トークンの最初の文字
var twoCharTokenChars = map[rune]bool{
	'=': true,
	'<': true,
	'>': true,
	'!': true,
	'|': true,
	'&': true,
	'#': true,
}

// 2文字トークンのチェック
func (l *Lexer) checkTwoCharToken(ch1 rune) Token {
	ch2 := l.peekChar()
	l.nextChar()

	var tok Token
	switch {
	// COMP
	case ch1 == '<' && ch2 == '=':
		tok = Token{TokenType: COMP, TokenSubType: LE, Literal: "<=", Context: l.lctx.toContext(l.start)}
	case ch1 == '<' && ch2 == '<':
		tok = Token{TokenType: SHIFT, TokenSubType: SL, Literal: "<<", Context: l.lctx.toContext(l.start)}
	case ch1 == '>' && ch2 == '=':
		tok = Token{TokenType: COMP, TokenSubType: GE, Literal: ">=", Context: l.lctx.toContext(l.start)}
	case ch1 == '>' && ch2 == '>':
		tok = Token{TokenType: SHIFT, TokenSubType: SR, Literal: ">>", Context: l.lctx.toContext(l.start)}
	case ch1 == '<' || ch1 == '>':
		tok = Token{TokenType: COMP, TokenSubType: TokenSubType(ch1), Literal: string(ch1), Context: l.lctx.toContext(l.start)}
	case ch1 == '=' && ch2 == '=':
		tok = Token{TokenType: COMP, TokenSubType: EQ, Literal: "==", Context: l.lctx.toContext(l.start)}
	case ch1 == '!' && ch2 == '=':
		tok = Token{TokenType: COMP, TokenSubType: NEQ, Literal: "!=", Context: l.lctx.toContext(l.start)}
	case ch1 == '!':
		return Token{TokenType: UNARY, TokenSubType: TokenSubType(ch1), Literal: string(ch1), Context: l.lctx.toContext(l.start)}
	case ch1 == '&' && ch2 == '&':
		tok = Token{TokenType: AND, TokenSubType: 0, Literal: "&&", Context: l.lctx.toContext(l.start)}
	case ch1 == '&':
		return Token{TokenType: MULDIV, TokenSubType: TokenSubType(ch1), Literal: string(ch1), Context: l.lctx.toContext(l.start)}
	case ch1 == '|' && ch2 == '|':
		tok = Token{TokenType: OR, TokenSubType: 0, Literal: "||", Context: l.lctx.toContext(l.start)}
	case ch1 == '|':
		return Token{TokenType: ADDSUB, TokenSubType: TokenSubType(ch1), Literal: string(ch1), Context: l.lctx.toContext(l.start)}
	case ch1 == '#' && ch2 == '#':
		tok = Token{TokenType: CONCAT, TokenSubType: 0, Literal: "##", Context: l.lctx.toContext(l.start)}
	default:
		// 1文字トークンを返す
		return Token{TokenType: TokenType(ch1), Literal: string(rune(ch1)), Context: l.lctx.toContext(l.start)}
	}
	l.nextChar()
	return tok
}

func (l *Lexer) nextChar() {
	if l.isEOF {
		l.lctx.curChar = EOF
		return
	}
	if l.lctx.index >= len(l.lctx.fileContent.Content) {
		if l.lctx.curChar != '\n' {
			l.lctx.curChar = '\n'
			return
		}
		l.lctx.curChar = EOF
		l.isEOF = true
		return
	}
	start := l.lctx.index
	l.lctx.index += l.charSize(l.lctx.fileContent.Content[l.lctx.index])
	l.lctx.curChar = []rune(string(l.lctx.fileContent.Content[start:l.lctx.index]))[0]
	if l.lctx.curChar == '\n' {
		l.lctx.lineNumber++
	}
}

func (l *Lexer) peekChar() rune {
	if l.lctx.index >= len(l.lctx.fileContent.Content) {
		return '\n'
	}
	end := l.lctx.index + l.charSize(l.lctx.fileContent.Content[l.lctx.index])
	return []rune(string(l.lctx.fileContent.Content[l.lctx.index:end]))[0]
}

func (l *Lexer) skipWhitespace() {
	for !l.isEOF && (l.lctx.curChar == ' ' || l.lctx.curChar == '\t' || l.lctx.curChar == '\r') {
		l.nextChar()
	}
}

func (l *Lexer) readString(quote rune) string {
	escapeChars := map[rune]rune{
		'\'': '\'',
		'"':  '"',
		'\\': '\\',
		'0':  '\x00',
		'a':  '\a',
		'b':  '\b',
		'f':  '\f',
		'n':  '\n',
		'r':  '\r',
		't':  '\t',
		'v':  '\v',
	}
	runes := []rune{}

	index := l.lctx.index
	var ch rune
	for {
		l.nextChar()
		ch = l.lctx.curChar
		if ch == quote {
			break
		}
		if ch == '\n' {
			l.logger.Error(errcode.ESTR_END_QUOTE, l.lctx.toContext(index))
			break
		}
		if ch < ' ' {
			l.logger.Error(errcode.ESTR_CTRL, l.lctx.toContext(index))
			break
		}
		if ch == '\\' {
			ec, ok := escapeChars[l.peekChar()]
			if ok {
				ch = ec
				l.nextChar()
			}
		}
		runes = append(runes, ch)
	}
	return string(runes)
}

func (l *Lexer) readWord() string {
	startIndex := l.lctx.index - 1
	for l.lctx.index < len(l.lctx.fileContent.Content) && l.isWordChar(rune(l.lctx.fileContent.Content[l.lctx.index])) {
		l.lctx.index++
	}
	return string(l.lctx.fileContent.Content[startIndex:l.lctx.index])
}

func (l *Lexer) readHexString() string {
	if !l.isHexChar(l.lctx.curChar) {
		return ""
	}
	startIndex := l.lctx.index - 1
	for l.lctx.index < len(l.lctx.fileContent.Content) && l.isHexChar(rune(l.lctx.fileContent.Content[l.lctx.index])) {
		l.lctx.index++
	}
	return string(l.lctx.fileContent.Content[startIndex:l.lctx.index])
}

func (l *Lexer) readBinString() string {
	if !l.isBinChar(l.lctx.curChar) {
		return ""
	}
	startIndex := l.lctx.index - 1
	for l.lctx.index < len(l.lctx.fileContent.Content) && l.isBinChar(rune(l.lctx.fileContent.Content[l.lctx.index])) {
		l.lctx.index++
	}
	return string(l.lctx.fileContent.Content[startIndex:l.lctx.index])
}

func (l *Lexer) readOctString() string {
	if !l.isOctChar(l.lctx.curChar) {
		return ""
	}
	startIndex := l.lctx.index - 1
	for l.lctx.index < len(l.lctx.fileContent.Content) && l.isOctChar(rune(l.lctx.fileContent.Content[l.lctx.index])) {
		l.lctx.index++
	}
	return string(l.lctx.fileContent.Content[startIndex:l.lctx.index])
}

func (l *Lexer) isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) isXDigit(ch rune) bool {
	return l.isDigit(ch) || 'a' <= ch && ch <= 'f' || 'A' <= ch && ch <= 'F'
}

func (l *Lexer) isHexChar(ch rune) bool {
	return l.isXDigit(ch) || ch == '_'
}

func (l *Lexer) isHexString(s string) bool {
	for _, c := range s {
		if !l.isHexChar(rune(c)) {
			return false
		}
	}
	return true
}

func (l *Lexer) isBinDigit(ch rune) bool {
	return ch == '0' || ch == '1'
}

func (l *Lexer) isBinChar(ch rune) bool {
	return ch == '0' || ch == '1' || ch == '_'
}

func (l *Lexer) isOctChar(ch rune) bool {
	return '0' <= ch && ch <= '7' || ch == '_'
}

func (l *Lexer) isAlpha(ch rune) bool {
	return 'A' <= ch && ch <= 'Z' || 'a' <= ch && ch <= 'z' || ch == '_'
}

func (l *Lexer) isWordChar(ch rune) bool {
	return l.isDigit(ch) || l.isAlpha(ch)
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
