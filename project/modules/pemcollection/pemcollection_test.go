package pemcollection

import (
	"dc.local/zlogger"
	"os"
	"path"
	"runtime"
	"testing"
)

var global_testdataPath string

func init() {
	global_testdataPath = testdataDir()
	zlogger.SetGlobalLevel(zlogger.WarnLevel)
}

func TestPemFile_New_DoesNotExist(t *testing.T) {
	pemPackagePath := path.Join(global_testdataPath, "client/packaged.pem")
	pf := new(PemCollection)
	err := pf.New(pemPackagePath, "VerySecretPassword")
	if err == nil {
		t.Errorf("Failed: FileNotFound expected")
	}
}

func testdataDir() string {
	_, filename, _, _ := runtime.Caller(0)
	testdataPath := path.Join(path.Dir(filename), "testdata")
	if _, err := os.Stat(testdataPath); os.IsNotExist(err) {
		return ""
	}
	return testdataPath
}
