package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
	"github.com/golang-jwt/jwt/v4"
)

var secret = os.Getenv("JWT_SECRET")

func CreateAccessToken(author *domain.Author, expired int) (accessToken string, err error) {

	exp := time.Now().Add(time.Hour * time.Duration(expired))

	// token := jwt.New(jwt.SigningMethodHS256)
	// // claims := token.Claims.(jwt.MapClaims) старая тема
	// // claims[""]

	claims := &domain.JwtCustomClaims{
		Username: author.Username,
		ID:       author.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err

	}

	return accToken, err
}

func CreateRefreshToken(author *domain.Author, expired int) (refreshToken string, err error) {
	claimRefresh := &domain.JwtCustomRefreshClaims{
		ID: author.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expired))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimRefresh)
	refToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return refToken, err

}

func IsAuthorized(requestedToken string) (bool, error) {
	_, err := jwt.Parse(requestedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func ExtractIDFromToken(requestedToken string) (string, error) {
	token, err := jwt.Parse(requestedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims["id"].(string), nil
}
