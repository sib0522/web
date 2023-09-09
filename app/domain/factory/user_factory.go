package factory

import "GoEcho/app/util/generator"

type IUserFactory interface {
	NewUserData() (string, error)
}

type UserFactory struct {
}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func NewUserData() (string, error) {
	uuid := generator.NewUuidV4String()
	return uuid, nil
}
