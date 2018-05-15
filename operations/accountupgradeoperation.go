package operations

//go:generate ffjson   $GOFILE

import (
	"github.com/denkhaus/bitshares/types"
	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

func init() {
	op := &AccountUpgradeOperation{}
	types.OperationMap[op.Type()] = op
}

type AccountUpgradeOperation struct {
	AccountToUpgrade        types.GrapheneID  `json:"account_to_upgrade"`
	Extensions              types.Extensions  `json:"extensions"`
	Fee                     types.AssetAmount `json:"fee"`
	UpgradeToLifetimeMember bool              `json:"upgrade_to_lifetime_member"`
}

func (p *AccountUpgradeOperation) ApplyFee(fee types.AssetAmount) {
	p.Fee = fee
}

func (p AccountUpgradeOperation) Type() types.OperationType {
	return types.OperationTypeAccountUpgrade
}

//TODO: validate order
func (p AccountUpgradeOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode operation id")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode Fee")
	}

	if err := enc.Encode(p.AccountToUpgrade); err != nil {
		return errors.Annotate(err, "encode AccountToUpgrade")
	}

	if err := enc.Encode(p.UpgradeToLifetimeMember); err != nil {
		return errors.Annotate(err, "encode UpgradeToLifetimeMember")
	}

	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}

	return nil
}

//NewAccountUpgradeOperation creates a new AccountUpgradeOperation
func NewAccountUpgradeOperation() *AccountUpgradeOperation {
	tx := AccountUpgradeOperation{
		Extensions: types.Extensions{},
	}
	return &tx
}
