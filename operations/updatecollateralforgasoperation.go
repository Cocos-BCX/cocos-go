package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeUpdateCollateralForGas] = func() types.Operation {
		op := &UpdateCollateralForGasOperation{}
		return op
	}
}

type UpdateCollateralForGasOperation struct {
	types.OperationFee
	Mortgager   types.AccountID `json:"mortgager"`
	Beneficiary types.AccountID `json:"beneficiary"`
	Collateral  types.Int64     `json:"collateral"`
}

func (p UpdateCollateralForGasOperation) Type() types.OperationType {
	return types.OperationTypeUpdateCollateralForGas
}

func (p UpdateCollateralForGasOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
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

func (p UpdateCollateralForGasOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.Mortgager); err != nil {
		return errors.Annotate(err, "encode Mortgager")
	}

	if err := enc.Encode(p.Beneficiary); err != nil {
		return errors.Annotate(err, "encode Beneficiary")
	}

	if err := enc.Encode(p.Collateral); err != nil {
		return errors.Annotate(err, "encode Collateral")
	}

	return nil
}
