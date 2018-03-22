package main


import (
	"fmt"
	"log"
	"context"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "./event_ipfs_storage"
)

func main() {

	fmt.Println("initiating ipc connection")
	client, err := ethclient.Dial("/home/solidity/.ethereum/geth.ipc")
	if err != nil {
		log.Fatal("error connecting to blockchain ", err)
	} else {
		fmt.Println("ipc connection successfully established")
	}

	contractAddress := common.HexToAddress("0x971ff3C7BA604fcaCbdBee5b56FE790f22258A83")

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
            fmt.Println(storage.A)
    	}
    }
}