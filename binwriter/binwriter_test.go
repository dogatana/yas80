package binwriter

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dogatana/yas80/errcode"
	"github.com/dogatana/yas80/internal/testutil"
	"github.com/dogatana/yas80/logging"
	"github.com/dogatana/yas80/object"
)

func TestBinWriterGapAndSortSegment(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
		fill  int
	}{
		{
			input: `org 0 \ ei \ org 4 \ ret `,
			code:  []byte{0xfb, 0x00, 0x00, 0x00, 0xc9},
			fill:  0,
		},
		{
			input: `org 0 \ ei \ org 4 \ ret `,
			code:  []byte{0xfb, 0xff, 0xff, 0xff, 0xc9},
			fill:  255,
		},
		{
			input: `org 4 \ ret \ org 0 \ ei `,
			code:  []byte{0xfb, 0x80, 0x80, 0x80, 0xc9},
			fill:  0x80,
		},
		{
			input: `ld hl, $1234 \ org 1 \ ret`,
			err:   errcode.EBW_OVERLAPPED,
		},
	}

	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New()
		prog, _ := evalInput(tt.input, logger, env)
		code, ok := binFromObj(prog, tt.fill, logger)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, logger)
			continue
		}

		if !ok {
			t.Errorf("[%d], codeFromObj %s", tn, logger.GetErrors()[0].String())
		}

		if err := testutil.BytesEqual(code, tt.code); err != nil {
			t.Errorf("[%d], %s", tn, err.Error())
		}

	}
}

func TestBinWriterAbsAndRel(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
		fill  int
	}{
		{
			input: `org 0 \ call $ \ org $1000, rel \ call $ \ org $2000, rel \ call $ \ org $10 \ ret`,
			code: []byte{
				0xcd, 0x00, 0x00, 0xcd, 0x00, 0x10, 0xcd, 0x00, 0x20,
				0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				0xc9},
			fill: 255,
		},
	}

	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New()
		prog, _ := evalInput(tt.input, logger, env)
		code, ok := binFromObj(prog, tt.fill, logger)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, logger)
			continue
		}

		if !ok {
			t.Errorf("[%d], codeFromObj %s", tn, logger.GetErrors()[0].String())
		}

		if err := testutil.BytesEqual(code, tt.code); err != nil {
			t.Errorf("[%d], %s", tn, err.Error())
		}

	}
}

func TestBinWriterMultiFiles(t *testing.T) {
	tests := []struct {
		input any
		code  []byte
		err   string
		fill  int
	}{
		{
			input: []string{
				`ld a, 1`,
				`ld a, 2`,
				`ld a, 3`},
			code: []byte{0x3e, 0x01, 0x3e, 0x02, 0x3e, 0x03},
		},
		{
			input: []string{
				`ld a, 1`,
				`org 4 \ ld a, 2`,
				`org 8 \ ld a, 3`},
			code: []byte{0x3e, 0x01, 0, 0, 0x3e, 0x02, 0, 0, 0x3e, 0x03},
		},
		{
			input: []string{
				`ld a, 1`,
				`org $1234, REL \ ld hl, $`,
				`org $5678, REL \ ld hl, $`,
				`org 8 \ ld hl, $`},
			code: []byte{0x3e, 0x01, 0x21, 0x34, 0x12, 0x21, 0x78, 0x56, 0x21, 0x08, 0x00},
		},
	}

	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New()
		prog, _ := evalInput(tt.input, logger, env)
		code, ok := binFromObj(prog, tt.fill, logger)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, logger)
			continue
		}

		if !ok {
			t.Errorf("[%d], codeFromObj %s", tn, logger.GetErrors()[0].String())
		}

		if err := testutil.BytesEqual(code, tt.code); err != nil {
			t.Errorf("[%d], %s", tn, err.Error())
		}

	}
}

func TestRemoveDuplicateMessages(t *testing.T) {
	tests := []struct {
		input any
		code  []byte
		err   string
		fill  int
	}{
		{
			input: `ld a, 1024`, code: []byte{0x3e, 0x00},
		},
	}
	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New()
		prog, _ := evalInput(tt.input, logger, env)
		code, ok := binFromObj(prog, tt.fill, logger)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, logger)
			continue
		}

		if !ok {
			t.Errorf("[%d], codeFromObj %s", tn, logger.GetErrors()[0].String())
		}

		if err := testutil.BytesEqual(code, tt.code); err != nil {
			t.Errorf("[%d], %s", tn, err.Error())
		}

		if _, w, _ := logger.Count(); w != 2 {
			t.Errorf("[%d], logger should have 2 messages. got %d", tn, w)
		}
		logger.RemoveDupe()
		if _, w, _ := logger.Count(); w != 1 {
			t.Errorf("[%d], logger should have 1 messages. got %d", tn, w)
		}
	}
}

