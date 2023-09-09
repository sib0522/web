package models

import (
	"GoEcho/database"
	"fmt"
	"github.com/labstack/gommon/log"
	"time"
)

type Gallery struct {
	Id        string    `db:"id"`
	Path      string    `db:"path"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (gallery *Gallery) Create() {
	db := database.ConnectDB()
	query := fmt.Sprintf("INSERT INTO gallery (path, created_at, updated_at) VALUES ('%v', '%v', '%v')",
		gallery.Path,
		gallery.CreatedAt,
		gallery.UpdatedAt,

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

// ReadColumns カラム名を取得する
func (gallery *Gallery) ReadColumns() []string {
	db := database.ConnectDB()
	query := fmt.Sprint("SELECT * FROM gallery")
	rows := db.Query(query)
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal("カラムの取得に失敗しました")
		return nil
	}

	return columns
}

type galleryHolder struct {
	id, path, createdAt, updatedAt []byte
}

// ReadAllToString 全データを取得しstring型のスライスに変換する
func (gallery *Gallery) ReadAllToString() map[uint][]string {
	db := database.ConnectDB()
	query := fmt.Sprint("SELECT * FROM gallery")
	rows := db.Query(query)
	result := make(map[uint][]string)
	var idx uint = 0

	for rows.Next() {
		holder := new(galleryHolder)

		err := rows.Scan(&holder.id, &holder.path, &holder.createdAt, &holder.updatedAt)
		if err != nil {
			log.Error("データの取得に失敗しました")
			return nil
		}

		result[idx] = []string{
			string(holder.id),
			string(holder.path),
			string(holder.createdAt),
			string(holder.updatedAt),
		}
		idx++
	}

	return result
}
