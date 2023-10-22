package model

import (
	"time"
)

type LogAccount struct {
	id        uint32
	email     string
	createdAt time.Time
}

func NewLogAccount(email string, t time.Time) *LogAccount {
	return &LogAccount{
		email:     email,
		createdAt: t,
	}
}

func (r *LogAccount) Id() uint32 {
	return r.id
}

func (r *LogAccount) Email() string {
	return r.email
}

func (r *LogAccount) CreatedAt() time.Time {
	return r.createdAt
}
