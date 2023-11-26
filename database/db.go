package database

import (
	"GoEcho/web/lib"
	"database/sql"
	"fmt"

	"gopkg.in/yaml.v3"

	_ "github.com/go-sql-driver/mysql"
)

type DBInfo struct {
	// NOTE:小文字にしちゃうと外から参照できないので注意
	DBName string `yaml:"root"`
	DBPass string `yaml:"password"`
	DBPort string `yaml:"port"`
}

type DBConfig struct {
	// NOTE:小文字にしちゃうと外から参照できないので注意
	Info DBInfo `yaml:"database"`
}

type DB struct {
	instance *sql.DB
}

var db DB

func Instance() *sql.DB {
	return db.instance
}

// DBに接続
func ConnectDB() error {
	dbConfig := &DBConfig{}
	b, err := lib.NewAWSService().DonwloadConfig()
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, dbConfig)
	if err != nil {
		return err
	}

	conf := dbConfig.Info
	fmt.Println("db connection start...")
	dbSource := fmt.Sprintf("%v:%v@(%v)/db?parseTime=true", conf.DBName, conf.DBPass, conf.DBPort)

	sqlDB, err := sql.Open("mysql", dbSource)
	if err != nil {
		return err
	}

	db = DB{instance: sqlDB}

	fmt.Println("db connection success.")

	return nil
}

// 該当するデータがない場合のエラーか
func IsErrorNoRows(err error) bool {
	if err == nil {
		return false
	}
	return err.Error() == "sql: no rows in result set"
}
