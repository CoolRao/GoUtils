package goutils

import (
	"lottery/conf"
	"time"
)

/**
当前时间的时间戳
*/
func NowUnix() int {
	return int(time.Now().In(SysTimeLocation).Unix())
}

/**
将unix时间戳格式化为yyyymmdd H:i:s格式字符串
*/
func FormatFromUnixTime(t int64) string {
	if t > 0 {
		return time.Unix(t, 0).Format(SysTimeform)
	} else {
		return time.Now().Format(SysTimeform)
	}
}

/**
将unix时间戳格式化为yyyymmdd格式字符串
*/
func FormatFromUnixTimeShort(t int64) string {
	if t > 0 {
		return time.Unix(t, 0).Format(SysTimeformShort)
	} else {
		return time.Now().Format(SysTimeformShort)
	}
}

/**
将字符串转成时间
*/
func ParseTime(str string) (time.Time, error) {
	return time.ParseInLocation(SysTimeform, str, conf.SysTimeLocation)
}

/**
得到当前时间到下一天零点的延时
*/
func NextDayDuration() time.Duration {
	year, month, day := time.Now().Add(time.Hour * 24).Date()
	next := time.Date(year, month, day, 0, 0, 0, 0, conf.SysTimeLocation)
	return next.Sub(time.Now())
}
