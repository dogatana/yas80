package object

import (
	"bytes"
	"fmt"
	"strings"
	"yas80/filecontent"
	"yas80/internal/util"
	"yas80/parser"
)

const (
	OBJ_NULL = iota + 1
	OBJ_ERROR
	OBJ_NUMBER
	OBJ_STRING
	OBJ_ENUM
	OBJ_REGISTER
	OBJ_REG_INDIRECT
	OBJ_ADDR_INDIRECT
	OBJ_ORG
	OBJ_CODE
	OBJ_NODE
	OBJ_NODES
	OBJ_RETURN
	OBJ_BLOCK
	OBJ_PROC
	OBJ_FUNC
	OBJ_CHARMAP
	OBJ_MACRO
	OBJ_REF_NOTFOUND
	OBJ_SYMBOL
	OBJ_EXITM
	OBJ_VALUE
	OBJ_TEXT // OBJ_VALUE を代替する Lister 用 Object
	OBJ_FILE // 入力ファイル（include含む）変更通知用 Object
	OBJ_COMMENT
	OBJ_ARRAY
)

// 同一判定のため定数として定義
var (
	NULL  = &NullObject{}
	ERROR = &ErrorObject{}
)

type ObjectType int

type Object interface {
	Type() ObjectType
	String() string
}

// value - list ファイル出力用
type ValueObject struct {
	Value   Object
	Context *filecontent.Context
}

func (o *ValueObject) Type() ObjectType { return OBJ_VALUE }
func (o *ValueObject) String() string {
	return fmt.Sprintf("VALUE(%s)", o.Value.String())
}

// text - list ファイル出力用
type TextObject struct {
	Text    string
	Context *filecontent.Context
}

func (o *TextObject) Type() ObjectType { return OBJ_TEXT }
func (o *TextObject) String() string {
	return fmt.Sprintf("TEXT %s", o.Text)
}

// file
type FileObject struct {
	Filename string
	Line     int
}

func (o *FileObject) Type() ObjectType { return OBJ_FILE }
func (o *FileObject) String() string {
	return fmt.Sprintf("FILE %q:%d", o.Filename, o.Line)
}

// value - list ファイル出力用
type CommentObject struct {
	Text    string
	Context *filecontent.Context
}

func (o *CommentObject) Type() ObjectType { return OBJ_COMMENT }
func (o *CommentObject) String() string {
	var out bytes.Buffer

	if o.Context != nil {
		out.WriteString(fmt.Sprintf("comt %2d:%2d", o.Context.Line, o.Context.Offset))
		if o.Context.Source == nil {
			out.WriteString("(  ) ")
		} else {
			out.WriteString(fmt.Sprintf("(%2d) ", o.Context.Source.Line))
		}
	} else {
		out.WriteString("comt   :  (  ) ")
	}
	out.WriteString(o.Text)
	return out.String()
}

// org
type OrgObject struct {
	Addr      int
	AllocType int
}

func (o *OrgObject) Type() ObjectType { return OBJ_ORG }
func (o *OrgObject) String() string {
	if o.AllocType == 0 {
		return fmt.Sprintf("ORG $%04x", o.Addr)
	} else {
		return fmt.Sprintf("ORG $%04x, REL", o.Addr)
	}
}

// code
type CodeObject struct {
	Addr    int
	Code    []byte
	CZ80    int
	CR800   int
	Context *filecontent.Context
}

func (o *CodeObject) Type() ObjectType { return OBJ_CODE }
func (o *CodeObject) String() string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("code %2d:%2d", o.Context.Line, o.Context.Offset))
	if o.Context.Source == nil {
		out.WriteString("(  )")
	} else {
		out.WriteString(fmt.Sprintf("(%2d)", o.Context.Source.Line))
	}
	out.WriteString(fmt.Sprintf(" :%04x: [%2d] ", o.Addr, o.CZ80))
	for _, b := range o.Code {
		out.WriteString(fmt.Sprintf("%02x ", b))
	}
	return out.String()
}
func (o *CodeObject) Size() int {
	return len(o.Code)
}

