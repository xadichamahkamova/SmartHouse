package hasher

import (
	"crypto/sha1"
	"fmt"
)

func HashUserPassword(pass string) (string, error) {

	hash := sha1.New()
	_, err := hash.Write([]byte(pass))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
