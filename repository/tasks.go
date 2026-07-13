package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type TaskStatus string

const (
	StatusInProgress TaskStatus = "IN_PROGRESS"
	StatusDone       TaskStatus = "DONE"
	StatusToDo       TaskStatus = "TO_DO"
)

type Tasks struct {
	Total int    `json:"total"`
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

const (
	filename = "./tasks.json"
)

func AddTask(task Task) (int, error) {
	tasks, err := findTasks()
	if err != nil {
		return 0, err
	}
	newTaskID := tasks.Total + 1
	task.ID = newTaskID
	tasks.Total++
	tasks.Tasks = append(tasks.Tasks, task)
	err = saveTasks(tasks)
	if err != nil {
		return 0, err
	}
	return newTaskID, nil
}

func UpdateTask(taskID int, task Task) error {
	tasks, err := findTasks()
	if err != nil {
		return err
	}
	for i, v := range tasks.Tasks {
		if v.ID == taskID {
			tasks.Tasks[i] = task
			break
		}
	}
	err = saveTasks(tasks)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTask(taskID int) error {
	tasks, err := findTasks()
	if err != nil {
		return err
	}
	var filteredTasks []Task
	for _, v := range tasks.Tasks {
		if v.ID == taskID {
			continue
		}
		filteredTasks = append(filteredTasks, v)
	}
	if len(tasks.Tasks) == len(filteredTasks) {
		return errors.New("failed to delete specified id - not found")
	}
	tasks.Tasks = filteredTasks
	err = saveTasks(tasks)
	if err != nil {
		return err
	}
	return nil
}

func GetTasks() (*Tasks, error) {
	return findTasks()
}

func GetTaskByID(taskID int) (Task, error) {
	tasks, err := findTasks()
	if err != nil {
		return Task{}, err
	}
	for _, v := range tasks.Tasks {
		if v.ID == taskID {
			return v, err
		}
	}
	return Task{}, fmt.Errorf("task %d not found", taskID)
}

func findTasks() (*Tasks, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return &Tasks{}, nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var tasks Tasks
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return nil, fmt.Errorf("error decoding json: %w", err)
	}

	return &tasks, nil
}

func saveTasks(tasks *Tasks) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		return err
	}
	return nil
}
