package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeWitnessCreate] = func() types.Operation {
		op := &WitnessCreateOperation{}
		return op
	}
}

type WitnessCreateOperation struct {
	types.OperationFee
	WitnessAccount  types.AccountID `json:"witness_account"`
	URL             string          `json:"url"`
	BlockSigningKey types.PublicKey `json:"block_signing_key"`
}

func (p WitnessCreateOperation) Type() types.OperationType {
	return types.OperationTypeWitnessCreate
}

//TODO: verify order
func (p WitnessCreateOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.WitnessAccount); err != nil {
		return errors.Annotate(err, "encode WitnessAccount")
	}

	if err := enc.Encode(p.URL); err != nil {
		return errors.Annotate(err, "encode URL")
	}

	if err := enc.Encode(p.BlockSigningKey); err != nil {
		return errors.Annotate(err, "encode BlockSigningKey")
	}

	return nil
}
