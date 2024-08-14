package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io"

	"golang.org/x/crypto/bcrypt"
)

const saltSize = 16

func GenerateSalt() (string, error) {
	salt := make([]byte, saltSize)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

func HashPassword(password, salt, localParam string) (string, error) {
	saltedPassword := localParam + password + salt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CompareHashAndPassword(hashedPassword, password, salt, localParam string) error {
	saltedPassword := localParam + password + salt
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(saltedPassword))
}

type PasswordHasher interface {
	GenerateSalt() (string, error)
	HashPassword(password, salt, localParam string) (string, error)
	CompareHashAndPassword(hashedPassword, password, salt, localParam string) error
}

type hashUtilsImpl struct {
}

// CompareHashAndPassword implements PasswordHasher.
func (h *hashUtilsImpl) CompareHashAndPassword(hashedPassword string, password string, salt string, localParam string) error {
	return CompareHashAndPassword(hashedPassword, password, salt, localParam)
}

// GenerateSalt implements PasswordHasher.
func (h *hashUtilsImpl) GenerateSalt() (string, error) {
	return GenerateSalt()
}

// HashPassword implements PasswordHasher.
func (h *hashUtilsImpl) HashPassword(password string, salt string, localParam string) (string, error) {
	return HashPassword(password, salt, localParam)
}

func NewHashUtils() PasswordHasher {
	return &hashUtilsImpl{}
}
