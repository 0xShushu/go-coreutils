package main

import (
	"os"
	"fmt"
	"strings"
)

//made by 0xShushu
//2022

func main() {
	//get args
	args := strings.Join(os.Args[1:], "")
	//help messages
	if args == "--help" {
		fmt.Println("Usage: yes [STRING]...\nRepeatedly output a line with all specified STRING(s), or 'y'.")
	} else if args == "" { //print y if args are empty
		for {
			fmt.Println("y")
		}
	} else {
		for {
			fmt.Println(args)
		}
	}
}