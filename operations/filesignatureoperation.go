package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeFileSignature] = func() types.Operation {
		op := &FileSignatureOperation{}
		return op
	}
}

type FileSignatureOperation struct {
	types.OperationFee
	SignatureAccount types.AccountID `json:"signature_account"`
	FileID           types.FileID    `json:"file_id"`
	Signature        string          `json:"signature"`
}

func (p FileSignatureOperation) Type() types.OperationType {
	return types.OperationTypeFileSignature
}

func (p FileSignatureOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
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

func (p FileSignatureOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.SignatureAccount); err != nil {
		return errors.Annotate(err, "encode SignatureAccount")
	}

	if err := enc.Encode(p.FileID); err != nil {
		return errors.Annotate(err, "encode FileID")
	}

	if err := enc.Encode(p.Signature); err != nil {
		return errors.Annotate(err, "encode Signature")
	}

	return nil
}
