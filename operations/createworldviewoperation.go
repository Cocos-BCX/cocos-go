package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeCreateWorldView] = func() types.Operation {
		op := &CreateWorldViewOperation{}
		return op
	}
}

type CreateWorldViewOperation struct {
	types.OperationFee
	FeePayingAccount types.AccountID `json:"fee_paying_account"`
	WorldView        string          `json:"world_view"`
}

func (p CreateWorldViewOperation) Type() types.OperationType {
	return types.OperationTypeCreateWorldView
}

func (p CreateWorldViewOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if fee, ok := params["fee"]; ok {
		if err := enc.Encode(types.UInt64(fee.(float64))); err != nil {
			return errors.Annotate(err, "encode Fee")
		}
	}

	return nil
}

func (p CreateWorldViewOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.FeePayingAccount); err != nil {
		return errors.Annotate(err, "encode FeePayingAccount")
	}

	if err := enc.Encode(p.WorldView); err != nil {
		return errors.Annotate(err, "encode WorldView")
	}

	return nil
}
