package transport

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/titanius-getha/gravitum-test-task/pkg/config"
)

type Server struct {
	e *gin.Engine
}

func NewServer(mode config.EnvMode) *Server {
	var e *gin.Engine
	if mode == config.EnvModeDev {
		e = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		e = gin.New()
	}

	e.SetTrustedProxies(nil)

	return &Server{e}
}

func (s *Server) Run(host string, port int) error {
	return s.e.Run(fmt.Sprintf("%s:%d", host, port))
}

func (s *Server) Group(basepath string, handler func(g *gin.RouterGroup)) {
	group := s.e.Group(basepath)
	handler(group)
}
