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
	<-client.OnConnect.C
	fmt.Println("Connected to RPC")
	return client
}

func handleMessage(client *nkn.Client) {
	msg := <-client.OnMessage.C
	fmt.Println(string(msg.Data))
}

func Connect(address string) {
	client := createClient()
	client.Send(nkn.NewStringArray(address), []byte("hello world! WTF"), nil)
	handleMessage(client)
}
