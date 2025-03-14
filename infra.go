package infra

import (
	db "github.com/alanwade2001/go-sepa-db"
	"github.com/alanwade2001/go-sepa-infra/routing"
	q "github.com/alanwade2001/go-sepa-q"
)

type Infra struct {
	Router  *routing.Router
	Stomp   *q.Stomp
	Persist *db.Persist
}

func NewInfra() *Infra {
	app := &Infra{}

	app.Stomp = q.NewStomp()
	app.Persist = db.NewPersist()
	app.Router = routing.NewRouter()

	return app
}

func (a *Infra) Run() {

	a.Router.Run()
}
