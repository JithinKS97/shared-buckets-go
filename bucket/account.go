package bucket

import (
	"encoding/hex"
	"fmt"

	nkn "github.com/nknorg/nkn-sdk-go"
)

func createAccount() string {
	account, err := nkn.NewAccount(nil)
	if err != nil {
		fmt.Println(
			"Unable to create account",
		)
		fmt.Println(err)
	}
	var seed = hex.EncodeToString(account.Seed())
	return seed
}
