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

type userHolder struct {
	uuid, id, name, level, exp, createdAt, updatedAt []byte
}

// Create 管理画面上では生成できない
func (user *User) Create() {}

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

// ReadAllToString 全データを取得しstring型のスライスに変換する
func (user *User) ReadAllToString() map[uint][]string {
	db := database.ConnectDB()
	query := fmt.Sprint("SELECT * FROM user")
	rows := db.Query(query)
	result := make(map[uint][]string)
	var idx uint = 0

	for rows.Next() {
		holder := new(userHolder)

		err := rows.Scan(&holder.uuid, &holder.id, &holder.name, &holder.level, &holder.exp, &holder.createdAt, &holder.updatedAt)
		if err != nil {
			log.Error("データの取得に失敗しました")
			return nil
		}

		result[idx] = []string{
			string(holder.uuid),
			string(holder.id),
			string(holder.name),
			string(holder.level),
			string(holder.exp),
			string(holder.createdAt),
			string(holder.updatedAt),
		}
		idx++
	}

	return result
}
