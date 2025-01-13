package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//var nFlag = flag.Int("n", 1234, "help message for flag n")
	//flag.Parse()
	//fmt.Println(*nFlag)
	for {
		getInput()
	}

}

func getInput() string {
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	line = strings.ToLower(strings.Trim(line, " \r\n"))
	if err != nil {
		log.Fatal(err)
	}

	slicedString := strings.SplitN(line, " ", 2)
	command := slicedString[0]

	switch command {
	case "end":
		fmt.Println("Terminating Application")
		os.Exit(0)
	case "add":
		tmp := capitalizeFirst(slicedString[1])
		fmt.Printf("Task '%s' added successfully (ID: 2).\n", tmp)
	case "delete":
		tmp := capitalizeFirst(slicedString[1])
		fmt.Printf("'%s' has been deleted from the list.\n", tmp)
	case "list":
		fmt.Println("Printing all tasks.")
	default:
		return line
	}

	return line
}

func capitalizeFirst(line string) string {
	tmp := strings.Split(line, "")
	fLetter := strings.ToUpper(tmp[0])
	tmp[0] = fLetter
	wholeLine := strings.Join(tmp, "")
	return wholeLine
}
