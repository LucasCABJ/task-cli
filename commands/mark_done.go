package commands

import (
	"errors"
	"strconv"

	"github.com/LucasCABJ/task-cli/repository"
)

func MarkDone(options []string) error {
	if len(options) < 1 {
		return errors.New("Usage: ./task-cli mark-done <id>")
	}

	taskID, err := strconv.Atoi(options[0])
	if err != nil {
		return err
	}

	task, err := repository.GetTaskByID(taskID)
	if err != nil {
		return err
	}

	task.Status = repository.StatusDone
	repository.UpdateTask(taskID, task)
	return nil
}
