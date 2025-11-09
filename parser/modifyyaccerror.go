package parser

import (
	"fmt"
)

// goyacc の生成メッセージを適切なものに置換する
func modifyYaccError(msg string, yyval, yylval *Token) string {
	return fmt.Sprintf("%s VAL(%s) lval(%s)", msg, yyval.Literal, yylval.Literal)
	// switch {
	// case msg == "unexpected NUMBER":
	// 	return fmt.Sprintf("[E]ここで数値 %s は使えません", yylval.Literal)
	// case msg == "unexpected IDENT":
	// 	return fmt.Sprintf("[E]ここで識別子 %s は使えません", yylval.Literal)
	// case strings.Contains(msg, ", expected ':'"):
	// 	return "[E]ラベル指定には ':' が必要です"
	// case strings.HasPrefix(msg, "unexpected IDENT,"):
	// 	return fmt.Sprintf("[E]ここで識別子 %s は使えません", yylval.Literal)
	// case strings.HasPrefix(msg, "unexpected EOL"):
	// 	return "[E]行が途中で終了しています"
	// case strings.HasPrefix(msg, "unexpected "):
	// 	return fmt.Sprintf("[E]ここで %s は使えません", yylval.Literal)
	// default:
	// 	return fmt.Sprintf("%q yyVAL(%s) yylval(%s)", msg, yyval.String(), yylval.String())
	// }
}
