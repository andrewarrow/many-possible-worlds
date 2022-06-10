package redis

import (
	"fmt"
	"time"
)

func BucketForHour(t time.Time) string {
	format := t.Format("2006010215")
	return fmt.Sprintf("%s", format)
}
