package model

type Model interface {
	Create()
	ReadColumns() []string
	ReadAllToString() map[uint][]string
}

type Table struct {
	Name    string
	Columns []string
	Rows    [][]string
}

type IRepository interface {
	ReadTable() (*Table, error)
}
