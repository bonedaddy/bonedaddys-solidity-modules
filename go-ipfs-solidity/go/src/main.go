package main

import (
	"fmt"
	"log"
	"strings"

	//ethereum "github.com/ethereum/go-ethereum"
	//"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	//"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/ethclient"
	"./event_ipfs_storage"

	"github.com/howeyc/gopass"
)

const key = `{"address":"d72f0d88384c05c3d95c870ba98ac2d606939c65","crypto":{"cipher":"aes-128-ctr","ciphertext":"589a88ccbdaa312595343c907e944c8b9d9e133d443b43d4efa71c6c7cea26d0","cipherparams":{"iv":"4429d785f61dd7d37d7813a8a422d941"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f92dbdb8c2c4686a839978d9dab36601a2e950d001b6d7131dd9a22c68f32da1"},"mac":"9037da8e700215e1d79043a4fcac847768d27e28dfcd3ce16f094eb1d837f1e1"},"id":"6472fa0e-80e4-475a-8f35-ede98c37641e","version":3}`

func main() {
	client, err := ethclient.Dial("/home/solidity/.ethereum/rinkeby/geth.ipc")
	if err != nil {
		log.Fatalf("error connecting")
	}

	fmt.Println("Enter password to unlock key")
	passwd, err := gopass.GetPasswd()
	passwdString := string(passwd)
	// authorize a conenction so we can deploy stuff and change states
	auth, err := bind.NewTransactor(strings.NewReader(key), passwdString)
	if err != nil {
		log.Fatalf("error unlocking account")
	}

	address, tx, eventIpfsStorage, err := event_ipfs_storage.DeployEventIpfsStorage(auth, client)
	if err != nil {
		log.Fatalf("error deploying contract")
	}
	fmt.Printf("Contract Address:: 0x%x\n", address)
	fmt.Printf("Transaction hash 0x%x\n", tx.Hash())
	fmt.Println("Press enter to continue")
	fmt.Scanln()

	for i := 0; i < 5; i ++ {
		_, err = eventIpfsStorage.EmitStringStorageEvent(auth)
		if err!= nil {
			log.Fatalf ("error emitting event")
		}
	}

	var ch = make(chan *event_ipfs_storage.EventIpfsStorageStringStorage)
	sub, err := eventIpfsStorage.WatchStringStorage(nil, ch)
	if err != nil {
		log.Fatalf("Error establishing event Subscription")
	}

	go eventParser(sub, ch)
	fmt.Println("Press enter to continue")
	fmt.Scanln()
}

func eventParser(sub event.Subscription, c chan *event_ipfs_storage.EventIpfsStorageStringStorage) {
	
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case log := <-c:
			fmt.Printf("unpacked data %vz\n", log) 
		}
	}
}