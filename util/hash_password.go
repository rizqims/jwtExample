package util
import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(pass string) (string, error){
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(passHash), nil
}

func ComparePasswordHash(passHash string, passFromDB string) error{
	return bcrypt.CompareHashAndPassword([]byte(passHash), []byte(passFromDB))
}