package binwriter

import (
	"testing"
	"yas80/errcode"
	"yas80/internal/testutil"
	"yas80/logging"
	"yas80/object"
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
		{
			input: ` `,
			err:   errcode.EBW_NULL,
		},
	}

	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New("<binwrite>")
		prog, _ := evalInput(tt.input, logger, env)
		code, ok := codeFromObj(prog, tt.fill, logger)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, logger)
			continue
		}

		if !ok {
			t.Errorf("[%d], codeFromObj %s", tn, logger.Errors[0].Error())
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
		logger := logging.New("<binwrite>")
		prog, _ := evalInput(tt.input, logger, env)
		code, ok := codeFromObj(prog, tt.fill, logger)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, logger)
			continue
		}

		if !ok {
			t.Errorf("[%d], codeFromObj %s", tn, logger.Errors[0].Error())
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
		logger := logging.New("<binwrite>")
		prog, _ := evalInput(tt.input, logger, env)
		code, ok := codeFromObj(prog, tt.fill, logger)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, logger)
			continue
		}

		if !ok {
			t.Errorf("[%d], codeFromObj %s", tn, logger.Errors[0].Error())
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
		logger := logging.New("<binwrite>")
		prog, _ := evalInput(tt.input, logger, env)
		code, ok := codeFromObj(prog, tt.fill, logger)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, logger)
			continue
		}

		if !ok {
			t.Errorf("[%d], codeFromObj %s", tn, logger.Errors[0].Error())
		}

		if err := testutil.BytesEqual(code, tt.code); err != nil {
			t.Errorf("[%d], %s", tn, err.Error())
		}

		if len(logger.Warnings) != 2 {
			t.Errorf("[%d], logger should have 2 messages. got %d", tn, len(logger.Warnings))
		}
		logger.RemoveDupe()
		if len(logger.Warnings) != 1 {
			t.Errorf("[%d], logger should have 1 messages. got %d", tn, len(logger.Warnings))
		}
	}
}