// NULL
type NullObject struct{}

func (o *NullObject) Type() ObjectType { return OBJ_NULL }
func (o *NullObject) String() string   { return "NULL" }

// Error
type ErrorObject struct {
}

func (o *ErrorObject) Type() ObjectType { return OBJ_ERROR }
func (o *ErrorObject) String() string   { return "ERROR" }

// 右辺値で識別子が見つからない場合に使用
type RefNotFoundObject struct {
	Names []string
}

func (o *RefNotFoundObject) Type() ObjectType { return OBJ_REF_NOTFOUND }
func (o *RefNotFoundObject) String() string {
	var out bytes.Buffer

	out.WriteString("REF_NOTFOUND(")
	out.WriteString(strings.Join(o.Names, ", "))
	out.WriteRune(')')
	return out.String()
}

// proc - Object interface と Environ interface を実装する
type ProcObject struct {
	Name string
	Addr int
	Env  Environment
}

// Object の実装
func (o *ProcObject) Type() ObjectType { return OBJ_PROC }
func (o *ProcObject) String() string   { return fmt.Sprintf("PROC(%s[0x%04x])", o.Name, o.Addr) }

// Environ interface の実装
func (o *ProcObject) EnvType() int                       { return o.Env.EnvType() }
func (o *ProcObject) Get(name string) (Object, bool)     { return o.Env.Get(name) }
func (o *ProcObject) Set(name string, obj Object) Object { return o.Env.Set(name, obj) }
func (o *ProcObject) Outer() Environment                 { return o.Env.Outer() }
func (o *ProcObject) Store() map[string]Object           { return o.Env.Store() }

// enum - Object interface と Environ interface を実装する
type EnumObject struct {
	Name string
	Env  Environment
}

// Object の実装
func (o *EnumObject) Type() ObjectType { return OBJ_ENUM }
func (o *EnumObject) String() string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("ENUM(%s) {\n", o.Name))
	for k, v := range o.Env.Store() {
		out.WriteString(fmt.Sprintf("%s=%s\n", k, v.String()))
	}
	out.WriteString("}")

	return out.String()
}

// Environ interface の実装
func (o *EnumObject) EnvType() int                       { return o.Env.EnvType() }
func (o *EnumObject) Get(name string) (Object, bool)     { return o.Env.Get(name) }
func (o *EnumObject) Set(name string, obj Object) Object { return o.Env.Set(name, obj) }
func (o *EnumObject) Outer() Environment                 { return o.Env.Outer() }
func (o *EnumObject) Store() map[string]Object           { return o.Env.Store() }

// exitm
type ExitmObject struct {
}

func (o *ExitmObject) Type() ObjectType { return OBJ_EXITM }
func (o *ExitmObject) String() string   { return "EXITM" }

// return
type ReturnObject struct {
	Value      Object
	LineNumber int
}

func (o *ReturnObject) Type() ObjectType { return OBJ_RETURN }
func (o *ReturnObject) String() string {
	if o.Value == NULL {
		return "RETURN"
	}
	return "RETURN " + o.Value.String()
}

// 数値
type NumberObject struct {
	Value     int
	ForceWord bool // true: word 強制
	Context   *filecontent.Context
}

func (o *NumberObject) Type() ObjectType { return OBJ_NUMBER }
func (o *NumberObject) String() string   { return fmt.Sprintf("%d(0x%x)", o.Value, o.Value) }

// 文字列
type StringObject struct {
	Value   string
	Context *filecontent.Context
}

func (o *StringObject) Type() ObjectType { return OBJ_STRING }
func (o *StringObject) String() string   { return fmt.Sprintf("%q", o.Value) }

// レジスタ間接
type RegIndirectObject struct {
	Register     int
	Displacement int
}

