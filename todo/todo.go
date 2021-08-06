package todo

import (
	"net/http"
	"todo/entities"
	"todo/repository"

	"github.com/gin-gonic/gin"
)

type NewTaskTodo struct {
	Task string `json:"task"`
}
type App struct {
	// db *gorm.DB
	// db   Inserter
	repo repository.Repositoryr
}

func NewApp(repo repository.Repositoryr) *App {
	return &App{
		repo: repo,
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
	// app.db.Insert(&Task{
	// 	Title: task.Task,
	// 	Done:  false,
	// })
	app.repo.NewTask(&entities.Task{
		Title: task.Task,
		Done:  false,
	})
}

func (app App) GetTask(c *gin.Context) {
	c.JSON(http.StatusOK, app.repo.GetAllTask())
}

func DoneTask(c *gin.Context) {
	// idStr := c.Param("id")
	// id, err := strconv.Atoi(idStr)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, nil)
	// 	return
	// }

	// tasks[id].Done = true
}

func List() map[int]*entities.Task {
	return nil
}

// func New(task string) {
// 	defer func() {
// 		index++
// 	}()
// 	tasks[index] = &Task{
// 		Title: task,
// 		Done:  false,
// 	}
// }
