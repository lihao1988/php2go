package datetime

import (
	"fmt"
	"testing"
)

func TestDatetime(t *testing.T) {
	fmt.Println("Time: ", Time())

	timeStr := "2018-01-18 08:08:08"
	layout := "2006-01-02 15:04:05"
	timestamp, err := StrToTime(timeStr, layout)
	fmt.Println("StrToTime: ", timestamp, ", error: ", err)

	fmt.Println("Date: ", Date(timestamp, layout))

	Usleep(3000)
	fmt.Println("....")
	Sleep(3)

	fmt.Println("CheckDate: ", CheckDate(2, 29, 2021))
}
