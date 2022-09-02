package security

import "fmt"

func Execute() {
	fmt.Println("=== Testing base64 ===")
	base64Func()
	fmt.Println("=== Testing bcrypt ===")
	bcryptFunc()
	fmt.Println("=== Testing hmac ===")
	hmacFunc()
	fmt.Println("=== Testing jwt ===")
	jwtFunc()
	fmt.Println("=== Testing aes ===")
	aesFunc()
	fmt.Println("=== Testing sha256 ===")
	sha256Func()
}
