package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeDeleteNhAsset] = func() types.Operation {
		op := &DeleteNhAssetOperation{}
		return op
	}
}

type DeleteNhAssetOperation struct {
	types.OperationFee
	FeePayingAccount types.AccountID `json:"fee_paying_account"`
	NHAsset          types.NHAssetID `json:"nh_asset"`
}

func (p DeleteNhAssetOperation) Type() types.OperationType {
	return types.OperationTypeDeleteNhAsset
}

func (p DeleteNhAssetOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if fee, ok := params["fee"]; ok {
		if err := enc.Encode(types.UInt64(fee.(float64))); err != nil {
			return errors.Annotate(err, "encode Fee")
		}
	}

	return nil
}

func (p DeleteNhAssetOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.FeePayingAccount); err != nil {
		return errors.Annotate(err, "encode FeePayingAccount")
	}

	if err := enc.Encode(p.NHAsset); err != nil {
		return errors.Annotate(err, "encode NHAsset")
	}

	return nil
}
