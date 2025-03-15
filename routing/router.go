package routing

import (
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
	address := utils.Getenv("ADDRESS", ":8080")
	certFile := utils.Getenv("CERT", ":8080")
	keyFile := utils.Getenv("KEY", ":8080")

	return s.Router.RunTLS(address, certFile, keyFile)
}
