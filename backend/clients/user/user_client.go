package user

import (
	"ucc-arq-soft-I/database"
	Model "ucc-arq-soft-I/model/user"
)

func Create(user *Model.User) error {
	return database.DB.Create(user).Error
}

func ExistByEmail(email string) (bool, error) {
	var count int64

	err := database.DB.Model(&Model.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
