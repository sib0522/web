package repo

import (
	"GoEcho/app/domain/model"
	"GoEcho/app/util/clock"
	"GoEcho/database"
	"fmt"
	"time"
)

type adminAccountRepo struct {
	Id        uint32    `db:"id"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Nickname  string    `db:"nickname"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

func NewAdminAccountRepo() *adminAccountRepo {
	return &adminAccountRepo{}
}

func (r *adminAccountRepo) TableName() string {
	return "admin_account"
}

func (r *adminAccountRepo) CreateByModel(model *model.AdminAccount) {
	db := database.ConnectDB()
	query := fmt.Sprintf("INSERT INTO %v (email, password, nickname, updated_at, created_at) VALUES ('%v','%v','%v', '%v', '%v')",
		r.TableName(),
		model.Email(),
		model.Password(),
		model.Nickname(),
		model.UpdatedAt().Format(clock.DateTimeFormat),
		model.CreatedAt().Format(clock.DateTimeFormat),
	)
	db.Query(query)
}

func (r *adminAccountRepo) ReadByEmail(email string) (*model.AdminAccount, error) {
	db := database.ConnectDB()
	query := fmt.Sprintf("SELECT * FROM %v WHERE email = '%v'", r.TableName(), email)
	rows := db.Query(query)

	for rows.Next() {
		if err := rows.Scan(&r.Id, &r.Email, &r.Password, &r.Nickname, &r.UpdatedAt, &r.CreatedAt); err != nil {
			return nil, err
		}
	}

	result := model.NewAdminAccountByRepo(r.Email, r.Password, r.Nickname, r.UpdatedAt, r.CreatedAt)
	return result, nil
}

func (r *adminAccountRepo) ReadById(id uint32) (*model.AdminAccount, error) {
	db := database.ConnectDB()
	query := fmt.Sprintf("SELECT * FROM %v WHERE id = '%v'", r.TableName(), id)
	rows := db.Query(query)

	for rows.Next() {
		if err := rows.Scan(&r.Id, &r.Email, &r.Password, &r.Nickname, &r.UpdatedAt, &r.CreatedAt); err != nil {
			return nil, err
		}
	}

	result := model.NewAdminAccountByRepo(r.Email, r.Password, r.Nickname, r.UpdatedAt, r.CreatedAt)
	return result, nil
}

func (r *adminAccountRepo) ReadTable() (*model.Table, error) {
	db := database.ConnectDB()
	query := fmt.Sprintf("SELECT * FROM %v", r.TableName())
	rows := db.Query(query)

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	type holder struct {
		Id        string
		Email     string
		Password  string
		Nickname  string
		UpdatedAt string
		CreatedAt string
	}

	values := make([][]string, 0, len(columns))

	for rows.Next() {
		model := &holder{}
		if err := rows.Scan(&model.Id, &model.Email, &model.Password, &model.Nickname, &model.UpdatedAt, &model.CreatedAt); err != nil {
			return nil, err
		}
		value := []string{model.Id, model.Email, model.Password, model.Nickname, model.UpdatedAt, model.CreatedAt}
		values = append(values, value)
	}

	table := &model.Table{
		Name:    r.TableName(),
		Columns: columns,
		Rows:    values,
	}

	return table, nil
}
