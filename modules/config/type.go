package config

import "time"

// Duration 自定持续时间
type Duration time.Duration

// UnmarshalText 解析字符串到定时
func (d *Duration) UnmarshalText(data []byte) error {
	duration, err := time.ParseDuration(string(data))
	if err == nil {
		*d = Duration(duration)
	}
	return err
}

// MarshalText 转码为字符串
func (d Duration) MarshalText() ([]byte, error) {
	return []byte(time.Duration(d).String()), nil
}
