package api

import (
	"time"

	"github.com/denkhaus/bitshares/client"
	"github.com/denkhaus/bitshares/config"
	"github.com/denkhaus/bitshares/crypto"
	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
	"github.com/pquerna/ffjson/ffjson"
)

const (
	InvalidApiID           = -1
	AssetsListAll          = -1
	AssetsMaxBatchSize     = 100
	GetCallOrdersLimit     = 100
	GetLimitOrdersLimit    = 100
	GetSettleOrdersLimit   = 100
	GetTradeHistoryLimit   = 100
	GetAccountHistoryLimit = 100
)

type BitsharesAPI interface {

	//Common functions
	SetDebug(debug bool)
	Debug(descr string, in interface{})
	CallWsAPI(apiID int, method string, args ...interface{}) (interface{}, error)
	Close() error
	Connect() error
	DatabaseAPIID() int
	CryptoAPIID() int
	HistoryAPIID() int
	BroadcastAPIID() int
	SetCredentials(username, password string)
	OnError(func(error))
	OnNotify(subscriberID int, notifyFn func(msg interface{}) error) error
	BuildAndSignTransaction(keyBag *crypto.KeyBag, feeAsset types.GrapheneObject, ops ...types.Operation) (*types.Transaction, error)
	SignTransaction(keyBag *crypto.KeyBag, trx *types.Transaction) (*types.Transaction, error)

	//Websocket API functions
	BroadcastTransaction(tx *types.Transaction) (string, error)
	CancelAllSubscriptions() error
	CancelOrder(orderID types.GrapheneObject, broadcast bool) (*types.Transaction, error)
	GetAccountBalances(account types.GrapheneObject, assets ...types.GrapheneObject) (types.AssetAmounts, error)
	GetAccountByName(name string) (*types.Account, error)
	GetAccountHistory(account types.GrapheneObject, stop types.GrapheneObject, limit int, start types.GrapheneObject) (types.OperationHistories, error)
	GetAccounts(accountIDs ...types.GrapheneObject) (types.Accounts, error)
	GetBlock(number uint64) (*types.Block, error)
	GetCallOrders(assetID types.GrapheneObject, limit int) (types.CallOrders, error)
	GetChainID() (string, error)
	GetDynamicGlobalProperties() (*types.DynamicGlobalProperties, error)
	GetLimitOrders(base, quote types.GrapheneObject, limit int) (types.LimitOrders, error)
	GetMarginPositions(accountID types.GrapheneObject) (types.CallOrders, error)
	GetObjects(objectIDs ...types.GrapheneObject) ([]interface{}, error)
	GetSettleOrders(assetID types.GrapheneObject, limit int) (types.SettleOrders, error)
	GetTradeHistory(base, quote types.GrapheneObject, toTime, fromTime time.Time, limit int) (types.MarketTrades, error)
	ListAssets(lowerBoundSymbol string, limit int) (types.Assets, error)
	SetSubscribeCallback(notifyID int, clearFilter bool) error
	SubscribeToMarket(notifyID int, base types.GrapheneObject, quote types.GrapheneObject) error
	UnsubscribeFromMarket(base types.GrapheneObject, quote types.GrapheneObject) error

	//Wallet API functions
	ListAccountBalances(account types.GrapheneObject) (types.AssetAmounts, error)
	WalletLock() error
	WalletUnlock(password string) error
	WalletIsLocked() (bool, error)
	BorrowAsset(account types.GrapheneObject, amountToBorrow string, symbolToBorrow types.GrapheneObject, amountOfCollateral string, broadcast bool) (*types.Transaction, error)
	Buy(account types.GrapheneObject, base, quote types.GrapheneObject, rate string, amount string, broadcast bool) (*types.Transaction, error)
	BuyEx(account types.GrapheneObject, base, quote types.GrapheneObject, rate float64, amount float64, broadcast bool) (*types.Transaction, error)
	Sell(account types.GrapheneObject, base, quote types.GrapheneObject, rate string, amount string, broadcast bool) (*types.Transaction, error)
	SellEx(account types.GrapheneObject, base, quote types.GrapheneObject, rate float64, amount float64, broadcast bool) (*types.Transaction, error)
	SellAsset(account types.GrapheneObject, amountToSell string, symbolToSell types.GrapheneObject, minToReceive string, symbolToReceive types.GrapheneObject, timeout uint32, fillOrKill bool, broadcast bool) (*types.Transaction, error)
	WalletSignTransaction(tx *types.Transaction, broadcast bool) (*types.Transaction, error)
	SerializeTransaction(tx *types.Transaction) (string, error)
}

