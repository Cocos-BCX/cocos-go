package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeAddFileRelateAccount] = func() types.Operation {
		op := &AddFileRelateAccountOperation{}
		return op
	}
}

type AddFileRelateAccountOperation struct {
	types.OperationFee
	FileOwner      types.AccountID      `json:"file_owner"`
	FileID         types.FileID         `json:"file_id"`
	RelatedAccount types.AccountIDArray `json:"related_account"`
}

func (p AddFileRelateAccountOperation) Type() types.OperationType {
	return types.OperationTypeAddFileRelateAccount
}

func (p AddFileRelateAccountOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
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

func (p AddFileRelateAccountOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode Type")
	}
	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode Fee")
	}
	if err := enc.Encode(p.FileOwner); err != nil {
		return errors.Annotate(err, "encode FileOwner")
	}

	if err := enc.Encode(p.FileID); err != nil {
		return errors.Annotate(err, "encode FileID")
	}

	if err := enc.Encode(p.RelatedAccount); err != nil {
		return errors.Annotate(err, "encode RelatedAccount")
	}

	return nil
}
