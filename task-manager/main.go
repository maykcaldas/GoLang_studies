package main

import (
	"fmt"
	"os"
	"strconv"
	"task-manager/cli"
	"task-manager/tasks"
)

func main() {
	taskList := tasks.TaskList{}
	taskList, err := taskList.Load("tasks.json")
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		os.Exit(1)
	}
	lastIDCreated := len(taskList)

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input>")
		cli.Help()
		os.Exit(2)
	}
	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Title is required for adding a task")
			os.Exit(2)
		}
		title := os.Args[2]
		description := ""
		if len(os.Args) > 3 {
			description = os.Args[3]
		}
		task := tasks.Task{
			ID:          lastIDCreated + 1,
			Title:       title,
			Description: description,
			Status:      "pending",
		}

		fmt.Println("Adding item...")
		taskList = taskList.Add(task)
		fmt.Println(taskList)
		taskList.Save("tasks.json")
	case "list":
		fmt.Println("Listing items...")
		fmt.Println(taskList)
	case "help":
		cli.Help()
	case "markComplete":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			os.Exit(2)
		}
		err = taskList.MarkComplete(id)
		if err != nil {
			fmt.Println("Error marking task as complete:", err)
			os.Exit(1)
		}
		fmt.Printf("Marked task %d as complete.\n", id)
		err = taskList.Save("tasks.json")
		if err != nil {
			fmt.Println("Error saving tasks:", err)
			os.Exit(1)
		}
		fmt.Println(taskList)
	case "save":
		filename := "tasks.json"
		if len(os.Args) > 3 {
			filename = os.Args[2]
		}
		fmt.Println("Saving items...")
		err := taskList.Save(filename)
		if err != nil {
			fmt.Println("Error saving items.")
			os.Exit(1)
		} else {
			fmt.Println("Items saved successfully.")
		}
	case "load":
		filename := "tasks.json"
		if len(os.Args) > 3 {
			filename = os.Args[2]
		}
		fmt.Println("Loading items...")
		taskList, err := taskList.Load(filename)
		if err != nil {
			fmt.Println("Error loading items.")
			os.Exit(1)
		} else {
			fmt.Println("Items loaded successfully.")
		}
		fmt.Println(taskList)
	case "clear":
		taskList = tasks.TaskList{}
		taskList.Save("tasks.json")
		fmt.Println("Cleared all items.")
	default:
		fmt.Printf("Unknown command: %s\n", command)
		cli.Help()
		os.Exit(2)
	}
}
