package util

import "time"

func TimeFormat(timeUnix int64) string {
	return time.Unix(timeUnix, 0).Format(time.DateTime)
}
