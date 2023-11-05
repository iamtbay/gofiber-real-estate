package helpers

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iamtbay/real-estate-api/models"
)

type LoginJWT struct {
	UserID  string `json:"_id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	jwt.RegisteredClaims
}

func CreateJWT(userInfo *models.User) string {
	//Create Claims
	claims := LoginJWT{
		userInfo.ID,
		userInfo.Name,
		userInfo.Surname,
		userInfo.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    userInfo.Email,
			Subject:   "accessToken",
		},
	}
	//Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println(err)
	}
	return ss
}

func ParseJWT(tokenString string) (string, error) {
	var userID string
	token, err := jwt.ParseWithClaims(tokenString, &LoginJWT{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return "", errors.New("Fatal error. Please login again")
	}
	if claims, ok := token.Claims.(*LoginJWT); ok && token.Valid {
		userID = claims.UserID
	} else {
		return userID, err
	}
	return userID, nil

}
