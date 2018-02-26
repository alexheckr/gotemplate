package template

import (
	"os"

	logging "github.com/op/go-logging"
)

const (
	logger      = "gotemplate"
	loggingBase = "Data conversion functions"
)

var loggingFuncs funcTableMap

func (t *Template) addLoggingFuncs() {
	if loggingFuncs == nil {
		loggingFuncs = funcTableMap{
			"fatal":    {func(args ...interface{}) string { return logBase(Log.Fatal, args...) }, loggingBase, nil, []string{}, ""},
			"fatalf":   {func(format string, args ...interface{}) string { return logBasef(Log.Fatalf, format, args...) }, loggingBase, nil, []string{}, ""},
			"error":    {func(args ...interface{}) string { return logBase(Log.Error, args...) }, loggingBase, nil, []string{}, ""},
			"errorf":   {func(format string, args ...interface{}) string { return logBasef(Log.Errorf, format, args...) }, loggingBase, nil, []string{}, ""},
			"warning":  {func(args ...interface{}) string { return logBase(Log.Warning, args...) }, loggingBase, nil, []string{}, ""},
			"warningf": {func(format string, args ...interface{}) string { return logBasef(Log.Warningf, format, args...) }, loggingBase, nil, []string{}, ""},
			"notice":   {func(args ...interface{}) string { return logBase(Log.Notice, args...) }, loggingBase, nil, []string{}, ""},
			"noticef":  {func(format string, args ...interface{}) string { return logBasef(Log.Noticef, format, args...) }, loggingBase, nil, []string{}, ""},
			"info":     {func(args ...interface{}) string { return logBase(Log.Info, args...) }, loggingBase, nil, []string{}, ""},
			"infof":    {func(format string, args ...interface{}) string { return logBasef(Log.Infof, format, args...) }, loggingBase, nil, []string{}, ""},
			"debug":    {func(args ...interface{}) string { return logBase(Log.Debug, args...) }, loggingBase, nil, []string{}, ""},
			"debugf":   {func(format string, args ...interface{}) string { return logBasef(Log.Debugf, format, args...) }, loggingBase, nil, []string{}, ""},
		}
	}

	t.AddFunctions(loggingFuncs)
}

func logBase(f func(...interface{}), args ...interface{}) string {
	f(args...)
	return ""
}

func logBasef(f func(string, ...interface{}), format string, args ...interface{}) string {
	f(format, args...)
	return ""
}

// Log is the logger used to log message during template processing
var Log = logging.MustGetLogger(logger)

func getLogLevel() logging.Level {
	return logging.GetLevel(logger)
}

// InitLogging allows configuration of the default logging level
func InitLogging(level logging.Level, simple bool) {
	format := `[%{module}] %{time:2006/01/02 15:04:05.000} %{color}%{level:-8s} %{message}%{color:reset}`
	if simple {
		format = `[%{level}] %{message}`
	}
	logging.SetBackend(logging.NewBackendFormatter(logging.NewLogBackend(os.Stderr, "", 0), logging.MustStringFormatter(format)))
	logging.SetLevel(level, logger)
}