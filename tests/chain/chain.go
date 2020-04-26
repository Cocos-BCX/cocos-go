package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"strconv"

	sdk "github.com/Cocos-BCX/cocos-go"
	"github.com/Cocos-BCX/cocos-go/config"
	"github.com/Cocos-BCX/cocos-go/crypto"
	"github.com/Cocos-BCX/cocos-go/types"
)

func getData(api sdk.WebsocketAPI) {
	accountID := types.NewAccountID("1.2.16") // nicotest 1.2.16
	coreAsset := types.NewAssetID("1.3.0")

	fmt.Println("\nGetAccountBalances: ")
	balances, err := api.GetAccountBalances(accountID, coreAsset)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("balances: %v", balances)
	}

	fmt.Println("\nGetBlock: ")
	block, err := api.GetBlock(uint64(100))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("block: %v", block)
	}

	fmt.Println("\nGetDynamicGlobalProperties: ")
	gdp, err := api.GetDynamicGlobalProperties()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("gdb: %v", gdp)
	}
}

func testTransfer(api sdk.WebsocketAPI) {
	privateKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	// publicKey := "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx"
	localKeyBag := crypto.NewKeyBag()
	localKeyBag.Add(privateKey)

	from := types.NewAccountID("1.2.16") // nicotest 1.2.16
	to := types.NewAccountID("1.2.5")
	coreAsset := types.NewAssetID("1.3.0")
	amount := types.AssetAmount{
		Amount: types.Int64(100),
		Asset:  types.AssetIDFromObject(coreAsset),
	}

	// 1. true
	memo1 := string("memo test true")
	err1 := api.Transfer(localKeyBag, from, to, amount, memo1, true)
	fmt.Println(err1)

	// 2. false
	fmt.Println("\n-------> 2. transfer memo false")
	memo2 := string("memo test false")
	err2 := api.Transfer(localKeyBag, from, to, amount, memo2, false)
	fmt.Println(err2)

	// 3. no memo
	err3 := api.Transfer(localKeyBag, from, to, amount, "", false)
	fmt.Println(err3)
}

func testListAssets(api sdk.WebsocketAPI) {
	assets, error := api.ListAssets("", 5)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(assets)
}

