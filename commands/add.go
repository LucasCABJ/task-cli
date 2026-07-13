package commands

import (
	"errors"
	"fmt"
	"time"

	"github.com/LucasCABJ/task-cli/repository"
)

// AddTask handles adding a to-do task
func AddTask(options []string) error {
	if len(options) == 0 {
		return errors.New("missing task description. Usage: ./task-cli add \"to-do description\"")
	}
	currentTime := time.Now()
	description := options[0]
	newTask := repository.Task{
		Description: description,
		Status:      repository.StatusToDo,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}
	taskID, err := repository.AddTask(newTask)
	if err != nil {
		return err
	}
	fmt.Printf("Created task successfully with ID: %d\n", taskID)
	return nil
}
