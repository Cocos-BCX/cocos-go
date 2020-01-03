package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAssetUpdateFeedProducers] = func() types.Operation {
		op := &AssetUpdateFeedProducersOperation{}
		return op
	}
}

type AssetUpdateFeedProducersOperation struct {
	types.OperationFee
	Issuer           types.AccountID  `json:"issuer"`
	AssetToUpdate    types.AssetID    `json:"asset_to_update"`
	NewFeedProducers types.AccountIDs `json:"new_feed_producers"`
	Extensions       types.Extensions `json:"extensions"`
}

func (p AssetUpdateFeedProducersOperation) Type() types.OperationType {
	return types.OperationTypeAssetUpdateFeedProducers
}

func (p AssetUpdateFeedProducersOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode Fee")
	}

	if err := enc.Encode(p.Issuer); err != nil {
		return errors.Annotate(err, "encode Issuer")
	}

	if err := enc.Encode(p.AssetToUpdate); err != nil {
		return errors.Annotate(err, "encode AssetToUpdate")
	}

	if err := enc.Encode(p.NewFeedProducers); err != nil {
		return errors.Annotate(err, "encode NewFeedProducers")
	}

	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode Extensions")
	}

	return nil
}
