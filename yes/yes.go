package main

import (
	"fmt"
	"os"
	"strings"
)

//made by 0xShushu
//2022

func main() {
	//check args
	if len(os.Args) >= 2 {
		if os.Args[1] == "--help" { //print help if args equals to --help
			fmt.Println("Usage: yes [STRING]...\nRepeatedly output a line with all specified STRING(s), or 'y'.")
		} else { //else print args
			for { 
				fmt.Println(strings.Join(os.Args[1:], " "))
			}
		}
	} else { //if args are empy print y
		for {
			fmt.Println("y")
		}
	}
}