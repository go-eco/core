package emgo

import (
	"flag"

	"github.com/go-eco/core"

	"github.com/globalsign/mgo"
)

type Emgo interface {
	Session() *mgo.Session
	DB() (*mgo.Database, func())
	C(name string) (*mgo.Collection, func())
	Cleanup()
	Configure()
	Init()
}
type emgoImpl struct {
	s         *mgo.Session
	app       core.Application
	uriPrefix string
	uri       string
}

func GetEmgo(uriPrefix string, app core.Application) Emgo {
	return &emgoImpl{
		uriPrefix: uriPrefix,
		app:       app,
	}
}

func (m *emgoImpl) Init() {
	flagName := m.uriPrefix + "mgo-uri"
	flag.StringVar(&m.uri, flagName, "localhost", "URI for dialing mongodb server")
}
func (m *emgoImpl) Configure() {
	log := m.app.GetLogger()
	log.Errorf("Dial MGO server at %s", m.uri)
	session, err := mgo.Dial(m.uri)
	if err != nil {
		log.Fatal(err)
	}
	m.s = session
}
func (m *emgoImpl) Cleanup() {
	m.s.Close()
}

func (m *emgoImpl) Session() *mgo.Session {
	return m.s.Clone()
}

func (m *emgoImpl) DB() (*mgo.Database, func()) {
	ns := m.Session()
	return ns.DB(""), ns.Close
}

func (m *emgoImpl) C(name string) (*mgo.Collection, func()) {
	db, close := m.DB()
	return db.C(name), close
}
