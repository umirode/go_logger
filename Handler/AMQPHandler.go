package Handler

import "github.com/umirode/go_logger"

type AMQPHandler struct {
}

func (*AMQPHandler) Handle(record *logger.Record) bool {
	panic("implement me")
}

func (*AMQPHandler) IsHandling(record *logger.Record) bool {
	panic("implement me")
}

func (*AMQPHandler) Close() {
	panic("implement me")
}
