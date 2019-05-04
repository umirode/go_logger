package logger

import "time"

type Record struct {
	Message   string
	Context   interface{}
	Level     uint
	LevelName string
	Channel   string
	DateTime  time.Time
}
