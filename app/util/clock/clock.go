package clock

import (
	"time"

	"github.com/jinzhu/now"
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

func Now() *Clock {
	return &Clock{now.New(time.Now())}
}

func NowTime() time.Time {
	v := time.Now()
	return v
}
