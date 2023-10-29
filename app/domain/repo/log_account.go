package repo

import (
	"GoEcho/app/domain/model"
	"GoEcho/app/util/clock"
	"GoEcho/database"
	"fmt"
	"time"
)

type logAccountRepo struct {
	Id        uint32    `db:"id"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
}

func NewLogConnectRepo() *logAccountRepo {
	return &logAccountRepo{}
}

func (r *logAccountRepo) TableName() string {
	return "log_account"
}

func (r *logAccountRepo) CreateByModel(model *model.LogAccount) error {
	query := fmt.Sprintf("INSERT INTO %v (email, created_at) VALUES ('%v', '%v')",
		r.TableName(),
		model.Email(),
		model.CreatedAt().Format(clock.DateTimeFormat),
	)

	err := database.Instance().QueryRow(query).Err()
	if err != nil {
		return err
	}

	return nil
}
