package main

import (
	"os"
	"fmt"
	"time"
	"math/rand"
)

//made by shushu
//2022

func main() {
	//check args
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "--help": // print help message
			fmt.Println("Create a temporary file or directory, safely, and print its name.\nUsage:\n\t--help: print this message\n\t-d: make a temporary directory")
		case "-d": //create temp directory
			name := "/tmp/tmp."+RandString()
			if err := os.Mkdir(name, 0750); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(name)
		}
	} else { //create temp file
		name := "/tmp/tmp."+RandString()
		file, err := os.Create(name)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()

		fmt.Println(name)
	}

}

//generate random string
func RandString() string {
	n := 6
	rand.Seed(time.Now().UnixNano())
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Int63() % int64(len(letterBytes))]
    }
    return string(b)
}
