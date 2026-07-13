package commands

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/LucasCABJ/task-cli/repository"
)

func UpdateTask(options []string) error {
	if len(options) < 2 {
		return errors.New("invalid parameters: ./task-cli update <id> <new-description>")
	}

	taskID, err := strconv.Atoi(options[0])
	if err != nil {
		return err
	}
	description := options[1]

	task, err := repository.GetTaskByID(taskID)
	if err != nil {
		return err
	}

	task.Description = description
	err = repository.UpdateTask(taskID, task)
	if err != nil {
		return err
	}

	fmt.Printf("\nUpdated task %d successfully", taskID)

	return nil
}
