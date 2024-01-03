package common

import "time"

func DaysBetweenNowAndDate(targetDate string) int {
	layout := "2006-01-02"
	targetTime, _ := time.Parse(layout, targetDate)
	now := time.Now()
	difference := targetTime.Sub(now)
	days := int(difference.Hours() / 24)
	return days
}

func DaysBetweenNowAndTimestamp(targetTimestamp int64) int {
	// 将时间戳转换为 time.Time 对象
	targetTime := time.Unix(targetTimestamp/1000, 0) // 时间戳是毫秒级别，除以1000转换为秒

	now := time.Now()
	difference := targetTime.Sub(now)
	days := int(difference.Hours() / 24)

	return days
}

func FormatTimestampToDateString(timestamp int64) string {
	// 将毫秒级别的时间戳转换为 time.Time 对象
	timeObj := time.Unix(timestamp/1000, 0) // 时间戳是毫秒级别，除以1000转换为秒

	// 格式化为 "2006-01-02" 格式的字符串
	dateString := timeObj.Format("2006-01-02")

	return dateString
}
