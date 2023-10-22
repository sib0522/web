package clock

import (
	"github.com/jinzhu/now"
	"time"
)

const (
	YearFormat      = "2006"
	MonthFormat     = "2006-01"
	DateFormat      = "2006-01-02"
	TimeFormat      = "15:04:05"
	DateTimeFormat  = "2006-01-02 15:04:05"
	MilliTimeFormat = "2006-01-02 15:04:05.000"
	MicroTimeFormat = "2006-01-02 15:04:05.000000"
	NanoTimeFormat  = "2006-01-02 15:04:05.000000000"
)

type Clock struct {
	*now.Now
}

func NowToString() string {
	// https://qiita.com/icbmuma/items/5617f3fc5bc0215aade2
	// golangで時間をフォーマット指定した文字列に変換する時には
	// %Yといったような制御文字で表すのではなくて
	// "2006/1/2 15:04:05"という決まった日付に対して、出力例を与えるような形になっている。
	// (ちなみに、この日以外の指定は受け付けないので注意。)
	return time.Now().Format("2006-01-02 15:04:05")
}

func Now() *Clock {
	return &Clock{now.New(time.Now())}
}
