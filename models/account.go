package models

import (
	"GoEcho/database"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/labstack/gommon/log"
)

type Account struct {
	Id        uint      `db:"ID"`
	Email     string    `db:"ACCOUNT_NAME"`
	Password  string    `db:"PASSWORD"`
	Nickname  string    `db:"NICKNAME"`
	CreatedAt time.Time `db:"CREATED_AT"`
	UpdatedAt time.Time `db:"UPDATED_AT"`
}

// Create 新しいアカウントを生成
func (account *Account) Create() {
	db := database.ConnectDB()
	query := fmt.Sprintf("INSERT INTO account (ACCOUNT_NAME, PASSWORD, NICKNAME, CREATED_AT, UPDATED_AT) VALUES ('%v', '%v', '%v', '%v', '%v')",
		account.Email,
		account.Password,
		account.Nickname,

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

// Read アカウントデータを読み込む
func (account *Account) Read() bool {
	db := database.ConnectDB()

	query := fmt.Sprintf("SELECT * FROM account WHERE ACCOUNT_NAME='%v'", account.Email)
	rows := db.Query(query)

	if rows == nil {
		log.Error("アカウント名またはパスワードを確認してください")
		return false
	}
	inputPassword := account.Password

	for rows.Next() {
		err := rows.Scan(&account.Id, &account.Email, &account.Password, &account.Nickname, &account.CreatedAt, &account.UpdatedAt)
		if err != nil {
			log.Error("このアカウントは無効です")
			return false
		}

		if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(inputPassword)); err != nil {
			log.Error("パスワードが違います")
			return false
		}

		// emailごとにIdは1個しか存在しないはずなので
		// わざわざbreakする必要はない？
		return true
	}
	return false
}

// ReadAll 全アカウントデータを取得
func (account *Account) ReadAll() []*Account {
	db := database.ConnectDB()
	query := fmt.Sprint("SELECT * FROM account")
	rows := db.Query(query)
	var allAccount []*Account

	for rows.Next() {
		acc := new(Account)
		err := rows.Scan(acc.Id, acc.Email, acc.Password, acc.Nickname, acc.CreatedAt, acc.UpdatedAt)
		if err != nil {
			log.Error("データの取得に失敗しました")
			return nil
		}
		allAccount = append(allAccount, acc)
	}

	return allAccount
}

// ReadAllToString 全アカウントデータを取得しstring型のスライスに変換する
func (account *Account) ReadAllToString() map[uint][]string {
	db := database.ConnectDB()
	query := fmt.Sprint("SELECT * FROM account")
	rows := db.Query(query)
	result := make(map[uint][]string)
	var idx uint = 0

	type accountHolder struct {
		Id        []byte
		Email     []byte
		Password  []byte
		Nickname  []byte
		CreatedAt []byte
		UpdatedAt []byte
	}

	for rows.Next() {
		holder := new(accountHolder)
		err := rows.Scan(&holder.Id, &holder.Email, &holder.Password, &holder.Nickname, &holder.CreatedAt, &holder.UpdatedAt)
		if err != nil {
			log.Error("データの取得に失敗しました")
			return nil
		}

		result[idx] = []string{
			string(holder.Id),
			string(holder.Email),
			string(holder.Password),
			string(holder.Nickname),
			string(holder.CreatedAt),
			string(holder.UpdatedAt),
		}
		idx++
	}

	return result
}

// ReadColumns カラム名を取得する
func (account *Account) ReadColumns() []string {
	db := database.ConnectDB()
	query := fmt.Sprint("SELECT * FROM account")
	rows := db.Query(query)
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal("カラムの取得に失敗しました")
		return nil
	}

	return columns
}

// Update アカウント情報を変更
func (account *Account) Update() {

}

// Delete アカウントを削除
func (account *Account) Delete() bool {
	db := database.ConnectDB()
	query := fmt.Sprintf("DELETE FROM account WHERE id=%v", account.Id)
	db.Query(query)
	return true
}
