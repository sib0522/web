package models

import (
	"GoEcho/database"
	"fmt"
	"github.com/labstack/gommon/log"
	"time"
)

type User struct {
	UUID      string
	Id        uint
	Name      string
	Level     uint
	Exp       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ReadAll 全ユーザーデータを取得
func (user *User) ReadAll() []*User {
	db := database.ConnectDB()
	query := fmt.Sprint("SELECT * FROM user")
	rows := db.Query(query)

	var allUsers []*User

	for rows.Next() {
		user := new(User)
		err := rows.Scan(user.UUID, user.Id, user.Name, user.Level, user.Exp, user.CreatedAt, user.UpdatedAt)
		if err != nil {
			log.Error("データの取得に失敗しました")
			return nil
		}
		allUsers = append(allUsers, user)
	}

	return allUsers
}

// ReadColumns カラム名を取得する
func (user *User) ReadColumns() []string {
	db := database.ConnectDB()
	query := fmt.Sprint("SELECT * FROM user")
	rows := db.Query(query)
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal("カラムの取得に失敗しました")
		return nil
	}

	return columns
}
