package controllers

import (
	"net/http"

	"github.com/bingoohuang/go-rest-template/internal/pkg/ginx"
	"github.com/bingoohuang/go-rest-template/internal/pkg/models/tasks"
	"github.com/bingoohuang/go-rest-template/internal/pkg/persist"
	"github.com/gin-gonic/gin"
)

// GetTaskById godoc
// @Summary Retrieves task based on given ID
// @Description get Task by ID
// @Produce json
// @Param id path integer true "Task ID"
// @Success 200 {object} tasks.Task
// @Router /api/tasks/{id} [get]
// @Security Authorization Token
func GetTaskById(c *gin.Context, id string) ginx.Render {
	task, err := persist.GetTaskRepository().Get(id)
	if err != nil {
		return ginx.New404Error("task not found", err)
	}

	return ginx.JSON(task)
}

// GetTasks godoc
// @Summary Retrieves tasks based on query
// @Description Get Tasks
// @Produce json
// @Param taskname query string false "Taskname"
// @Param firstname query string false "Firstname"
// @Param lastname query string false "Lastname"
// @Success 200 {array} []tasks.Task
// @Router /api/tasks [get]
// @Security Authorization Token
func GetTasks(c *gin.Context, bind interface{}) ginx.Render {
	s := persist.GetTaskRepository()
	q := bind.(tasks.Task)

	t, err := s.Query(&q)
	if err != nil {
		return ginx.New404Error("tasks not found", err)
	}

	return ginx.JSON(t)
}

func CreateTask(c *gin.Context, bindJSON interface{}) ginx.Render {
	s := persist.GetTaskRepository()
	taskInput := bindJSON.(tasks.Task)
	if err := s.Add(&taskInput); err != nil {
		return ginx.New400Error("", err)
	}

	return ginx.StatusJSON(http.StatusCreated, taskInput)
}

func UpdateTask(c *gin.Context, id string, bindJSON interface{}) ginx.Render {
	s := persist.GetTaskRepository()
	taskInput := bindJSON.(tasks.Task)
	if _, err := s.Get(id); err != nil {
		return ginx.New404Error("tasks not found", err)
	}

	if err := s.Update(&taskInput); err != nil {
		return ginx.New404Error("", err)
	}

	return ginx.JSON(taskInput)
}

func DeleteTask(c *gin.Context, id string) ginx.Render {
	s := persist.GetTaskRepository()
	task, err := s.Get(id)
	if err != nil {
		return ginx.New404Error("task not found", err)
	}

	if err := s.Delete(task); err != nil {
		return ginx.New404Error("", err)
	}

	return ginx.StatusJSON(http.StatusNoContent, "")
}
