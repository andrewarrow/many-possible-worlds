package server

import (
	"fmt"
	"math"
	"time"
)

func DeltaAgo(t int64) string {
	delta := float64(time.Now().Unix() - t)
	d := delta / 86400.0
	days := math.Floor(d)
	hours := 24 * (d - days)
	fullHours := math.Floor(hours)
	mins := (hours - fullHours) * 60
	return fmt.Sprintf("%d day(s), %d hour(s), %d min(s)", int(days), int(fullHours), int(mins))
}
