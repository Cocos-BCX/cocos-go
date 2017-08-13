package objects

import (
	"strconv"
	"strings"

	"github.com/denkhaus/bitshares/util"
	"github.com/juju/errors"
)

const (
	InstanceUndefined int64 = -1
)

type GrapheneID struct {
	id         ObjectID
	spaceType  SpaceType
	objectType ObjectType
	instance   int64
}

//Id returns the objects ObjectID
func (p GrapheneID) Id() ObjectID {
	return p.id
}

//Type returns the objects ObjectType
func (p GrapheneID) Type() ObjectType {
	if !p.valid() {
		if err := p.FromString(string(p.id)); err != nil {
			panic(err.Error())
		}
	}

	switch p.spaceType {
	case SpaceTypeProtocol:
		switch p.objectType {
		case 1:
			return ObjectTypeBase
		case 2:
			return ObjectTypeAccount
		case 3:
			return ObjectTypeAsset
		case 4:
			return ObjectTypeFORCE_SETTLEMENT_OBJECT
		case 5:
			return ObjectTypeCOMMITTEE_MEMBER_OBJECT
		case 6:
			return ObjectTypeWITNESS_OBJECT
		case 7:
			return ObjectTypeLimitOrder
		case 8:
			return ObjectTypeCallOrder
		case 9:
			return ObjectTypeCUSTOM_OBJECT
		case 10:
			return ObjectTypePROPOSAL_OBJECT
		case 11:
			return ObjectTypeOPERATION_HISTORY_OBJECT
		case 12:
			return ObjectTypeWITHDRAW_PERMISSION_OBJECT
		case 13:
			return ObjectTypeVESTING_BALANCE_OBJECT
		case 14:
			return ObjectTypeWORKER_OBJECT
		case 15:
			return ObjectTypeBALANCE_OBJECT
		}

	case SpaceTypeImplementation:
		switch p.objectType {
		case 0:
			return ObjectTypeGLOBAL_PROPERTY_OBJECT
		case 1:
			return ObjectTypeDYNAMIC_GLOBAL_PROPERTY_OBJECT
		case 3:
			return ObjectTypeASSET_DYNAMIC_DATA
		case 4:
			return ObjectTypeAssetBitAssetData
		case 5:
			return ObjectTypeACCOUNT_BALANCE_OBJECT
		case 6:
			return ObjectTypeACCOUNT_STATISTICS_OBJECT
		case 7:
			return ObjectTypeTRANSACTION_OBJECT
		case 8:
			return ObjectTypeBLOCK_SUMMARY_OBJECT
		case 9:
			return ObjectTypeACCOUNT_TRANSACTION_HISTORY_OBJECT
		case 10:
			return ObjectTypeBLINDED_BALANCE_OBJECT
		case 11:
			return ObjectTypeCHAIN_PROPERTY_OBJECT
		case 12:
			return ObjectTypeWITNESS_SCHEDULE_OBJECT
		case 13:
			return ObjectTypeBUDGET_RECORD_OBJECT
		case 14:
			return ObjectTypeSPECIAL_AUTHORITY_OBJECT
		}
	}

	return ObjectTypeUndefined
}

//NewGrapheneID creates an new GrapheneID object
func NewGrapheneID(id ObjectID) *GrapheneID {
	gid := &GrapheneID{
		spaceType:  SpaceTypeUndefined,
		objectType: ObjectTypeUndefined,
		instance:   InstanceUndefined,
	}

	if err := gid.FromString(string(id)); err != nil {
		panic(err.Error())
	}

	return gid
}

func (p GrapheneID) valid() bool {
	return p.spaceType != SpaceTypeUndefined &&
		p.objectType != ObjectTypeUndefined &&
		p.instance != InstanceUndefined
}

func (p *GrapheneID) FromString(in string) error {
	parts := strings.Split(in, ".")

	if len(parts) == 3 {
		p.id = ObjectID(in)
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

		inst, err := strconv.ParseInt(parts[2], 10, 64)
		if err != nil {
			return errors.Errorf("unable to parse GrapheneID [instance] from %s", in)
		}

		p.instance = inst
		return nil
	}

	return errors.Errorf("unable to parse GrapheneID from %s", in)
}

func (p *GrapheneID) UnmarshalJSON(s []byte) error {
	str := string(s)

	if len(str) > 0 && str != "null" {
		q, err := util.SafeUnquote(str)
		if err != nil {
			return err
		}

		if err := p.FromString(q); err != nil {
			return err
		}

		return nil
	}

	return errors.Errorf("unable to unmarshal GrapheneID from %s", str)
}
