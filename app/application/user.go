package application

import (
	"GoEcho/app/api/ApiUserLogin"
	"GoEcho/app/api/ApiUserObatain"
	"GoEcho/app/constants"
	"GoEcho/app/domain/model"
	"GoEcho/app/domain/repo"
	"GoEcho/app/util/clock"
)

type UserObatainService struct {
	userStatusRepo repo.UserStatusRepo
}

func NewUserObatainService(userStatusRepo *repo.UserStatusRepo) *UserObatainService {
	return &UserObatainService{}
}

func (r *UserObatainService) UserObatainService(req *ApiUserObatain.Request) (*ApiUserObatain.Response, error) {
	res := &ApiUserObatain.Response{}

	t := clock.Now().Time

	userStatus, err := r.userStatusRepo.ReadByUuid(req.Uuid)
	if err != nil {
		userStatus = model.NewUserStatus(t)

		// 新しく生成したユーザーデータをDBに保存
		r.userStatusRepo.CreateOrUpdateByModel(userStatus)
		return res, err
	}

	if req.Reason == constants.OBATAIN_REASON_KILL_ENEMY {
		// Value is EnemyId
		// enemyId := req.Value

		// 正しくはマスタテーブルからenemyIdの報酬データを獲得するべき
		// 仮データ
		res.GetExp = 10
		res.GetMoney = 50
	}

	e := userStatus.Entity()
	e.Exp += res.GetExp
	e.Money += res.GetMoney

	e.Apply(userStatus)

	if err := r.userStatusRepo.CreateOrUpdateByModel(userStatus); err != nil {
		return nil, err
	}

	return res, nil
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
