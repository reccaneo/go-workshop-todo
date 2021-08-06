package repository

import (
	"errors"
	"sync"
	"todo/entities"
)

var index uint
var tasks map[uint]*entities.Task = make(map[uint]*entities.Task)
var mutex = sync.Mutex{}

type MemoryRepository struct {
	db map[uint]*entities.Task
}

func NewMemoryRepository() MemoryRepository {
	return MemoryRepository{
		db: tasks,
	}
}

func (repo MemoryRepository) NewTask(v interface{}) error {
	if cache, ok := v.(*entities.Task); ok {
		mutex.Lock()
		defer func() {
			index++
			mutex.Unlock()
		}()
		repo.db[index] = cache
	}

	return nil
}

func (repo MemoryRepository) TaskDone(index uint) error {
	if tasks == nil {
		return errors.New("nil ref")
	}
	if tasks[index] == nil {
		return errors.New("nil ref")
	}
	mutex.Lock()
	defer func() {
		mutex.Unlock()
	}()

	tasks[index].Done = true

	return nil
}

func (repo MemoryRepository) GetAllTask() (map[uint]*entities.Task, error) {
	return repo.db, nil
}
