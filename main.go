package main

import (
	"fmt"

	nkn "github.com/nknorg/nkn-sdk-go"
)

func main() {
	create()
}

func create() {
	account, err := nkn.NewAccount(nil)
	if err != nil {
		fmt.Println(
			"Unable to create account",
		)
		fmt.Println(err)
	}
	var seed = account.Seed()
	fmt.Println(seed)
}
