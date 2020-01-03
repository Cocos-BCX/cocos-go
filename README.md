# cocos-go

A cocos-go API consuming a websocket connection to an active full node or a RPC connection to your `cli_wallet`. 
Look for several examples in [examples](/examples) and [tests](/tests) folder. This is work in progress. To mitigate breaking changes, please use tagged branches. New tagged branches will be created for breaking changes. No additional cgo dependencies for transaction signing required. Use it at your own risk. 

## install

```bash
go get -u github.com/Cocos-BCX/cocos-go
```

Install dev-dependencies with

```bash
make init
```

This API uses [ffjson](https://github.com/pquerna/ffjson). 
If you change types or operations you have to regenerate the required static `MarshalJSON` and `UnmarshalJSON` functions for the new/changed code.

```bash
make generate
```

If you encounter any errors, try: 

```bash
make generate_new
```

to generate ffjson helpers from scratch.

## generate operation samples
To generate op samples for testing, go to [gen](/gen) package.
Generated operation samples get injected automatically while running operation tests.

## code

```go
import (
	sdk "github.com/Cocos-BCX/cocos-go"
)

wsURL := "ws://test.cocosbcx.net"

api := sdk.NewWebsocketAPI(wsURL)
if err := api.Connect(); err != nil {
	log.Println(err)
}

accountID := types.NewAccountID("1.2.16") // nicotest 1.2.16
coreAsset := types.NewAssetID("1.3.0")

balances, err := api.GetAccountBalances(accountID, coreAsset)
if err != nil {
	log.Fatal(err)
}
log.Printf("balances: %v", balances)
```

If you need wallet functions, use:

```go

import (
	sdk "github.com/Cocos-BCX/cocos-go"
)

walletURL := "http://127.0.0.1:8048"

walletAPI := sdk.NewWalletAPI(walletURL)
if err := walletAPI.Connect(); err != nil {
	log.Println(err)
}

info, err := walletAPI.Info()
if err != nil {
	log.Fatal(err)
}
log.Printf("info: %v", info)

...
```

For a long application lifecycle, you can use an API instance with latency tester that connects to the most reliable node.
Note: Because the tester takes time to unleash its magic, use the above-mentioned constructor for quick in and out.

```go
import (
	sdk "github.com/Cocos-BCX/cocos-go"
)

wsFullApiUrl := "ws://test.cocosbcx.net"

//wsFullApiUrl serves as "quick startup" fallback endpoint here, 
//until the latency tester provides the first results.

api, err := sdk.NewWithAutoEndpoint(wsFullApiUrl)
if err != nil {
	log.Fatal(err)
}

if err := api.Connect(); err != nil {
	log.Fatal(err)
}

api.OnError(func(err error) {
	log.Fatal(err)
})

...
```

## implemented and tested (serialize/unserialize) operations

- [x] OperationTypeTransfer                              
- [x] OperationTypeLimitOrderCreate                      
- [x] OperationTypeLimitOrderCancel                      
- [x] OperationTypeCallOrderUpdate                       
- [x] OperationTypeFillOrder                             
- [x] OperationTypeAccountCreate                         
- [x] OperationTypeAccountUpdate                         
- [x] OperationTypeAccountUpgrade                        
- [x] OperationTypeAssetCreate                           
- [x] OperationTypeAssetUpdate                           
- [x] OperationTypeAssetUpdateRestricted                 
- [x] OperationTypeAssetUpdateBitasset                   
- [x] OperationTypeAssetUpdateFeedProducers              
- [x] OperationTypeAssetIssue                            
- [x] OperationTypeAssetReserve                          
- [x] OperationTypeAssetSettle                           
- [x] OperationTypeAssetGlobalSettle                     
- [x] OperationTypeAssetPublishFeed                      
- [x] OperationTypeWitnessCreate                         
- [x] OperationTypeWitnessUpdate                         
- [x] OperationTypeProposalCreate                        
- [x] OperationTypeProposalUpdate                        
- [x] OperationTypeProposalDelete                        
- [x] OperationTypeCommitteeMemberCreate                 
- [x] OperationTypeCommitteeMemberUpdate                 
- [x] OperationTypeCommitteeMemberUpdateGlobalParameters  
- [x] OperationTypeVestingBalanceCreate                  
- [x] OperationTypeVestingBalanceWithdraw                
- [x] OperationTypeWorkerCreate                          
- [x] OperationTypeBalanceClaim                          
- [x] OperationTypeAssetSettleCancel                     
- [x] OperationTypeAssetClaimFees                        
- [x] OperationTypeBidCollateral                         
- [x] OperationTypeExecuteBid                            
- [x] OperationTypeContractCreate                        
- [x] OperationTypeCallContractFunction                  
- [x] OperationTypeTemporaryAuthorityChange              
- [x] OperationTypeRegisterNhAssetCreator                
- [x] OperationTypeCreateWorldView                       
- [x] OperationTypeRelateWorldView                       
- [x] OperationTypeCreateNhAsset                         
- [x] OperationTypeDeleteNhAsset                         
- [x] OperationTypeTransferNhAsset                       
- [x] OperationTypeCreateNhAssetOrder                    
- [x] OperationTypeCancelNhAssetOrder                    
- [x] OperationTypeFillNhAssetOrder                      
- [x] OperationTypeCreateFile                            
- [x] OperationTypeAddFileRelateAccount                  
- [x] OperationTypeFileSignature                         
- [x] OperationTypeRelateParentFile                      
- [x] OperationTypeReviseContract                        
- [ ] OperationTypeCrontabCreate                         
- [ ] OperationTypeCrontabCancel                         
- [ ] OperationTypeCrontabRecover                        
- [x] OperationTypeUpdateCollateralForGas                
- [x] OperationTypeAccountAuthentication                 

## todo   
* add missing operations   
* add convenience functions   

Have fun and feel free to contribute needed operations and tests.   