type bitsharesAPI struct {
	wsClient       client.WebsocketClient
	rpcClient      client.RPCClient
	username       string
	password       string
	databaseAPIID  int
	historyAPIID   int
	cryptoAPIID    int
	broadcastAPIID int
	debug          bool
}

func (p *bitsharesAPI) getApiID(identifier string) (int, error) {
	defer p.SetDebug(false)
	resp, err := p.wsClient.CallAPI(1, identifier, types.EmptyParams)
	if err != nil {
		return InvalidApiID, errors.Annotate(err, identifier)
	}

	p.Debug("getApiID <", resp)

	return int(resp.(float64)), nil
}

// login
func (p *bitsharesAPI) login() (bool, error) {
	defer p.SetDebug(false)
	resp, err := p.wsClient.CallAPI(1, "login", p.username, p.password)
	if err != nil {
		return false, err
	}

	p.Debug("login <", resp)

	return resp.(bool), nil
}

// SetSubscribeCallback
func (p *bitsharesAPI) SetSubscribeCallback(notifyID int, clearFilter bool) error {
	defer p.SetDebug(false)
	// returns nil if successfull
	_, err := p.wsClient.CallAPI(p.databaseAPIID, "set_subscribe_callback", notifyID, clearFilter)
	if err != nil {
		return err
	}

	return nil
}

// SubscribeToMarket
func (p *bitsharesAPI) SubscribeToMarket(notifyID int, base types.GrapheneObject, quote types.GrapheneObject) error {
	defer p.SetDebug(false)
	// returns nil if successfull
	_, err := p.wsClient.CallAPI(p.databaseAPIID, "subscribe_to_market", notifyID, base.Id(), quote.Id())
	if err != nil {
		return err
	}

	return nil
}

// UnsubscribeFromMarket
func (p *bitsharesAPI) UnsubscribeFromMarket(base types.GrapheneObject, quote types.GrapheneObject) error {
	defer p.SetDebug(false)
	// returns nil if successfull
	_, err := p.wsClient.CallAPI(p.databaseAPIID, "unsubscribe_from_market", base.Id(), quote.Id())
	if err != nil {
		return err
	}

	return nil
}

// CancelAllSubscriptions
func (p *bitsharesAPI) CancelAllSubscriptions() error {
	defer p.SetDebug(false)
	// returns nil
	_, err := p.wsClient.CallAPI(p.databaseAPIID, "cancel_all_subscriptions", types.EmptyParams)
	if err != nil {
		return err
	}

	return nil
}

//Broadcast a transaction to the network.
//The transaction will be checked for validity in the local database prior to broadcasting.
//If it fails to apply locally, an error will be thrown and the transaction will not be broadcast.
func (p *bitsharesAPI) BroadcastTransaction(tx *types.Transaction) (string, error) {
	defer p.SetDebug(false)
	resp, err := p.wsClient.CallAPI(p.broadcastAPIID, "broadcast_transaction", tx)
	if err != nil {
		return "", err
	}

	p.Debug("broadcast_transaction <", resp)

	return resp.(string), nil
}

//SignTransaction signes a given transaction assuming valid block properties and fees.
func (p *bitsharesAPI) SignTransaction(keyBag *crypto.KeyBag, tx *types.Transaction) (*types.Transaction, error) {
	defer p.SetDebug(false)

	pubKeys, err := p.GetPotentialSignatures(tx)
	if err != nil {
		return nil, errors.Annotate(err, "GetPotentialSignatures")
	}

	p.Debug("potential pubkeys <", pubKeys)

	signedTrx := crypto.NewSignedTransaction(tx)
	privKeys := keyBag.GetPotentialPrivKeys(pubKeys)

	if err := signedTrx.Sign(privKeys, config.CurrentConfig()); err != nil {
		return nil, errors.Annotate(err, "Sign")
	}

	return tx, nil

}

