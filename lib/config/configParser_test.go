package config

import (
	"hcc/horn/lib/logger"
	"testing"
)

func Test_Init(t *testing.T) {
	err = logger.Init()
	if err != nil {
		t.Fatal()
	}

	defer func() {
		_ = logger.FpLog.Close()
	}()

	Init()
}
