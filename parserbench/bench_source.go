package parserbench

import (
	"fmt"
	"strings"
)

func GenerateLargeSource(lines int) string {
	var b strings.Builder
	b.Grow(lines * 40)

	for i := 0; i < lines; i++ {
		// ラベル + 命令 + 即値 + コメント
		b.WriteString(fmt.Sprintf("label_%d: ld a, 0x%02x ; comment\n", i, i&0xff))
	}

	return b.String()
}
