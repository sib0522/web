package model

import (
	"GoEcho/app/util/generator"
	"time"
)

type UserStatus struct {
	id        uint32
	uuid      string
	level     uint32
	exp       uint64
	createdAt time.Time
	updatedAt time.Time
}

func NewUserStatus(t time.Time) *UserStatus {
	return &UserStatus{
		uuid:      generator.NewUuidV4String(),
		level:     1,
		exp:       0,
		createdAt: t,
		updatedAt: t,
	}
}

func NewUserStatusByRepo(uuid string, level uint32, exp uint64, createdAt, updatedAt time.Time) *UserStatus {
	return &UserStatus{
		uuid:      uuid,
		level:     level,
		exp:       exp,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (r *UserStatus) Id() uint32 {
	return r.id
}

func (r *UserStatus) Uuid() string {
	return r.uuid
}

func (r *UserStatus) Level() uint32 {
	return r.level
}

func (r *UserStatus) Exp() uint64 {
	return r.exp
}

func (r *UserStatus) CreatedAt() time.Time {
	return r.createdAt
}

func (r *UserStatus) UpdatedAt() time.Time {
	return r.updatedAt
}
