package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeCancelNhAssetOrder] = func() types.Operation {
		op := &CancelNhAssetOrderOperation{}
		return op
	}
}

type CancelNhAssetOrderOperation struct {
	types.OperationFee
	Order            types.NHAssetOrderID `json:"order"`
	FeePayingAccount types.AccountID      `json:"fee_paying_account"`
	Extensions       types.Extensions     `json:"extensions"`
}

func (p CancelNhAssetOrderOperation) Type() types.OperationType {
	return types.OperationTypeCancelNhAssetOrder
}

func (p CancelNhAssetOrderOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if fee, ok := params["fee"]; ok {
		if err := enc.Encode(types.UInt64(fee.(float64))); err != nil {
			return errors.Annotate(err, "encode Fee")
		}
	}

	return nil
}

func (p CancelNhAssetOrderOperation) Marshal(enc *util.TypeEncoder) error {
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

	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode Extensions")
	}

	return nil
}
