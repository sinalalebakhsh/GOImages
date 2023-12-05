package logging

type LogLevel int

const (
	Trace LogLevel = iota
	Debug
	Information
	Warning
	Fatal
	None
)

// This file defines the Logger interface, which specifies methods 
// for logging messages with different levels of severity, 
// which is set using a LogLevel value ranging from Trace to Fatal. 
// There is also a None level that specifies no logging output. 
// For each level of severity, the Logger interface defines 
// one method that accepts a simple string 
// and one method that accepts a template string and placeholder values.
type Logger interface {
	Trace(string)
	Tracef(string, ...interface{})
	Debug(string)
	Debugf(string, ...interface{})
	Info(string)
	Infof(string, ...interface{})
	Warn(string)
	Warnf(string, ...interface{})
	Panic(string)
	Panicf(string, ...interface{})
}

