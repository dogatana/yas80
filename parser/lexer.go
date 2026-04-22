package parser

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/dogatana/yas80/errcode"
	"github.com/dogatana/yas80/filecontent"
	"github.com/dogatana/yas80/intern"
	"github.com/dogatana/yas80/internal/util"
	"github.com/dogatana/yas80/logging"
)

const EOF = 0

type LexerContext struct {
	filename    string
	lineNumber  int
	index       int
	fileContent *filecontent.FileContent
	curChar     rune
}

func (ctx *LexerContext) toContext(start int) filecontent.Context {
	return filecontent.Context{FileContent: ctx.fileContent, Line: uint32(ctx.lineNumber), Index: uint32(start)}
}

// 最低限必要な構造体を定義
type Lexer struct {
	callback func() *filecontent.FileContent
	isEOF    bool
	start    int // token 開始 index
	logger   *logging.Logger
	program  *BlockStatement
	lctx     *LexerContext

	lexState    int
	fileContent *filecontent.FileContent

	stack     util.Stack[*LexerContext]           // include ファイル処理用
	including map[string]bool                     // 循環 include 検査用
	FcMap     map[string]*filecontent.FileContent // lister 用
}

func NewLexer(logger *logging.Logger, callback func() *filecontent.FileContent) *Lexer {
	l := &Lexer{
		logger:   logger,
		callback: callback,
		program:  &BlockStatement{},

		stack:     util.Stack[*LexerContext]{},
		including: map[string]bool{},
		FcMap:     map[string]*filecontent.FileContent{},
	}
	return l
}

// include "name" をパースした際に parser から呼ばれる
func (l *Lexer) Push(filename string, fc *filecontent.FileContent, ctx *filecontent.Context) error {
	name := filepath.FromSlash(fc.Filename)
	abs, err := filepath.Abs(name)
	if err != nil {
		return fmt.Errorf(errcode.EFILE_ERR, fc.Filename, err.Error())
	}
	if l.including[abs] {
		return fmt.Errorf(errcode.EINCLUDE_CYCLIC, filename)
	}
	l.including[abs] = true
	l.stack.Push(l.lctx)
	l.fileContent = fc
	l.registerFileContent()
	l.lexState = 1
	return nil
}

// Lister 用に include を含むアセンブルファイルを map へ登録
func (l *Lexer) registerFileContent() {
	name := util.SlashPath(l.fileContent.Filename) // 登録キーは  / 区切りの path とする
	l.FcMap[name] = l.fileContent
}

// var fileContent *filecontent.FileContent
// var stack = &stackType{}

// yyLexer インターフェースメソッド
func (l *Lexer) Lex(lval *yySymType) int {

	for {
		switch l.lexState {
		case 0:
			// 次のファイルを取得
			l.fileContent = l.callback()
			if l.fileContent == nil {
				// これ以上 filecontent が得られない場合 EOF を返すステートへ移行
				l.lexState = 9
				break
			}
			l.registerFileContent()
			l.lexState++

		case 1:
			// setup LexerContext
			lctx := &LexerContext{filename: l.fileContent.Filename, lineNumber: 1, fileContent: l.fileContent}
			l.stack.Push(lctx)

			l.lctx = lctx
			l.isEOF = false

			l.lexState++

		case 2:
			// pop
			if lctx, ok := l.stack.Pop(); !ok {
				l.lexState = 0
			} else {
				l.isEOF = false
				l.lctx = lctx

				// return FILE
				tok := Token{TokenType: FILE, TokenSubType: TokenSubType(lctx.lineNumber), SymbolID: intern.InternString(lctx.filename)}
				lval.token = tok

				// lctx がファイル先頭なら 1文字進める
				if l.lctx.index == 0 {
					l.nextChar()
				}
				l.lexState++
				return int(tok.TokenType)
			}

		case 3:
			tok := l.NextToken()
			if tok.TokenType == EOF {
				abs, _ := filepath.Abs(filepath.FromSlash(l.lctx.filename))
				l.including[abs] = false
				l.lexState = 2
				break
			}
			lval.token = tok
			return int(tok.TokenType)

		case 9:
			tok := Token{TokenType: 0}
			lval.token = tok
			return int(tok.TokenType)
		}
	}
}

