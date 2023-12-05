package logging

import (
	"log"
	"os"
)


// The NewDefaultLogger function creates a DefaultLogger 
// with a minimum severity level and log.
// Loggers that write messages to standard out.
func NewDefaultLogger(level LogLevel) Logger {
	flags := log.Lmsgprefix | log.Ltime
	return &DefaultLogger{
		minLevel: level,
		loggers: map[LogLevel]*log.Logger{
			Trace:       log.New(os.Stdout, "TRACE ", flags),
			Debug:       log.New(os.Stdout, "DEBUG ", flags),
			Information: log.New(os.Stdout, "INFO ", flags),
			Warning:     log.New(os.Stdout, "WARN ", flags),
			Fatal:       log.New(os.Stdout, "FATAL ", flags),
		},
		triggerPanic: true,
	}
}
