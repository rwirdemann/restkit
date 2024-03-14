package adapter

import "time"

type Time struct{}

func (Time) TS() string {
	return time.Now().Format("20060315150405")
}
