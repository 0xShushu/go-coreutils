package main

import (
	"fmt"
	"os"
	"os/user"
)

//made by shushu
//2022

func main() {
	//get args
	args := os.Args[1]
	if args == "--help" { //check for help
		fmt.Println("Prints the current username")
		os.Exit(0)
	}

	//get user
	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//print user
	fmt.Println(user.Username)
}