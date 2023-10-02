package base64OeapSha256

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"dc.local/zlogger"
	"encoding/base64"
	"github.com/pkg/errors"
)

func B64Decode(cipherTextB64Str string) ([]byte, error) {
	logger := zlogger.GetLogger("decryptor.B64Decode.Decrypt")
	cipherTextB64 := []byte(cipherTextB64Str)
	cipherText := make([]byte, base64.StdEncoding.DecodedLen(len(cipherTextB64)))
	n, e := base64.StdEncoding.Decode(cipherText, cipherTextB64)
	if e != nil {
		err := errors.Wrap(e, "base64 decoding failed")
		logger.Error().Err(err).Msg("")
		return nil, err
	}
	cipherText = cipherText[:n]

	return cipherText, nil
}

func Decrypt(cipherTextB64Str string, key *rsa.PrivateKey) (string, error) {
	logger := zlogger.GetLogger("decryptor.base64OeapSha256.Decrypt")
	cipherText, e := B64Decode(cipherTextB64Str) // Hier wird B64Decode aufgerufen
	clearText, e := rsa.DecryptOAEP(sha256.New(), nil, key, cipherText, []byte(""))
	if e != nil {
		err := errors.Wrap(e, "base64 decoding failed")
		logger.Error().Err(err).Msg("")
		return "", err
	}
	ctx := string(clearText)
	return ctx, nil
}

func Encrypt(clearText string, key *rsa.PublicKey) (string, error) {
	logger := zlogger.GetLogger("decryptor.base64OeapSha256.EnCrypt")
	decClearText := []byte(clearText)
	cipherText, e := rsa.EncryptOAEP(sha256.New(), rand.Reader, key, decClearText, []byte(""))
	if e != nil {
		err := errors.Wrap(e, "encryption failed")
		logger.Error().Err(err).Msg("")
		return "", err
	}
	cipherTextB64 := base64.StdEncoding.EncodeToString(cipherText)
	return cipherTextB64, nil
}
