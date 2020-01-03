package operations

//go:generate ffjson  $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationTypeFillOrder] = func() types.Operation {
		op := &FillOrderOperation{}
		return op
	}
}

//virtual order
type FillOrderOperation struct {
	types.OperationFee
	OrderID   types.ObjectID    `json:"order_id"`
	AccountID types.AccountID   `json:"account_id"`
	Pays      types.AssetAmount `json:"pays"`
	Receives  types.AssetAmount `json:"receives"`
	FillPrice types.Price       `json:"fill_price"`
	IsMaker   bool              `json:"is_maker"`
}

func (p FillOrderOperation) Type() types.OperationType {
	return types.OperationTypeFillOrder
}

func (p FillOrderOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	return nil
}

func (p FillOrderOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode Fee")
	}

	if err := enc.Encode(p.OrderID); err != nil {
		return errors.Annotate(err, "encode OrderID")
	}

	if err := enc.Encode(p.AccountID); err != nil {
		return errors.Annotate(err, "encode AccountID")
	}

	if err := enc.Encode(p.Pays); err != nil {
		return errors.Annotate(err, "encode Pays")
	}

	if err := enc.Encode(p.Receives); err != nil {
		return errors.Annotate(err, "encode Receives")
	}

	if err := enc.Encode(p.FillPrice); err != nil {
		return errors.Annotate(err, "encode fillprice")
	}

	if err := enc.Encode(p.IsMaker); err != nil {
		return errors.Annotate(err, "encode ismaker")
	}

	return nil
}
