package utils

import (
	"fmt"
	"strings"
	"time"
	"workout-tracker/m/v0.0.0/src/configs"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(configs.GetDotEnvVariable("SECRET_KEY"))

func CreateAccessToken(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id"			: id,
		"exp"			: time.Now().Add(time.Hour * 6).Unix(),
	})

	accessToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func VerifyAccessToken(c *gin.Context) error {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Unix(int64(exp), 0).Before(time.Now()) {
				return fmt.Errorf("token has expired")
			}
		} else {
			return fmt.Errorf("token expiration time is invalid")
		}
	} else {
		return fmt.Errorf("invalid token claims")
	}

	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID (c *gin.Context) (uint, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if id, ok := claims["id"].(float64); ok {
			return uint(id), nil
		}
	}

	return 0, fmt.Errorf("invalid token claims")
}