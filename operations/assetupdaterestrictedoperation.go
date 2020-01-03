package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAssetUpdateRestricted] = func() types.Operation {
		op := &AssetUpdateRestrictedOperation{}
		return op
	}
}

type ObjectIDs []types.ObjectID

type AssetUpdateRestrictedOperation struct {
	types.OperationFee
	Payer          types.AccountID  `json:"payer"`
	TargetAsset    types.AssetID    `json:"target_asset"`
	IsAdd          bool             `json:"isadd"`
	RestrictedType types.UInt8      `json:"restricted_type"`
	RestrictedList ObjectIDs        `json:"restricted_List"`
	Extensions     types.Extensions `json:"extensions"`
}

func (p AssetUpdateRestrictedOperation) Type() types.OperationType {
	return types.OperationTypeAssetUpdateRestricted
}

func (p AssetUpdateRestrictedOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
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

func (p AssetUpdateRestrictedOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.Payer); err != nil {
		return errors.Annotate(err, "encode Payer")
	}

	if err := enc.Encode(p.TargetAsset); err != nil {
		return errors.Annotate(err, "encode TargetAsset")
	}

	if err := enc.Encode(p.IsAdd); err != nil {
		return errors.Annotate(err, "encode IsAdd")
	}

	if err := enc.Encode(p.RestrictedType); err != nil {
		return errors.Annotate(err, "encode RestrictedType")
	}

	if err := enc.Encode(p.RestrictedList); err != nil {
		return errors.Annotate(err, "encode RestrictedList")
	}

	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}

	return nil
}
