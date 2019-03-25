package anyrpc

// null logger
var defaultLogger = func(tmpl string, s ...interface{}) {}
var errorLogger = func(tmpl string, s ...interface{}) {}

type loggerType func(tmpl string, s ...interface{})

func SetLogger(logger loggerType) {
	defaultLogger = logger
}

func SetErrorLogger(logger loggerType) {
	errorLogger = logger
}
