package redis

import (
	"fmt"
	"time"
)

func BucketForHour(t time.Time) string {
	format := t.Format("2006010215")
	return fmt.Sprintf("%s", format)
}

func BucketForDay(t time.Time) string {
	format := t.Format("20060102")
	return fmt.Sprintf("%s", format)
}

func BucketForMonth(t time.Time) string {
	format := t.Format("200601")
	return fmt.Sprintf("%s", format)
}
