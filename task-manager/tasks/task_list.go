package tasks

import (
	"encoding/json"
	"os"
)

type TaskList []Task

func (tl TaskList) String() string {
	result := ""
	for _, task := range tl {
		result += task.String() + "\n"
	}
	return result
}

func (tl *TaskList) Add(task Task) TaskList {
	*tl = append(*tl, task)
	return *tl
}

func (tl *TaskList) MarkComplete(id int) error {
	for i, task := range *tl {
		if task.ID == id {
			(*tl)[i].Complete()
			return nil
		}
	}
	return TaskNotFoundError{ID: id}
}

func (tl *TaskList) Save(filename string) error {
	data, err := json.Marshal(*tl)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (tl *TaskList) Load(filename string) (TaskList, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return TaskList{}, nil // Return empty list if file doesn't exist
		}
		return TaskList{}, err
	}
	tasks := TaskList{}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}
