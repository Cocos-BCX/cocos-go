package types

import (
	"strconv"
	"strings"

	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
	"github.com/pquerna/ffjson/ffjson"
)

type GrapheneObject interface {
	util.TypeMarshaller
	Id() string
	Type() ObjectType
	Equals(id GrapheneObject) bool
	Valid() bool
}

type GrapheneObjects []GrapheneObject

func (p GrapheneObjects) ToStrings() []string {
	ids := make([]string, len(p))
	for idx, o := range p {
		ids[idx] = o.Id()
	}

	return ids
}

type GrapheneIDs []GrapheneID

func (p GrapheneIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode GrapheneID")
		}
	}

	return nil
}

type GrapheneID struct {
	id         string
	spaceType  SpaceType
	objectType ObjectType
	instance   UInt64
}

func (p GrapheneID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.instance)); err != nil {
		return errors.Annotate(err, "encode instance")
	}

	return nil
}

func (p GrapheneID) MarshalJSON() ([]byte, error) {
	return ffjson.Marshal(p.id)
}

func (p *GrapheneID) UnmarshalJSON(s []byte) error {
	str := string(s)

	if len(str) > 0 && str != "null" {
		q, err := util.SafeUnquote(str)
		if err != nil {
			return errors.Annotate(err, "SafeUnquote")
		}

		if err := p.FromString(q); err != nil {
			return errors.Annotate(err, "FromString")
		}

		return nil
	}

	return errors.Errorf("unable to unmarshal GrapheneID from %s", str)
}

func (p GrapheneID) Equals(o GrapheneObject) bool {
	return p.id == o.Id()
}

func (p GrapheneID) EqualsID(o string) bool {
	return p.id == o
}

//Id returns the objects ID
func (p GrapheneID) Id() string {
	return p.id
}

//Type returns the objects ObjectType
func (p GrapheneID) Type() ObjectType {
	if !p.Valid() {
		if err := p.FromString(p.id); err != nil {
			panic(errors.Annotate(err, "from string").Error())
		}
	}
	return p.objectType
}

//Space returns the objects SpaceType
func (p GrapheneID) Space() SpaceType {
	if !p.Valid() {
		if err := p.FromString(p.id); err != nil {
			panic(errors.Annotate(err, "from string").Error())
		}
	}
	return p.spaceType
}

//NewGrapheneID creates an new GrapheneID object
func NewGrapheneID(id string) *GrapheneID {
	gid := &GrapheneID{
		spaceType:  SpaceTypeUndefined,
		objectType: ObjectTypeUndefined,
	}

	if err := gid.FromString(id); err != nil {
		panic(err.Error())
	}

	return gid
}

func (p GrapheneID) String() string {
	return p.Id()
}

func (p GrapheneID) Valid() bool {
	return p.id != "" &&
		p.spaceType != SpaceTypeUndefined &&
		p.objectType != ObjectTypeUndefined
}

func (p *GrapheneID) FromRawData(in interface{}) error {
	o, ok := in.(map[string]interface{})
	if !ok {
		return errors.New("input is not map[string]interface{}")
	}

	if id, ok := o["id"]; ok {
		return p.FromString(id.(string))
	}

	return errors.New("input is no graphene object")
}

func (p *GrapheneID) FromString(in string) error {
	parts := strings.Split(in, ".")

	if len(parts) == 3 {
		p.id = in
		space, err := strconv.Atoi(parts[0])
		if err != nil {
			return errors.Errorf("unable to parse GrapheneID [space] from %s", in)
		}

		p.spaceType = SpaceType(space)

		typ, err := strconv.Atoi(parts[1])
		if err != nil {
			return errors.Errorf("unable to parse GrapheneID [type] from %s", in)
		}

		p.objectType = ObjectType(typ)

		inst, err := strconv.ParseUint(parts[2], 10, 64)
		if err != nil {
			return errors.Errorf("unable to parse GrapheneID [instance] from %s", in)
		}

		p.instance = UInt64(inst)
		return nil
	}

	return errors.Errorf("unable to parse GrapheneID from %s", in)
}
