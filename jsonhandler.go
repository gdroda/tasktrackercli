package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func (app *application) saveList() {
	userprofile := os.Getenv("USERPROFILE")
	file, err := os.Create(path.Join(userprofile, "tasks.json"))
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
}

func (app *application) loadList() {
	userprofile := os.Getenv("USERPROFILE")
	file, err := os.Open(path.Join(userprofile, "tasks.json"))
	if file == nil {
		app.saveList()
		return
	}
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
