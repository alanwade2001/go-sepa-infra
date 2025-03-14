package routing

import (
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

func (s *Router) Run(address string) {
	s.Router.Run(address)
}
