package test

import (
	"testing"

	"github.com/go-eco/core"
)

type appTestService struct {
	initCall    bool
	runningCall bool
	cleanupCall bool
}

func (m *appTestService) Init() {
	m.initCall = true
}
func (m *appTestService) Run() {
	m.runningCall = true
}
func (m *appTestService) Cleanup() {
	m.cleanupCall = true
}

func TestRunApp(t *testing.T) {
	app := core.NewApp()
	runner := appTestService{}
	app.RegisterRunner(&runner)
	app.Run()

	if !runner.initCall {
		t.Fatal("Method Init isn't called.")
	}
	if !runner.runningCall {
		t.Fatal("Method Run isn't called.")
	}
	if !runner.cleanupCall {
		t.Fatal("Method Cleanup isn't called.")
	}
}

func TestRegisterService(t *testing.T) {
	app := core.NewApp()
	service := appTestService{}
	app.RegisterService(&service)
	runner := appTestService{}
	app.RegisterRunner(&runner)
	app.Run()

	if !service.initCall {
		t.Fatal("Method Init isn't called.")
	}
	if !service.cleanupCall {
		t.Fatal("Method Cleanup isn't called.")
	}
}
