package main

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
)

func main() {
	someText := "mmatos"
	hash, err := hashTextTo32Bytes(someText)
	fmt.Printf("%s\n %s", hash, err)
}

func hashTextTo32Bytes(hashThis string) (hashed string, err error) {

	if len(hashThis) == 0 {
		return "", errors.New("No input supplied")
	}

	hasher := sha256.New()
	hasher.Write([]byte(hashThis))

	stringToSHA256 := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	// Cut the length down to 32 bytes and return.
	return stringToSHA256[:32], nil
}
