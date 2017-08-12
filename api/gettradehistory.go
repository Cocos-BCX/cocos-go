package api

import (
	"time"

	"github.com/denkhaus/bitshares/objects"
	"github.com/juju/errors"
	"github.com/pquerna/ffjson/ffjson"
)

const (
	GetTradeHistoryLimit = 100
)

//GetTradeHistory returns MarketTrade object.
func (p *BitsharesApi) GetTradeHistory(base, quote string, toTime, fromTime time.Time, limit int) ([]objects.MarketTrade, error) {
	if limit > GetTradeHistoryLimit {
		limit = GetTradeHistoryLimit
	}

	resp, err := p.client.CallApi(0, "get_trade_history", base, quote, toTime, fromTime, limit)
	if err != nil {
		return nil, errors.Annotate(err, "get_trade_history")
	}

	//spew.Dump(resp)
	data := resp.([]interface{})
	ret := make([]objects.MarketTrade, len(data))

	for idx, a := range data {
		if err := ffjson.Unmarshal(toBytes(a), &ret[idx]); err != nil {
			return nil, errors.Annotate(err, "unmarshal MarketTrade")
		}
	}

	return ret, nil
}
