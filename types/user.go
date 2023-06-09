package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" `
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
}

type CreateUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUserFromDto(dto CreateUserDto) (*User, error) {

	encryptedPassword, encryptErr := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)
	if encryptErr != nil {
		return nil, encryptErr
	}
	return &User{Email: dto.Email, Password: string(encryptedPassword)}, nil
}

func IsEmailValid(email string) bool {
	emailRegex, _ := regexp.Compile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$")
	return emailRegex.MatchString(email)
}
