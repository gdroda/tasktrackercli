package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
	if len(inp) < 2 && inp[0] != "list" {
		invalidCommand()
		return
	}

	var t task

	switch inp[0] {

	case "add":
		tmp := capitalizeFirst(inp[1])
		var tmpId int
		if len(taskList) > 0 {
			tmpId = taskList[len(taskList)-1].Id + 1
		} else {
			tmpId = 1
		}

		tmpTime := timeStruct{int64(time.Now().Year()), int64(time.Now().Month()), int64(time.Now().Day()), int64(time.Now().Hour()), int64(time.Now().Minute()), int64(time.Now().Second())}
		getTime(tmpTime)
		t = task{tmpId, inp[1], "todo", tmpTime, tmpTime}
		taskList = append(taskList, t)
		fmt.Printf("Task '%s' added successfully (ID: %d).\n", tmp, tmpId)

	case "delete":
		ind, err := strconv.Atoi(inp[1])
		if err != nil {
			fmt.Println("hi")
		}
		for i := 0; i < len(taskList); i++ {
			if taskList[i].Id == ind {
				taskList = append(taskList[:i], taskList[i+1:]...)
			}
		}

	case "mark-in-progress":
		ind, err := strconv.Atoi(inp[1])
		if err != nil {
			fmt.Println("Could not parse to int.")
		}
		for i := 0; i < len(taskList); i++ {
			if taskList[i].Id == ind {
				taskList[i].Status = "in-progress"
			}
		}

	case "mark-done":
		ind, err := strconv.Atoi(inp[1])
		if err != nil {
			fmt.Println("Could not parse to int", err)
		}
		for i := 0; i < len(taskList); i++ {
			if taskList[i]. Id = ind {
				taskList[i].Status = "done"
			}
		}

	case "list":
		if len(inp) < 2 {
			inp = append(inp, "")
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

		fmt.Println(" ID    Description               Status")

		// For switch for list printing
		for i := 0; i < len(taskList); i++ {
			switch inp[1] {
			case "done":
				fmt.Println("Printing all done tasks.")
			case "todo":
				if taskList[i].Status == "todo" {
					listPrint(taskList[i])
				}
			case "in-progress":
				if taskList[i].Status == "in-progress" {
					listPrint(taskList[i])
				}
			case "":
				listPrint(taskList[i])
			default:
				invalidCommand()
			}
		}
		return

	default:
		invalidCommand()
	}

	app := &application{}
	app.saveList()

}

func invalidCommand() {
	fmt.Println("- Invalid command -\n Use these commands: \n  add - Adds new task\n  delete - Deletes a task\n  list - Lists all tasks\n  list done - Lists all done tasks\n  list todo - Lists all to-do tasks")
}

func listPrint(t task) {
	fmt.Printf(" %03d   %-23s   %-15s\n", t.Id, t.Descr, t.Status)
}

func getTime(t timeStruct) {
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", t.year, t.month, t.day, t.hour, t.minute, t.day)
}

func capitalizeFirst(line string) string {
	tmp := strings.Split(line, "")
	fLetter := strings.ToUpper(tmp[0])
	tmp[0] = fLetter
	wholeLine := strings.Join(tmp, "")
	return wholeLine
}

type timeStruct struct {
	year   int64
	month  int64
	day    int64
	hour   int64
	minute int64
	second int64
}

type task struct {
	Id          int
	Descr       string
	Status      string
	TimeCreated timeStruct
	TimeUpdated timeStruct
}
