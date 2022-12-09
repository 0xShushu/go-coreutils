package main

import (
	"fmt"
	"os"
	"os/user"
)

//made by shushu
//2022

func main() {
	//check for help message
	if len(os.Args) >= 2 && os.Args[1] == "--help" {
		fmt.Println("Print current user")
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