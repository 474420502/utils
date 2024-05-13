// time truncate with Location
package truncate

import "time"

// 以时间片截断时间带time.Location决定 类似time.Truncate 但是支持国际时间. 例如 day的截断范围有值的差异
func TruncateZone(target time.Time, dur time.Duration, in *time.Location) time.Time {
	target = target.In(in)
	_, offset := target.Zone()
	offsetDuration := time.Second * time.Duration(offset)
	return target.Add(offsetDuration).Truncate(dur).Add(-offsetDuration)
}

// 以时间片截断时间 类似time.Truncate 但是支持国际时间. 例如 day的截断范围有值的差异
func TruncateLocalZone(target time.Time, dur time.Duration) time.Time {
	_, offset := target.Zone()
	offsetDuration := time.Second * time.Duration(offset)
	return target.Add(offsetDuration).Truncate(dur).Add(-offsetDuration)
}
