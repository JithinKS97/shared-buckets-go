package main

import (
	"os"
	"shared-buckets/bucket"
)

func executeCommand() {
	command := os.Args[1]
	switch command {
	case "create":
		bucket.CreateBucket()
	case "start":
		bucket.Start()
	}
}