// yyLexer インターフェースメソッド
func (l *Lexer) Error(msg string, c *filecontent.Context) {
	ctx := *c
	if strings.HasPrefix(msg, "[I]") {
		l.logger.Info(msg[3:], &ctx)
	} else if strings.HasPrefix(msg, "[W]") {
		l.logger.Warning(msg[3:], &ctx)
	} else if strings.HasPrefix(msg, "[E]") {
		l.logger.Error(msg[3:], &ctx)
	} else if msg[0] == '[' {
		l.logger.Error(msg, &ctx)
	} else {
		l.logger.Error(msg, &ctx)
	}
}

func (l *Lexer) NextToken() Token {
	var literal string
	var ch rune
	var tok Token

LINE_CONT:
	// 空白をスキップ
	l.skipWhitespace()
	l.start = l.lctx.index - 1

	// var tokType int
	switch {
	case l.lctx.curChar == EOF:
		// EOF
		return Token{TokenType: 0, Context: l.lctx.toContext(l.start)}

	case l.lctx.curChar == ';':
		// コメント
		for l.lctx.curChar != '\n' && l.lctx.curChar != EOF {
			l.nextChar()
		}
		tok = Token{TokenType: EOL, Context: l.lctx.toContext(l.start)}
		l.nextChar()
		l.lctx.lineNumber++
		return tok

	case l.lctx.curChar == '\\' && l.peekChar() == '\n':
		// 行継続
		l.nextChar()
		l.nextChar()
		l.lctx.lineNumber++
		goto LINE_CONT

	case l.lctx.curChar == '\\':
		// マルチステートメント
		tok := Token{TokenType: EOL, Context: l.lctx.toContext(l.start)}
		l.nextChar()
		return tok

	case l.lctx.curChar == '\n':
		// EOL
		tok = Token{TokenType: EOL, Context: l.lctx.toContext(l.start)}
		l.nextChar()
		l.lctx.lineNumber++
		return tok

	case l.lctx.curChar == '"' || l.lctx.curChar == '\'' || l.lctx.curChar == '`':
		// 文字列リテラル, raw string リテラル
		s := l.readString(l.lctx.curChar)
		if l.lctx.curChar != '\n' {
			// ESTR_END_QUOTE の場合、nextChar すると EOL トークンを出力しない
			l.nextChar()
		}
		return Token{TokenType: STRING, SymbolID: intern.InternString(s), Context: l.lctx.toContext(l.start)}

	case l.lctx.curChar == '+' || l.lctx.curChar == '-' || l.lctx.curChar == '^':
		// ADDSUB
		ch := l.lctx.curChar
		tok = Token{TokenType: ADDSUB, TokenSubType: TokenSubType(ch), Context: l.lctx.toContext(l.start)}
		l.nextChar()
		return tok

	case l.lctx.curChar == '*' || l.lctx.curChar == '/':
		// MULDIV（%は数値リテラルとの識別のため別処理）
		ch := l.lctx.curChar
		tok = Token{TokenType: MULDIV, TokenSubType: TokenSubType(ch), Context: l.lctx.toContext(l.start)}
		l.nextChar()
		return tok

	case twoCharTokenChars[l.lctx.curChar]:
		// 2 文字トークン
		tok := l.checkTwoCharToken(l.lctx.curChar)
		tok.Context = l.lctx.toContext(l.start)
		return tok

	case oneCharTokenChars[l.lctx.curChar]:
		// 1文字トークン
		ch = l.lctx.curChar
		tok = Token{TokenType: TokenType(ch), Context: l.lctx.toContext(l.start)}
		l.nextChar()
		return tok

	case l.lctx.curChar == '~':
		// 単項演算子（ここでは ~ のみ。'!', '-' はパーサ側で判定
		ch = l.lctx.curChar
		tok = Token{TokenType: UNARY, TokenSubType: TokenSubType(ch), Context: l.lctx.toContext(l.start)}
		l.nextChar()
		return tok

	case l.lctx.curChar == '$' && l.isWordChar(l.peekChar()):
		// システム識別子($)
		// l.nextChar()
		// literal = "$" + l.readWord()
		b := l.readWord()
		if l.isHexString(string(b[1:])) && b[1] != '_' {
			tok = Token{TokenType: NUMBER, SymbolID: intern.InternString(string(b)), Context: l.lctx.toContext(l.start)}
		} else {
			tok = Token{TokenType: IDENT, SymbolID: intern.InternBytes(b), Context: l.lctx.toContext(l.start)}
		}
		l.nextChar()
		return tok

	case l.lctx.curChar == '$':
		if l.peekChar() == '$' { // $$
			tok = Token{TokenType: IDENT, SymbolID: intern.ID_ALOC, Context: l.lctx.toContext(l.start)}
			l.nextChar()
			l.nextChar()
			return tok
		}
		tok = Token{TokenType: IDENT, SymbolID: intern.ID_LOC, Context: l.lctx.toContext(l.start)} // $
		l.nextChar()
		return tok

	case l.lctx.curChar == '%' && l.isBinDigit(l.peekChar()):
		// 2進数リテラル(%)
		l.nextChar()
		literal = "%" + l.readBinString()
		tok = Token{TokenType: NUMBER, SymbolID: intern.InternString(literal), Context: l.lctx.toContext(l.start)}
		l.nextChar()
		return tok

	case l.lctx.curChar == '%':
		// % 演算子
		ch := l.lctx.curChar
		tok = Token{TokenType: MULDIV, TokenSubType: TokenSubType(ch), Context: l.lctx.toContext(l.start)}
		l.nextChar()
		return tok

	case l.lctx.curChar == '0' && (l.peekChar() == 'x' || l.peekChar() == 'X'):
		// 16進数リテラル(0x)
		l.nextChar() // '0'をスキップ
		literal = "0" + string(l.lctx.curChar)
		l.nextChar() // 'x'または'X'をスキップ
		if !l.isHexChar(l.lctx.curChar) {
			ctx := l.lctx.toContext(l.start)
			l.logger.Error(fmt.Sprintf(errcode.ENUMBER_LITERAL, literal), &ctx)
			return Token{TokenType: NUMBER, SymbolID: intern.ID_STR_ZERO, Context: l.lctx.toContext(l.start)}
		}
		literal += l.readHexString()
		tok = Token{TokenType: NUMBER, SymbolID: intern.InternString(literal), Context: l.lctx.toContext(l.start)}
		l.nextChar()
		return tok

	case l.lctx.curChar == '0' && (l.peekChar() == 'b' || l.peekChar() == 'B'):
		// 2数リテラル(0b)
		l.nextChar() // '0'をスキップ
		literal = "0" + string(l.lctx.curChar)
		l.nextChar() // 'b'または'B'をスキップ

		ch := l.lctx.curChar
		// 0bh
		if ch == 'h' || ch == 'H' {
			tok = Token{TokenType: NUMBER, SymbolID: intern.InternString("$" + literal), Context: l.lctx.toContext(l.start)}
			l.nextChar() // skip 'h' / 'H'
			return tok
		}
		// 0b[\x]
		if l.isHexChar(ch) {
			hex := l.readHexString()
			l.nextChar()
			ch = l.lctx.curChar
			if ch == 'h' || ch == 'H' {
				tok = Token{TokenType: NUMBER, SymbolID: intern.InternString("$" + literal + hex), Context: l.lctx.toContext(l.start)}
				l.nextChar()
				return tok
			}
			if l.isBinString(hex) {
				return Token{TokenType: NUMBER, SymbolID: intern.InternString(literal + hex), Context: l.lctx.toContext(l.start)}
			}
			ctx := l.lctx.toContext(l.start)
			l.logger.Error(fmt.Sprintf(errcode.ENUMBER_LITERAL, literal+hex), &ctx)
			return Token{TokenType: NUMBER, SymbolID: intern.ID_STR_ZERO, Context: l.lctx.toContext(l.start)}
		}
		ctx := l.lctx.toContext(l.start)
		l.logger.Error(fmt.Sprintf(errcode.ENUMBER_LITERAL, literal), &ctx)
		return Token{TokenType: NUMBER, SymbolID: intern.ID_STR_ZERO, Context: l.lctx.toContext(l.start)}

	case l.lctx.curChar == '0' && (l.peekChar() == 'o' || l.peekChar() == 'O'):
		// 8数リテラル(0o)
		l.nextChar() // '0'をスキップ
		literal = "0" + string(l.lctx.curChar)
		l.nextChar() // 'o'または'O'をスキップ
		if !l.isOctChar(l.lctx.curChar) {
			ctx := l.lctx.toContext(l.start)
			l.logger.Error(fmt.Sprintf(errcode.ENUMBER_LITERAL, literal), &ctx)
			return Token{TokenType: NUMBER, SymbolID: intern.ID_STR_ZERO, Context: l.lctx.toContext(l.start)}
		}
		literal += l.readOctString()
		tok = Token{TokenType: NUMBER, SymbolID: intern.InternString(literal), Context: l.lctx.toContext(l.start)}
		l.nextChar()
		return tok

	case l.isDigit(l.lctx.curChar):
		// 10進リテラル, 16進リテラル（末尾 h）
		literal = l.readHexString()
		l.nextChar()
		ch := l.lctx.curChar
		if l.isHexString(literal) && (ch == 'h' || ch == 'H') {
			literal += string(ch)
			tok = Token{TokenType: NUMBER, SymbolID: intern.InternString(literal), Context: l.lctx.toContext(l.start)}
			l.nextChar()
			return tok
		}
		return Token{TokenType: NUMBER, SymbolID: intern.InternString(literal), Context: l.lctx.toContext(l.start)}

	case l.isAlpha(l.lctx.curChar):
		// IDENT, DOT_IDENT, 予約語
		b := l.readWord()
		if l.peekChar() == '.' {
			// LABEL abc.def
			l.nextChar()
			b = append(b, l.readWord()...)
			id := intern.InternBytes(b)
			tok = Token{TokenType: DOT_IDENT, SymbolID: id, Context: l.lctx.toContext(l.start)}
			l.nextChar()
			return tok
		}
		// AF' の処理
		if len(b) == 2 && (b[0] == 'a' || b[0] == 'A') && (b[1] == 'f' || b[1] == 'F') && l.peekChar() == '\'' {
			b = append(b, '\'')
			l.nextChar()
		}

		// intern literal
		id := intern.InternBytes(b)
		// z80 予約語
		tok, ok := z80ReservedWords[id]
		if ok {
			tok.Context = l.lctx.toContext(l.start)
			l.nextChar()
			return tok
		}
		// yas80 予約語
		tok, ok = reservedWords[id]
		if ok {
			tok.Context = l.lctx.toContext(l.start)
			l.nextChar()
			return tok
		}

		// これ以外は IDENT
		tok = Token{TokenType: IDENT, SymbolID: id, Context: l.lctx.toContext(l.start)}
		l.nextChar()
		return tok

	case l.lctx.curChar == '.':
		// LOCAL_IDENT
		if !l.isWordChar(l.peekChar()) {
			// "." のケース
			tok = Token{TokenType: LOCAL_IDENT, SymbolID: intern.InternString("."), Context: l.lctx.toContext(l.start)}
			l.nextChar()
			return tok
		}
		b := l.readWord()
		l.nextChar()
		id := intern.InternBytes(b)
		return Token{TokenType: LOCAL_IDENT, SymbolID: id, Context: l.lctx.toContext(l.start)}

	case l.lctx.curChar == '@':
		// AT_IDENT, ANON_IDENT
		ch := l.peekChar()
		switch {
		case ch == '@': // @@
			l.nextChar()
			l.nextChar()
			id := intern.InternString("@@")
			return Token{TokenType: ANON_IDENT, SymbolID: id, Context: l.lctx.toContext(l.start)}

		case ch == 'f' || ch == 'F' || ch == 'b' || ch == 'B': // @F, @B
			l.nextChar()
			l.nextChar()
			id := intern.InternBytes([]byte{'@', byte(ch)})
			return Token{TokenType: ANON_IDENT, SymbolID: id, Context: l.lctx.toContext(l.start)}

		case '0' <= ch && ch <= '9': // @n @nF @nB
			b := []byte{'@', byte(ch)}
			l.nextChar()
			ch = l.peekChar()
			if ch == 'f' || ch == 'F' || ch == 'b' || ch == 'B' {
				b = append(b, byte(ch))
				l.nextChar()
			}
			l.nextChar()
			id := intern.InternBytes(b)
			return Token{TokenType: ANON_IDENT, SymbolID: id, Context: l.lctx.toContext(l.start)}

		}

		b := l.readWord()
		l.nextChar()

		id := intern.InternBytes(b)
		return Token{TokenType: AT_IDENT, SymbolID: id, Context: l.lctx.toContext(l.start)}
	default:
		literal = string(l.lctx.curChar)
		tok = Token{TokenType: INVALID, SymbolID: intern.InternString(literal), Context: l.lctx.toContext(l.start)}
		l.nextChar()
		return tok
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
		tok = Token{TokenType: COMP, TokenSubType: LE, Context: l.lctx.toContext(l.start)}
	case ch1 == '<' && ch2 == '<':
		tok = Token{TokenType: SHIFT, TokenSubType: SL, Context: l.lctx.toContext(l.start)}
	case ch1 == '>' && ch2 == '=':
		tok = Token{TokenType: COMP, TokenSubType: GE, Context: l.lctx.toContext(l.start)}
	case ch1 == '>' && ch2 == '>':
		tok = Token{TokenType: SHIFT, TokenSubType: SR, Context: l.lctx.toContext(l.start)}
	case ch1 == '<' || ch1 == '>':
		tok = Token{TokenType: COMP, TokenSubType: TokenSubType(ch1), Context: l.lctx.toContext(l.start)}
	case ch1 == '=' && ch2 == '=':
		tok = Token{TokenType: COMP, TokenSubType: EQ, Context: l.lctx.toContext(l.start)}
	case ch1 == '!' && ch2 == '=':
		tok = Token{TokenType: COMP, TokenSubType: NEQ, Context: l.lctx.toContext(l.start)}
	case ch1 == '!':
		return Token{TokenType: UNARY, TokenSubType: TokenSubType(ch1), Context: l.lctx.toContext(l.start)}
	case ch1 == '&' && ch2 == '&':
		tok = Token{TokenType: AND, TokenSubType: 0, Context: l.lctx.toContext(l.start)}
	case ch1 == '&':
		return Token{TokenType: MULDIV, TokenSubType: TokenSubType(ch1), Context: l.lctx.toContext(l.start)}
	case ch1 == '|' && ch2 == '|':
		tok = Token{TokenType: OR, TokenSubType: 0, Context: l.lctx.toContext(l.start)}
	case ch1 == '|':
		return Token{TokenType: ADDSUB, TokenSubType: TokenSubType(ch1), Context: l.lctx.toContext(l.start)}
	case ch1 == '#' && ch2 == '#':
		tok = Token{TokenType: CONCAT, TokenSubType: 0, Context: l.lctx.toContext(l.start)}
	default:
		// 1文字トークンを返す
		return Token{TokenType: TokenType(ch1), Context: l.lctx.toContext(l.start)}
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
	// if l.lctx.curChar == '\n' {
	// 	l.lctx.lineNumber++
	// }
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

var escapeChars = map[rune]rune{
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

func (l *Lexer) readString(quote rune) string {
	runes := []rune{}

	index := l.lctx.index
	var ch rune
	for {
		l.nextChar()
		ch = l.lctx.curChar
		if ch == quote {
			break
		}
		ctx := l.lctx.toContext(index)
		if ch == '\n' {
			l.logger.Error(errcode.ESTR_END_QUOTE, &ctx)
			break
		}
		if ch < ' ' {
			l.logger.Error(errcode.ESTR_CTRL, &ctx)
			break
		}
		if quote != '`' && ch == '\\' { // raw string 以外はエスケープシーケンスを処理
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

func (l *Lexer) readWord() []byte {
	startIndex := l.lctx.index - 1
	for l.lctx.index < len(l.lctx.fileContent.Content) && l.isWordChar(rune(l.lctx.fileContent.Content[l.lctx.index])) {
		l.lctx.index++
	}
	return l.lctx.fileContent.Content[startIndex:l.lctx.index]
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

func (l *Lexer) isBinString(s string) bool {
	for _, c := range s {
		if !l.isBinChar(rune(c)) {
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
