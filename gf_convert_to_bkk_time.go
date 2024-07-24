package helpers

import "time"

func ConvertToBKKTime(t time.Time) time.Time {
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return time.Time{}
	}
	return t.In(location) // return current time in Bangkok
}
