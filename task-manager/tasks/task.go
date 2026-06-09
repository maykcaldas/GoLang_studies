package tasks

import (
	"fmt"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (t Task) String() string {
	return fmt.Sprintf("\tID: %d,\n\tTitle: %s,\n\tDescription: %s,\n\tStatus: %s", t.ID, t.Title, t.Description, t.Status)
}

func (t *Task) Complete() {
	t.Status = "complete"
}
