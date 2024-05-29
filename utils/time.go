package utils

import "time"

type _Time struct{}

var Time _Time

func (c _Time) Timestamp2RFC3339(timestamp int64) string {
	t := time.Unix(timestamp/1000, (timestamp%1000)*1000000) // 使用原始时间戳,而不是除以 1000
	return t.UTC().Format(time.RFC3339)                      // 将时间转换为 UTC 时区
}
