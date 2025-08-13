package utils

// LoginService to provide user login with JWT token support
import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKeyToken        = []byte(os.Getenv("JWT_SECRET"))
	secretKeyRefreshToken = []byte(os.Getenv("JWT_REFRESH_SECRET"))
)

func CreateToken(username string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"role":     role,
			"exp":      time.Now().Add(time.Second * 15).Unix(),
		})
	tokenString, err := token.SignedString(secretKeyToken)
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
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(secretKeyRefreshToken)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKeyToken, nil
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
		return secretKeyRefreshToken, nil
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
