package models

type Model interface {
	Create()
	ReadColumns() []string
	ReadAllToString() map[uint][]string
}
