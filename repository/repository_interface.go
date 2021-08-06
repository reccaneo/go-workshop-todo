package repository

import "todo/entities"

type Repositoryr interface {
	NewTask(interface{}) error
	GetAllTask() (map[uint]*entities.Task, error)
	TaskDone(uint) error
}
