package security

import (
	"os"
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
)

var (
	JwtSecretKey     = []byte(os.Getenv("JWT_SECRET_KEY"))
	JwtSigningMethod = jwt.SigningMethodHS256.Name
)

func NewToken(userId string) (string, error) {
	claims := jwt.StandardClaims{
		Id:        userId,
		Issuer:    userId,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecretKey)
}