func testCreateAsset(api sdk.WebsocketAPI, name string) {
	privateKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	// publicKey := "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx"
	keyBag := crypto.NewKeyBag()
	keyBag.Add(privateKey)
	// localKeyBag.Add("5Jdvatdk3qpZ8Ek9tQyqh3QwQ5mWNZ7kfnwSVwMUsLLmdUAfUwo")
	// if err != nil {
	// 	log.Println(err)
	// }

	issuer := types.NewAccountID("1.2.16") // nicotest
	coreAsset := types.NewAssetID("1.3.0")

	rate := types.Price{
		Base: types.AssetAmount{
			Amount: types.Int64(1),
			Asset:  types.AssetIDFromObject(types.NewAssetID("1.3.3")),
		},
		Quote: types.AssetAmount{
			Amount: types.Int64(1),
			Asset:  types.AssetIDFromObject(coreAsset),
		},
	}

	common := types.AssetOptions{
		MaxSupply:        210000000000,
		CoreExchangeRate: &rate,
	}
	err := api.CreateAsset(keyBag, issuer, name, 5, common, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func testGetBlock(api sdk.WebsocketAPI, from, to uint64) {
	for i := from; i <= to; i++ {
		log.Printf("block: %v ", i)
		block, err := api.GetBlock(i)
		if err != nil {
			log.Fatal(err)
			continue
		}
		log.Printf("block %v: %v", i, block)
		log.Printf("block operations: %v\n", block.Transactions)
		for _, trx := range block.Transactions {
			fmt.Printf("trx id: %v, SignedTransaction: %v\n", trx.TransactionId, trx.SignedTransaction)
		}
	}
}

func testHeadBlockPrefix(api sdk.WebsocketAPI) {
	fmt.Println("\nGetDynamicGlobalProperties: ")
	gdp, err := api.GetDynamicGlobalProperties()
	if err != nil {
		fmt.Println(err)
		return
	}

	rawBlockID, err := hex.DecodeString(gdp.HeadBlockID)
	if err != nil {
		return
	}
	fmt.Println("HeadBlockID: ", gdp.HeadBlockID, ", rawBlockID: ", rawBlockID)
	if len(rawBlockID) < 8 {
		return
	}

	rawPrefix := rawBlockID[4:8]
	fmt.Println("rawPrefix: ", rawPrefix)

	var prefix uint32
	//binary.LittleEndian 小端序
	if err := binary.Read(bytes.NewReader(rawPrefix), binary.LittleEndian, &prefix); err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("prefix: ", types.UInt32(prefix))
	fmt.Println("prefix: ", prefix)
}

func testBroadcastTrx(api sdk.WebsocketAPI) {
	fmt.Println("\nUnmarshal trx")
	trxJSON := `
	{
		"ref_block_num": 3019,
		"ref_block_prefix": 1345571086,
		"expiration": "2019-12-25T09:21:20",
		"operations": [[
			0,{
			  "fee": {
				"amount": 21054,
				"asset_id": "1.3.0"
			  },
			  "from": "1.2.15",
			  "to": "1.2.5",
			  "amount": {
				"amount": 1000000,
				"asset_id": "1.3.0"
			  },
			  "memo": {
				"from": "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx",
				"to": "COCOS5TrJztVAY5F9aWDw5KtDHfdrffQn7F3sjgbL8YyssiKhVCLNf7",
				"nonce": "6558818321903381236",
				"message": "1a87bb599789b28987809e51bdfbe69e03d1456a9392108c717a78ef238374ed"
			  },
			  "extensions": []
			}
		  ]
		],
		"extensions": [],
		"signatures": [
		  "207f701184579cdc559cc782ad12df2f7ff345d874afca790c7b79cf8fb6faa9ab3e06a8cfc12b4af1bc582d6acb02655e6347f764653cb5c33db54e511eb9d26f"
		]
	  }
	`
	var signedTrx types.SignedTransaction
	err := json.Unmarshal([]byte(trxJSON), &signedTrx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(signedTrx)

	if err := api.BroadcastTransaction(&signedTrx); err != nil {
		fmt.Println(err)
	}
}

///// WebsocketAPI test
func testGetAccountBalances(api sdk.WebsocketAPI) {
	name := "nicotest"
	account, err := api.GetAccountByName(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	// accountID := account.ID
	fmt.Printf("%s account id: %v\n", name, account.ID)
	accountID := types.NewAccountID(account.ID.String())
	accountBalances, err := api.GetAccountBalances(accountID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(accountBalances)
}

func testGetAccountByName(api sdk.WebsocketAPI) {
	name := "nicotest"
	account, err := api.GetAccountByName(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(account)
}

func testAPIID(api sdk.WebsocketAPI) {
	history := api.HistoryAPIID()
	fmt.Println("history api id: ", history)

	database := api.DatabaseAPIID()
	fmt.Println("database api id: ", database)

	broadcast := api.BroadcastAPIID()
	fmt.Println("broadcast api id: ", broadcast)
}

func testGetAccountHistory(api sdk.WebsocketAPI) {
	name := "nicotest"
	account, err := api.GetAccountByName(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s account id: %v\n", name, account.ID)
	accountID := types.NewAccountID(account.ID.String())
	fmt.Println(accountID)
	startObject := types.NewObjectID("1.11.5")
	stopObject := types.NewObjectID("1.11.20") // history: 1.11.xxx
	history, err := api.GetAccountHistory(accountID, stopObject, 10, startObject)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(history)
}

func testGetAccounts(api sdk.WebsocketAPI) {
	a1 := types.NewObjectID("1.2.16")
	a2 := types.NewObjectID("1.2.15")
	a3 := types.NewObjectID("1.2.14")
	accounts, err := api.GetAccounts(a1, a2, a3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(accounts)
	fmt.Println(len(accounts))
}

func testGetBlockHeader(api sdk.WebsocketAPI, num int) {
	blockHeader, err := api.GetBlockHeader(uint64(num))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(blockHeader)
}

func testRegisterAccount(api sdk.WebsocketAPI, name string) {
	// "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx",
	// "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	pubkey := "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx"
	priKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	keyBag := crypto.NewKeyBag()
	keyBag.Add(priKey)

	pub, err := types.NewPublicKeyFromString(pubkey)
	if err != nil {
		fmt.Println(err)
		return
	}
	// name := "tester2"
	register := "nicotest"
	err = api.RegisterAccount(keyBag, name, pub, pub, register)
	fmt.Println(err)
}

func testUpgradeAccount(api sdk.WebsocketAPI, name string) {
	// name := "tester2"
	priKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	keyBag := crypto.NewKeyBag()
	keyBag.Add(priKey)

	err := api.UpgradeAccount(keyBag, name)
	fmt.Println(err)
}

func testGetWitness(api sdk.WebsocketAPI) {
	witness, err := api.GetWitness("init5")
	fmt.Println(witness, err)

	witness, err = api.GetWitness("1.2.5")
	fmt.Println(witness, err)
}

func testGetCommitteeMember(api sdk.WebsocketAPI) {
	committee, err := api.GetCommitteeMember("1.2.5")
	fmt.Println(committee, err)
}

func testCreateContract(api sdk.WebsocketAPI) {
	name := "nicotest"
	account, err := api.GetAccountByName(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	priKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	pubKey := "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx"
	keyBag := crypto.NewKeyBag()
	keyBag.Add(priKey)

	publicKey, err := types.NewPublicKeyFromString(pubKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	contractName := "contract.debug.hellotest"
	contractCode := "function hello(str) chainhelper:log('create contract test, str: ' .. str) end"
	error := api.ContractCreate(keyBag, account, contractName, contractCode, publicKey)
	fmt.Println(error)
}

func testCallContract(api sdk.WebsocketAPI) {
	fmt.Printf(">>> %v\n", runFuncName())
	name := "nicotest"
	account, err := api.GetAccountByName(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	priKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	// pubKey := "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx"
	keyBag := crypto.NewKeyBag()
	keyBag.Add(priKey)

	// publicKey, err := types.NewPublicKeyFromString(pubKey)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	contractName := "contract.debug.param0"
	contract, err := api.GetContract(contractName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// func (p *websocketAPI) CallContract(keyBag *crypto.KeyBag, caller, creator *types.Account,
	// contractID types.ContractID, function string, valueList []types.LuaType, amount float64)

	functionName := "param"
	// arg := types.LuaString{
	// 	V: "Call Contract test",
	// }
	valueList := make([]types.LuaTypeParam, 0)
	// valueList[0] = arg
	error := api.CallContract(keyBag, account, contract.ID, functionName, valueList)
	fmt.Println(error)
}

func testCreateContractFromFile(api sdk.WebsocketAPI) {
	name := "nicotest"
	account, err := api.GetAccountByName(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	priKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	pubKey := "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx"
	keyBag := crypto.NewKeyBag()
	keyBag.Add(priKey)

	publicKey, err := types.NewPublicKeyFromString(pubKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	contractName := "contract.debug.param0" // contract.debug.test0248
	// contractCode := "function hello() chainhelper:log('create contract test') end"
	filename := "test1.lua"
	error := api.ContractCreateFromFile(keyBag, account, contractName, filename, publicKey)
	fmt.Println(error)
}

func testReveseContractFromFile(api sdk.WebsocketAPI) {
	fmt.Printf(">>> %v\n", runFuncName())
	name := "nicotest"
	account, err := api.GetAccountByName(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	priKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	// pubKey := "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx"
	keyBag := crypto.NewKeyBag()
	keyBag.Add(priKey)

	// publicKey, err := types.NewPublicKeyFromString(pubKey)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	contractName := "contract.debug.param0" // contract.debug.test0248
	contract, err := api.GetContract(contractName)
	if err != nil {
		fmt.Println(err)
		return
	}

	filename := "test1.lua"
	// ReviseContractFromFile(keyBag *crypto.KeyBag, reviser *types.Account, contractID types.ContractID, filename string)
	error := api.ReviseContractFromFile(keyBag, account, contract.ID, filename)
	fmt.Println(error)
}

func testGetContract(api sdk.WebsocketAPI, name string) {
	// contractName := "contract.debug.test0248" // contract.debug.test0248
	// contractName := "1.16.2" // contract.debug.test0248
	contract, error := api.GetContract(name)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(contract)
	fmt.Printf("contractData: %v\n", contract.ContractData)
	fmt.Printf("contractABI: %v\n", contract.ContractABI)
}

func testGetContracts(api sdk.WebsocketAPI) {
	// contractName := "contract.debug.test0248" // contract.debug.test0248
	// contractName := "1.16.2" // contract.debug.test0248

	max := 66
	for i := 1; i < max; i++ {
		name := "1.16." + strconv.Itoa(i)
		contract, error := api.GetContract(name)
		if error != nil {
			fmt.Printf("get_contract [%v] error. %v\n", name, error)
			return
		}
		// fmt.Printf("name: %v, contractData: %v\n", name, contract.ContractData)
		fmt.Printf("name: %v, contractABI: %v\n\n", name, contract.ContractABI)
	}

}

func testGetVestingBalances(api sdk.WebsocketAPI) {
	name := "nicotest"
	account, err := api.GetAccountByName(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := api.GetVestingBalances(account)
	fmt.Println(result, err)
}

func testGetConnectedPeers(api sdk.WebsocketAPI) {
	result, err := api.GetConnectedPeers()
	fmt.Println(result, err)
}

func testGetInfo(api sdk.WebsocketAPI) {
	result, err := api.Info()
	fmt.Println(result, err)
}

func testGetGlobalProperties(api sdk.WebsocketAPI) {
	result, err := api.GetGlobalProperties()
	fmt.Println(result, err)
}

func testGetChainProperties(api sdk.WebsocketAPI) {
	result, err := api.GetChainProperties()
	fmt.Println(result, err)
}

func testCreateContractFromFile2(api sdk.WebsocketAPI) {
	fmt.Printf(">>> %v\n", runFuncName())
	name := "nicotest"
	account, err := api.GetAccountByName(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	priKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	pubKey := "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx"
	keyBag := crypto.NewKeyBag()
	keyBag.Add(priKey)

	publicKey, err := types.NewPublicKeyFromString(pubKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	contractName := "contract.debug.param2" // contract.debug.test0248
	// contractCode := "function hello() chainhelper:log('create contract test') end"
	filename := "test2.lua"
	error := api.ContractCreateFromFile(keyBag, account, contractName, filename, publicKey)
	fmt.Println(error)
}

func testReviseContractFromFile2(api sdk.WebsocketAPI) {
	fmt.Printf(">>> %v\n", runFuncName())
	name := "nicotest"
	account, err := api.GetAccountByName(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	priKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	// pubKey := "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx"
	keyBag := crypto.NewKeyBag()
	keyBag.Add(priKey)

	contractName := "contract.debug.param2" // contract.debug.test0248
	contract, err := api.GetContract(contractName)
	if err != nil {
		return
	}
	filename := "test2.lua"
	error := api.ReviseContractFromFile(keyBag, account, contract.ID, filename)
	fmt.Println(error)
}

func testCallContract2(api sdk.WebsocketAPI) {
	fmt.Printf(">>> %v\n", runFuncName())
	name := "nicotest"
	account, err := api.GetAccountByName(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	priKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	// pubKey := "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx"
	keyBag := crypto.NewKeyBag()
	keyBag.Add(priKey)

	// publicKey, err := types.NewPublicKeyFromString(pubKey)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	contractName := "contract.debug.param2"
	contract, err := api.GetContract(contractName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// func (p *websocketAPI) CallContract(keyBag *crypto.KeyBag, caller, creator *types.Account,
	// contractID types.ContractID, function string, valueList []types.LuaType, amount float64)

	functionName := "param1"
	arg := types.LuaString{
		V: "Call Contract test param1",
	}
	// valueList := make([]types.LuaTypeParam, 1)
	// param1 := types.LuaTypeParam{
	// 	types.LuaTypeString,
	// 	arg,
	// }
	// valueList[0] = param1
	valueList := []types.LuaTypeParam{
		types.LuaTypeParam{
			types.LuaTypeString,
			arg,
		},
	}
	error := api.CallContract(keyBag, account, contract.ID, functionName, valueList)
	fmt.Println(error)
}

func testCallContract3(api sdk.WebsocketAPI) {
	fmt.Printf(">>> %v\n", runFuncName())
	name := "nicotest"
	account, err := api.GetAccountByName(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	priKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	// pubKey := "COCOS56a5dTnfGpuPoWACnYj65dahcXMpTrNQkV3hHWCFkLxMF5mXpx"
	keyBag := crypto.NewKeyBag()
	keyBag.Add(priKey)

	// publicKey, err := types.NewPublicKeyFromString(pubKey)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	contractName := "contract.debug.param2"
	contract, err := api.GetContract(contractName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// func (p *websocketAPI) CallContract(keyBag *crypto.KeyBag, caller, creator *types.Account,
	// contractID types.ContractID, function string, valueList []types.LuaType, amount float64)

	functionName := "param2"
	arg := types.LuaString{
		V: "Call Contract test param1",
	}
	// valueList := make([]types.LuaTypeParam, 1)
	// param1 := types.LuaTypeParam{
	// 	types.LuaTypeString,
	// 	arg,
	// }
	// valueList[0] = param1
	valueList := []types.LuaTypeParam{
		types.LuaTypeParam{
			types.LuaTypeString,
			arg,
		},
		types.LuaTypeParam{
			types.LuaTypeBool,
			types.LuaBool{
				V: true,
			},
		},
	}
	error := api.CallContract(keyBag, account, contract.ID, functionName, valueList)
	fmt.Println(error)
}

func testCallContract4(api sdk.WebsocketAPI) {
	fmt.Printf(">>> %v\n", runFuncName())
	name := "nicotest"
	account, err := api.GetAccountByName(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	priKey := "5J2SChqa9QxrCkdMor9VC2k9NT4R4ctRrJA6odQCPkb3yL89vxo"
	keyBag := crypto.NewKeyBag()
	keyBag.Add(priKey)

	contractName := "contract.debug.param2"
	contract, err := api.GetContract(contractName)
	if err != nil {
		fmt.Println(err)
		return
	}

	functionName := "test_luatype"
	// test luatype string
	// valueList := []types.LuaTypeParam{
	// 	types.LuaTypeParam{
	// 		types.LuaTypeString,
	// 		types.LuaString{
	// 			V: "test LuaString Type",
	// 		},
	// 	},
	// }
	// error := api.CallContract(keyBag, account, contract.ID, functionName, valueList)
	// fmt.Println(error)

	// // test luatype bool
	// valueList = []types.LuaTypeParam{
	// 	types.LuaTypeParam{
	// 		types.LuaTypeBool,
	// 		types.LuaBool{
	// 			V: true,
	// 		},
	// 	},
	// }
	// error = api.CallContract(keyBag, account, contract.ID, functionName, valueList)
	// fmt.Println(error)

	// test luatype Int
	// valueList := []types.LuaTypeParam{
	// 	types.LuaTypeParam{
	// 		types.LuaTypeInt,
	// 		types.LuaInt{
	// 			V: 200,
	// 		},
	// 	},
	// }
	// error := api.CallContract(keyBag, account, contract.ID, functionName, valueList)
	// fmt.Println(error)

	// test Luatype number
	valueList := []types.LuaTypeParam{
		types.LuaTypeParam{
			types.LuaTypeNumber,
			types.LuaNumber{
				V: 3.14,
			},
		},
	}
	error := api.CallContract(keyBag, account, contract.ID, functionName, valueList)
	fmt.Println(error)

	// test luatype table
	// valueList := []types.LuaTypeParam{
	// 	types.LuaTypeParam{
	// 		types.LuaTypeTable,
	// 		types.LuaTable{
	// 			Value: types.LuaMap{
	// 				types.LuaTypeItem{
	// 					types.LuaKey{
	// 						Key: types.LuaType{
	// 							types.LuaString{V: "key1"},
	// 						},
	// 					},
	// 					types.LuaString{V: "value1"},
	// 				},
	// 				types.LuaTypeItem{
	// 					types.LuaKey{
	// 						Key: types.LuaType{
	// 							types.LuaInt{V: 200},
	// 						},
	// 					},
	// 					types.LuaString{V: "value2"},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// error := api.CallContract(keyBag, account, contract.ID, functionName, valueList)
	// fmt.Println(error)
}

func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func main() {
	fmt.Printf(">>> %v\n", runFuncName())
	// config.SetCurrent(config.ChainIDTestnet)
	// wsURL := "wss://test.cocosbcx.net"

	// config.SetCurrent(config.ChainIDMainnet)
	// wsURL := "wss://api.cocosbcx.net"

	// config.SetCurrent(config.ChainIDLocal)
	// wsURL := "ws://127.0.0.1:8049"

	chainID := config.ChainIDLocal
	wsURL := "ws://127.0.0.1:8049"
	config.SetCurrent(chainID)

	// chain api 测试
	log.Println("------- chain api test ----------")
	api := sdk.NewWebsocketAPI(wsURL)
	if err := api.Connect(); err != nil {
		log.Println(err)
	}

	// get chain data
	// testAPIID(api)
	// getData(api)
	// testListAssets(api)
	// testHeadBlockPrefix(api)
	// testGetAccountByName(api)
	// testGetAccountBalances(api)
	// testGetAccountHistory(api)
	// testGetAccounts(api)
	// testGetBlock(api, 28735, 28739)
	// testGetBlockHeader(api, 28738)
	// testBroadcastTrx(api)

	// operation
	// testTransfer(api)

	// asset_name := "USDG"
	// testCreateAsset(api, asset_name)

	// new_account_name := "tester3"
	// testRegisterAccount(api, new_account_name)
	// testUpgradeAccount(api, new_account_name)

	// testGetWitness(api)
	// testGetCommitteeMember(api)
	// testCreateContract(api)
	// testGetVestingBalances(api)

	// testGetConnectedPeers(api)
	// testGetInfo(api)
	// testGetGlobalProperties(api)
	// testGetChainProperties(api)

	// contract
	// fmt.Println("\n\n--------------- create contract by file")
	// testCreateContractFromFile(api)

	// fmt.Println("\n\n--------------- get contract")
	// testGetContract(api, "1.16.1")

	// testGetContract(api, "1.16.2")
	// testGetContract(api, "1.16.3")
	// testGetContract(api, "1.16.5")
	// testGetContract(api, "1.16.9")
	// testGetContracts(api)

	// testReveseContractFromFile(api)
	// testCallContract(api)

	// testCreateContractFromFile2(api)
	// testCallContract2(api)
	// testCallContract3(api)

	// testReviseContractFromFile2(api)
	testCallContract4(api)

}
