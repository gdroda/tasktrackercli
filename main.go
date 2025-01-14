package main

import (
	"fmt"
	"os"
	"time"
)

type application struct {
}

var taskList []task

func main() {
	app := &application{}

	app.loadList()

	arg := os.Args[1:]
	getInput(arg)

}

func getInput(inp []string) {
	if len(inp) < 1 {
		invalidCommand()
		return
	} else if len(inp) < 2 && inp[0] != "list" {
		invalidCommand()
		return
	} else if len(inp) > 2 {
		invalidCommand()
		return
	}

	app := &application{}

	switch inp[0] {

	// Add command
	case "add":
		app.addTask(inp)

	// Delete command
	case "delete":
		app.deleteTask(inp)

	// Mark To Do
	case "mark-todo", "mark-in-progress", "mark-done":
		app.markTask(inp)

	// List Command
	case "list":
		app.listTask(inp)

	// A just in case invalid command message
	default:
		invalidCommand()
	}

	app.saveList()

}

func invalidCommand() {
	fmt.Println("   - Invalid Command -\n   Available commands: \n\n   add \"task_description\" - Creates a New Task")
	fmt.Println("   delete task_id - Deletes a task\n\n   mark-todo task_id - Changes Task Status to 'To-Do'")
	fmt.Println("   mark-in-progress task_id - Changes Task Status to 'In Progress'\n   mark-done task_id - Changes Task Status to 'Done'")
	fmt.Println("\n   list - Lists all tasks")
	fmt.Println("   list done - Lists all 'Done' tasks\n   list todo - Lists all 'To-Do' tasks")
	fmt.Println("   list in-progress - Lists all 'In Progress' tasks")
}

func listPrint(t task) {
	fmt.Printf(" %03d   %-25s   %-15s   ", t.Id, t.Descr, t.Status)
	fmt.Printf("%02d-%02d-%d %02d:%02d:%02d          ", t.TimeCreated.Day(), t.TimeCreated.Month(), t.TimeCreated.Year(), t.TimeCreated.Hour(), t.TimeCreated.Minute(), t.TimeCreated.Second())
	fmt.Printf("%02d-%02d-%d %02d:%02d:%02d\n", t.TimeUpdated.Day(), t.TimeUpdated.Month(), t.TimeUpdated.Year(), t.TimeUpdated.Hour(), t.TimeUpdated.Minute(), t.TimeUpdated.Second())

}

type task struct {
	Id          int
	Descr       string
	Status      string
	TimeCreated time.Time
	TimeUpdated time.Time
}
