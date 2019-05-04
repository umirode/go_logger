package Handler

import "github.com/umirode/go_logger"

type AMQPHandler struct {
}

func (*AMQPHandler) Handle(record *go_logger.Record) bool {
	panic("implement me")
}

func (*AMQPHandler) IsHandling(record *go_logger.Record) bool {
	panic("implement me")
}

func (*AMQPHandler) Close() {
	panic("implement me")
}
