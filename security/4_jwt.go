package security

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// UserClaims struct
type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

// JWT func
func jwtFunc() {
	pass := UserClaims{
		SessionID: 10,
	}

	hashedPass, err := createToken(&pass)
	fmt.Println(hashedPass)
	if err != nil {
		panic(err)
	}

	result, err := parseToken(hashedPass)
	fmt.Println(result)
}

// Valid func
func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("Token has expired")
	}

	if u.SessionID == 0 {
		return fmt.Errorf("Invalid session ID")
	}

	return nil
}

func createToken(c *UserClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	signedToken, err := t.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("Error in createToken when signing token: %w", err)
	}
	return signedToken, nil
}

func parseToken(signedToken string) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(signedToken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("Invalid signing algorithm")
		}

		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("Error in parseToken while parsing token: %w", err)
	}

	if !t.Valid {
		return nil, fmt.Errorf("Error in parseToken, token is not valid")
	}

	return t.Claims.(*UserClaims), nil
}
