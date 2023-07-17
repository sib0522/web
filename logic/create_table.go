package logic

import (
	"GoEcho/models"
)

type Table struct {
	IsLogin bool
	Name    string
	Columns []string
	RowsMap map[uint][]string
}

var tableMap = map[string]models.Model{
	"account": &models.Account{},
	"user":    &models.User{},
}

func CreateTable(tableName string) Table {
	m := tableMap[tableName]
	columnNames := m.ReadColumns()
	rowsMap := m.ReadAllToString()

	tb := &Table{
		IsLogin: true,
		Name:    tableName,
		Columns: columnNames,
		RowsMap: rowsMap,
	}

	return *tb
}
