package bucket

import "fmt"

type Bucket struct {
	Seed string
}

func CreateBucket() {
	seed := createAccount()
	bucket := Bucket{seed}
	saveToFile(bucket)
}

func Start() {
	bucket := readFromFile()
	fmt.Printf("%v", bucket)
}
