package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeReviseContract] = func() types.Operation {
		op := &ReviseContractOperation{}
		return op
	}
}

type ReviseContractOperation struct {
	types.OperationFee
	Reviser    types.AccountID  `json:"reviser"`
	ContractID types.ContractID `json:"contract_id"`
	Data       string           `json:"data"`
	Extensions types.Extensions `json:"extensions"`
}

func (p ReviseContractOperation) Type() types.OperationType {
	return types.OperationTypeReviseContract
}

func (p ReviseContractOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
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

func (p ReviseContractOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.Reviser); err != nil {
		return errors.Annotate(err, "encode Reviser")
	}

	if err := enc.Encode(p.ContractID); err != nil {
		return errors.Annotate(err, "encode ContractID")
	}

	if err := enc.Encode(p.Data); err != nil {
		return errors.Annotate(err, "encode Data")
	}

	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}

	return nil
}
