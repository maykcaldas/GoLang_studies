package tasks

import "fmt"

type TaskNotFoundError struct {
	ID int
}

func (err TaskNotFoundError) Error() string {
	return fmt.Sprintf("Task with ID %d not found", err.ID)
}