//BuildAndSignTransaction builds a new Transaction by the given operation(s), applies fees, current block data and signes the transaction.
func (p *bitsharesAPI) BuildAndSignTransaction(keyBag *crypto.KeyBag, feeAsset types.GrapheneObject, ops ...types.Operation) (*types.Transaction, error) {
	defer p.SetDebug(false)

	operations := types.Operations(ops)
	fees, err := p.GetRequiredFees(operations, feeAsset)
	if err != nil {
		return nil, errors.Annotate(err, "GetRequiredFees")
	}

	if err := operations.ApplyFees(fees); err != nil {
		return nil, errors.Annotate(err, "ApplyFees")
	}

	props, err := p.GetDynamicGlobalProperties()
	if err != nil {
		return nil, errors.Annotate(err, "GetDynamicGlobalProperties")
	}

	tx, err := types.NewTransactionWithBlockData(props)
	if err != nil {
		return nil, errors.Annotate(err, "NewTransaction")
	}

	tx.Operations = operations

	pubKeys, err := p.GetPotentialSignatures(tx)
	if err != nil {
		return nil, errors.Annotate(err, "GetPotentialSignatures")
	}

	p.Debug("potential pubkeys <", pubKeys)

	signedTrx := crypto.NewSignedTransaction(tx)

	privKeys := keyBag.GetPotentialPrivKeys(pubKeys)
	if len(privKeys) == 0 {
		return nil, types.ErrNoSigningKeyFound
	}

	if err := signedTrx.Sign(privKeys, config.CurrentConfig()); err != nil {
		return nil, errors.Annotate(err, "Sign")
	}

	return tx, nil
}

//GetPotentialSignatures will return the set of all public keys that could possibly sign for a given transaction.
//This call can be used by wallets to filter their set of public keys to just the relevant subset prior to calling
//GetRequiredSignatures to get the minimum subset.
func (p *bitsharesAPI) GetPotentialSignatures(tx *types.Transaction) (types.PublicKeys, error) {
	defer p.SetDebug(false)

	resp, err := p.wsClient.CallAPI(p.databaseAPIID, "get_potential_signatures", tx)
	if err != nil {
		return nil, err
	}

	p.Debug("get_potential_signatures <", resp)

	ret := types.PublicKeys{}
	if err := ffjson.Unmarshal(util.ToBytes(resp), &ret); err != nil {
		return nil, errors.Annotate(err, "unmarshal PublicKeys")
	}

	return ret, nil
}

//GetBlock returns a Block by block number.
func (p *bitsharesAPI) GetBlock(number uint64) (*types.Block, error) {
	defer p.SetDebug(false)
	resp, err := p.wsClient.CallAPI(0, "get_block", number)
	if err != nil {
		return nil, err
	}

	p.Debug("get_block <", resp)

	ret := types.Block{}
	if err := ffjson.Unmarshal(util.ToBytes(resp), &ret); err != nil {
		return nil, errors.Annotate(err, "unmarshal Block")
	}

	return &ret, nil
}

//GetAccountByName returns a Account object by username
func (p *bitsharesAPI) GetAccountByName(name string) (*types.Account, error) {
	defer p.SetDebug(false)
	resp, err := p.wsClient.CallAPI(0, "get_account_by_name", name)
	if err != nil {
		return nil, err
	}

	p.Debug("get_account_by_name <", resp)

	ret := types.Account{}
	if err := ffjson.Unmarshal(util.ToBytes(resp), &ret); err != nil {
		return nil, errors.Annotate(err, "unmarshal Account")
	}

	return &ret, nil
}

// GetAccountHistory returns OperationHistory object(s).
// account: The account whose history should be queried
// stop: ID of the earliest operation to retrieve
// limit: Maximum number of operations to retrieve (must not exceed 100)
// start: ID of the most recent operation to retrieve
func (p *bitsharesAPI) GetAccountHistory(account types.GrapheneObject, stop types.GrapheneObject, limit int, start types.GrapheneObject) (types.OperationHistories, error) {
	defer p.SetDebug(false)

	if limit > GetAccountHistoryLimit {
		limit = GetAccountHistoryLimit
	}

	resp, err := p.wsClient.CallAPI(p.historyAPIID, "get_account_history", account.Id(), stop.Id(), limit, start.Id())
	if err != nil {
		return nil, err
	}

	p.Debug("get_account_history <", resp)

	ret := types.OperationHistories{}
	if err := ffjson.Unmarshal(util.ToBytes(resp), &ret); err != nil {
		return nil, errors.Annotate(err, "unmarshal Histories")
	}

	return ret, nil
}

