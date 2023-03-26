package logger

type Logger interface {
	Debug(string)
	Info(string)
	Error(error)
	Fatal(error)
}
