package go_logger

type IFormatter interface {
	Format(message interface{}) string
}