//GetAccounts returns a list of accounts by ID.
func (p *bitsharesAPI) GetAccounts(accounts ...types.GrapheneObject) (types.Accounts, error) {
	defer p.SetDebug(false)

	ids := types.GrapheneObjects(accounts).ToStrings()
	resp, err := p.wsClient.CallAPI(0, "get_accounts", ids)
	if err != nil {
		return nil, err
	}

	p.Debug("get_accounts <", resp)

	ret := types.Accounts{}
	if err := ffjson.Unmarshal(util.ToBytes(resp), &ret); err != nil {
		return nil, errors.Annotate(err, "unmarshal Accounts")
	}

	return ret, nil
}

//GetDynamicGlobalProperties
func (p *bitsharesAPI) GetDynamicGlobalProperties() (*types.DynamicGlobalProperties, error) {
	defer p.SetDebug(false)

	resp, err := p.wsClient.CallAPI(0, "get_dynamic_global_properties", types.EmptyParams)
	if err != nil {
		return nil, err
	}

	p.Debug("get_dynamic_global_properties <", resp)

	var ret types.DynamicGlobalProperties
	if err := ffjson.Unmarshal(util.ToBytes(resp), &ret); err != nil {
		return nil, errors.Annotate(err, "unmarshal DynamicGlobalProperties")
	}

	return &ret, nil
}

//GetAccountBalances retrieves AssetAmount objects by given AccountID
func (p *bitsharesAPI) GetAccountBalances(account types.GrapheneObject, assets ...types.GrapheneObject) (types.AssetAmounts, error) {
	defer p.SetDebug(false)

	ids := types.GrapheneObjects(assets).ToStrings()
	resp, err := p.wsClient.CallAPI(0, "get_account_balances", account.Id(), ids)
	if err != nil {
		return nil, err
	}

	p.Debug("get_account_balances <", resp)

	ret := types.AssetAmounts{}
	if err := ffjson.Unmarshal(util.ToBytes(resp), &ret); err != nil {
		return nil, errors.Annotate(err, "unmarshal AssetAmounts")
	}

	return ret, nil
}

//ListAssets retrieves assets
//@param lowerBoundSymbol: Lower bound of symbol names to retrieve
//@param limit: Maximum number of assets to fetch, if the constant AssetsListAll is passed, all existing assets will be retrieved.
func (p *bitsharesAPI) ListAssets(lowerBoundSymbol string, limit int) (types.Assets, error) {
	defer p.SetDebug(false)

	if limit > AssetsMaxBatchSize || limit == AssetsListAll {
		limit = AssetsMaxBatchSize
	}

	resp, err := p.wsClient.CallAPI(0, "list_assets", lowerBoundSymbol, limit)
	if err != nil {
		return nil, err
	}

	p.Debug("list_assets <", resp)

	ret := types.Assets{}
	if err := ffjson.Unmarshal(util.ToBytes(resp), &ret); err != nil {
		return nil, errors.Annotate(err, "unmarshal Assets")
	}

	return ret, nil
}

//GetRequiredFees calculates the required fee for each operation in the specified asset type.
//If the asset type does not have a valid core_exchange_rate
func (p *bitsharesAPI) GetRequiredFees(ops types.Operations, feeAsset types.GrapheneObject) (types.AssetAmounts, error) {
	defer p.SetDebug(false)

	resp, err := p.wsClient.CallAPI(0, "get_required_fees", ops.Types(), feeAsset.Id())
	if err != nil {
		return nil, err
	}

	p.Debug("get_required_fees <", resp)

	ret := types.AssetAmounts{}
	if err := ffjson.Unmarshal(util.ToBytes(resp), &ret); err != nil {
		return nil, errors.Annotate(err, "unmarshal AssetAmounts")
	}

	return ret, nil
}

