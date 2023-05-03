package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateJWT(username, password, role string) (string, error) {
	viper.SetConfigFile("../.env")
	value := viper.GetString("SECRET_KEY")
	if role != "abrigo" && role != "tutor" {
		return "", fmt.Errorf("role must be \"abrigo\" or \"tutor\"")
	}

	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	claims["role"] = role // abrigo ou tutor

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(value))
	if err != nil {
		return "", err
	}
	return signedToken, err
}

func ValidateAuth(tkstr string) (string, error) {
	token, err := jwt.ParseWithClaims(tkstr, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method %v", t.Header["alg"])
		}
		viper.SetConfigFile("../.env")
		return []byte(viper.GetString("SECRET_KEY")), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*jwt.MapClaims)
	fmt.Printf("%T\n", claims)
	if !ok {
		return "", fmt.Errorf("invalid claims")
	}
	role, ok := (*claims)["role"].(string)
	if !ok {
		return "", fmt.Errorf("invalid claim")
	}
	if token.Valid {
		return role, nil
	}
	return "", fmt.Errorf("invalid token")

}
