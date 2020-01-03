package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAssetGlobalSettle] = func() types.Operation {
		op := &AssetGlobalSettleOperation{}
		return op
	}
}

type AssetGlobalSettleOperation struct {
	types.OperationFee
	Issuer        types.AccountID  `json:"issuer"`
	AssetToSettle types.AssetID    `json:"asset_to_settle"`
	SettlePrice   types.Price      `json:"settle_price"`
	Extensions    types.Extensions `json:"extensions"`
}

func (p AssetGlobalSettleOperation) Type() types.OperationType {
	return types.OperationTypeAssetGlobalSettle
}

func (p AssetGlobalSettleOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode Fee")
	}
	if err := enc.Encode(p.Issuer); err != nil {
		return errors.Annotate(err, "encode Issuer")
	}
	if err := enc.Encode(p.AssetToSettle); err != nil {
		return errors.Annotate(err, "encode AssetToSettle")
	}
	if err := enc.Encode(p.SettlePrice); err != nil {
		return errors.Annotate(err, "encode SettlePrice")
	}
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode Extensions")
	}
	return nil
}
