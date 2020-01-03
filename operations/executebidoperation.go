package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeExecuteBid] = func() types.Operation {
		op := &ExecuteBidOperation{}
		return op
	}
}

type ExecuteBidOperation struct {
	types.OperationFee
	Bidder     types.AccountID   `json:"bidder"`
	Debt       types.AssetAmount `json:"debt"`
	Collateral types.AssetAmount `json:"collateral"`
}

func (p ExecuteBidOperation) Type() types.OperationType {
	return types.OperationTypeExecuteBid
}

func (p ExecuteBidOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	return nil
}

func (p ExecuteBidOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode Fee")
	}

	if err := enc.Encode(p.Bidder); err != nil {
		return errors.Annotate(err, "encode Bidder")
	}

	if err := enc.Encode(p.Debt); err != nil {
		return errors.Annotate(err, "encode Debt")
	}

	if err := enc.Encode(p.Collateral); err != nil {
		return errors.Annotate(err, "encode Collateral")
	}

	return nil
}
