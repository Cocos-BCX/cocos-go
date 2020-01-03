package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeCreateNhAsset] = func() types.Operation {
		op := &CreateNhAssetOperation{}
		return op
	}
}

type CreateNhAssetOperation struct {
	types.OperationFee
	FeePayingAccount types.AccountID `json:"fee_paying_account"`
	Owner            types.AccountID `json:"owner"`
	AssetID          string          `json:"asset_id"`
	WorldView        string          `json:"world_view"`
	BaseDescribe     string          `json:"base_describe"`
	Collateral       types.Int64     `json:"collateral"`
}

func (p CreateNhAssetOperation) Type() types.OperationType {
	return types.OperationTypeCreateNhAsset
}

func (p CreateNhAssetOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if fee, ok := params["fee"]; ok {
		if err := enc.Encode(types.UInt64(fee.(float64))); err != nil {
			return errors.Annotate(err, "encode Fee")
		}
	}

	if ppk, ok := params["price_per_kbyte"]; ok {
		if err := enc.Encode(types.UInt32(ppk.(float64))); err != nil {
			return errors.Annotate(err, "encode PricePerKByte")
		}
	}

	return nil
}

func (p CreateNhAssetOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.FeePayingAccount); err != nil {
		return errors.Annotate(err, "encode FeePayingAccount")
	}

	if err := enc.Encode(p.Owner); err != nil {
		return errors.Annotate(err, "encode Owner")
	}

	if err := enc.Encode(p.AssetID); err != nil {
		return errors.Annotate(err, "encode AssetID")
	}

	if err := enc.Encode(p.WorldView); err != nil {
		return errors.Annotate(err, "encode WorldView")
	}

	if err := enc.Encode(p.BaseDescribe); err != nil {
		return errors.Annotate(err, "encode BaseDescribe")
	}

	return nil
}
