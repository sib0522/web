package repo

import (
	"GoEcho/app/domain/model"
	"GoEcho/app/util/clock"
	"GoEcho/database"
	"fmt"
	"time"
)

type UserStatusRepo struct {
	Id        uint32    `db:"id"`
	Uuid      string    `db:"uuid"`
	Level     uint32    `db:"level"`
	Exp       uint64    `db:"exp"`
	Money     uint64    `db:"money"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewUserStatusRepo() *UserStatusRepo {
	return &UserStatusRepo{}
}

func (r *UserStatusRepo) TableName() string {
	return "user_status"
}

func (r *UserStatusRepo) CreateOrUpdateByModel(model *model.UserStatus) error {
	query := fmt.Sprintf("INSERT INTO %v (uuid, level, exp, money, created_at, updated_at) VALUES ('%v', '%v', '%v', '%v', '%v', '%v')",
		r.TableName(),
		model.Uuid(),
		model.Level(),
		model.Exp(),
		model.Money(),
		model.CreatedAt().Format(clock.DateTimeFormat),
		model.UpdatedAt().Format(clock.DateTimeFormat),
	)
	// アップデート
	query += "ON DUPLICATE KEY UPDATE level = VALUES(level), exp = VALUES(exp), money = VALUES(money), updated_at = VALUES(updated_at);"

	_, err := database.Instance().Query(query)
	if err != nil {
		return err
	}

	return nil
}

// データがなければErrNoRowsエラーを返す
func (r *UserStatusRepo) ReadByUuid(uuid string) (*model.UserStatus, error) {
	query := fmt.Sprintf("SELECT * FROM %v WHERE uuid = %v", r.TableName(), uuid)
	err := database.Instance().QueryRow(query).Scan(&r.Id, &r.Uuid, &r.Level, &r.Exp, &r.Money, &r.CreatedAt, &r.UpdatedAt)
	if err != nil {
		return nil, err
	}

	result := model.NewUserStatusByRepo(r.Uuid, r.Level, r.Exp, r.Money, r.CreatedAt, r.UpdatedAt)
	return result, nil
}

// データがなければErrNoRowsエラーを返す
func (r *UserStatusRepo) ReadById(id uint32) (*model.UserStatus, error) {
	query := fmt.Sprintf("SELECT * FROM %v WHERE id = %v", r.TableName(), id)
	err := database.Instance().QueryRow(query).Scan(&r.Id, &r.Uuid, &r.Level, &r.Exp, &r.Money, &r.CreatedAt, &r.UpdatedAt)
	if err != nil {
		return nil, err
	}

	result := model.NewUserStatusByRepo(r.Uuid, r.Level, r.Exp, r.Money, r.CreatedAt, r.UpdatedAt)
	return result, nil
}

// テーブルを取得
func (r *UserStatusRepo) ReadTable() (*model.Table, error) {
	query := fmt.Sprintf("SELECT * FROM %v", r.TableName())
	rows, err := database.Instance().Query(query)
	if err != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	type holder struct {
		Id        string
		Uuid      string
		Level     string
		Exp       string
		Money     string
		CreatedAt string
		UpdatedAt string
	}

	values := make([][]string, 0, len(columns))

	for rows.Next() {
		holder := &holder{}
		if err := rows.Scan(&holder.Id, &holder.Uuid, &holder.Level, &holder.Exp, &holder.Money, &holder.CreatedAt, &holder.UpdatedAt); err != nil {
			return nil, err
		}
		value := []string{holder.Id, holder.Uuid, holder.Level, holder.Exp, holder.Money, holder.CreatedAt, holder.UpdatedAt}
		values = append(values, value)
	}

	table := &model.Table{
		Name:    r.TableName(),
		Columns: columns,
		Rows:    values,
	}

	return table, nil
}
