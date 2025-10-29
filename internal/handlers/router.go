package handlers

import (
	"github.com/gin-gonic/gin"
	"task_manager/internal/models"
	"task_manager/internal/config"
)

func CreateRouter(lib *models.TaskLib, cfg *config.Config) *gin.Engine {
	setupGin(cfg)

	router := gin.Default()
	
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