//GetLimitOrders returns a slice of LimitOrder types.
func (p *bitsharesAPI) GetLimitOrders(base, quote types.GrapheneObject, limit int) (types.LimitOrders, error) {
	defer p.SetDebug(false)

	if limit > GetLimitOrdersLimit {
		limit = GetLimitOrdersLimit
	}

	resp, err := p.wsClient.CallAPI(0, "get_limit_orders", base.Id(), quote.Id(), limit)
	if err != nil {
		return nil, err
	}

	p.Debug("get_limit_orders <", resp)

	ret := types.LimitOrders{}
	if err := ffjson.Unmarshal(util.ToBytes(resp), &ret); err != nil {
		return nil, errors.Annotate(err, "unmarshal LimitOrders")
	}

	return ret, nil
}

//GetSettleOrders returns a slice of SettleOrder types.
func (p *bitsharesAPI) GetSettleOrders(assetID types.GrapheneObject, limit int) (types.SettleOrders, error) {
	defer p.SetDebug(false)

	if limit > GetSettleOrdersLimit {
		limit = GetSettleOrdersLimit
	}

	resp, err := p.wsClient.CallAPI(0, "get_settle_orders", assetID.Id(), limit)
	if err != nil {
		return nil, err
	}

	p.Debug("get_settle_orders <", resp)

	ret := types.SettleOrders{}
	if err := ffjson.Unmarshal(util.ToBytes(resp), &ret); err != nil {
		return nil, errors.Annotate(err, "unmarshal SettleOrders")
	}

	return ret, nil
}

//GetCallOrders returns a slice of CallOrder types.
func (p *bitsharesAPI) GetCallOrders(assetID types.GrapheneObject, limit int) (types.CallOrders, error) {
	defer p.SetDebug(false)

	if limit > GetCallOrdersLimit {
		limit = GetCallOrdersLimit
	}

	resp, err := p.wsClient.CallAPI(0, "get_call_orders", assetID.Id(), limit)
	if err != nil {
		return nil, err
	}

	p.Debug("get_call_orders <", resp)

	ret := types.CallOrders{}
	if err := ffjson.Unmarshal(util.ToBytes(resp), &ret); err != nil {
		return nil, errors.Annotate(err, "unmarshal CallOrders")
	}

	return ret, nil
}

//GetMarginPositions returns a slice of CallOrder objects for the debug specified account.
func (p *bitsharesAPI) GetMarginPositions(accountID types.GrapheneObject) (types.CallOrders, error) {
	defer p.SetDebug(false)

	resp, err := p.wsClient.CallAPI(0, "get_margin_positions", accountID.Id())
	if err != nil {
		return nil, err
	}

	p.Debug("get_margin_positions <", resp)

	ret := types.CallOrders{}
	if err := ffjson.Unmarshal(util.ToBytes(resp), &ret); err != nil {
		return nil, errors.Annotate(err, "unmarshal CallOrders")
	}

	return ret, nil
}

//GetTradeHistory returns MarketTrade object.
func (p *bitsharesAPI) GetTradeHistory(base, quote types.GrapheneObject, toTime, fromTime time.Time, limit int) (types.MarketTrades, error) {
	defer p.SetDebug(false)

	if limit > GetTradeHistoryLimit {
		limit = GetTradeHistoryLimit
	}

	resp, err := p.wsClient.CallAPI(0, "get_trade_history", base.Id(), quote.Id(), toTime, fromTime, limit)
	if err != nil {
		return nil, err
	}

	p.Debug("get_trade_history <", resp)

	ret := types.MarketTrades{}
	if err := ffjson.Unmarshal(util.ToBytes(resp), &ret); err != nil {
		return nil, errors.Annotate(err, "unmarshal MarketTrades")
	}

	return ret, nil
}

//GetChainID returns the ID of the chain we are connected to.
func (p *bitsharesAPI) GetChainID() (string, error) {
	defer p.SetDebug(false)

	resp, err := p.wsClient.CallAPI(p.databaseAPIID, "get_chain_id", types.EmptyParams)
	if err != nil {
		return "", err
	}

	p.Debug("get_chain_id <", resp)

	return resp.(string), nil
}

