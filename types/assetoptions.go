package types

import (
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

//go:generate ffjson $GOFILE

type AssetOptions struct {
	MaxSupply         Int64      `json:"max_supply"`
	MarketFeePercent  UInt16     `json:"market_fee_percent"`
	MaxMarketFee      Int64      `json:"max_market_fee"`
	IssuerPermissions UInt16     `json:"issuer_permissions"`
	Flags             UInt16     `json:"flags"`
	CoreExchangeRate  *Price     `json:"core_exchange_rate,omitempty"`
	Description       String     `json:"description"`
	Extensions        Extensions `json:"extensions"`
}

func (p AssetOptions) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(p.MaxSupply); err != nil {
		return errors.Annotate(err, "encode MaxSupply")
	}

	if err := enc.Encode(p.MarketFeePercent); err != nil {
		return errors.Annotate(err, "encode MarketFeePercent")
	}

	if err := enc.Encode(p.MaxMarketFee); err != nil {
		return errors.Annotate(err, "encode MaxMarketFee")
	}

	if err := enc.Encode(p.IssuerPermissions); err != nil {
		return errors.Annotate(err, "encode IssuerPermissions")
	}

	if err := enc.Encode(p.Flags); err != nil {
		return errors.Annotate(err, "encode Flags")
	}

	if err := enc.Encode(p.CoreExchangeRate != nil); err != nil {
		return errors.Annotate(err, "encode have CoreExchangeRate")
	}

	if err := enc.Encode(p.CoreExchangeRate); err != nil {
		return errors.Annotate(err, "encode CoreExchangeRate")
	}

	if err := enc.Encode(p.Description); err != nil {
		return errors.Annotate(err, "encode Description")
	}

	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode extensions")
	}

	return nil
}
