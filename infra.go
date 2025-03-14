package infra

import (
	db "github.com/alanwade2001/go-sepa-db"
	"github.com/alanwade2001/go-sepa-infra/routing"
	q "github.com/alanwade2001/go-sepa-q"
	utils "github.com/alanwade2001/go-sepa-utils"
)

type Infra struct {
	Router  *routing.Router
	Stomp   *q.Stomp
	Persist *db.Persist
}

func NewInfra() *Infra {
	app := &Infra{}

	app.Stomp = q.NewStomp()
	schema := utils.Getenv("DB_SCHEMA", "postgres")
	app.Persist = db.NewPersist(schema)
	app.Router = routing.NewRouter()

	return app
}

func (a *Infra) Run() {
	address := utils.Getenv("ADDRESS", ":8080")
	a.Router.Run(address)
}
