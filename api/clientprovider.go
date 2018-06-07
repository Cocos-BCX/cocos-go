package api

import (
	"log"
	"sync"

	"github.com/denkhaus/bitshares/client"
	"github.com/denkhaus/bitshares/latency"
	"github.com/juju/errors"
)

type ClientProvider interface {
	OnError(fn client.ErrorFunc)
	OnNotify(subscriberID int, fn client.NotifyFunc) error
	CallAPI(apiID int, method string, args ...interface{}) (interface{}, error)
	SetDebug(debug bool)
	Close() error
}

type SimpleClientProvider struct {
	client.WebsocketClient
}

func NewSimpleClientProvider(endpointURL string) *SimpleClientProvider {
	sim := SimpleClientProvider{
		WebsocketClient: client.NewWebsocketClient(endpointURL),
	}

	return &sim
}

func (p *SimpleClientProvider) CallAPI(apiID int, method string, args ...interface{}) (interface{}, error) {
	if !p.WebsocketClient.IsConnected() {
		if err := p.Connect(); err != nil {
			return nil, errors.Annotate(err, "Connect")
		}
	}

	return p.WebsocketClient.CallAPI(apiID, method, args...)
}

type BestNodeClientProvider struct {
	mu sync.Mutex
	client.WebsocketClient
	tester latency.LatencyTester
}

func NewBestNodeClientProvider(endpointURL string) (*BestNodeClientProvider, error) {
	tester, err := latency.NewLatencyTester(endpointURL)
	if err != nil {
		return nil, errors.Annotate(err, "NewLatencyTester")
	}

	tester.Start()
	pr := &BestNodeClientProvider{
		tester:          tester,
		WebsocketClient: tester.TopNodeClient(),
	}

	tester.OnTopNodeChanged(pr.onTopNodeChanged)
	return pr, nil
}

func (p *BestNodeClientProvider) onTopNodeChanged(newEndpoint string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.WebsocketClient.IsConnected() {
		p.WebsocketClient.Close()
	}

	p.WebsocketClient = p.tester.TopNodeClient()
	log.Println("top node client changed")
}

func (p *BestNodeClientProvider) CallAPI(apiID int, method string, args ...interface{}) (interface{}, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.WebsocketClient.IsConnected() {
		if err := p.Connect(); err != nil {
			return nil, errors.Annotate(err, "Connect")
		}
	}

	return p.WebsocketClient.CallAPI(apiID, method, args...)
}

func (p *BestNodeClientProvider) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.WebsocketClient.IsConnected() {
		p.WebsocketClient.Close()
		if err := p.WebsocketClient.Close(); err != nil {
			return errors.Annotate(err, "Close [client]")
		}
	}

	if err := p.tester.Close(); err != nil {
		return errors.Annotate(err, "Close [tester]")
	}

	return nil
}
