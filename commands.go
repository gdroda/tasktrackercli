package main

import (
	"fmt"
	"strconv"
	"time"
)

// Add Task Function
func (app *application) addTask(inp []string) {
	var t task
	var tmpId int
	if len(inp[1]) > 25 {
		fmt.Println("Error: Description Length over maximum length (25)")
		fmt.Println("Please try again.")
		return
	}
	if len(taskList) > 0 {
		tmpId = taskList[len(taskList)-1].Id + 1
	} else {
		tmpId = 1
	}
	t = task{tmpId, inp[1], "To-Do", time.Now(), time.Now()}
	taskList = append(taskList, t)
	fmt.Printf("Task '%s' added successfully (ID: %d).\n", inp[1], tmpId)
}

// Delete Task Function
func (app *application) deleteTask(inp []string) {
	var found bool = false
	ind, err := strconv.Atoi(inp[1])
	if err != nil {
		fmt.Println("Error: Please input a valid ID integer")
		return
	}
	for i := 0; i < len(taskList); i++ {
		if taskList[i].Id == ind {
			taskList = append(taskList[:i], taskList[i+1:]...)
			found = true
		}
	}
	if !found {
		fmt.Println("Error: ID not found, please enter a valid ID integer")
		return
	} else {
		fmt.Printf("Task with ID: %d has been deleted", ind)
	}
}

// Mark Function
func (app *application) markTask(inp []string) {
	switch inp[0] {
	// Mark To Do
	case "mark-todo":

		var found bool = false
		ind, err := strconv.Atoi(inp[1])
		if err != nil {
			fmt.Println("Error: Please input a valid ID integer")
			return
		}
		for i := 0; i < len(taskList); i++ {
			if taskList[i].Id == ind {
				taskList[i].Status = "To-Do"
				taskList[i].TimeUpdated = time.Now()
				found = true
			}
		}
		if !found {
			fmt.Println("Error: ID not found, please enter a valid ID integer")
			return
		} else {
			fmt.Printf("Task with ID: %d has changed status to 'To-Do'", ind)
		}

	// Mark In Progress
	case "mark-in-progress":
		var found bool = false
		ind, err := strconv.Atoi(inp[1])
		if err != nil {
			fmt.Println("Error: Please input a valid ID integer")
			return
		}
		for i := 0; i < len(taskList); i++ {
			if taskList[i].Id == ind {
				taskList[i].Status = "In Progress"
				taskList[i].TimeUpdated = time.Now()
				found = true
			}
		}
		if !found {
			fmt.Println("Error: ID not found, please enter a valid ID integer")
			return
		} else {
			fmt.Printf("Task with ID: %d has changed status to 'In Progress'", ind)
		}

	// Mark Done
	case "mark-done":
		var found bool = false
		ind, err := strconv.Atoi(inp[1])
		if err != nil {
			fmt.Println("Error: Please input a valid ID integer")
			return
		}
		for i := 0; i < len(taskList); i++ {
			if taskList[i].Id == ind {
				taskList[i].Status = "Done"
				taskList[i].TimeUpdated = time.Now()
				found = true
			}
		}
		if !found {
			fmt.Println("Error: ID not found, please enter a valid ID integer")
			return
		} else {
			fmt.Printf("Task with ID: %d has changed status to 'Done'", ind)
		}
	}
}

// List Function
func (app *application) listTask(inp []string) {
	if len(inp) < 2 {
		inp = append(inp, "") //Adding space for default list print
	}

	// Switch for print title
	switch inp[1] {
	case "done":
		fmt.Println("Printing all done tasks.")
	case "todo":
		fmt.Println("- Printing TO-DO Tasks -")
	case "in-progress":
		fmt.Println("- Printing In Progress Tasks -")
	case "":
		fmt.Println("- Printing All Tasks -")
	}

	// List titles
	fmt.Println(" ID    Description                 Status            Created                      Updated")

	// For switch for list printing
	for i := 0; i < len(taskList); i++ {
		switch inp[1] {
		case "done":
			if taskList[i].Status == "Done" {
				listPrint(taskList[i])
			}
		case "todo":
			if taskList[i].Status == "To-Do" {
				listPrint(taskList[i])
			}
		case "in-progress":
			if taskList[i].Status == "In Progress" {
				listPrint(taskList[i])
			}
		case "":
			listPrint(taskList[i])
		default:
			invalidCommand()
		}
	}
}
