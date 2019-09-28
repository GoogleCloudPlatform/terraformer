package datahub

import "time"

func Uint64ToTimeString(t uint64) string {
	return time.Unix(int64(t), 0).Format("2006-01-02 15:04:05")
}
