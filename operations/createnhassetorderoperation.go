package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeCreateNhAssetOrder] = func() types.Operation {
		op := &CreateNhAssetOrderOperation{}
		return op
	}
}

type CreateNhAssetOrderOperation struct {
	types.OperationFee
	Seller           types.AccountID   `json:"seller"`
	OTCAccount       types.AccountID   `json:"otcaccount"`
	PendingOrdersFee types.AssetAmount `json:"pending_orders_fee"`
	NHAsset          types.NHAssetID   `json:"nh_asset"`
	Memo             string            `json:"memo"`
	Price            types.AssetAmount `json:"price"`
	Expiration       types.Time        `json:"expiration"`
}

func (p CreateNhAssetOrderOperation) Type() types.OperationType {
	return types.OperationTypeCreateNhAssetOrder
}

func (p CreateNhAssetOrderOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if fee, ok := params["fee"]; ok {
		if err := enc.Encode(types.UInt64(fee.(float64))); err != nil {
			return errors.Annotate(err, "encode Fee")
		}
	}

	return nil
}

func (p CreateNhAssetOrderOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.Seller); err != nil {
		return errors.Annotate(err, "encode Seller")
	}

	if err := enc.Encode(p.OTCAccount); err != nil {
		return errors.Annotate(err, "encode OTCAccount")
	}

	if err := enc.Encode(p.PendingOrdersFee); err != nil {
		return errors.Annotate(err, "encode PendingOrdersFee")
	}

	if err := enc.Encode(p.NHAsset); err != nil {
		return errors.Annotate(err, "encode NHAsset")
	}

	if err := enc.Encode(p.Memo); err != nil {
		return errors.Annotate(err, "encode Memo")
	}

	if err := enc.Encode(p.Price); err != nil {
		return errors.Annotate(err, "encode Price")
	}

	if err := enc.Encode(p.Expiration); err != nil {
		return errors.Annotate(err, "encode Expiration")
	}

	return nil
}
