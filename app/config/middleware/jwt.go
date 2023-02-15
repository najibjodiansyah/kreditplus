package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: viper.GetString("JWT_SIGNING_METHOD"),
		SigningKey:    []byte(viper.GetString("JWT_SECRET")),
	})
}

func CreateToken(nik string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["nik"] = nik
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 24 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(viper.GetString("JWT_SECRET")))
}

func GetNik(e echo.Context) (string, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		nik := claims["nik"].(string)
		if nik == "" {
			return nik, fmt.Errorf("empty nik")
		}
		return nik, nil
	}
	return "", fmt.Errorf("invalid user")
}

func ExtractToken(e echo.Context) (string, error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		nik := claims["nik"].(string)
		if nik == "" {
			return "", fmt.Errorf("empty nik")
		}
	}
	return "", fmt.Errorf("invalid user")
}
