package routing

import (
	"log"

	utils "github.com/alanwade2001/go-sepa-utils"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Router *gin.Engine
}

func NewRouter() *Router {
	s := &Router{
		Router: gin.Default(),
	}

	return s
}

func (s *Router) Run() error {
	address := utils.Getenv("ADDRESS", ":8080")
	return s.Router.Run(address)
}

func (s *Router) RunWithTLS() error {
	address := utils.Getenv("ADDRESS", ":8443")
	certFile := utils.Getenv("CERT", "server.crt")
	keyFile := utils.Getenv("KEY", "server.key")

	log.Printf("Address:[%s]", address)
	log.Printf("cert:[%s]", certFile)
	log.Printf("key:[%s]", keyFile)

	return s.Router.RunTLS(address, certFile, keyFile)
}
