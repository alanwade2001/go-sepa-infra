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

func (s *Router) Run() {
	address := utils.Getenv("ADDRESS", ":8080")
	s.Router.Run(address)
}
