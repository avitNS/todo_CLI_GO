package main

import (
	"fmt"
	"os"
)

const JsonPath = "tasks.json"

var (
	cmd   Command
	tasks []Task
	err   error
)

func main() {

	if tasks, err = loadTasks(); err != nil {
		fmt.Println("Failed to load tasks: ", err)
		return
	}

	if cmd, err = parseArgs(os.Args[1:]); err != nil {
		fmt.Println(err)
		return
	}
	tasks, mutated, err := cmd.Execute(tasks)

	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println("Available commands:\nadd <title>\ndone <id>\nremove <id>\nlist")
		return
	}

	if mutated {
		if err = saveTasks(tasks); err != nil {
			fmt.Println("Failed to save tasks: ", err)
		}

	}
}
