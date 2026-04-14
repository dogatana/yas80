package intern

type SymbolID int

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

// Reverse lookup: SymbolID → string
func Lookup(id SymbolID) string {
	return idToStr[id]
}

var ID_LOC SymbolID  // $
var ID_ALOC SymbolID // $$

func init() {
	ID_LOC = Intern("$")
	ID_ALOC = Intern("$$")
}
