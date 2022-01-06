package logger

import (
	"testing"
)

func Test_Logger_Prepare(t *testing.T) {
	err := Init()
	if err != nil {
		t.Fatal()
	}
	defer func() {
		_ = FpLog.Close()
	}()
}
