package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateJWT(username, password string) (string, error) {
	viper.SetConfigFile("../.env")
	value := viper.GetString("SECRET_KEY")

	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(value))
	if err != nil {
		return "", err
	}
	return signedToken, err
}
