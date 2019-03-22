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
}
type emgoImpl struct {
	s *mgo.Session
}

func GetEmgo(uriPrefix string) Emgo {
	flagName := uriPrefix + "mgo-url"
	var url string
	flag.StringVar(&url, flagName, "localhost", "URL for dialing mongodb server")
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	return &emgoImpl{
		s: session,
	}
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
