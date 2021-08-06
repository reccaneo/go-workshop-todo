package todo

import (
	"todo/entities"
	"todo/repository"
)

//	Business Domain aka Usecase
type Service struct {
	repo repository.Repositoryr
}

func NewService(repo repository.Repositoryr) Service {
	return Service{
		repo: repo,
	}
}

func (srv Service) Add(task entities.Task) error {
	return srv.repo.NewTask(&task)
}

func (srv Service) Done(id uint) error {
	return srv.repo.TaskDone(id)
}

func (srv Service) List() (map[uint]*entities.Task, error) {
	return srv.repo.GetAllTask()
}
