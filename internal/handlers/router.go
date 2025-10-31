package handlers

import (
	"github.com/gin-gonic/gin"
	"task_manager/internal/models"
	"task_manager/internal/config"
	"github.com/gin-contrib/cors"
	"time"
)

func CreateRouter(lib *models.TaskLib, cfg *config.Config) *gin.Engine {
	setupGin(cfg)

	router := gin.Default()
	
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5500", "http://127.0.0.1:5500", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Static("/static", "./web/static")
	router.LoadHTMLFiles("./web/index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
        "API_URL": "http://" + cfg.Server.Host + ":" + cfg.Server.Port + "/api",
    })
	})

	api := router.Group("/api") 
	{
		api.GET("/health", healthCheck())
		api.GET("/tasks", getTasks(lib))
		api.GET("/tasks/:id", getTask(lib))
		api.POST("/tasks", createTask(lib))
		api.PUT("/tasks/:id", updateTask(lib))
		api.PATCH("/tasks/:id", patchTask(lib))
		api.DELETE("/tasks/:id", deleteTask(lib))
		api.POST("/tasks/:id", endTask(lib))
	}

	return router
}

func setupGin(cfg *config.Config) {
    if cfg.Gin.Mode == "release" {
        gin.SetMode(gin.ReleaseMode)
    } else {
        gin.SetMode(gin.DebugMode)
    }
}