package api

import (
	"github.com/denkhaus/bitshares/objects"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func (p *bitsharesAPI) Broadcast(privKeys [][]byte, feeAsset objects.GrapheneObject, ops ...objects.Operation) (string, error) {

	operations := objects.Operations(ops)
	fees, err := p.GetRequiredFees(operations, feeAsset)
	if err != nil {
		return "", errors.Annotate(err, "GetRequiredFees")
	}

	if err := operations.ApplyFees(fees); err != nil {
		return "", errors.Annotate(err, "ApplyFees")
	}

	prop, err := p.GetDynamicGlobalProperties()
	if err != nil {
		return "", errors.Annotate(err, "GetDynamicGlobalProperties")
	}

	tx := objects.NewTransaction()
	tx.Operations = operations

	if err := tx.Sign(privKeys, prop, p.chainConfig.Id()); err != nil {
		return "", errors.Annotate(err, "Sign")
	}

	util.DumpJSON("tx >", tx)

	resp, err := p.BroadcastTransaction(tx)
	if err != nil {
		return "", errors.Annotate(err, "BroadcastTransaction")
	}

	return resp, err
}
