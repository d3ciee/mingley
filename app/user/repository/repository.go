package user

import (
	"errors"
	userModels "mingley/app/user/models"
	userValidators "mingley/app/user/validators"

	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) Create(
	username string,
	email string,
	password string) (*userModels.User, error) {

	var err error

	_, err = userValidators.ValidateUsername(username)
	_, err = userValidators.ValidateEmail(email)
	_, err = userValidators.ValidatePassword(password)

	passwordHash, bcryptErr := bcrypt.GenerateFromPassword([]byte(password), 10)

	if bcryptErr != nil {
		err = errors.New("internal error")
		log.Error(bcryptErr)
	}

	if err != nil {
		return nil, err
	}

	u := userModels.User{
		PasswordHash: passwordHash,
		Username:     username,
		Email:        email,
	}

	r.db.Create(&u)
	return &u, nil

}
