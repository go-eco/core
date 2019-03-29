package emgo

import (
	"flag"
	"log"

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
	uriPrefix string
	uri       string
}

func GetEmgo(uriPrefix string) Emgo {
	return &emgoImpl{
		uriPrefix: uriPrefix,
	}
}

func (m *emgoImpl) Init() {
	flagName := m.uriPrefix + "mgo-uri"
	flag.StringVar(&m.uri, flagName, "localhost", "URI for dialing mongodb server")
}
func (m *emgoImpl) Configure() {
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
