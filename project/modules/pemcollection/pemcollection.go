package pemcollection

import (
	"crypto/tls"
	"path/filepath"
)

type CollectionEntry struct {
	filePath string
	cert     tls.Certificate
}

func (ce *CollectionEntry) New(filePath string) error {
	ce.filePath = filepath.Clean(filePath)
	_ = ce.readFromFile()
	ce.cert, _ = tls.LoadX509KeyPair(ce.filePath, ce.filePath)
}
