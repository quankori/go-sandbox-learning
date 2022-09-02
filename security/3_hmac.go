package security

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"log"
)

var key = []byte{}

// Hmac func
func hmacFunc() {
	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}

	hashedPass, err := signMessage(key)
	fmt.Println(hashedPass)
	if err != nil {
		panic(err)
	}

	result, err := checkSig(key, hashedPass)
	if err != nil {
		log.Fatalln("Not logged in")
	}
	log.Println(result)
	log.Println("Logged in!")
}

func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)

	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("Error in signMessage while hashing message: %w", err)
	}

	signature := h.Sum(nil)
	return signature, nil
}

func checkSig(msg, sig []byte) (bool, error) {
	newSig, err := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("Error in checkSig while getting signature of message: %w", err)
	}

	same := hmac.Equal(newSig, sig)
	return same, nil
}
