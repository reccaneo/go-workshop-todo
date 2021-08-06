package repository

import "todo/entities"

type Repositoryr interface {
	NewTask(*entities.Task) error
	GetAllTask() map[int]*entities.Task
}
