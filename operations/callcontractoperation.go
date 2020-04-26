package operations

//go:generate ffjson $GOFILE

import (
	"fmt"

	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeCallContractFunction] = func() types.Operation {
		op := &CallContractFunction{}
		return op
	}
}

// CallContractFunction ...
type CallContractFunction struct {
	types.OperationFee
	Caller       types.AccountID      `json:"caller"`
	ContractID   types.ContractID     `json:"contract_id"`
	FunctionName string               `json:"function_name"`
	ValueList    []types.LuaTypeParam `json:"value_list"` // [[index, {"v":value}], ....]
	Extensions   types.Extensions     `json:"extensions"`
}

// Type ...
func (p CallContractFunction) Type() types.OperationType {
	return types.OperationTypeCallContractFunction
}

// MarshalFeeScheduleParams ...
func (p CallContractFunction) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
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

// Marshal ...
func (p CallContractFunction) Marshal(enc *util.TypeEncoder) error {
	fmt.Printf("--> type: %v\n", p.Type())
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	fmt.Printf("--> Caller: %v\n", p.Caller)
	if err := enc.Encode(p.Caller); err != nil {
		return errors.Annotate(err, "encode Caller")
	}

	fmt.Printf("--> ContractID: %v\n", p.ContractID)
	if err := enc.Encode(p.ContractID); err != nil {
		return errors.Annotate(err, "encode ContractID")
	}

	fmt.Printf("--> FunctionName: %v\n", p.FunctionName)
	if err := enc.Encode(p.FunctionName); err != nil {
		return errors.Annotate(err, "encode FunctionName")
	}

	fmt.Printf("--> ValueList size: %v\n", uint64(len(p.ValueList)))
	if err := enc.EncodeUVarint(uint64(len(p.ValueList))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	fmt.Printf("--> ValueList: %v\n", p.ValueList)
	for _, item := range p.ValueList {
		if err := enc.Encode(item); err != nil {
			return errors.Annotate(err, "encode LuaType item")
		}
	}

	fmt.Printf("--> Extensions: %v\n", p.Extensions)
	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}

	return nil
}
