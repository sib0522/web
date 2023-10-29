package model

import (
	"GoEcho/app/util/generator"
	"time"
)

// ユーザー基本データ
type UserStatus struct {
	id        uint32
	uuid      string
	level     uint32
	exp       uint64
	money     uint64
	createdAt time.Time
	updatedAt time.Time
}

func NewUserStatus(t time.Time) *UserStatus {
	return &UserStatus{
		uuid:      generator.NewUuidV4String(),
		level:     1,
		exp:       0,
		money:     0,
		createdAt: t,
		updatedAt: t,
	}
}

func NewUserStatusByRepo(uuid string, level uint32, exp, money uint64, createdAt, updatedAt time.Time) *UserStatus {
	return &UserStatus{
		uuid:      uuid,
		level:     level,
		exp:       exp,
		money:     money,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (r *UserStatus) Update(level uint32, exp, money uint64, t time.Time) {
	r.level = level
	r.exp = exp
	r.money = money
	r.updatedAt = t
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

func (r *UserStatus) Money() uint64 {
	return r.money
}

func (r *UserStatus) CreatedAt() time.Time {
	return r.createdAt
}

func (r *UserStatus) UpdatedAt() time.Time {
	return r.updatedAt
}
