package meetup

import (
	"fmt"
	"math"
	"time"
)

// Define the WeekSchedule type here.
type WeekSchedule int

const (
	First  = iota
	Second = iota
	Third  = iota
	Fourth = iota
	Last   = iota
	Teenth = iota
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	if wSched == First {
		firstDayOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
		fmt.Println(firstDayOfMonth, firstDayOfMonth.Weekday(), wDay)
		wDayDiff := math.Abs(float64(wDay - firstDayOfMonth.Weekday()))
		fmt.Println(wDayDiff)
		result := firstDayOfMonth.Add(time.Hour * 24 * time.Duration(wDayDiff))
		fmt.Println(result)
		return result.Day()
	}
	return 0
	//first || jump || last || teenth
	//get the first matching weekday
	//jump to the right
}
