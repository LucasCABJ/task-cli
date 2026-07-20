package commands

import (
	"fmt"
	"time"

	"github.com/LucasCABJ/task-cli/repository"
)

func ListTasks(options []string) error {
	tasks, err := repository.GetTasks()
	if err != nil {
		return err
	}
	if len(options) == 0 {
		printTasks(tasks.Tasks)
		return nil
	}
	switch options[0] {
	case "todo":
		printTasksByStatus(tasks.Tasks, repository.StatusToDo)
	case "in-progress":
		printTasksByStatus(tasks.Tasks, repository.StatusInProgress)
	case "done":
		printTasksByStatus(tasks.Tasks, repository.StatusDone)
	default:
		return fmt.Errorf("invalid status: %s", options[0])
	}
	return nil
}

func printTasksByStatus(tasks []repository.Task, status repository.TaskStatus) {
	var filteredTasks []repository.Task
	for _, task := range tasks {
		if task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}
	printTasks(filteredTasks)
}

func printTasks(tasks []repository.Task) {
	totalTasks := len(tasks)
	if totalTasks == 0 {
		fmt.Println("\nYou have 0 tasks!")
		fmt.Println()
		return
	}
	fmt.Println("\n============================================================")
	fmt.Printf("Total Tasks: %d\n", totalTasks)
	fmt.Println("============================================================")
	for _, task := range tasks {
		printTask(&task)
	}
	fmt.Println("============================================================")
	fmt.Println()
}

func printTask(task *repository.Task) {
	fmt.Printf("ID: %d - Description: %s - Status: %s - Created At: %s - Updated At : %s\n",
		task.ID, task.Description, task.Status, task.CreatedAt.Format(time.DateTime), task.UpdatedAt.Format(time.DateTime))
}
