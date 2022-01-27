package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func saveToFile(object Bucket) {
	fmt.Println(object)
	file, _ := json.Marshal(object)
	_ = ioutil.WriteFile("file.json", file, 0644)
}

func readFromFile() Bucket {
	file, _ := ioutil.ReadFile("file.json")
	data := Bucket{}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}
