package todo

import (
	"net/http"
	"strconv"
	"todo/entities"
	"todo/repository"

	"github.com/gin-gonic/gin"
)

type NewTaskTodo struct {
	Task string `json:"task"`
}
type App struct {
	repo repository.Repositoryr
}

func NewApp(repo repository.Repositoryr) *App {
	return &App{
		repo: repo,
	}
}

func (app App) AddTask(c *gin.Context) {
	var task NewTaskTodo
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := app.repo.NewTask(&entities.Task{
		Title: task.Task,
		Done:  false,
	}); err != nil {
		c.JSON(http.StatusInsufficientStorage, nil)
		return
	}
}

func (app App) GetTask(c *gin.Context) {
	result, err := app.repo.GetAllTask()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (app App) DoneTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if err = app.repo.TaskDone(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Status(http.StatusOK)
}
