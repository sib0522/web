package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
)

type DB struct {
	instance *sql.DB
}

type DBInfo struct {
	DBName string `yaml:"root"`
	DBPass string `yaml:"password"`
	DBPort string `yaml:"port"`
}

type DBConfig struct {
	Info DBInfo `yaml:"database"`
}

func loadDBInfo() *DBInfo {
	dbConfig := DBConfig{}
	b, _ := os.ReadFile("config.yaml")
	err := yaml.Unmarshal(b, &dbConfig)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &dbConfig.Info
}

func ConnectDB() *DB {
	conf := loadDBInfo()
	fmt.Println("db connection start...")
	dbSource := fmt.Sprintf("%v:%v@(%v)/db?parseTime=true", conf.DBName, conf.DBPass, conf.DBPort)

	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	dbs := DB{instance: db}

	return &dbs
}

func (db *DB) Query(query string) *sql.Rows {
	res, err := db.instance.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return res
}
