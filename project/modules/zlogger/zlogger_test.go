package zlogger

import (
	"bytes"
	"github.com/rs/zerolog"
	"testing"
)

func TestSetGlobalLevel(t *testing.T) {
	SetGlobalLevel(DebugLevel)
	if zerolog.GlobalLevel() != DebugLevel {
		t.Errorf("Failed to set DebugLevel")
	}
	SetGlobalLevel(InfoLevel)
	if zerolog.GlobalLevel() != InfoLevel {
		t.Errorf("Failed to set InfoLevel")
	}
	SetGlobalLevel(WarnLevel)
	if zerolog.GlobalLevel() != WarnLevel {
		t.Errorf("Failed to set InfoLevel")
	}
}

func TestGetLogger(t *testing.T) {
	SetGlobalLevel(DebugLevel)
	mylogger := GetLogger("foobar.test")
	out := &bytes.Buffer{}
	newlogger := mylogger.Output(out)
	newlogger.Error().Msg("footext")
	got := string(out.Bytes())
	if got == "{\"level\":\"error\",\"component\":\"foobar.test\",\"time\":1680920893,\"message\":\"footext\"}\n" {
		return
	}
	t.Errorf("logger did not log expected line")
}
