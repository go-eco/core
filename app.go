package core

import (
	"log"
)

type Application interface {
	RegisterService(Service)
	RegisterRunner(Runable)
	Run()
	Shutdown()
}
type Runable interface {
	Service
	Run()
}
type Service interface {
	Init()
	Cleanup()
}

type applicationImpl struct {
	services []Service
	runner   Runable
}

func NewApp() Application {
	return &applicationImpl{
		services: []Service{},
	}
}

func (m *applicationImpl) RegisterService(s Service) {
	s.Init()
	m.services = append(m.services, s)
}
func (m *applicationImpl) RegisterRunner(r Runable) {
	r.Init()
	m.runner = r
}
func (m *applicationImpl) Run() {
	if m.runner == nil {
		log.Fatal("No runner is registered. Service stops.")
	}
	log.Println("Start main runner")
	m.runner.Run()
	m.Shutdown()
}
func (m *applicationImpl) Shutdown() {
	m.runner.Cleanup()
	for _, s := range m.services {
		s.Cleanup()
	}
}
