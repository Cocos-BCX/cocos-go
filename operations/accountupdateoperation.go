package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAccountUpdate] = func() types.Operation {
		op := &AccountUpdateOperation{}
		return op
	}
}

type AccountUpdateOperation struct {
	types.OperationFee
	LockWithVote *types.LockWithVotePairType   `json:"lock_with_vote,omitempty"`
	Account      types.AccountID               `json:"account"`
	Owner        *types.Authority              `json:"owner,omitempty"`
	Active       *types.Authority              `json:"active,omitempty"`
	NewOptions   *types.AccountOptions         `json:"new_options,omitempty"`
	Extensions   types.AccountUpdateExtensions `json:"extensions"`
}

func (p AccountUpdateOperation) Type() types.OperationType {
	return types.OperationTypeAccountUpdate
}

func (p AccountUpdateOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
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

func (p AccountUpdateOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode Fee")
	}

	if err := enc.Encode(p.Account); err != nil {
		return errors.Annotate(err, "encode Account")
	}

	if err := enc.Encode(p.Owner != nil); err != nil {
		return errors.Annotate(err, "encode have Owner")
	}

	if err := enc.Encode(p.Owner); err != nil {
		return errors.Annotate(err, "encode Owner")
	}

	if err := enc.Encode(p.Active != nil); err != nil {
		return errors.Annotate(err, "encode have Active")
	}

	if err := enc.Encode(p.Active); err != nil {
		return errors.Annotate(err, "encode Active")
	}

	if err := enc.Encode(p.NewOptions != nil); err != nil {
		return errors.Annotate(err, "encode have NewOptions")
	}

	if err := enc.Encode(p.NewOptions); err != nil {
		return errors.Annotate(err, "encode NewOptions")
	}

	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}

	return nil
}
