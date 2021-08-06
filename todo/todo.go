package todo

import (
	"net/http"
	"strconv"
	"todo/entities"

	"github.com/gin-gonic/gin"
)

type Servicer interface {
	Add(entities.Task) error
	Done(id uint) error
	List() (map[uint]*entities.Task, error)
}

type App struct {
	// repo repository.Repositoryr
	srv Servicer
}

func NewApp(srv Servicer) *App {
	return &App{
		srv: srv,
	}
}

func (app App) AddTask(c *gin.Context) {
	var task NewTaskTodo
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	if err := app.srv.Add(entities.Task{
		Title: task.Task,
		Done:  false,
	}); err != nil {
		c.JSON(http.StatusInsufficientStorage, nil)
		return
	}

	// if err := app.repo.NewTask(&entities.Task{
	// 	Title: task.Task,
	// 	Done:  false,
	// }); err != nil {
	// 	c.JSON(http.StatusInsufficientStorage, nil)
	// 	return
	// }
}

func (app App) GetTask(c *gin.Context) {
	result, err := app.srv.List()
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

	if err = app.srv.Done(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Status(http.StatusOK)
}
