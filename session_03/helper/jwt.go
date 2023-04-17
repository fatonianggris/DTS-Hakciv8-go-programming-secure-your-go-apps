package helper

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	secret = "@pr1j@l9h1y@sset1@w@n"
)

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	claimsToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := claimsToken.SignedString([]byte(secret))

	return signedToken
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	err := errors.New("login to proceed")

	headerAuthorization := ctx.Request.Header.Get("Authorization")
	if !strings.HasPrefix(headerAuthorization, "Bearer ") {
		return nil, err
	}

	token := strings.Split(headerAuthorization, " ")[1]
	parsedToken, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(secret), nil
	})

	if _, ok := parsedToken.Claims.(jwt.MapClaims); !ok && !parsedToken.Valid {
		return nil, err
	}

	return parsedToken.Claims.(jwt.MapClaims), nil
}

func GetUserIdFromClaims(ctx *gin.Context) uint {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	return userId
}
