package main


import (
	"fmt"
	"log"
	//"context"
	"strings"

//	ethereum "github.com/ethereum/go-ethereum"
	//"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
  //  "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

    "./event_ipfs_storage"
)

const key = `{"address":"d72f0d88384c05c3d95c870ba98ac2d606939c65","crypto":{"cipher":"aes-128-ctr","ciphertext":"589a88ccbdaa312595343c907e944c8b9d9e133d443b43d4efa71c6c7cea26d0","cipherparams":{"iv":"4429d785f61dd7d37d7813a8a422d941"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f92dbdb8c2c4686a839978d9dab36601a2e950d001b6d7131dd9a22c68f32da1"},"mac":"9037da8e700215e1d79043a4fcac847768d27e28dfcd3ce16f094eb1d837f1e1"},"id":"6472fa0e-80e4-475a-8f35-ede98c37641e","version":3}`

func main() {

	fmt.Println("initiating ipc connection")
	client, err := ethclient.Dial("/home/solidity/.ethereum/geth.ipc")
	if err != nil {
		log.Fatal("error connecting to blockchain ", err)
	} else {
		fmt.Println("ipc connection successfully established")
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), "password123")
	if err != nil {
		log.Fatalf("error unlocking account")
	}


	address, tx, eventIpfsStorage, err := event_ipfs_storage.DeployEventIpfsStorage(auth, client)
	if err != nil {
		log.Fatal("error instantiating instance of contract", err)
	}
	fmt.Printf("Contract address 0x%x\n", address)
	fmt.Printf("Transaction hash 0x%x\n", tx.Hash())
	fmt.Println("contract deployed, waiting, press enter to continue")
	fmt.Scanln()

	tx, err = eventIpfsStorage.EmitStringStorageEvent(auth)
	if err != nil {
		log.Fatal("error emitting event", err)
	}

	var ch = make(chan *event_ipfs_storage.EventIpfsStorageTextData)
	sub, err := eventIpfsStorage.WatchTextData(nil, ch)
	if err != nil {
		log.Fatal("error establishing event subscription", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal("error parsing event", err)
		case log := <-ch:
			fmt.Println("log", log.A)
		}
	}


/*
    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
    }

    contractAbi, err := abi.JSON(strings.NewReader(event_ipfs_storage.EventIpfsStorageABI))

    if err != nil {
    	log.Fatal("error loading contract abi")
    } else {
    	fmt.Println("contract abi successfully loaded")
    }

    var ch = make(chan types.Log)
    ctx := context.Background()

    sub, err := client.SubscribeFilterLogs(ctx, query, ch)
    if err != nil {
    	log.Fatal("error creating filter log", err)
    }
    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case log := <-ch:
            fmt.Println("Log:", log.Data)
            var storage struct {
				A   string
				Raw types.Log // Blockchain specific contextual infos
			}
            contractAbi.Unpack(&storage, "StringStorage", log.Data)
            fmt.Println(common.BytesToHash(log.Topics[0].Bytes()).Hex())
            fmt.Println("Log string ", log.String())
    	}
    }*/
}