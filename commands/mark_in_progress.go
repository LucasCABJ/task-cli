package commands

import (
	"errors"
	"strconv"

	"github.com/LucasCABJ/task-cli/repository"
)

func MarkInProgress(options []string) error {
	if len(options) < 1 {
		return errors.New("Usage: ./task-cli mark-in-progress <id>")
	}
	
	taskID, err := strconv.Atoi(options[0])
	if err != nil {
		return err
	}

	task, err := repository.GetTaskByID(taskID)
	if err != nil {
		return err
	}

	task.Status = repository.StatusInProgress
	repository.UpdateTask(taskID, task)
	return nil
}