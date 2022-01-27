package bucket

import "fmt"

type Bucket struct {
	Seed string
}

func createBucket() {
	seed := createAccount()
	bucket := Bucket{seed}
	saveToFile(bucket)
}

func start() {
	bucket := readFromFile()
	fmt.Printf("%v", bucket)
}
