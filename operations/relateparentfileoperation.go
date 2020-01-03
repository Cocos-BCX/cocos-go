package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeRelateParentFile] = func() types.Operation {
		op := &RelateParentFileOperation{}
		return op
	}
}

type RelateParentFileOperation struct {
	types.OperationFee
	SubFileOwner    types.AccountID `json:"sub_file_owner"`
	ParentFile      types.FileID    `json:"parent_file"`
	ParentFileOwner types.AccountID `json:"parent_file_owner"`
	SubFile         types.FileID    `json:"sub_file"`
}

func (p RelateParentFileOperation) Type() types.OperationType {
	return types.OperationTypeRelateParentFile
}

func (p RelateParentFileOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if fee, ok := params["fee"]; ok {
		if err := enc.Encode(types.UInt64(fee.(float64))); err != nil {
			return errors.Annotate(err, "encode Fee")
		}
	}

	return nil
}

func (p RelateParentFileOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.SubFileOwner); err != nil {
		return errors.Annotate(err, "encode SubFileOwner")
	}

	if err := enc.Encode(p.ParentFile); err != nil {
		return errors.Annotate(err, "encode ParentFile")
	}

	if err := enc.Encode(p.ParentFileOwner); err != nil {
		return errors.Annotate(err, "encode ParentFileOwner")
	}

	if err := enc.Encode(p.SubFile); err != nil {
		return errors.Annotate(err, "encode SubFile")
	}

	return nil
}
