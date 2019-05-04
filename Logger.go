package go_logger

import "time"

type Processor func(record *Record) *Record

type Logger struct {
	Name       string
	Handlers   []IHandler
	Processors []Processor
}

func NewLogger(name string, handlers []IHandler, processors []Processor) *Logger {
	return &Logger{
		Name:       name,
		Handlers:   handlers,
		Processors: processors,
	}
}

func (l *Logger) AddHandler(handler IHandler) {
	l.Handlers = append(l.Handlers, handler)
}

func (l *Logger) Close(handler IHandler) {
	for _, handler := range l.Handlers {
		handler.Close()
	}
}

func (l *Logger) addRecord(message string, context interface{}, level uint) bool {
	record := &Record{}
	record.Channel = l.Name
	record.Context = context
	record.DateTime = time.Now()
	record.Level = level
	record.LevelName = Levels[level]

	availableHandlers := make([]IHandler, 0)
	for _, handler := range l.Handlers {
		if handler.IsHandling(record) {
			availableHandlers = append(availableHandlers, handler)
		}
	}

	if len(availableHandlers) == 0 {
		return false
	}

	for _, processor := range l.Processors {
		record = processor(record)
	}

	for _, handler := range availableHandlers {
		handler.Handle(record)
	}

	return true
}

func (l *Logger) Emergency(message string, context interface{}) {
	l.addRecord(message, context, Emergency)
}

func (l *Logger) Alert(message string, context interface{}) {
	l.addRecord(message, context, Alert)
}

func (l *Logger) Critical(message string, context interface{}) {
	l.addRecord(message, context, Critical)
}

func (l *Logger) Error(message string, context interface{}) {
	l.addRecord(message, context, Error)
}

func (l *Logger) Warning(message string, context interface{}) {
	l.addRecord(message, context, Warning)
}

func (l *Logger) Notice(message string, context interface{}) {
	l.addRecord(message, context, Notice)
}

func (l *Logger) Info(message string, context interface{}) {
	l.addRecord(message, context, Info)
}

func (l *Logger) Debug(message string, context interface{}) {
	l.addRecord(message, context, Debug)
}
