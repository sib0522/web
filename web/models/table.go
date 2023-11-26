package models

import (
	"GoEcho/app/domain/model"
	"GoEcho/app/domain/repo"
)

// テーブルを作る
func CreateTable(tableName string) (*model.Table, error) {
	switch tableName {
	case "account":
		return readTable(repo.NewAdminAccountRepo())
	case "user":
		return readTable(repo.NewUserStatusRepo())
	}
	return nil, nil
}

func readTable(repo model.IRepository) (*model.Table, error) {
	table, err := repo.ReadTable()
	if err != nil {
		return nil, err
	}
	return table, nil
}
