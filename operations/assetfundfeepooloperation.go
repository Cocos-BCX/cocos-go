package operations

//go:generate ffjson   $GOFILE

import (
	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	op := &AssetFundFeePoolOperation{}
	types.OperationMap[op.Type()] = op
}

type AssetFundFeePoolOperation struct {
	Amount      types.UInt64      `json:"amount"`
	AssetID     types.GrapheneID  `json:"asset_id"`
	Extensions  types.Extensions  `json:"extensions"`
	Fee         types.AssetAmount `json:"fee"`
	FromAccount types.GrapheneID  `json:"from_account"`
}

func (p *AssetFundFeePoolOperation) ApplyFee(fee types.AssetAmount) {
	p.Fee = fee
}

func (p AssetFundFeePoolOperation) Type() types.OperationType {
	return types.OperationTypeAssetFundFeePool
}

func (p AssetFundFeePoolOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode operation id")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.FromAccount); err != nil {
		return errors.Annotate(err, "encode new options")
	}

	if err := enc.Encode(p.AssetID); err != nil {
		return errors.Annotate(err, "encode asset id")
	}

	if err := enc.Encode(p.Amount); err != nil {
		return errors.Annotate(err, "encode amount")
	}

	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}

	return nil
}

//NewAssetFundFeePoolOperation creates a new AssetFundFeePoolOperation
func NewAssetFundFeePoolOperation() *AssetFundFeePoolOperation {
	tx := AssetFundFeePoolOperation{
		Extensions: types.Extensions{},
	}
	return &tx
}
