package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey       = []byte(os.Getenv("JWT_SECRET_KEY"))
	refreshKey      = []byte(os.Getenv("JWT_REFRESH_KEY"))
	minutesCount, _ = strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))
	hourCount, _    = strconv.Atoi(os.Getenv("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT"))
)

func CreateToken(username string, role string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"role":     role,
			"exp":      time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix(), //test for 15s
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func CreateRefreshToken(username string, role string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"role":     role,
			"exp":      time.Now().Add(time.Hour * time.Duration(hourCount)).Unix(),
		})
	tokenString, err := token.SignedString(refreshKey)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("Invalid token")
	}
	return nil
}

func VerifyRefreshToken(refreshTokenString string) (string, string, error) {
	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return refreshKey, nil
	})
	if err != nil {
		return "", "", err
	}
	if !token.Valid {
		return "", "", fmt.Errorf("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", fmt.Errorf("invalid token claims")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "", "", fmt.Errorf("missing username in token claims")
	}

	role, ok := claims["role"].(string)
	if !ok {
		role = "user"
	}

	return username, role, nil
}
