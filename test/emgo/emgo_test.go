package emgo

import (
	"testing"

	"github.com/globalsign/mgo/bson"
	"github.com/go-eco/core"
	"github.com/go-eco/core/emgo"
)

type appTestMgo struct {
	mgo emgo.Emgo
	t   *testing.T
}

func (m *appTestMgo) Init()      {}
func (m *appTestMgo) Configure() {}
func (m *appTestMgo) Cleanup()   {}
func (m *appTestMgo) Run() {
	db, close := m.mgo.C("user")
	defer close()
	err := db.Insert(bson.M{"name": "Nhu", "description": "A software developer"})
	if err != nil {
		m.t.Fatal(err)
	}
}

func TestEmgo(t *testing.T) {
	app := core.NewApp()
	mgo := emgo.GetEmgo("", app)
	app.RegisterService(mgo)

	appTest := appTestMgo{
		mgo: mgo,
		t:   t,
	}
	app.RegisterRunner(&appTest)

	app.Run()
}
