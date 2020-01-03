package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeContractCreate] = func() types.Operation {
		op := &ContractCreateOperation{}
		return op
	}
}

type ContractCreateOperation struct {
	types.OperationFee
	Owner             types.AccountID  `json:"owner"`
	Name              string           `json:"name"`
	Data              string           `json:"data"`
	ContractAuthority types.PublicKey  `json:"contract_authority"`
	Extensions        types.Extensions `json:"extensions"`
}

func (p ContractCreateOperation) Type() types.OperationType {
	return types.OperationTypeContractCreate
}

func (p ContractCreateOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
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

func (p ContractCreateOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.Owner); err != nil {
		return errors.Annotate(err, "encode Owner")
	}

	if err := enc.Encode(p.Name); err != nil {
		return errors.Annotate(err, "encode Name")
	}

	if err := enc.Encode(p.Data); err != nil {
		return errors.Annotate(err, "encode Data")
	}

	if err := enc.Encode(p.ContractAuthority); err != nil {
		return errors.Annotate(err, "encode ContractAuthority")
	}

	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}

	return nil
}
