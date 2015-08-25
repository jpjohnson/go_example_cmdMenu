package main

import (
	"fmt"
	"os"
	exec "os/exec"
)

func menu() {
	fmt.Println("\n-------- Menu ------")
	fmt.Println("Select choice and press 'Enter'")
	fmt.Println("1: System Date")
	fmt.Println("2: list Dir")
	fmt.Println("x: Exit")
	fmt.Println("---------------------")
}

func main() {
	var input string

	// loop for input
	for {
		// show menu
		menu()
		// get input
		fmt.Scanln(&input)

		// process choice
		switch input {
		case "1":
			showDate()
		case "2":
			fmt.Println(runCmd("ls"))
		case "x":
			os.Exit(0)
		}
	}
}
func showDate() {
	fmt.Println(runCmd("date"))
}
func runCmd(cmd string) string {
	out, err := exec.Command(cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(out)
}
