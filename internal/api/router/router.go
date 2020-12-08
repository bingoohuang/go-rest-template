package router

import (
	"fmt"
	"io"
	"os"

	"github.com/bingoohuang/go-rest-template/internal/api/controllers"
	"github.com/bingoohuang/go-rest-template/internal/api/middlewares"
	"github.com/bingoohuang/go-rest-template/internal/pkg/ginx"
	"github.com/bingoohuang/go-rest-template/internal/pkg/models/tasks"
	models "github.com/bingoohuang/go-rest-template/internal/pkg/models/users"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	app := gin.New()

	// Logging to a file.
	f, _ := os.Create("log/api.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(p gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			p.ClientIP, p.TimeStamp.Format("2006-01-02 15:04:05 -0700"),
			p.Method, p.Path, p.Request.Proto, p.StatusCode, p.Latency, p.Request.UserAgent(), p.ErrorMessage)
	}))
	app.Use(gin.Recovery(), middlewares.CORS())
	app.NoRoute(middlewares.NoRouteHandler())

	// Routes
	// ================== Login Routes
	app.POST("/api/login", ginx.BindJSON(controllers.Login, controllers.LoginInput{}))
	// ================== Docs Routes
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// ================== User Routes
	app.GET("/api/users", ginx.Bind(controllers.GetUsers, models.User{}))
	app.GET("/api/users/:id", ginx.Param(controllers.GetUserById, "id"))
	app.POST("/api/users", ginx.BindJSON(controllers.CreateUser, controllers.UserInput{}))
	app.PUT("/api/users/:id", ginx.ParamBindJSON(controllers.UpdateUser, "id", controllers.UserInput{}))
	app.DELETE("/api/users/:id", ginx.Param(controllers.DeleteUser, "id"))
	// ================== Tasks Routes
	app.GET("/api/tasks/:id", ginx.Param(controllers.GetTaskById, "id"))
	app.GET("/api/tasks", ginx.Bind(controllers.GetTasks, tasks.Task{}))
	app.POST("/api/tasks", ginx.BindJSON(controllers.CreateTask, tasks.Task{}))
	app.PUT("/api/tasks/:id", ginx.ParamBindJSON(controllers.UpdateTask, "id", tasks.Task{}))
	app.DELETE("/api/tasks/:id", ginx.Param(controllers.DeleteTask, "id"))

	return app
}
