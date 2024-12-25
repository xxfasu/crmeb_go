package util

import (
	"crmeb_go/constants"
	"time"
)

func TimeFormat(timeUnix int64) string {
	return time.Unix(timeUnix, 0).Format(time.DateTime)
}

func CalculateDateRange(data string) (int64, int64) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	var startTime, endTime int64

	switch data {
	case constants.SearchDateDay:
		startTime, endTime = getDayStartEnd(now)
	case constants.SearchDateYesterday:
		startTime, endTime = getDayStartEnd(now.AddDate(0, 0, -1))
	case constants.SearchDateLately7:
		startTime = today.AddDate(0, 0, -7).Unix()
		endTime = today.Add(-1 * time.Second).Unix()
	case constants.SearchDateWeek:
		// 获取本周的开始时间，周日到周六
		startTime = today.AddDate(0, 0, -int(today.Weekday())).Unix()
		endTime = now.Unix()
	case constants.SearchDatePreWeek:
		// 获取上周的开始和结束时间
		startTime = today.AddDate(0, 0, -int(today.Weekday())-6).Unix()
		endTime = today.AddDate(0, 0, -int(today.Weekday())).Unix()
	case constants.SearchDateLately30:
		startTime = today.AddDate(0, 0, -30).Unix()
		endTime = now.Unix()
	case constants.SearchDateMonth:
		// 当前月的第一天
		startTime = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Unix()
		endTime = now.Unix()
	case constants.SearchDatePreMonth:
		// 上个月的第一天和最后一天
		startTime = time.Date(now.Year(), now.Month()-1, 1, 0, 0, 0, 0, now.Location()).Unix()
		endTime = time.Date(now.Year(), now.Month(), 0, 23, 59, 59, 999999999, now.Location()).Unix()
	case constants.SearchDateYear:
		// 当前年份的第一天
		startTime = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location()).Unix()
		endTime = now.Unix()
	case constants.SearchDatePreYear:
		// 去年的第一天和最后一天
		startTime = time.Date(now.Year()-1, 1, 1, 0, 0, 0, 0, now.Location()).Unix()
		endTime = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location()).Unix() - 1
	default:
		// 默认今天
		startTime, endTime = getDayStartEnd(now)
	}

	return startTime, endTime
}

// getDayStartEnd 获取当天零点和24点的 Unix 时间戳
func getDayStartEnd(t time.Time) (int64, int64) {
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
	end := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location()).Unix()
	return start, end
}

// GetListDate 获取某一时间段内的时间集合
func GetListDate(data string) []string {
	today := time.Now()
	dateString := make([]string, 0, 32)
	switch data {
	case constants.SearchDateLately30:
		startDate := today.AddDate(0, 0, -30)
		for date := startDate; date.Unix() <= today.Unix(); date = date.Add(24 * time.Hour) {
			dateString = append(dateString, date.Format("01-02"))
		}
	case constants.SearchDateMonth:
		year, month, _ := today.Date()
		location := today.Location()
		firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, location)
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
		for day := firstOfMonth; day.Before(lastOfMonth) || day.Equal(lastOfMonth); day = day.AddDate(0, 0, 1) {
			startOfDay := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, location)
			dateString = append(dateString, startOfDay.Format("2"))
		}
	case constants.SearchDateWeek:
		dateString = []string{"星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"}
	case constants.SearchDateYear:
		dateString = []string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"}
	}

	return dateString
}
