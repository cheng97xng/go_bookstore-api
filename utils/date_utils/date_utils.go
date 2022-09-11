package date_utils

import "time"

const (
	apiDateLayout = "2006-1-02T15:04:05.000Z"
)

func GetNow() time.Time {
	return time.Now()
}
func GetNowString() string {
	return GetNow().Format(apiDateLayout)

}
