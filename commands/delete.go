package commands

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/LucasCABJ/task-cli/repository"
)

func DeleteTask(options []string) error {
	if len(options) < 1 {
		return errors.New("invalid parameters: ./task-cli <id>")
	}

	taskID, err := strconv.Atoi(options[0])
	if err != nil {
		return err
	}

	err = repository.DeleteTask(taskID)
	if err != nil {
		return err
	}

	fmt.Printf("Deleted task with ID %d successfully\n", taskID)

	return nil
}
