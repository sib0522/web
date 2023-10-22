package repo

import (
	"GoEcho/app/domain/model"
	"GoEcho/database"
	"fmt"
	"time"
)

type UserStatusRepo struct {
	Id        uint32    `db:"id"`
	Uuid      string    `db:"uuid"`
	Level     uint32    `db:"level"`
	Exp       uint64    `db:"exp"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewUserStatusRepo() *UserStatusRepo {
	return &UserStatusRepo{}
}

func (r *UserStatusRepo) TableName() string {
	return "user_status"
}

func (r *UserStatusRepo) CreateOrUpdateByModel(model *model.UserStatus) {
	db := database.ConnectDB()
	query := fmt.Sprintf("INSERT INTO %v (uuid, level, exp, created_at, updated_at) VALUES ('%v', '%v', '%v', '%v', '%v')",
		r.TableName(),
		model.Uuid(),
		model.Level(),
		model.Exp(),
		model.CreatedAt(),
		model.UpdatedAt(),
	)
	db.Query(query)
}

func (r *UserStatusRepo) ReadByUuid(uuid string) (*model.UserStatus, error) {
	db := database.ConnectDB()
	query := fmt.Sprintf("SELECT * FROM %v WHERE uuid = %v", r.TableName(), uuid)
	rows := db.Query(query)

	result := &model.UserStatus{}
	if err := rows.Scan(result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *UserStatusRepo) ReadTable() (*model.Table, error) {
	db := database.ConnectDB()
	query := fmt.Sprintf("SELECT * FROM %v", r.TableName())
	rows := db.Query(query)

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	type holder struct {
		Uuid      string
		Level     string
		Exp       string
		CreatedAt string
		UpdatedAt string
	}

	values := make([][]string, 0, len(columns))

	for rows.Next() {
		holder := &holder{}
		if err := rows.Scan(&holder.Uuid, &holder.Level, &holder.Exp, &holder.CreatedAt, &holder.UpdatedAt); err != nil {
			return nil, err
		}
		value := []string{holder.Uuid, holder.Level, holder.Exp, holder.CreatedAt, holder.UpdatedAt}
		values = append(values, value)
	}

	table := &model.Table{
		Name:    r.TableName(),
		Columns: columns,
		Rows:    values,
	}

	return table, nil
}
