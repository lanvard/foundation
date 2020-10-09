package loggers

import (
	"errors"
	"github.com/lanvard/contract/inter"
)

type Stack struct {
	app     inter.App
	Loggers []string
}

func (s *Stack) SetApp(app inter.App) {
	s.app = app
}

func (s Stack) Log(severity inter.Severity, message string) {
	for _, logger := range s.getLoggers() {
		logger.Log(severity, message)
	}
}

func (s Stack) LogWith(severity inter.Severity, message string, data interface{}) {
	// 	for _, logger := range s.getLoggers() {
	// 		logger.LogWith(severity, message, data)
	// 	}
}

func (s Stack) Emergency(message string) {
	// 	for _, logger := range s.getLoggers() {
	// 		logger.Emergency(message)
	// 	}
}

func (s Stack) EmergencyWith(message string, data interface{}) {
	// 	for _, logger := range s.getLoggers() {
	// 		logger.EmergencyWith(message, data)
	// 	}
}

func (s Stack) Alert(message string) {
	// 	for _, logger := range s.getLoggers() {
	// 		logger.Alert(message)
	// 	}
}

func (s Stack) AlertWith(message string, data interface{}) {
	// 	for _, logger := range s.getLoggers() {
	// 		logger.AlertWith(message, data)
	// 	}
}

func (s Stack) Critical(message string) {
	// 	for _, logger := range s.getLoggers() {
	// 		logger.Critical(message)
	// 	}
}

func (s Stack) CriticalWith(message string, data interface{}) {
	// 	for _, logger := range s.getLoggers() {
	// 		logger.CriticalWith(message, data)
	// 	}
}

func (s Stack) Error(message string) {
	// 	for _, logger := range s.getLoggers() {
	// 		logger.Error(message)
	// 	}
}

func (s Stack) ErrorWith(message string, data interface{}) {
	// 	for _, logger := range s.getLoggers() {
	// 		logger.ErrorWith(message, data)
	// 	}
}

func (s Stack) Warning(message string) {
	// 	for _, logger := range s.getLoggers() {
	// 		logger.Warning(message)
	// 	}
}

func (s Stack) WarningWith(message string, data interface{}) {
	// 	for _, logger := range s.getLoggers() {
	// 		logger.WarningWith(message, data)
	// 	}
}

func (s Stack) Notice(message string) {
	// 	for _, logger := range s.getLoggers() {
	// 		logger.Notice(message)
	// 	}
}

func (s Stack) NoticeWith(message string, data interface{}) {
	// 	for _, logger := range s.getLoggers() {
	// 		logger.NoticeWith(message, data)
	// 	}
}

func (s Stack) Info(message string) {
	// 	for _, logger := range s.getLoggers() {
	// 		logger.Info(message)
	// 	}
}

func (s Stack) InfoWith(message string, data interface{}) {
	// 	for _, logger := range s.getLoggers() {
	// 		logger.InfoWith(message, data)
	// 	}
}

func (s Stack) Debug(message string) {
	// 	for _, logger := range s.getLoggers() {
	// 		logger.Debug(message)
	// 	}
}

func (s Stack) DebugWith(message string, data interface{}) {
	// 	for _, logger := range s.getLoggers() {
	// 		logger.DebugWith(message, data)
	// 	}
}

func (s Stack) getLoggers() []inter.Logger {
	var loggers []inter.Logger
	allChannels := s.app.Make("config.Logging.Loggers").(map[string]inter.Logger)
	for _, loggerName := range s.Loggers {
		logger, ok := allChannels[loggerName]
		if !ok {
			panic(errors.New("no logger found by: " + loggerName))
		}
		loggers = append(loggers, logger)
	}
	return loggers
}