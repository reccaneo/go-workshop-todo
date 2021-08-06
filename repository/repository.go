package repository

import (
	"todo/entities"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var index int
var tasks map[int]*entities.Task = make(map[int]*entities.Task)

// func NewGormInsert() GormInsert {
// 	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	db.AutoMigrate(&entities.Task{})

// 	return GormInsert{
// 		db: db,
// 	}
// }

type Repository struct {
	inserter Inserter
}

func NewRepotitory(inserter Inserter) Repository {
	return Repository{
		inserter: inserter,
	}
}

type Inserter interface {
	Insert(interface{}) error
}

func (repo Repository) NewTask(task *entities.Task) error {
	return repo.inserter.Insert(task)
}

func (repo Repository) GetAllTask() map[int]*entities.Task {
	return tasks
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
	db map[int]*entities.Task
}

func (insert MemInsert) Insert(v interface{}) error {
	if cache, ok := v.(*entities.Task); ok {
		defer func() {
			index++
		}()
		insert.db[index] = cache
	}

	return nil
}
