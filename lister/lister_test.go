package lister

import (
	"testing"
	"yas80/errcode"
	"yas80/internal/testutil"
	"yas80/logging"
	"yas80/object"
)

func TestLister(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		{input: "zilog", err: errcode.EINCLUDE_CYCLIC},
	}

	for tn, tt := range tests {
		path := testutil.GetTestFilePath(t, tt.input+".asm")
		env := object.NewEnvironment(nil)
		logger := logging.New()

		code, _ := evalFile(path, logger, env)
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, logger)
			continue
		}

		expected := tt.code
		if expected == nil {
			expected = testutil.ReadTestDataFile(t, tt.input+".bin")
		}

		if err := testutil.BytesEqual(code, expected); err != nil {
			t.Errorf("[%d] generated code diff %s", tn, err.Error())
		}
	}
}
