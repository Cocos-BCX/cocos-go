package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAccountAuthentication] = func() types.Operation {
		op := &AccountAuthenticationOperation{}
		return op
	}
}

type AccountAuthenticationOperation struct {
	types.OperationFee
	AccountID  types.AccountID  `json:"account_id"`
	Data       types.String     `json:"data"`
	Extensions types.Extensions `json:"extensions"`
}

func (p AccountAuthenticationOperation) Type() types.OperationType {
	return types.OperationTypeAccountAuthentication
}

func (p AccountAuthenticationOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if bfee, ok := params["basic_fee"]; ok {
		if err := enc.Encode(types.UInt64(bfee.(float64))); err != nil {
			return errors.Annotate(err, "encode BasicFee")
		}
	}
	if pfee, ok := params["premium_fee"]; ok {
		if err := enc.Encode(types.UInt64(pfee.(float64))); err != nil {
			return errors.Annotate(err, "encode PremiumFee")
		}
	}
	if ppk, ok := params["price_per_kbyte"]; ok {
		if err := enc.Encode(types.UInt32(ppk.(float64))); err != nil {
			return errors.Annotate(err, "encode PricePerKByte")
		}
	}

	return nil
}

func (p AccountAuthenticationOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode Type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode Fee")
	}
	if err := enc.Encode(p.AccountID); err != nil {
		return errors.Annotate(err, "encode AccountID")
	}

	if err := enc.Encode(p.Data); err != nil {
		return errors.Annotate(err, "encode Data")
	}

	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode Extensions")
	}

	return nil
}
