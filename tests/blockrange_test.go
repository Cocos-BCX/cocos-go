package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/bradhe/stopwatch"
	"github.com/stretchr/testify/assert"
	// register operations

	_ "github.com/denkhaus/bitshares/operations"
	"github.com/denkhaus/bitshares/util"
)

func TestBlockRange(t *testing.T) {
	api := NewTestAPI(t, WsFullApiUrl)
	defer api.Close()

	block := uint64(26878913)

	for {
		bl, err := api.GetBlock(block)
		if err != nil {
			assert.FailNow(t, err.Error(), "GetBlock")
		}

		nTrx := uint64(len(bl.Transactions))
		fmt.Printf("block %d: binary compare %d transactions\n", block, nTrx)
		watch := stopwatch.Start()

		for _, tx := range bl.Transactions {
			time.Sleep(300 * time.Millisecond) // to avoid EOF from client
			ref, test, err := CompareTransactions(api, &tx, false)
			if err != nil {
				util.Dump("trx", tx)
				assert.FailNow(t, err.Error(), "CompareTransactions")
				return
			}

			if !assert.Equal(t, ref, test) {
				util.Dump("trx", tx)
				return
			}
		}

		watch.Stop()
		fmt.Printf("ms/trx:%v\n", time.Duration(uint64(watch.Milliseconds())/nTrx))
		block++
	}

	//util.Dump("get_block >", res)
}
