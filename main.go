package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LucasCABJ/task-cli/commands"
)

type HandleOptions func([]string) error

var commandsMap = map[string]HandleOptions{
	"add":    commands.AddTask,
	"list":   commands.ListTasks,
	"update": commands.UpdateTask,
	"delete": commands.DeleteTask,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid args, please use: ./task-cli <command> <params>")
		os.Exit(1)
	}

	command := os.Args[1]
	options := os.Args[2:]

	err := commandsMap[command](options)
	if err != nil {
		log.Fatal(err)
	}
}
