package pemcollection

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"github.com/youmark/pkcs8"
	"os"
	"strings"
)

func LoadPemCollection(certFile string, password string) (tls.Certificate, error) {
	fileContent, err := os.ReadFile(certFile)
	decryptedContent, err := DecryptPrivateKey(fileContent, password)
	if err != nil {
		return tls.Certificate{}, err
	}
	otherContent := decryptedContent
	cert, err := tls.X509KeyPair(decryptedContent, otherContent)
	return cert, err
}

func DecryptPrivateKey(certPEMBlock []byte, password string) ([]byte, error) {
	var newPem []byte
	var err error
	for {
		var certDERBlock *pem.Block
		certDERBlock, certPEMBlock = pem.Decode(certPEMBlock)
		if certDERBlock == nil {
			break
		}
		var keyPem []byte
		if certDERBlock.Type == "PRIVATE KEY" || strings.HasSuffix(certDERBlock.Type, " PRIVATE KEY") {
			keyPem, err = processKey(certDERBlock, password)
			newPem = append(newPem, keyPem...)
		} else {
			keyPem = pem.EncodeToMemory(certDERBlock)
			newPem = append(newPem, keyPem...)
		}
	}
	return newPem, err
}

func processKey(certDERBlock *pem.Block, password string) ([]byte, error) {
	var keyBlk []uint8
	var err error
	if strings.Contains(certDERBlock.Type, "ENCRYPTED") {
		keyBlk, err = decryptKey(certDERBlock.Bytes, password)
	} else {
		keyBlk, err = copyKey(certDERBlock.Bytes)
	}
	keyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: keyBlk,
	})
	return keyPem, err
}

func decryptKey(block []byte, password string) ([]byte, error) {
	pkey, _ := pkcs8.ParsePKCS8PrivateKey(block, []byte(password))
	keyBlk, err := x509.MarshalPKCS8PrivateKey(pkey)
	return keyBlk, err
}

func copyKey(block []byte) ([]byte, error) {
	pkey := block
	keyBlk, err := x509.MarshalPKCS8PrivateKey(pkey)
	return keyBlk, err
}
