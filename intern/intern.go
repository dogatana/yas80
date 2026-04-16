package intern

type SymbolID uint32

func (id SymbolID) String() string {
	return idToStr[id]
}

const symbolCap = 10000

var strToID map[string]SymbolID
var idToStr []string

func init() {
	strToID = make(map[string]SymbolID, symbolCap)
	idToStr = make([]string, 0, symbolCap)
}

// Intern: string → SymbolID
func Intern(s string) SymbolID {
	if id, ok := strToID[s]; ok {
		return id
	}
	id := SymbolID(len(idToStr))
	strToID[s] = id
	idToStr = append(idToStr, s)
	return id
}

// 大文字化 buffer
var upperBuf []byte

func InternBytes(b []byte) SymbolID {
	// 必要に応じて拡大
	if cap(upperBuf) < len(b) {
		upperBuf = make([]byte, len(b))
	}
	upperBuf = upperBuf[:len(b)]

	// 大文字化（コピーは upperBuf のみ）
	for i, c := range b {
		if 'a' <= c && c <= 'z' {
			upperBuf[i] = c - 32
		} else {
			upperBuf[i] = c
		}
	}

	// unsafe で string 化 - これはできた文字列の内容を壊すのでNG
	// s := *(*string)(unsafe.Pointer(&upperBuf)
	// string でコピーを作成する
	s := string(upperBuf)
	return Intern(s)
}

// Reverse lookup: SymbolID → string
func Lookup(id SymbolID) string {
	return idToStr[id]
}

// Token NUMBER, STRING のLiteral の intern
// 登録済みチェック、大文字化なし
func InternString(s string) SymbolID {
	id := SymbolID(len(idToStr))
	strToID[s] = id
	idToStr = append(idToStr, s)
	return id
}

var ID_LOC SymbolID      // $
var ID_ALOC SymbolID     // $$
var ID_PASS SymbolID     // $PASS
var ID_STAGE2 SymbolID   // $STAGE2
var ID_STR_ZERO SymbolID // "0"

func init() {
	ID_LOC = Intern("$")
	ID_ALOC = Intern("$$")
	ID_PASS = Intern("$PASS")
	ID_STAGE2 = Intern("$STAGE2")
	ID_STR_ZERO = InternString("0")
}
