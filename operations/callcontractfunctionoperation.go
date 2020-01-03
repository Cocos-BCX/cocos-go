package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeCallContractFunction] = func() types.Operation {
		op := &ReviseContractOperation{}
		return op
	}
}

type LuaTypes struct {
}

type LuaTypesVec []LuaTypes

type CallContractFunctionOperation struct {
	types.OperationFee
	Caller       types.AccountID  `json:"caller"`
	ContractID   types.ContractID `json:"contract_id"`
	FunctionName string           `json:"function_name"`
	ValueList    LuaTypesVec      `json:"value_list"`
	Extensions   types.Extensions `json:"extensions"`
}

func (p CallContractFunctionOperation) Type() types.OperationType {
	return types.OperationTypeCallContractFunction
}

func (p CallContractFunctionOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
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

func (p CallContractFunctionOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.Caller); err != nil {
		return errors.Annotate(err, "encode Caller")
	}

	if err := enc.Encode(p.ContractID); err != nil {
		return errors.Annotate(err, "encode ContractID")
	}

	if err := enc.Encode(p.FunctionName); err != nil {
		return errors.Annotate(err, "encode FunctionName")
	}

	if err := enc.Encode(p.ValueList); err != nil {
		return errors.Annotate(err, "encode ValueList")
	}

	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}

	return nil
}
