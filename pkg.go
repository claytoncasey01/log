package log

// singletons ftw?
var Log Interface = &Logger{
	Handler: HandlerFunc(handleStdLog),
}

// SetHandler sets the handler. This is not thread-safe.
// The default handler outputs to the stdlib log.
func SetHandler(h Handler) {
	if logger, ok := Log.(*Logger); ok {
		logger.Handler = h
	}
}

// WithFields returns a new entry with `fields` set.
func WithFields(fields Fielder) *Entry {
	return Log.WithFields(fields)
}

// WithField returns a new entry with the `key` and `value` set.
func WithField(key string, value interface{}) *Entry {
	return Log.WithField(key, value)
}

// WithError returns a new entry with the "error" set to `err`.
func WithError(err error) *Entry {
	return Log.WithError(err)
}

// Debug level message.
func Debug(msg string) {
	Log.Debug(msg)
}

// Info level message.
func Info(msg string) {
	Log.Info(msg)
}

// Warn level message.
func Warn(msg string) {
	Log.Warn(msg)
}

// Error level message.
func Error(msg string) {
	Log.Error(msg)
}

// Fatal level message, followed by an exit.
func Fatal(msg string) {
	Log.Fatal(msg)
}

// Debugf level formatted message.
func Debugf(msg string, v ...interface{}) {
	Log.Debugf(msg, v...)
}

// Infof level formatted message.
func Infof(msg string, v ...interface{}) {
	Log.Infof(msg, v...)
}

// Warnf level formatted message.
func Warnf(msg string, v ...interface{}) {
	Log.Warnf(msg, v...)
}

// Errorf level formatted message.
func Errorf(msg string, v ...interface{}) {
	Log.Errorf(msg, v...)
}

// Fatalf level formatted message, followed by an exit.
func Fatalf(msg string, v ...interface{}) {
	Log.Fatalf(msg, v...)
}

// Trace returns a new entry with a Stop method to fire off
// a corresponding completion log, useful with defer.
func Trace(msg string) *Entry {
	return Log.Trace(msg)
}
