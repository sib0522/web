package application

import (
	"GoEcho/app/api/ApiUserLogin"
	"GoEcho/app/domain/model"
	"GoEcho/app/domain/repo"
	"GoEcho/app/util/clock"
)

type UserStatusUpdateService struct {
	userStatusRepo repo.UserStatusRepo
}

func NewUserStatusUpdateService() *UserStatusUpdateService {
	return &UserStatusUpdateService{}
}

func (service *UserStatusUpdateService) UserStatusUpdate() {
}

type UserLoginService struct {
	userStatusRepo repo.UserStatusRepo
}

func NewUserLoginService(userStatusRepo *repo.UserStatusRepo) *UserLoginService {
	return &UserLoginService{
		userStatusRepo: *userStatusRepo,
	}
}

func (r *UserLoginService) UserLoginService(req *ApiUserLogin.Request) (*ApiUserLogin.Response, error) {
	res := &ApiUserLogin.Response{}

	t := clock.Now().Time

	userStatus, err := r.userStatusRepo.ReadByUuid(req.Uuid)
	if err != nil {
		userStatus = model.NewUserStatus(t)

		// 新しく生成したユーザーデータをDBに保存
		r.userStatusRepo.CreateOrUpdateByModel(userStatus)
		return res, err
	}
	res.Uuid = userStatus.Uuid()
	res.Level = userStatus.Level()
	res.Exp = userStatus.Exp()

	return res, nil
}
