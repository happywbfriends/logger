package logger

type Logger interface {
	With(k, v string) Logger
	Errorf(fmt string, p ...any)
	Warnf(fmt string, p ...any)
	Infof(fmt string, p ...any)
	Debugf(fmt string, p ...any)
}

type NoLogger struct{}

func (l *NoLogger) With(_, _ string) Logger {
	return l
}
func (l *NoLogger) Errorf(string, ...any) {}
func (l *NoLogger) Warnf(string, ...any)  {}
func (l *NoLogger) Infof(string, ...any)  {}
func (l *NoLogger) Debugf(string, ...any) {}
