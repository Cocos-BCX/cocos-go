package api

import (
	"time"

	"github.com/denkhaus/bitshares/objects"
	"github.com/denkhaus/bitshares/rpc"
	"github.com/juju/errors"
)

const (
	InvalidApiID = -1
)

var (
	EmptyParams = []interface{}{}
)

type BitsharesAPI interface {
	Close() error
	Connect() error
	DatabaseApiID() int
	CryptoApiID() int
	HistoryApiID() int
	NetworkApiID() int
	SetCredentials(username, password string)
	OnError(func(error))
	OnNotify(subscriberID int, notifyFn func(msg interface{}) error) error
	CallAPI(apiID int, method string, args ...interface{}) (interface{}, error)
	SetSubscribeCallback(notifyID int, clearFilter bool) error
	CancelAllSubscriptions() error
	SubscribeToMarket(notifyID int, base objects.GrapheneObject, quote objects.GrapheneObject) error
	UnsubscribeFromMarket(base objects.GrapheneObject, quote objects.GrapheneObject) error
	GetAccountBalances(account objects.GrapheneObject, assets ...objects.GrapheneObject) ([]objects.AssetAmount, error)
	GetAccountByName(name string) (*objects.Account, error)
	GetAccounts(accountIDs ...objects.GrapheneObject) ([]objects.Account, error)
	GetCallOrders(assetID objects.GrapheneObject, limit int) ([]objects.CallOrder, error)
	GetLimitOrders(base, quote objects.GrapheneObject, limit int) ([]objects.LimitOrder, error)
	GetObjects(objectIDs ...objects.GrapheneObject) ([]interface{}, error)
	GetSettleOrders(assetID objects.GrapheneObject, limit int) ([]objects.SettleOrder, error)
	GetTradeHistory(base, quote string, toTime, fromTime time.Time, limit int) ([]objects.MarketTrade, error)
	ListAssets(lowerBoundSymbol string, limit int) ([]objects.Asset, error)
	GetChainID() (string, error)
}

type bitsharesAPI struct {
	client        rpc.WebsocketClient
	chainConfig   *ChainConfig
	username      string
	password      string
	databaseApiID int
	historyApiID  int
	cryptoApiID   int
	networkApiID  int
}

func (p *bitsharesAPI) DatabaseApiID() int {
	return p.databaseApiID
}

func (p *bitsharesAPI) NetworkApiID() int {
	return p.networkApiID
}

func (p *bitsharesAPI) HistoryApiID() int {
	return p.historyApiID
}

func (p *bitsharesAPI) CryptoApiID() int {
	return p.cryptoApiID
}

func (p *bitsharesAPI) getApiID(identifier string) (int, error) {
	resp, err := p.client.CallAPI(1, identifier, EmptyParams)
	if err != nil {
		return InvalidApiID, errors.Annotate(err, identifier)
	}

	//util.Dump(identifier+" in", resp)
	return int(resp.(float64)), nil
}

func (p *bitsharesAPI) login() (bool, error) {
	resp, err := p.client.CallAPI(1, "login", p.username, p.password)
	if err != nil {
		return false, errors.Annotate(err, "login")
	}

	//util.Dump("login in", resp)
	return resp.(bool), nil
}

func (p *bitsharesAPI) SetSubscribeCallback(notifyID int, clearFilter bool) error {

	// returns nil if successfull
	_, err := p.client.CallAPI(p.databaseApiID, "set_subscribe_callback", notifyID, clearFilter)
	if err != nil {
		return errors.Annotate(err, "set_subscribe_callback")
	}

	return nil
}

func (p *bitsharesAPI) CancelAllSubscriptions() error {

	// returns nil
	_, err := p.client.CallAPI(p.databaseApiID, "cancel_all_subscriptions", EmptyParams)
	if err != nil {
		return errors.Annotate(err, "cancel_all_subscriptions")
	}
	
	return nil
}

func (p *bitsharesAPI) CallAPI(apiID int, method string, args ...interface{}) (interface{}, error) {
	return p.client.CallAPI(apiID, method, args...)
}

func (p *bitsharesAPI) OnNotify(subscriberID int, notifyFn func(msg interface{}) error) error {
	return p.client.OnNotify(subscriberID, notifyFn)
}

func (p *bitsharesAPI) OnError(errorFn func(err error)) {
	p.client.OnError(errorFn)
}

//SetCredentials defines username and password for login.
func (p *bitsharesAPI) SetCredentials(username, password string) {
	p.username = username
	p.password = password
}

func (p *bitsharesAPI) Connect() (err error) {
	if err := p.client.Connect(); err != nil {
		return errors.Annotate(err, "connect")
	}

	if ok, err := p.login(); err != nil || !ok {
		if err != nil {
			return errors.Annotate(err, "login")
		}
		return errors.New("login not successful")
	}

	p.databaseApiID, err = p.getApiID("database")
	if err != nil {
		return errors.Annotate(err, "get database API ID")
	}

	p.historyApiID, err = p.getApiID("history")
	if err != nil {
		return errors.Annotate(err, "get history API ID")
	}

	p.networkApiID, err = p.getApiID("network_broadcast")
	if err != nil {
		return errors.Annotate(err, "get network API ID")
	}

	p.cryptoApiID, err = p.getApiID("crypto")
	if err != nil {
		return errors.Annotate(err, "get crypto API ID")
	}

	chainID, err := p.GetChainID()
	if err != nil {
		return errors.Annotate(err, "get chain ID")
	}

	p.chainConfig, err = p.GetChainConfig(chainID)
	if err != nil {
		return errors.Annotate(err, "get chain config")
	}

	return nil
}

//Close() shuts the api down and closes the underlying connection.
func (p *bitsharesAPI) Close() error {
	var err error
	if p.client != nil {
		err = p.client.Close()
		p.client = nil
	}

	return err
}

//New creates a new BitsharesAPI interface.
func New(endpointURL string) BitsharesAPI {
	client := rpc.NewWebsocketClient(endpointURL)

	api := bitsharesAPI{
		client:        client,
		databaseApiID: InvalidApiID,
		historyApiID:  InvalidApiID,
		networkApiID:  InvalidApiID,
		cryptoApiID:   InvalidApiID,
	}

	return &api
}
