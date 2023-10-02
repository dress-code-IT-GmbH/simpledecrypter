package pemcollection

import (
	"crypto/tls"
	"path/filepath"
)

type PemCollection struct {
	filePath string
	cert     tls.Certificate
}

func (ce *PemCollection) New(filePath string, password string) error {
	ce.filePath = filepath.Clean(filePath)
	ce.cert, _ = LoadPemCollection(ce.filePath, password)
	return nil
}

func (ce *PemCollection) Keys()
