package application

import (
	"GoEcho/app/domain/factory"
	"GoEcho/app/domain/repository"
	"fmt"
)

type UserCreateService struct {
	userRepo    repository.UserDataRepo
	userFactory factory.IUserFactory
}

func NewUserCreateService() *UserCreateService {
	return &UserCreateService{}
}

func (service *UserCreateService) UserCreate() (string, error) {
	uuid, err := service.userFactory.NewUserData()
	if err != nil {

		// 何かエラーを返す
		return "", nil
	}
	// user repositoryを通じてDBにuserのuuidを保存する
	service.userRepo.Update(uuid)
	fmt.Println("user create")

	// jsonにシリアライズしてクライアントへ返す

	return uuid, nil
}

type UserUpdateService struct {
}

func NewUserUpdateService() *UserUpdateService {
	return &UserUpdateService{}
}

func (service *UserUpdateService) UserUpdate() {
	fmt.Println("user update")
}

type UserLoadService struct {
}

func NewUserLoadService() *UserLoadService {
	return &UserLoadService{}
}

func (service *UserLoadService) UserLoad() {
	fmt.Println("user load")
}
