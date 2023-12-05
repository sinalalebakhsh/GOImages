package main
import (
	//"fmt"
	"platform/logging"
)
func writeMessage(logger logging.Logger) {
	logger.Info("Hello, Platform")
}
func main() {
	var logger logging.Logger = logging.NewDefaultLogger(logging.Information)
	writeMessage(logger)
}