//GetObjects returns a list of Graphene Objects by ID.
func (p *bitsharesAPI) GetObjects(ids ...types.GrapheneObject) ([]interface{}, error) {
	defer p.SetDebug(false)

	params := types.GrapheneObjects(ids).ToStrings()
	resp, err := p.wsClient.CallAPI(0, "get_objects", params)
	if err != nil {
		return nil, err
	}

	p.Debug("get_objects <", resp)

	data := resp.([]interface{})
	ret := make([]interface{}, 0)
	id := types.GrapheneID{}

	for _, obj := range data {
		if obj == nil {
			continue
		}

		if err := id.FromRawData(obj); err != nil {
			return nil, errors.Annotate(err, "from raw data")
		}

		b := util.ToBytes(obj)

		//TODO: implement
		// ObjectTypeBase
		// ObjectTypeWitness
		// ObjectTypeCustom
		// ObjectTypeProposal
		// ObjectTypeWithdrawPermission
		// ObjectTypeVestingBalance
		// ObjectTypeWorker
		switch id.Space() {
		case types.SpaceTypeProtocol:
			switch id.Type() {
			case types.ObjectTypeAccount:
				acc := types.Account{}
				if err := ffjson.Unmarshal(b, &acc); err != nil {
					return nil, errors.Annotate(err, "unmarshal Account")
				}
				ret = append(ret, acc)
			case types.ObjectTypeAsset:
				ass := types.Asset{}
				if err := ffjson.Unmarshal(b, &ass); err != nil {
					return nil, errors.Annotate(err, "unmarshal Asset")
				}
				ret = append(ret, ass)
			case types.ObjectTypeForceSettlement:
				set := types.SettleOrder{}
				if err := ffjson.Unmarshal(b, &set); err != nil {
					return nil, errors.Annotate(err, "unmarshal SettleOrder")
				}
				ret = append(ret, set)
			case types.ObjectTypeLimitOrder:
				lim := types.LimitOrder{}
				if err := ffjson.Unmarshal(b, &lim); err != nil {
					return nil, errors.Annotate(err, "unmarshal LimitOrder")
				}
				ret = append(ret, lim)
			case types.ObjectTypeCallOrder:
				cal := types.CallOrder{}
				if err := ffjson.Unmarshal(b, &cal); err != nil {
					return nil, errors.Annotate(err, "unmarshal CallOrder")
				}
				ret = append(ret, cal)
			case types.ObjectTypeCommiteeMember:
				mem := types.CommiteeMember{}
				if err := ffjson.Unmarshal(b, &mem); err != nil {
					return nil, errors.Annotate(err, "unmarshal CommiteeMember")
				}
				ret = append(ret, mem)
			case types.ObjectTypeOperationHistory:
				hist := types.OperationHistory{}
				if err := ffjson.Unmarshal(b, &hist); err != nil {
					return nil, errors.Annotate(err, "unmarshal OperationHistory")
				}
				ret = append(ret, hist)
			case types.ObjectTypeBalance:
				bal := types.Balance{}
				if err := ffjson.Unmarshal(b, &bal); err != nil {
					return nil, errors.Annotate(err, "unmarshal Balance")
				}
				ret = append(ret, bal)

			default:
				util.DumpUnmarshaled(id.Type().String(), b)
				return nil, errors.Errorf("unable to parse GrapheneObject with ID %s", id)
			}

		case types.SpaceTypeImplementation:
			switch id.Type() {
			case types.ObjectTypeAssetBitAssetData:
				bit := types.BitAssetData{}
				if err := ffjson.Unmarshal(b, &bit); err != nil {
					return nil, errors.Annotate(err, "unmarshal BitAssetData")
				}
				ret = append(ret, bit)

			default:
				util.DumpUnmarshaled(id.Type().String(), b)
				return nil, errors.Errorf("unable to parse GrapheneObject with ID %s", id)
			}
		}
	}

	return ret, nil
}

// CancelOrder cancels an order given by orderID
func (p *bitsharesAPI) CancelOrder(orderID types.GrapheneObject, broadcast bool) (*types.Transaction, error) {
	defer p.SetDebug(false)

	resp, err := p.wsClient.CallAPI(0, "cancel_order", orderID.Id(), broadcast)
	if err != nil {
		return nil, err
	}

	p.Debug("cancel_order <", resp)

	ret := types.Transaction{}
	if err := ffjson.Unmarshal(util.ToBytes(resp), &ret); err != nil {
		return nil, errors.Annotate(err, "unmarshal Transaction")
	}

	return &ret, nil
}

