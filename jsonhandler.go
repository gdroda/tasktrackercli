package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func (app *application) saveList() {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	b, err := json.MarshalIndent(taskList, "", "")
	if err != nil {
		fmt.Println("Error marshaling file:", err)
		return
	}

	_, err = file.Write(b)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
	fmt.Println("Data saved")
}

func (app *application) loadList() {
	file, err := os.Open("tasks.json")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&taskList)
	if err != nil {
		fmt.Println("Error decoding JSON", err)
		return
	}
}
