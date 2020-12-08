package api

import (
	"fmt"
	"io"
	"os"

	"github.com/bingoohuang/go-rest-template/internal/api/controllers"
	"github.com/bingoohuang/go-rest-template/internal/api/middlewares"
	"github.com/bingoohuang/go-rest-template/internal/pkg/ginx"
	"github.com/bingoohuang/go-rest-template/internal/pkg/models/tasks"
	models "github.com/bingoohuang/go-rest-template/internal/pkg/models/users"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/bingoohuang/go-rest-template/pkg/helpers"

	"github.com/bingoohuang/go-rest-template/internal/pkg/conf"
	"github.com/bingoohuang/go-rest-template/internal/pkg/db"
	"github.com/gin-gonic/gin"
)

func setConfiguration(configPath string) {
	conf.Setup(configPath)
	db.SetupDB()
	gin.SetMode(conf.GetConf().Server.Mode)
}

func Run(configPath ...string) {
	setConfiguration(helpers.DefaultTo(configPath, "data/conf.yml"))
	c := conf.GetConf()
	web := routerSetup()

	fmt.Println("REST API Running on port", c.Server.Port)
	_ = web.Run(fmt.Sprintf(":%d", c.Server.Port))
}

func routerSetup() *gin.Engine {
	// Logging to a file.
	f, _ := os.Create("log/api.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	app := gin.New()

	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(p gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			p.ClientIP, p.TimeStamp.Format("2006-01-02 15:04:05"),
			p.Method, p.Path, p.Request.Proto, p.StatusCode, p.Latency, p.Request.UserAgent(), p.ErrorMessage)
	}))
	app.Use(gin.Recovery(), middlewares.CORS())
	app.NoRoute(middlewares.NoRouteHandler())

	// Routes
	// Docs Routes, http://localhost:3000/docs/index.html
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Login Routes
	app.POST("/api/login", ginx.FnJSON(controllers.Login, controllers.LoginInput{}))
	// User Routes
	app.GET("/api/users", ginx.FnBind(controllers.GetUsers, models.User{}))
	app.GET("/api/users/:id", ginx.FnParam(controllers.GetUserById, "id"))
	app.POST("/api/users", ginx.FnJSON(controllers.CreateUser, controllers.UserInput{}))
	app.PUT("/api/users/:id", ginx.FnParamJSON(controllers.UpdateUser, "id", controllers.UserInput{}))
	app.DELETE("/api/users/:id", ginx.FnParam(controllers.DeleteUser, "id"))
	// Tasks Routes
	app.GET("/api/tasks/:id", ginx.FnParam(controllers.GetTaskById, "id"))
	app.GET("/api/tasks", ginx.FnBind(controllers.GetTasks, tasks.Task{}))
	app.POST("/api/tasks", ginx.FnJSON(controllers.CreateTask, tasks.Task{}))
	app.PUT("/api/tasks/:id", ginx.FnParamJSON(controllers.UpdateTask, "id", tasks.Task{}))
	app.DELETE("/api/tasks/:id", ginx.FnParam(controllers.DeleteTask, "id"))

	return app
}
