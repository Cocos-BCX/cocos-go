// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package types

import (
	"fmt"

	"github.com/Cocos-BCX/cocos-go/logging"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

type NHAssetOrderID struct {
	ObjectID
}

func (p NHAssetOrderID) Marshal(enc *util.TypeEncoder) error {
	n, err := enc.EncodeUVarintByByte(uint64(p.Instance()))
	if err != nil {
		return errors.Annotate(err, "encode instance")
	}

	for i := 0; i < 8-n; i++ {
		if err := enc.EncodeUVarint(uint64(0)); err != nil {
			return errors.Annotate(err, "encode zero")
		}
	}

	return nil
}

func (p *NHAssetOrderID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeNHAssetOrder) << 48) | instance)
	return nil
}

type NHAssetOrderIDs []NHAssetOrderID

func (p NHAssetOrderIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode NHAssetOrderID")
		}
	}

	return nil
}

func NHAssetOrderIDFromObject(ob GrapheneObject) NHAssetOrderID {
	id, ok := ob.(*NHAssetOrderID)
	if ok {
		return *id
	}

	p := NHAssetOrderID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeNHAssetOrder {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeNHAssetOrder'", p.ID()))
	}

	return p
}

//NewNHAssetOrderID creates an new NHAssetOrderID object
func NewNHAssetOrderID(id string) GrapheneObject {
	gid := new(NHAssetOrderID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"NHAssetOrderID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeNHAssetOrder {
		logging.Errorf(
			"NHAssetOrderID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeNHAssetOrder'", id),
		)
		return nil
	}

	return gid
}
