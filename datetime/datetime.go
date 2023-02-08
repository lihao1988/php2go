package datetime

import (
	"fmt"
	"time"
)

// Time get timestamp
// php time
func Time() int64 {
	return time.Now().Unix()
}

// StrToTime parse time_string to timestamp
// php strtotime
func StrToTime(timeStr, layout string) (int64, error) {
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// Date parse timestamp to time string
// php date
func Date(timestamp int64, layout string) string {
	return time.Unix(timestamp, 0).Format(layout)
}

// CheckDate check date is right
// php checkdate
func CheckDate(month, day, year uint) bool {
	layout := "2006-01-02"
	timeStr := fmt.Sprintf("%d-%02d-%02d", year, month, day)
	_, err := time.Parse(layout, timeStr)
	if err != nil {
		return false
	}

	return true
}

// Sleep pauses the current goroutine for seconds
// php sleep
func Sleep(seconds int64) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

// Usleep pauses the current goroutine for microseconds
// php usleep
func Usleep(microseconds int64) {
	time.Sleep(time.Duration(microseconds) * time.Microsecond)
}
