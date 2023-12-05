package main

import (
	//"fmt"
	"platform/config"
	"platform/logging"
)


// The configuration is loaded from the config.json file, 
// and the Configuration implementation is
// passed to the NewDefaultLogger function, 
// which uses it to read the log level setting.


// The writeMessage function demonstrates the use of a configuration section, 
// which can be a good way to provide a component with the settings it needs, 
// especially if multiple instances are required with different settings, 
// each of which can be defined in its own section.
func writeMessage(logger logging.Logger, cfg config.Configuration) {
	section, ok := cfg.GetSection("main")
	if ok {
		message, ok := section.GetString("message")
		if ok {
			logger.Info(message)
		} else {
			logger.Panic("Cannot find configuration setting")
		}
	} else {
		logger.Panic("Config section not found")
	}
}

func main() {
	var cfg config.Configuration
	var err error
	cfg, err = config.Load("config.json")
	if err != nil {
		panic(err)
	}
	var logger logging.Logger = logging.NewDefaultLogger(cfg)
	writeMessage(logger, cfg)
}

