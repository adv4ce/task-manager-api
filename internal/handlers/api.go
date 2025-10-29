package handlers

import (
	"net/http"
	"strconv"
	"task_manager/internal/models"
	"task_manager/internal/services"

	"github.com/gin-gonic/gin"
)

func healthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, services.APIResponse{
			Status: "healthy",
		})
	}
}

func createTask(lib *models.TaskLib) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var cTask services.Task

		if err := ctx.ShouldBindJSON(&cTask); err != nil {
			ctx.JSON(http.StatusBadRequest, services.APIResponse{
				Status: "error",
				Error:  err.Error(),
			})
			return
		}

		lib.Create(cTask.Title, cTask.Description, cTask.Priority)

		ctx.JSON(http.StatusOK, services.APIResponse{
			Status:  "successful",
			Data:    cTask,
			Message: "Task created",
		})
	}
}

func getTask(lib *models.TaskLib) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			ctx.JSON(http.StatusNotFound, services.APIResponse{
				Status:  "error",
				Message: "Invalid task ID",
				Error:   err.Error(),
			})
			return
		}

		task, err := lib.Get(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, services.APIResponse{
				Status: "error",
				Error:  err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, services.APIResponse{
			Status: "successful",
			Data:   task,
		})
	}
}

func updateTask(lib *models.TaskLib) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var updTask services.Task

		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, services.APIResponse{
				Status:  "error",
				Message: "Invalid task ID",
				Error:   err.Error(),
			})
			return
		}
		
		if err := ctx.ShouldBindJSON(&updTask); err != nil {
			ctx.JSON(http.StatusBadRequest, services.APIResponse{
				Status: "error",
				Error:  err.Error(),
			})
			return
		}

		err = lib.Update(id, updTask.Title, updTask.Description, updTask.Priority)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, services.APIResponse{
				Status: "error",
				Error:  err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, services.APIResponse{
			Status:  "successful",
			Data:    updTask,
			Message: "Updated successfully",
		})
	}
}

func getTasks(lib *models.TaskLib) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, services.APIResponse{
			Status: "successful",
			Data:   lib,
		})
	}
}

func deleteTask(lib *models.TaskLib) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, services.APIResponse{
				Status:  "error",
				Message: "Invalid task ID",
				Error:   err.Error(),
			})
			return
		}

		err = lib.Delete(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, services.APIResponse{
				Status:  "error",
				Message: "Invalid task ID",
				Error:   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, services.APIResponse{
			Status: "successful",
			Data: gin.H{
				"deletedID": id,
			},
		})
	}
}

func patchTask(lib *models.TaskLib) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pthT services.PatchTask

		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, services.APIResponse{
				Status:  "error",
				Message: "Invalid task ID",
				Error:   err.Error(),
			})
			return
		}

		if err := ctx.ShouldBindJSON(&pthT); err != nil {
			ctx.JSON(http.StatusBadRequest, services.APIResponse{
				Status: "error",
				Error:  err.Error(),
			})
			return
		}

		chgVals := make(map[string]string)
		services.SturcToMap(&pthT, &chgVals)

		err = lib.Patch(id, chgVals)

		if err != nil {
			ctx.JSON(http.StatusNotFound, services.APIResponse{
				Status: "error",
				Error:  err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, services.APIResponse{
			Status: "successful",
			Data: gin.H{
				"changeVals": pthT,
			},
			Message: "Updated successfully",
		})
	}
}

func endTask(lib *models.TaskLib) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, _ := strconv.Atoi(idStr)

		err := lib.End(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, services.APIResponse{
				Status: "error",
				Error: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, services.APIResponse{
			Status: "successful",
			Data: gin.H{
				"id": id,
			},
			Message: "Task closed",
		})
	}
}