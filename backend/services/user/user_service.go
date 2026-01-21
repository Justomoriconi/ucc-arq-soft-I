package user

import (
	"errors"
	"strings"
	Client "ucc-arq-soft-I/clients/user"
	Domain "ucc-arq-soft-I/domain/user"
	Model "ucc-arq-soft-I/model/user"
	"ucc-arq-soft-I/utils"
)

func CreateUser(req Domain.User) (Domain.User, error) {
	if strings.TrimSpace(req.Name) == "" {
		return Domain.User{}, errors.New("The name is required")

	}

	if strings.TrimSpace(req.Password) == "" {
		return Domain.User{}, errors.New("The password is required")
	}

	if strings.TrimSpace(req.Email) == "" {
		return Domain.User{}, errors.New("The email is required")
	}

	if strings.TrimSpace(req.LastName) == "" {
		return Domain.User{}, errors.New("The last name is required")
	}

	role := strings.TrimSpace(req.Role)
	if role == "" {
		role = "USER"
	}

	exist, err := Client.ExistByEmail(req.Email)

	if err != nil {
		return Domain.User{}, err
	}

	if exist {
		return Domain.User{}, errors.New("User already exists")
	}

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return Domain.User{}, err
	}

	user := Model.User{
		Name:     req.Name,
		LastName: req.LastName,
		Role:     role,
		Email:    req.Email,
		Password: hashed,
	}

	if err := Client.Create(&user); err != nil {
		return Domain.User{}, err
	}

	return Domain.User{
		IDUser: user.IDUser,
	}, err
}
