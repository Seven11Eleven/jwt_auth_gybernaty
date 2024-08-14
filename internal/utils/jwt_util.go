package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/Seven11Eleven/jwt_auth_gybernaty/domain"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

var secret = os.Getenv("JWT_SECRET")

type JWTUtils interface {
	CreateAccessToken(author *domain.Author, expired int) (string, error)
	CreateRefreshToken(author *domain.Author, expired int) (string, error)
	IsAuthorized(requestedToken string) (bool, error)
	ExtractIDFromToken(requestedToken string) (uuid.UUID, error)
}

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

func ExtractIDFromToken(requestedToken string) (uuid.UUID, error) {
	token, err := jwt.Parse(requestedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return uuid.Nil, fmt.Errorf("invalid token")
	}
	id := claims["id"].(string)
	uuid := uuid.MustParse(id)
	return uuid, nil
}

type jwtUtilsImpl struct{

}

func NewJWTUtils() JWTUtils {
	return &jwtUtilsImpl{}
}

func (u *jwtUtilsImpl) CreateAccessToken(author *domain.Author, expired int) (string, error) {
	return CreateAccessToken(author, expired)
}

func (u *jwtUtilsImpl) CreateRefreshToken(author *domain.Author, expired int) (string, error) {
	return CreateRefreshToken(author, expired)
}

func (u *jwtUtilsImpl) IsAuthorized(requestedToken string) (bool, error) {
	return IsAuthorized(requestedToken)
}

func (u *jwtUtilsImpl) ExtractIDFromToken(requestedToken string) (uuid.UUID, error) {
	return ExtractIDFromToken(requestedToken)
}
