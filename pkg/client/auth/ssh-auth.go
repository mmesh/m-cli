package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"hash"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
	"mmesh.dev/m-lib/pkg/errors"
)

func (api *API) SSHAuth() bool {
	if _, err := getSSHPrivKey(); err != nil {
		return false
	}

	return true
}

// getSSHPrivKey gets the configured ssh key
func getSSHPrivKey() (string, error) {
	privKeyFile := viper.GetString("auth.ssh.privateKey")

	if len(privKeyFile) == 0 {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", errors.Wrapf(err, "[%v] function os.UserHomeDir()", errors.Trace())
		}

		privKeyFile = filepath.Join(homeDir, ".mmesh", "id_rsa")
	}

	if !fileExists(privKeyFile) {
		return "", errors.New("missing private key")
	}

	privkey, err := ioutil.ReadFile(privKeyFile)
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function ioutil.ReadFile()", errors.Trace())
	}

	return string(privkey), nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func rsaDecrypt(data string) (string, error) {
	priv, err := getSSHPrivKey()
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function getSSHPrivKey()", errors.Trace())
	}

	sshkey, err := ssh.ParseRawPrivateKey([]byte(priv))
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function ssh.ParseRawPrivateKey()", errors.Trace())
	}

	key := sshkey.(*rsa.PrivateKey)

	str, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function base64.StdEncoding.DecodeString()", errors.Trace())
	}

	decrypted, err := decryptOAEP(sha256.New(), rand.Reader, key, str, nil)
	if err != nil {
		return "", errors.Wrapf(err, "[%v] function decryptOAEP()", errors.Trace())
	}

	return string(decrypted), nil
}

func decryptOAEP(hash hash.Hash, random io.Reader, private *rsa.PrivateKey, msg []byte, label []byte) ([]byte, error) {
	msgLen := len(msg)
	step := private.PublicKey.Size()
	var decryptedBytes []byte

	for start := 0; start < msgLen; start += step {
		finish := start + step
		if finish > msgLen {
			finish = msgLen
		}

		decryptedBlockBytes, err := rsa.DecryptOAEP(hash, random, private, msg[start:finish], label)
		if err != nil {
			return nil, errors.Wrapf(err, "[%v] function rsa.DecryptOAEP()", errors.Trace())
		}

		decryptedBytes = append(decryptedBytes, decryptedBlockBytes...)
	}

	return decryptedBytes, nil
}
