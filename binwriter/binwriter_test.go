package binwriter

import (
	"testing"
	"yas80/internal/testutil"
	"yas80/logging"
	"yas80/object"
)

func TestGapAndSortSegment(t *testing.T) {
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
	}

	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New("<binwrite>")
		prog, _ := evalInput(tt.input, logger, env)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, logger)
			continue
		}

		code, err := codeFromObj(prog, tt.fill)

		if err != nil {
			t.Errorf("[%d], codeFromObj %s", tn, err.Error())
		}

		if err = testutil.BytesEqual(code, tt.code); err != nil {
			t.Errorf("[%d], %s", tn, err.Error())
		}

	}
}
