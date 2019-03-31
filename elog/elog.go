package elog

import (
	"flag"

	"github.com/sirupsen/logrus"
)

var logLevels map[string]logrus.Level = map[string]logrus.Level{
	"trace": logrus.TraceLevel,
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
	"fatal": logrus.FatalLevel,
	"panic": logrus.PanicLevel,
}

type Elog interface {
	Init()
	Configure()
	Cleanup()
	GetLogger() *logrus.Logger
}

type elogImpl struct {
	logger    *logrus.Logger
	logLevel  string
	logFormat string
}

func NewLogger() Elog {
	return &elogImpl{
		logger: logrus.New(),
	}
}

func (m *elogImpl) Init() {
	flag.StringVar(&m.logLevel, "LOG_LEVEL", "WARNING", "Logrus level")
	flag.StringVar(&m.logFormat, "LOG_FORMAT", "text", "text|json")
}
func (m *elogImpl) Configure() {
	m.configureLogLevel()
}
func (m *elogImpl) GetLogger() *logrus.Logger {
	return m.logger
}
func (m *elogImpl) Cleanup() {}

func (m *elogImpl) configureLogLevel() {
	if level, ok := logLevels[m.logLevel]; ok {
		m.logger.SetLevel(level)
	} else {
		m.logger.SetLevel(logrus.WarnLevel)
		m.logger.Warnf("Logger: Invalid log level %s", m.logLevel)
	}
}
func (m *elogImpl) configureLogFormat() {
	if m.logFormat == "json" {
		m.logger.SetFormatter(&logrus.JSONFormatter{})
		return
	} else if m.logFormat != "text" {
		m.logger.Warnf("Logger: Invalid log format %s", m.logFormat)
	}
	m.logger.SetFormatter(&logrus.TextFormatter{})
}
