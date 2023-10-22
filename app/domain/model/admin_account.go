package model

import (
	"time"
)

type AdminAccount struct {
	id        uint32
	email     string
	password  string
	nickname  string
	updatedAt time.Time
	createdAt time.Time
}

func NewAdminAccount(email, password, nickname string, t time.Time) *AdminAccount {
	return &AdminAccount{
		email:     email,
		password:  password,
		nickname:  nickname,
		updatedAt: t,
		createdAt: t,
	}
}

func NewAdminAccountByRepo(email, password, nickname string, updatedAt, createdAt time.Time) *AdminAccount {
	return &AdminAccount{
		email:     email,
		password:  password,
		nickname:  nickname,
		updatedAt: updatedAt,
		createdAt: createdAt,
	}
}

func (r *AdminAccount) Id() uint32 {
	return r.id
}

func (r *AdminAccount) Email() string {
	return r.email
}

func (r *AdminAccount) Password() string {
	return r.password
}

func (r *AdminAccount) Nickname() string {
	return r.nickname
}

func (r *AdminAccount) UpdatedAt() time.Time {
	return r.updatedAt
}

func (r *AdminAccount) CreatedAt() time.Time {
	return r.createdAt
}
