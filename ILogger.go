package logger

const Emergency = 600
const Alert = 550
const Critical = 500
const Error = 400
const Warning = 300
const Notice = 250
const Info = 200
const Debug = 100

var Levels = map[uint]string{
	Emergency: "EMERGENCY",
	Alert:     "ALERT",
	Critical:  "CRITICAL",
	Error:     "ERROR",
	Warning:   "WARNING",
	Notice:    "NOTICE",
	Info:      "INFO",
	Debug:     "DEBUG",
}

type ILogger interface {
	Emergency(message string, context interface{})
	Alert(message string, context interface{})
	Critical(message string, context interface{})
	Error(message string, context interface{})
	Warning(message string, context interface{})
	Notice(message string, context interface{})
	Info(message string, context interface{})
	Debug(message string, context interface{})
}
