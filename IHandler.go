package go_logger

type IHandler interface {
	Handle(record *Record) bool
	IsHandling(record *Record) bool
	Close()
}
