package main

import (
	"fmt"
	"os"
)

func executeCommand() {
	command := os.Args[1]
	switch command {
	case "create":
		seed := CreateAccount()
		bucket := createBucket(seed)
		saveToFile(bucket)
	case "start":
		bucket := readFromFile()
		fmt.Println(bucket)
	}
}
