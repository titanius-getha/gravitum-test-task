package app

import (
	"github.com/gin-gonic/gin"
	"github.com/titanius-getha/gravitum-test-task/domain/user"
	"github.com/titanius-getha/gravitum-test-task/pkg/config"
	"github.com/titanius-getha/gravitum-test-task/pkg/database"
	transport "github.com/titanius-getha/gravitum-test-task/pkg/transport/http"
	"github.com/titanius-getha/gravitum-test-task/pkg/transport/http/userhandler"
)

type app struct {
	s *transport.Server
	c *config.AppConfig
}

func New(conf *config.AppConfig) *app {
	s := transport.NewServer(conf.Env)

	db, err := database.NewPostgres(database.GetPostgresDsn(conf.Db.Name, conf.Db.User, conf.Db.Password, conf.Db.Host, conf.Db.Port))
	if err != nil {
		panic("cannot connect to database: " + err.Error())
	}

	us, err := user.NewService(user.NewPostgresRepository(db))
	if err != nil {
		panic("user service initialization error: " + err.Error())
	}
	uh := userhandler.New(us)

	s.Group("users", func(g *gin.RouterGroup) {
		g.GET(":id", uh.GetUser)
		g.POST("", uh.CreateUser)
		g.PUT(":id", uh.UpdateUser)
	})

	return &app{s, conf}
}

func (a *app) Start() {
	a.s.Run(a.c.Host, a.c.Port)
}
