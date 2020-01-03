package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeRelateWorldView] = func() types.Operation {
		op := &RelateWorldViewOperation{}
		return op
	}
}

type RelateWorldViewOperation struct {
	types.OperationFee
	RelatedAccount types.AccountID `json:"related_account"`
	WorldView      string          `json:"world_view"`
	ViewOwner      types.AccountID `json:"view_owner"`
}

func (p RelateWorldViewOperation) Type() types.OperationType {
	return types.OperationTypeRelateWorldView
}

func (p RelateWorldViewOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if fee, ok := params["fee"]; ok {
		if err := enc.Encode(types.UInt64(fee.(float64))); err != nil {
			return errors.Annotate(err, "encode Fee")
		}
	}

	return nil
}

func (p RelateWorldViewOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.RelatedAccount); err != nil {
		return errors.Annotate(err, "encode RelatedAccount")
	}

	if err := enc.Encode(p.WorldView); err != nil {
		return errors.Annotate(err, "encode WorldView")
	}

	if err := enc.Encode(p.ViewOwner); err != nil {
		return errors.Annotate(err, "encode ViewOwner")
	}

	return nil
}
