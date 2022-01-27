package bucket

import (
	"encoding/hex"
	"fmt"

	nkn "github.com/nknorg/nkn-sdk-go"
)

type Bucket struct {
	Seed   string
	isHost bool
}

func CreateBucket() {
	seed := createAccount()
	bucket := Bucket{seed, true}
	saveToFile(bucket)
}

func Start() {
	bucket := readFromFile()
	client := createClientFromSeed(bucket.Seed)
	handleMessage(client)
}

func createClientFromSeed(seed string) *nkn.Client {
	// Seed obtained from
	seedByte, _ := hex.DecodeString(seed)
	account, _ := nkn.NewAccount(seedByte)
	client, _ := nkn.NewClient(account, "", nil)
	<-client.OnConnect.C
	fmt.Println("Connected to RPC")
	fmt.Println("Address of this node:", client.Address())
	return client
}

func createClient() *nkn.Client {
	account, _ := nkn.NewAccount(nil)
	client, _ := nkn.NewClient(account, "", nil)
	return client
}

func handleMessage(client *nkn.Client) <-chan string {

	r := make(chan string)

	go func() {
		defer close(r)
		msg := <-client.OnMessage.C
		println("handle message func")
		println("Message")
		fmt.Println(msg.Data)
	}()

	return r
}

func Connect(address string) {
	client := createClient()
	print(address)
	client.Send(nkn.NewStringArray(address), []byte("hello world!"), nil)
	handleMessage(client)
}
