package types

//go:generate ffjson $GOFILE

import (
	"math"

	"github.com/gkany/graphSDK/util"
	"github.com/juju/errors"
)

type AssetAmounts []AssetAmount

type AssetAmount struct {
	Amount Int64   `json:"amount"`
	Asset  AssetID `json:"asset_id"`
}

func (p AssetAmount) Valid() bool {
	return p.Asset.Valid() && p.Amount != 0
}

func (p AssetAmount) Rate(prec float64) float64 {
	return float64(p.Amount) / math.Pow(10, prec)
}

func (p AssetAmount) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(p.Amount); err != nil {
		return errors.Annotate(err, "encode Amount")
	}

	if err := enc.Encode(p.Asset); err != nil {
		return errors.Annotate(err, "encode Asset")
	}

	return nil
}

func NewZeroCoreAsset() AssetAmount {
	return AssetAmount{
		Amount: Int64(0),
		Asset:  AssetIDFromObject(NewAssetID("1.3.0")),
	}
}

func NewAsset(amount int64, assetID string) AssetAmount {
	return AssetAmount{
		Amount: Int64(amount),
		Asset:  AssetIDFromObject(NewAssetID(assetID)),
	}
}
