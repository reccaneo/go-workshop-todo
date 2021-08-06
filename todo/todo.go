package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var index int
var tasks map[int]*Task = make(map[int]*Task)

type Task struct {
	gorm.Model
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type NewTaskTodo struct {
	Task string `json:"task"`
}

type Inserter interface {
	Insert(interface{}) error
}

type App struct {
	// db *gorm.DB
	db Inserter
}

type GormInsert struct {
	db *gorm.DB
}

func (insert GormInsert) Insert(v interface{}) error {
	return errors.WithMessage(insert.db.Create(v).Error, "gorm insert")
}

func NewMemInsert() MemInsert {
	return MemInsert{
		db: tasks,
	}
}

type MemInsert struct {
	db map[int]*Task
}

func (insert MemInsert) Insert(v interface{}) error {
	if cache, ok := v.(*Task); ok {
		defer func() {
			index++
		}()
		insert.db[index] = cache
	}

	return nil
}

func NewApp(db Inserter) *App {
	return &App{
		db: db,
	}
}

func (app App) AddTask(c *gin.Context) {
	// fdb, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fdb.AutoMigrate(&Task{})

	var task NewTaskTodo
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// New(task.Task)
	// app.db.Create(&Task{
	// 	Title: task.Task,
	// 	Done:  false,
	// })
	app.db.Insert(&Task{
		Title: task.Task,
		Done:  false,
	})
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
