package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAssetClaimFees] = func() types.Operation {
		op := &AssetClaimFeesOperation{}
		return op
	}
}

type AssetClaimFeesOperation struct {
	types.OperationFee
	Issuer        types.AccountID               `json:"issuer"`
	AmountToClaim types.AssetAmount             `json:"amount_to_claim"`
	Extensions    types.AccountUpdateExtensions `json:"extensions"`
}

func (p AssetClaimFeesOperation) Type() types.OperationType {
	return types.OperationTypeAssetClaimFees
}

func (p AssetClaimFeesOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
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

func (p AssetClaimFeesOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode Fee")
	}

	if err := enc.Encode(p.Issuer); err != nil {
		return errors.Annotate(err, "encode Issuer")
	}

	if err := enc.Encode(p.AmountToClaim); err != nil {
		return errors.Annotate(err, "encode AmountToClaim")
	}

	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}

	return nil
}
