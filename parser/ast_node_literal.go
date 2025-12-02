package parser

var nodeLiteralTable map[int]string = map[int]string{
	NODE_NODE:               "NODE_NODE",
	NODE_PROGRAM:            "NODE_PROGRAM",
	NODE_ERROR:              "NODE_ERROR",
	NODE_STMT:               "NODE_STMT",
	NODE_DELETED_STMT:       "NODE_DELETED_STMT",
	NODE_LABEL_STMT:         "NODE_LABEL_STMT",
	NODE_EXPR_STMT:          "NODE_EXPR_STMT",
	NODE_CONST_STMT:         "NODE_CONST_STMT",
	NODE_VAR_STMT:           "NODE_VAR_STMT",
	NODE_ASIGN_STMT:         "NODE_ASIGN_STMT",
	NODE_ENUM_STMT:          "NODE_ENUM_STMT",
	NODE_ENUM_ELEMENTS_STMT: "NODE_ENUM_ELEMENTS_STMT",
	NODE_REPT_STMT:          "NODE_REPT_STMT",
	NODE_IF_STMT:            "NODE_IF_STMT",
	NODE_BLOCK_STMT:         "NODE_BLOCK_STMT",
	NODE_FUNC_STMT:          "NODE_FUNC_STMT",
	NODE_EXITM_STMT:         "NODE_EXITM_STMT",
	NODE_RETURN_STMT:        "NODE_RETURN_STMT",
	NODE_PROC_STMT:          "NODE_PROC_STMT",
	NODE_MACRO_STMT:         "NODE_MACRO_STMT",
	NODE_MACRO_CALL_STMT:    "NODE_MACRO_CALL_STMT",
	NODE_EXPR:               "NODE_EXPR",
	NODE_ENUM_ELEMENT:       "NODE_ENUM_ELEMENT",
	NODE_NUMBER:             "NODE_NUMBER",
	NODE_STRING:             "NODE_STRING",
	NODE_IDENT:              "NODE_IDENT",
	NODE_LOCAL_IDENT:        "NODE_LOCAL_IDENT",
	NODE_DOT_IDENT:          "NODE_DOT_IDENT",
	NODE_ARRAY:              "NODE_ARRAY",
	NODE_INDEXED_EXPR:       "NODE_INDEXED_EXPR",
	NODE_INFIX_EXPR:         "NODE_INFIX_EXPR",
	NODE_PREFIX_EXPR:        "NODE_PREFIX_EXPR",
	NODE_CALL:               "NODE_CALL",
	NODE_EXPR_LIST:          "NODE_EXPR_LIST",
	NODE_LABEL:              "NODE_LABEL",
	NODE_LOCAL_LABEL:        "NODE_LOCAL_LABEL",
	NODE_AT_LABEL:           "NODE_AT_LABEL",
	NODE_INDIRECT:           "NODE_INDIRECT",
}

func nodeLiteral(id int) string {
	if s, ok := nodeLiteralTable[id]; ok {
		return s
	}
	return "NODE_???"
}
