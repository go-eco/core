package core

import (
	"flag"
	"log"

	"github.com/sirupsen/logrus"

	"github.com/go-eco/core/elog"

	"github.com/facebookgo/flagenv"
	// Load environment variables from .env file
	_ "github.com/joho/godotenv/autoload"
)

// Application interface represent a main application object.
type Application interface {
	RegisterService(Service)
	RegisterRunner(Runnable)
	Run()
	Shutdown()
	GetLogger() *logrus.Logger
}

// Runnable interface represent a service that can be executed by method Run()
type Runnable interface {
	Service
	Run()
}

// Service interface represent a dependency which is used in the application
type Service interface {
	// Init method is called when registered to application.
	// It should be used to define flags.
	// Those flags when will have their values binded when the application run.
	// Right before the method Configure is called.
	Init()
	// Configure method is called when application start to run.
	// It is called after all flags are binded with their values.
	// This method is suitable for create connections, open files, create internal variables.
	Configure()
	// Cleanup method is called when shutdown signal is sent to application.
	// The service should close their resources in this method.
	Cleanup()
}

type applicationImpl struct {
	services []Service
	runner   Runnable
	logger   elog.Elog
}

// NewApp initiate and return an empty Application interface
func NewApp() Application {
	app := &applicationImpl{
		services: []Service{},
		logger:   elog.NewLogger(),
	}
	app.RegisterService(app.logger)

	return app
}

// RegisterService bind a service to the application
func (m *applicationImpl) RegisterService(s Service) {
	s.Init()
	m.services = append(m.services, s)
}
func (m *applicationImpl) RegisterRunner(r Runnable) {
	r.Init()
	m.runner = r
}
func (m *applicationImpl) Run() {
	if m.runner == nil {
		log.Fatal("No runner is registered. Service stops.")
	}

	flagenv.Prefix = ""
	flagenv.Parse()
	flag.Parse()
	m.configure()

	log.Println("Start main runner")
	m.runner.Run()
	m.Shutdown()
}
func (m *applicationImpl) configure() {
	for _, s := range m.services {
		s.Configure()
	}
	m.runner.Configure()
}
func (m *applicationImpl) Shutdown() {
	m.runner.Cleanup()
	for _, s := range m.services {
		s.Cleanup()
	}
}

func (m *applicationImpl) GetLogger() *logrus.Logger {
	return m.logger.GetLogger()
}
