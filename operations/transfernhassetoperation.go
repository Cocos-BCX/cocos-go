package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeTransferNhAsset] = func() types.Operation {
		op := &CreateNhAssetOperation{}
		return op
	}
}

type TransferNhAssetOperation struct {
	types.OperationFee
	From    types.AccountID `json:"from"`
	To      types.AccountID `json:"to"`
	NHAsset types.NHAssetID `json:"nh_asset"`
}

func (p TransferNhAssetOperation) Type() types.OperationType {
	return types.OperationTypeTransferNhAsset
}

func (p TransferNhAssetOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if fee, ok := params["fee"]; ok {
		if err := enc.Encode(types.UInt64(fee.(float64))); err != nil {
			return errors.Annotate(err, "encode Fee")
		}
	}

	return nil
}

func (p TransferNhAssetOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.From); err != nil {
		return errors.Annotate(err, "encode From")
	}

	if err := enc.Encode(p.To); err != nil {
		return errors.Annotate(err, "encode To")
	}

	if err := enc.Encode(p.NHAsset); err != nil {
		return errors.Annotate(err, "encode NHAsset")
	}

	return nil
}
