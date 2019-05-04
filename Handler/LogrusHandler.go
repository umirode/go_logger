package Handler

import "github.com/umirode/go_logger"

type LogrusHandler struct {
}

func (*LogrusHandler) Handle(record *go_logger.Record) bool {
	panic("implement me")
}

func (*LogrusHandler) IsHandling(record *go_logger.Record) bool {
	panic("implement me")
}

func (*LogrusHandler) Close() {
	panic("implement me")
}
