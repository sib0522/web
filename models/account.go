package models

import (
	"GoEcho/database"
	"fmt"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Account struct {
	Id        uint
	Email     string
	Password  string
	Nickname  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Create 新しいアカウントを生成
func (account *Account) Create() {
	db := database.ConnectDB()

	/*
		// 新しいidを生成
		id, err := uuid.NewRandom()
		if err != nil {
			err.Error()
			return
		}
	*/

	query := fmt.Sprintf("INSERT INTO account VALUES (%v, '%v', '%v', '%v', '%v', '%v')", 2, account.Email, account.Password, account.Nickname, time.Now().String(), time.Now().String())
	db.Query(query)
}

// Read アカウントデータを読み込む
func (account *Account) Read() (string, string) {
	db := database.ConnectDB()

	query := fmt.Sprintf("SELECT ACCOUNT_NAME, NICKNAME, PASSWORD FROM account WHERE ACCOUNT_NAME='%v'", account.Email)
	rows := db.Query(query)

	if rows == nil {
		log.Error("アカウントが存在しません")
		return "", ""
	}

	var (
		email          string
		nickName       string
		hashedPassword string
	)

	for rows.Next() {
		err := rows.Scan(&email, &nickName, &hashedPassword)
		if err != nil {
			log.Error("このアカウントは無効です")
			return "", ""
		}

		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(account.Password))
		if err != nil {
			log.Error("パスワードが違います")
			return "", ""
		}

		// emailごとにIdは1個しか存在しないはずなので
		break
	}

	return email, nickName
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
func (account *Account) Delete() {

}