func (o *RegIndirectObject) Type() ObjectType { return OBJ_REG_INDIRECT }
func (o *RegIndirectObject) String() string {
	if o.Displacement != 0 {
		return fmt.Sprintf("(%s%+d)", parser.TokenLiteral(o.Register), o.Displacement)
	} else {
		return fmt.Sprintf("(%s)", parser.TokenLiteral(o.Register))
	}
}

// アドレス間接
type AddrIndirectObject struct {
	Address int
}

func (o *AddrIndirectObject) Type() ObjectType { return OBJ_ADDR_INDIRECT }
func (o *AddrIndirectObject) String() string {
	return fmt.Sprintf("($%x)", o.Address)
}

// Node
type StatementObject struct {
	Statement parser.Statement
}

func (n *StatementObject) Type() ObjectType { return OBJ_NODE }
func (n *StatementObject) String() string   { return n.Statement.String() }

// Nodes
type StatemetnsObject struct {
	Statements []parser.Statement
}

func (n *StatemetnsObject) Type() ObjectType { return OBJ_NODES }
func (n *StatemetnsObject) String() string   { return fmt.Sprintf("NODES(%v)", n.Statements) }

// Block
type BlockObject struct {
	Block []Object
}

func (o *BlockObject) Type() ObjectType { return OBJ_BLOCK }
func (o *BlockObject) String() string {
	strs := []string{}

	for _, o := range o.Block {
		strs = append(strs, o.String())
	}
	return strings.Join(strs, "\n")
}

// charmap
type CharamapObject struct {
	Name    string
	DefChar int
	Cmap    map[string][]byte
	Context *filecontent.Context
}

func (o *CharamapObject) Type() ObjectType { return OBJ_CHARMAP }
func (o *CharamapObject) String() string {
	return fmt.Sprintf("CHARAMAP %s (%d chars)", o.Name, len(o.Cmap))
}

// Function
type FunctionObject struct {
	Name    string
	Params  []string
	Body    parser.Node
	Env     Environment
	Context *filecontent.Context
}

func (o *FunctionObject) Type() ObjectType { return OBJ_FUNC }
func (o *FunctionObject) String() string {
	var out bytes.Buffer

	out.WriteString(o.Name + " FUNC")
	if len(o.Params) > 0 {
		out.WriteRune(' ')
		out.WriteString(strings.Join(o.Params, ", "))
	}
	out.WriteRune('\n')
	out.WriteString(o.Body.String() + "\n")
	out.WriteString("ENDF")

	return out.String()
}

// Macro
type MacroObject struct {
	Name   string
	Params []string
	End    int
	Body   *parser.BlockStatement
}

func (o *MacroObject) Type() ObjectType { return OBJ_MACRO }
func (o *MacroObject) String() string {
	var out bytes.Buffer

	out.WriteString(o.Name + " MACRO")
	if len(o.Params) > 0 {
		out.WriteRune(' ')
		out.WriteString(strings.Join(o.Params, ", "))
	}
	out.WriteString(" \\ " + o.Body.String() + " \\ ENDM")

	return out.String()
}

// 配列
type ArrayObject struct {
	Values      []Object
	Expressions []parser.Expression
}

func (o *ArrayObject) Type() ObjectType { return OBJ_ARRAY }
func (o *ArrayObject) String() string {

	values := []string{}
	for _, v := range o.Values {
		values = append(values, v.String())
	}

	var out bytes.Buffer
	out.WriteRune('[')
	out.WriteString(strings.Join(values, ","))
	out.WriteRune(']')

	return out.String()
}

// prog(*object.BlockObject) を単一階層の slice に変換
func FlattenObject(obj Object) []Object {
	objs := []Object{}

	switch obj := obj.(type) {
	case *BlockObject:
		for _, o := range obj.Block {
			objs = append(objs, FlattenObject(o)...)
		}
	default:
		objs = append(objs, obj)
	}
	return objs
}

