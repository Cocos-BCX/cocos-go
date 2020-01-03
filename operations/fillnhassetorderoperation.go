package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeFillNhAssetOrder] = func() types.Operation {
		op := &FillNhAssetOrderOperation{}
		return op
	}
}

type FillNhAssetOrderOperation struct {
	types.OperationFee
	Order            types.NHAssetOrderID `json:"order"`
	FeePayingAccount types.AccountID      `json:"fee_paying_account"`
	Seller           types.AccountID      `json:"seller"`
	NHAsset          types.NHAssetID      `json:"nh_asset"`
	PriceAmount      string               `json:"price_amount"`
	PriceAssetID     types.AssetID        `json:"price_asset_id"`
	PriceAssetSymbol string               `json:"price_asset_symbol"`
	Expiration       types.Time           `json:"expiration"`
}

func (p FillNhAssetOrderOperation) Type() types.OperationType {
	return types.OperationTypeFillNhAssetOrder
}

func (p FillNhAssetOrderOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if fee, ok := params["fee"]; ok {
		if err := enc.Encode(types.UInt64(fee.(float64))); err != nil {
			return errors.Annotate(err, "encode Fee")
		}
	}

	return nil
}

func (p FillNhAssetOrderOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.Order); err != nil {
		return errors.Annotate(err, "encode Order")
	}

	if err := enc.Encode(p.FeePayingAccount); err != nil {
		return errors.Annotate(err, "encode FeePayingAccount")
	}

	if err := enc.Encode(p.Seller); err != nil {
		return errors.Annotate(err, "encode Seller")
	}

	if err := enc.Encode(p.NHAsset); err != nil {
		return errors.Annotate(err, "encode NHAsset")
	}

	if err := enc.Encode(p.PriceAmount); err != nil {
		return errors.Annotate(err, "encode PriceAmount")
	}

	if err := enc.Encode(p.PriceAssetID); err != nil {
		return errors.Annotate(err, "encode PriceAssetID")
	}

	if err := enc.Encode(p.PriceAssetSymbol); err != nil {
		return errors.Annotate(err, "encode PriceAssetSymbol")
	}

	if err := enc.Encode(p.Expiration); err != nil {
		return errors.Annotate(err, "encode Expiration")
	}

	return nil
}
