package main

import (
	"fmt"
	"log"

	"github.com/Cocos-BCX/cocos-go/config"
	"github.com/Cocos-BCX/cocos-go/crypto"
	"github.com/Cocos-BCX/cocos-go/types"
)

func main() {
	config.SetCurrent(config.ChainIDBCXTest)

	// testAccount := "nicotest"
	privateKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	publicKey := "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx"

	// key bag 测试
	log.Println("------- key bag test ----------")
	localKeyBag := crypto.NewKeyBag()
	err := localKeyBag.Add(privateKey)
	if err != nil {
		log.Println(err)
		return
	}

	publics := localKeyBag.Publics()
	fmt.Printf("publics: %v %v\n", publics, publicKey)

	privates := localKeyBag.Privates()
	fmt.Printf("privates: %v\n", privates)

	sPubKey, err := types.NewPublicKeyFromString(publicKey)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("sPubKey: %v\n", sPubKey.String())
		// priKey := localKeyBag.PrivatesByPublics(sPubKey)
		// fmt.Printf("publicKey: %v, priKey: %v\n", publicKey, priKey)
	}

}
