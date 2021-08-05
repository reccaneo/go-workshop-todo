package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var index int
var tasks map[int]*Task = make(map[int]*Task)

type Task struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type NewTaskTodo struct {
	Task string `json:"task"`
}

func AddTask(c *gin.Context) {

	var task NewTaskTodo
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	New(task.Task)
}

func GetTask(c *gin.Context) {
	c.JSON(http.StatusOK, List())
}

func DoneTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	tasks[id].Done = true
}

func List() map[int]*Task {
	return tasks
}

func New(task string) {
	defer func() {
		index++
	}()
	tasks[index] = &Task{
		Title: task,
		Done:  false,
	}
}
