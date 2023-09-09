package models

import (
	"GoEcho/database"
	"fmt"
	"github.com/labstack/gommon/log"
	"time"
)

type User struct {
	UUID      string
	Level     uint
	Exp       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type userHolder struct {
	uuid, level, exp, createdAt, updatedAt []byte
}

// Create 管理画面上では生成できない
func (user *User) Create() {
	db := database.ConnectDB()
	query := fmt.Sprintf("INSERT INTO user (UUID, USER_LEVEL, USER_EXP, CREATED_AT, UPDATED_AT) VALUES ('%v', '%v', '%v', '%v', '%v')",
		user.UUID,
		user.Level,
		user.Exp,

		// https://qiita.com/icbmuma/items/5617f3fc5bc0215aade2
		// golangで時間をフォーマット指定した文字列に変換する時には
		// %Yといったような制御文字で表すのではなくて
		// "2006/1/2 15:04:05"という決まった日付に対して、出力例を与えるような形になっている。
		// (ちなみに、この日以外の指定は受け付けないので注意。)
		time.Now().Format("2006-01-02 15:04:05"),
		time.Now().Format("2006-01-02 15:04:05"),
	)
	db.Query(query)
}

// ReadAll 全ユーザーデータを取得
func (user *User) ReadAll() []*User {
	db := database.ConnectDB()
	query := fmt.Sprint("SELECT * FROM user")
	rows := db.Query(query)

	var allUsers []*User

	for rows.Next() {
		user := new(User)
		err := rows.Scan(user.UUID, user.Level, user.Exp, user.CreatedAt, user.UpdatedAt)
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

// ReadAllToString 全データを取得しstring型のスライスに変換する(管理画面用）
func (user *User) ReadAllToString() map[uint][]string {
	db := database.ConnectDB()
	query := fmt.Sprint("SELECT * FROM user")
	rows := db.Query(query)
	result := make(map[uint][]string)
	var idx uint = 0

	for rows.Next() {
		holder := new(userHolder)

		err := rows.Scan(&holder.uuid, &holder.level, &holder.exp, &holder.createdAt, &holder.updatedAt)
		if err != nil {
			log.Error("データの取得に失敗しました")
			return nil
		}

		result[idx] = []string{
			string(holder.uuid),
			string(holder.level),
			string(holder.exp),
			string(holder.createdAt),
			string(holder.updatedAt),
		}
		idx++
	}

	return result
}
