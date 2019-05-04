package Handler

import "github.com/umirode/go_logger"

type LogrusHandler struct {
}

func (*LogrusHandler) Handle(record *logger.Record) bool {
	panic("implement me")
}

func (*LogrusHandler) IsHandling(record *logger.Record) bool {
	panic("implement me")
}

func (*LogrusHandler) Close() {
	panic("implement me")
}
