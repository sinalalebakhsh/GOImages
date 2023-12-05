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
	// To get implementations of the Logger and Configuration interfaces, 
	// the code in the main function needs to
	// know how to create instances of structs that implement those interfaces:
	// This is a workable approach, 
	// but it undermines the purpose of defining an interface, 
	// it requires care to ensure that instances are created consistently, 
	// and it complicates the process of replacing one interface
	// implementation with another.
	// My preferred approach is to use dependency injection (DI), in which code that depends on an interface
	// can obtain an implementation without needing to select an underlying type or create an instance directly. I
	// am going to start with service location, which will provide the foundation for more advanced features later.
	// During application startup, the interfaces defined by the application will be added to a register, along
	// with a factory function that creates instances of an implementation struct. So, for example, the platform.
	// logger.Logger interface will be registered with a factory function that invokes the NewDefaultLogger
	// function. When an interface is added to the register, it is known as a service.
	// During execution, application components that need the features described by the service go to the
	// registry and request the interface they want. The registry invokes the factory function and returns the struct
	// that is created, which allows the application component to use the interface features without knowing or
	// specifying which implementation struct will be used or how it is created. Don’t worry if this doesn’t make
	// sense—this can be a difficult topic to understand, and it becomes easier once you see it in action.
	cfg, err = config.Load("config.json")
	if err != nil {
		panic(err)
	}
	var logger logging.Logger = logging.NewDefaultLogger(cfg)
	writeMessage(logger, cfg)
}
