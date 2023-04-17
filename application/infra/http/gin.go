package infrahttp

import (
	"fd-test/application/adapter"
	"fd-test/application/controller"
	"fd-test/application/database"
	"fd-test/application/repository"
	"fd-test/application/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	port   string
	db     *database.DB
	middle Middleware
}

func NewRouter(port string, db *database.DB) Router {
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

	userAdapter := adapter.NewUserAdapter()
	userRepo := repository.NewUserRepo(r.db)
	userService := service.NewUserService(userRepo, userAdapter)
	userController := controller.NewUserController(userService, r.middle)

	userController.RegisterRoute(r.router)

	r.router.Run(fmt.Sprintf(":%s", r.port))
}

func (r Router) SetMiddleware() Router {
	mid := NewMiddleware()
	r.middle = mid
	return r
}
