package types

import (
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

//go:generate ffjson $GOFILE

type PriceFeed struct {
	SettlementPrice            Price  `json:"settlement_price"`
	MaintenanceCollateralRatio UInt16 `json:"maintenance_collateral_ratio"`
	MaximumShortSqueezeRatio   UInt16 `json:"maximum_short_squeeze_ratio"`
}

func (p PriceFeed) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(p.SettlementPrice); err != nil {
		return errors.Annotate(err, "encode SettlementPrice")
	}

	if err := enc.Encode(p.MaintenanceCollateralRatio); err != nil {
		return errors.Annotate(err, "encode MaintenanceCollateralRatio")
	}

	if err := enc.Encode(p.MaximumShortSqueezeRatio); err != nil {
		return errors.Annotate(err, "encode MaximumShortSqueezeRatio")
	}
	return nil
}
