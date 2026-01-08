package testutil

import (
	"regexp"
	"strings"
	"testing"
	"yas80/logging"
)

func TestLogMessage(t *testing.T, tn int, err string, logger *logging.Logger) {
	ename := ErrcodeNames[err]

	var msgs []logging.LogMessage
	switch ename[0] {
	case 'E':
		msgs = logger.Errors
		if len(msgs) == 0 {
			t.Fatalf("[%d] no error", tn)
			return
		}
	case 'W':
		msgs = logger.Warnings
		if len(msgs) == 0 {
			t.Fatalf("[%d] no warning", tn)
			return
		}
	case 'I':
		msgs = logger.Infomation
		if len(msgs) == 0 {
			t.Fatalf("[%d] no information", tn)
			return
		}
	}

	if !hasMessage(msgs, err) {
		t.Errorf("[%d] not [%s] \"%s\" => \"%s\"",
			tn,
			ename,
			err,
			msgs[0])
	}
}

func hasMessage(messages []logging.LogMessage, expected string) bool {
	re := regexp.MustCompile(`\.?%.\.?`)
	ss := re.Split(expected, -1)

	for _, emsg := range messages {
		result := true
		for _, s := range ss {
			if !strings.Contains(emsg.Message(), s) {
				result = false
				break
			}
		}
		if result {
			return result
		}
	}
	return false
}
