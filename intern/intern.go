package intern

import "unsafe"

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

func Size() int { return len(idToStr) }

// Token NUMBER, STRING のLiteral の intern
// 大文字化なし
func InternString(s string) SymbolID {
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
	// readWord で新規確保した buf を使用するの安全にゼロコピーできる
	s := unsafe.String(&b[0], len(b))

	// InternString と同内容
	if id, ok := strToID[s]; ok {
		return id
	}
	id := SymbolID(len(idToStr))
	strToID[s] = id
	idToStr = append(idToStr, s)
	return id
}

// Reverse lookup: SymbolID → string
func Lookup(id SymbolID) string {
	return idToStr[id]
}

var ID_LOC SymbolID      // $
var ID_ALOC SymbolID     // $$
var ID_PASS SymbolID     // $PASS
var ID_STAGE2 SymbolID   // $STAGE2
var ID_STR_ZERO SymbolID // "0"
var ID_FILL SymbolID     // $FILL

func init() {
	ID_LOC = InternString("$")
	ID_ALOC = InternString("$$")
	ID_PASS = InternString("$PASS")
	ID_FILL = InternString("$FILL")
	ID_STAGE2 = InternString("$STAGE2")
	ID_STR_ZERO = InternString("0")
}
