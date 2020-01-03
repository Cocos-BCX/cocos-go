package main

import (
	"fmt"
	"log"

	sdk "github.com/Cocos-BCX/cocos-go"
	"github.com/Cocos-BCX/cocos-go/config"
	"github.com/Cocos-BCX/cocos-go/crypto"
)

func main() {
	// config.SetCurrent(config.ChainIDBCXDev)
	// wsURL := "ws://127.0.0.1:8049"
	// // wsURL := "ws://test.cocosbcx.net"
	// walletURL := "http://127.0.0.1:8048"

	config.SetCurrent(config.ChainIDBCXTest)
	walletURL := "http://127.0.0.1:8048"

	// testAccount := "nicotest"
	privateKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	publicKey := "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx"

	// 1. key bag 测试
	log.Println("------- key bag test ----------")
	localKeyBag := crypto.NewKeyBag()
	err := localKeyBag.Add(privateKey)
	if err != nil {
		log.Println(err)
	} else {
		publics := localKeyBag.Publics()
		fmt.Printf("publics: %v %v\n", publics, publicKey)
	}

	// wallet api 测试
	log.Println("------- wallet api test ----------")
	walletAPI := sdk.NewWalletAPI(walletURL)
	if err := walletAPI.Connect(); err != nil {
		log.Println(err)
	}

	info, err := walletAPI.Info()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("info: %v", info)

	block, err := walletAPI.GetBlock(uint64(4497))
	log.Printf("block: %v, err: %v\n", block, err)
}
