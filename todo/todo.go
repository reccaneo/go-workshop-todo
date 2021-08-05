package todo

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func AddTask(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var task NewTaskTodo
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		rw.WriteHeader((http.StatusBadRequest))
		return
	}

	New(task.Task)
}

func GetTask(rw http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(rw).Encode(List()); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func DoneTask(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
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
