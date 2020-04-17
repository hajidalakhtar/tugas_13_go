package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "Hajid Al akhtar"

	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	var sha = sha1.New()
	sha.Write([]byte(encodedString))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	fmt.Println(encryptedString)
}
