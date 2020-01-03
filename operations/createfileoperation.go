package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeCreateFile] = func() types.Operation {
		op := &CreateFileOperation{}
		return op
	}
}

type CreateFileOperation struct {
	types.OperationFee
	FileOwner   types.AccountID `json:"file_owner"`
	FileName    string          `json:"file_name"`
	FileContent string          `json:"file_content"`
}

func (p CreateFileOperation) Type() types.OperationType {
	return types.OperationTypeCreateFile
}

func (p CreateFileOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
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

func (p CreateFileOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.FileOwner); err != nil {
		return errors.Annotate(err, "encode FileOwner")
	}

	if err := enc.Encode(p.FileName); err != nil {
		return errors.Annotate(err, "encode FileName")
	}

	if err := enc.Encode(p.FileContent); err != nil {
		return errors.Annotate(err, "encode FileContent")
	}

	return nil
}
