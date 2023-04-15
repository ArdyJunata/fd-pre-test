package infrahttp

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	port   string
	db     *sql.DB
}

func NewRouter(port string, db *sql.DB) Router {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})

	return Router{
		router: router,
		port:   port,
		db:     db,
	}
}

func (r Router) Run() {
	fmt.Println("server running at port", r.port)

	r.router.Run(fmt.Sprintf(":%s", r.port))
}