func TestMztBinarySize(t *testing.T) {
	tests := []struct {
		input any
		code  []byte
		name  string
		start int
		load  int
	}{
		{
			input: `nop`, code: []byte{0},
		},
	}
	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New()
		prog, _ := evalInput(tt.input, logger, env)

		mzt, ok := mztFromObj(prog, 0, logger, tt.name, 0)
		if !ok {
			t.Errorf("[%d], codeFromObj %s", tn, logger.GetErrors()[0].String())
		}

		if err := testutil.BytesEqual(mzt[0x80:], tt.code); err != nil {
			t.Errorf("[%d], %s", tn, err.Error())
		}

		if len(mzt) != len(tt.code)+0x80 {
			t.Errorf("[%d] mzt size not 0x%x. got 0x%x\n", tn, len(tt.code)+0x80, len(mzt))
			continue
		}
		size := int(mzt[0x13])*256 + int(mzt[0x12])
		if size != len(tt.code) {
			t.Errorf("[%d] code size in mzt not 0x%x. got 0x%x", tn, len(tt.code), size)
			continue
		}
	}
}

func TestMztLoadAddr(t *testing.T) {
	tests := []struct {
		input any
		code  []byte
		name  string
		start int
		load  int
	}{
		{input: `nop`, load: 0},
		{input: `org $1234 \ nop`, load: 0x1234},
	}
	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New()
		prog, _ := evalInput(tt.input, logger, env)

		mzt, ok := mztFromObj(prog, 0, logger, tt.name, 0)
		if !ok {
			t.Errorf("[%d], codeFromObj %s", tn, logger.GetErrors()[0].String())
		}

		load := int(mzt[0x15])*256 + int(mzt[0x14])
		if load != tt.load {
			t.Errorf("[%d] load addr not 0x%x. got 0x%x", tn, tt.load, load)
			continue
		}
	}
}

func TestMztStartAddr(t *testing.T) {
	tests := []struct {
		input any
		code  []byte
		name  string
		start int
		load  int
	}{
		{input: `nop`, start: 0},
		{input: `org $1234 \ nop`, start: 0x1234},
		{input: `org $1234 \ nop \ end $5678`, start: 0x5678},
		{input: `org $1234 \ nop \ entry $5678`, start: 0x5678},
		{input: `entry $1234 \ org $1234 \ nop \ entry $5678`, start: 0x5678},
	}
	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New()
		prog, _ := evalInput(tt.input, logger, env)

		mzt, ok := mztFromObj(prog, 0, logger, tt.name, 0)
		if !ok {
			t.Errorf("[%d], codeFromObj %s", tn, logger.GetErrors()[0].String())
		}

		end := int(mzt[0x17])*256 + int(mzt[0x16])
		if end != tt.load {
			t.Errorf("[%d] start addr not 0x%x. got 0x%x", tn, tt.start, end)
			continue
		}
	}
}

func TestMztLoadName(t *testing.T) {
	tests := []struct {
		input any
		code  []byte
		name  string
		start int
		load  int
	}{
		{input: `nop`, name: "short"},
		{input: `nop`, name: "0123456789012345"},
		{input: `nop`, name: "01234567890123456"},
		{input: `nop`, name: "012345678901234567"},
		{input: `nop`, name: "012345678901234568"},
	}
	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New()
		prog, _ := evalInput(tt.input, logger, env)

		mzt, ok := mztFromObj(prog, 0, logger, tt.name, 0)
		if !ok {
			t.Errorf("[%d], codeFromObj %s", tn, logger.GetErrors()[0].String())
		}

		mztName := mzt[1:0x12]

		// 16 文字以内 + 0d
		ttName := []byte(tt.name)
		ttName = append(ttName, 0x0d)
		ttName = append(ttName, bytes.Repeat([]byte{0x20}, 16)...)
		ttName = ttName[:16]
		ttName = append(ttName, 0x0d)

		if err := testutil.BytesEqual(mztName, ttName); err != nil {
			t.Errorf("[%d] name not %q. got %q", tn, string(ttName), string(mztName))
		}
	}
}

func TestT88LoadNameAndAddr(t *testing.T) {
	tests := []struct {
		input any
		code  []byte
		name  string
		start int
		load  int
	}{
		{input: `nop`, name: "t88", load: 0},
		{input: `org $1234 \ nop`, name: "t88", load: 0x1234},
	}
	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New()
		prog, _ := evalInput(tt.input, logger, env)

		ttname := ("$$$" + tt.name + strings.Repeat(" ", 6))[:9]

		name, load, err := t88FromObj(prog, 0, logger, tt.name)
		if err != nil {
			t.Errorf("[%d] t88FromObj error %s", tn, err.Error())
			continue
		}
		if name != ttname {
			t.Errorf("[%d] name not %q. got %q", tn, ttname, name)
			continue
		}

		if load != tt.load {
			t.Errorf("[%d] load addr not 0x%x. got 0x%x", tn, tt.load, load)
			continue
		}
	}
}