func (p bitsharesAPI) Debug(descr string, in interface{}) {
	if p.debug {
		util.Dump(descr, in)
	}
}

func (p *bitsharesAPI) SetDebug(debug bool) {
	p.debug = debug

	if p.wsClient != nil {
		p.wsClient.SetDebug(p.debug)
	}

	if p.rpcClient != nil {
		p.rpcClient.SetDebug(p.debug)
	}

}

func (p *bitsharesAPI) DatabaseAPIID() int {
	return p.databaseAPIID
}

func (p *bitsharesAPI) BroadcastAPIID() int {
	return p.broadcastAPIID
}

func (p *bitsharesAPI) HistoryAPIID() int {
	return p.historyAPIID
}

func (p *bitsharesAPI) CryptoAPIID() int {
	return p.cryptoAPIID
}

func (p *bitsharesAPI) CallWsAPI(apiID int, method string, args ...interface{}) (interface{}, error) {
	defer p.SetDebug(false)
	return p.wsClient.CallAPI(apiID, method, args...)
}

//OnError hook your notify callback here
func (p *bitsharesAPI) OnNotify(subscriberID int, notifyFn func(msg interface{}) error) error {
	return p.wsClient.OnNotify(subscriberID, notifyFn)
}

//OnError hook your error callback here
func (p *bitsharesAPI) OnError(errorFn func(err error)) {
	p.wsClient.OnError(errorFn)
}

//SetCredentials defines username and password for login.
func (p *bitsharesAPI) SetCredentials(username, password string) {
	p.username = username
	p.password = password
}

// Connect initializes the API and connects underlying clients
func (p *bitsharesAPI) Connect() (err error) {
	if err := p.wsClient.Connect(); err != nil {
		return errors.Annotate(err, "ws connect")
	}

	if p.rpcClient != nil {
		if err := p.rpcClient.Connect(); err != nil {
			return errors.Annotate(err, "rpc connect")
		}
	}

	if ok, err := p.login(); err != nil || !ok {
		if err != nil {
			return errors.Annotate(err, "login")
		}
		return errors.New("login failed")
	}

	if err := p.getAPIIDs(); err != nil {
		return errors.Annotate(err, "getApiIDs")
	}

	chainID, err := p.GetChainID()
	if err != nil {
		return errors.Annotate(err, "GetChainID")
	}

	if err = config.SetCurrentConfig(chainID); err != nil {
		return errors.Annotate(err, "SetCurrentConfig")
	}

	return nil
}

func (p *bitsharesAPI) getAPIIDs() (err error) {
	p.databaseAPIID, err = p.getApiID("database")
	if err != nil {
		return errors.Annotate(err, "database")
	}

	p.historyAPIID, err = p.getApiID("history")
	if err != nil {
		return errors.Annotate(err, "history")
	}

	p.broadcastAPIID, err = p.getApiID("network_broadcast")
	if err != nil {
		return errors.Annotate(err, "network")
	}

	p.cryptoAPIID, err = p.getApiID("crypto")
	if err != nil {
		return errors.Annotate(err, "crypto")
	}

	return nil
}

//Close() shuts down the api and closes underlying clients.
func (p *bitsharesAPI) Close() error {
	if p.rpcClient != nil {
		if err := p.rpcClient.Close(); err != nil {
			return errors.Annotate(err, "close rpc client")
		}
		p.rpcClient = nil
	}

	if p.wsClient != nil {
		if err := p.wsClient.Close(); err != nil {
			return errors.Annotate(err, "close ws client")
		}
		p.wsClient = nil
	}

	return nil
}

//New creates a new BitsharesAPI interface.
//Param wsEndpointURL is mandatory.
//Param rpcEndpointURL is optional
func New(wsEndpointURL, rpcEndpointURL string) BitsharesAPI {
	var rpcClient client.RPCClient
	if rpcEndpointURL != "" {
		rpcClient = client.NewRPCClient(rpcEndpointURL)
	}

	api := bitsharesAPI{
		wsClient:       client.NewWebsocketClient(wsEndpointURL),
		rpcClient:      rpcClient,
		databaseAPIID:  InvalidApiID,
		historyAPIID:   InvalidApiID,
		broadcastAPIID: InvalidApiID,
		cryptoAPIID:    InvalidApiID,
		debug:          false,
	}

	return &api
}
