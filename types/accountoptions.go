package types

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

type AccountOptions struct {
	MemoKey    PublicKey  `json:"memo_key"`
	Votes      Votes      `json:"votes"`
	Extensions Extensions `json:"extensions"`
}

func (p AccountOptions) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(p.MemoKey); err != nil {
		return errors.Annotate(err, "encode MemoKey")
	}

	if err := enc.Encode(p.Votes); err != nil {
		return errors.Annotate(err, "encode Votes")
	}

	if err := enc.Encode(p.Extensions); err != nil {
		return errors.Annotate(err, "encode Extensions")
	}

	return nil
}