func IsTruthy(obj Object) bool {
	switch obj := obj.(type) {
	case *NumberObject:
		return obj.Value != 0
	case *StringObject:
		return obj.Value != ""
	default:
		return false
	}
}

func CollectCode(objects []Object) []byte {
	var result []byte

	for _, obj := range objects {
		switch obj := obj.(type) {
		case *CodeObject:
			result = append(result, obj.Code...)
		case *BlockObject:
			result = append(result, CollectCode(obj.Block)...)
		}
	}
	return result
}

// map[ファイル名][行番号] []Object
type FileMap map[string]*util.OrderedMap[int, []Object]

func (fm FileMap) Print() {
	for name, lm := range fm {
		for _, line := range lm.Keys() {
			objs, _ := lm.Get(line)
			fmt.Printf("%s [%d] %d objects\n", name, line, len(objs))
			for _, o := range objs {
				fmt.Println(o.String())
			}
		}
	}
}

// ファイルを Grouup の map
func BuildGroupMap(objects []Object) FileMap {
	fmap := map[string]*util.OrderedMap[int, []Object]{}

	var name string
	for _, obj := range objects {
		switch obj := obj.(type) {
		case *FileObject:
			name = obj.Filename
			if _, ok := fmap[name]; !ok {
				fmap[name] = util.NewOrderedMap[int, []Object]()
			}
		case *CodeObject:
			fm, _ := fmap[name]
			if objs, ok := fm.Get(obj.Context.Line); !ok {
				fm.Set(obj.Context.Line, []Object{obj})
			} else {
				objs = append(objs, obj)
				fm.Set(obj.Context.Line, []Object{obj})
			}
		case *CommentObject:
			fm, _ := fmap[name]
			if objs, ok := fm.Get(obj.Context.Line); !ok {
				fm.Set(obj.Context.Line, []Object{obj})
			} else {
				objs = append(objs, obj)
				fm.Set(obj.Context.Line, []Object{obj})
			}
		}
	}
	return fmap
}

type FileBlock struct {
	Filename    string
	Line        int
	LineObjects *util.OrderedMap[int, []Object]
}

func (fb *FileBlock) Print() {
	fmt.Printf("%s: %d\n", fb.Filename, fb.Line)
	for _, line := range fb.LineObjects.Keys() {
		objs, _ := fb.LineObjects.Get(line)
		fmt.Printf("%d: %d objects\n", line, len(objs))
		for _, o := range objs {
			fmt.Println(o.String())
		}
	}
}

// 評価結果を FileBlock のスライスにして返す
func BuildFileBlock(objects []Object) []*FileBlock {
	blocks := []*FileBlock{}

	var fb *FileBlock
	for _, obj := range objects {
		switch obj := obj.(type) {
		case *FileObject:
			if fb != nil {
				blocks = append(blocks, fb)
			}
			lo := util.NewOrderedMap[int, []Object]()
			name := util.SlashPath(obj.Filename) // path を / 区切りとする
			fb = &FileBlock{Filename: name, Line: obj.Line, LineObjects: lo}

		case *CodeObject:
			line := obj.Context.Line
			if objs, ok := fb.LineObjects.Get(line); !ok {
				fb.LineObjects.Set(line, []Object{obj})
			} else {
				objs = append(objs, obj)
				fb.LineObjects.Set(line, objs)
			}

		case *CommentObject:
			line := obj.Context.Line
			if objs, ok := fb.LineObjects.Get(line); !ok {
				fb.LineObjects.Set(line, []Object{obj})
			} else {
				objs = append(objs, obj)
				fb.LineObjects.Set(line, objs)
			}

		case *TextObject:
			line := obj.Context.Line
			if objs, ok := fb.LineObjects.Get(line); !ok {
				fb.LineObjects.Set(line, []Object{obj})
			} else {
				objs = append(objs, obj)
				fb.LineObjects.Set(line, objs)
			}
		}
	}
	blocks = append(blocks, fb)
	return blocks
}
