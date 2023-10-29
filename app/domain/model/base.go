package model

type Table struct {
	Name    string
	Columns []string
	Rows    [][]string
}

type IRepository interface {
	ReadTable() (*Table, error)
}
