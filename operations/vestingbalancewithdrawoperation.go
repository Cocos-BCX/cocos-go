package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeVestingBalanceWithdraw] = func() types.Operation {
		op := &VestingBalanceWithdrawOperation{}
		return op
	}
}

type VestingBalanceWithdrawOperation struct {
	types.OperationFee
	VestingBalance types.VestingBalanceID `json:"vesting_balance"`
	Owner          types.AccountID        `json:"owner"`
	Amount         types.AssetAmount      `json:"amount"`
}

func (p VestingBalanceWithdrawOperation) Type() types.OperationType {
	return types.OperationTypeVestingBalanceWithdraw
}

func (p VestingBalanceWithdrawOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.VestingBalance); err != nil {
		return errors.Annotate(err, "encode VestingBalance")
	}

	if err := enc.Encode(p.Owner); err != nil {
		return errors.Annotate(err, "encode Owner")
	}

	if err := enc.Encode(p.Amount); err != nil {
		return errors.Annotate(err, "encode Amount")
	}

	return nil
}
