package utils

// TimeToNumber 将特定小时分钟转换为数字
func TimeToNumber(hour int, minute int) float64 {
	if hour > 23 || hour < 0 || minute < 0 || minute > 59 {
		return 0
	}
	if minute == 0 {
		return float64(hour)
	}
	return float64(hour) + float64(minute)/60.0
}
