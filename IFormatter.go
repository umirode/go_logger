package logger

type IFormatter interface {
	Format(message interface{}) string
}
