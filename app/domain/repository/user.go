package repository

import "GoEcho/models"

type UserDataRepo struct {
}

func NewUserDataRepo() *UserDataRepo {
	return &UserDataRepo{}
}

func (r *UserDataRepo) Update(uuid string) {
	u := &models.User{
		UUID:  uuid,
		Level: 1,
		Exp:   0,
	}
	u.Create()
}
