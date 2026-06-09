package cli

import (
	"fmt"
)

func Help() {
	fmt.Println("Available commands:")
	fmt.Println("  help        		 - Show this help message")
	fmt.Println("  add <task>  		 - Add task to the manager")
	fmt.Println("  list        		 - List tasks in the manager")
	fmt.Println("  markComplete <id> - Mark a task as complete by ID")
	fmt.Println("  save [filename] 	 - Save tasks to a file (default: tasks.json)")
	fmt.Println("  load [filename] 	 - Load tasks from a file (default: tasks.json)")
	fmt.Println("  clear       		 - Clear all tasks")
}
