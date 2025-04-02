package logging_test

import (
	"gosingleton/internal/logging"
	"testing"
)

func Test_NewLogger(t *testing.T) {
	t.Run("should be a single instance", func(t *testing.T) {
		l1 := logging.NewLogger()
		l2 := logging.NewLogger()

		if &l1 != &l2 {
			t.Errorf("Expected %v to equal %v", &l1, &l2)
		}
	})
}

func Test_Logger_Info(t *testing.T) {
	l := logging.NewLogger(logging.WithJsonHandler(), logging.WithJsonHandler())
	l.Info("This should be a json log", "key", "value")
}
