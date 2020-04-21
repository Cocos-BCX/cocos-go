package types

//go:generate ffjson $GOFILE

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gkany/cocos-go/util"
	"github.com/juju/errors"
	"github.com/pquerna/ffjson/ffjson"
)

// LuaType ... lua type index number
const (
	LuaTypeInt      = iota // 0
	LuaTypeNumber          // 1
	LuaTypeString          // 2
	LuaTypeBool            // 3
	LuaTypeTable           // 4
	LuaTypeFunction        // 5
	// LuaTypeUserData
)

// LuaKey ... lua key index number
const (
	LuaKeyInt    = iota // 0
	LuaKeyNumber        // 1
	LuaKeyString        // 2
	LuaKeyBool          // 3
)

// LuaInt ... type LuaInt uint64  eg:[0, {"v": 100}]
type LuaInt struct {
	V uint64 `json:"v"`
}

// Marshal ...
func (o LuaInt) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(LuaTypeInt); err != nil {
		return errors.Annotate(err, "encode type")
	}

	if err := enc.Encode(o.V); err != nil {
		return errors.Annotate(err, "encode data")
	}
	return nil
}

// MarshalJSON ...
// func (o LuaInt) MarshalJSON() ([]byte, error) {
// 	return ffjson.Marshal([]interface{}{
// 		LuaTypeInt,
// 		o,
// 	})
// }

// LuaNumber ... float64 eg:[1, {"v": 3.14}]
type LuaNumber struct {
	V float64 `json:"v"`
}

// Marshal ...
func (o LuaNumber) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(LuaTypeNumber); err != nil {
		return errors.Annotate(err, "encode type")
	}

	if err := enc.Encode(o.V); err != nil {
		return errors.Annotate(err, "encode data")
	}
	return nil
}

// MarshalJSON ...
// func (o LuaNumber) MarshalJSON() ([]byte, error) {
// 	return ffjson.Marshal([]interface{}{
// 		LuaTypeNumber,
// 		o,
// 	})
// }

// LuaString ... string eg:[2, {"v": "hello, Lua contract"}]
type LuaString struct {
	V string `json:"v"`
}

// Marshal ...
func (o LuaString) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(LuaTypeString); err != nil {
		return errors.Annotate(err, "encode type")
	}

	if err := enc.Encode(o.V); err != nil {
		return errors.Annotate(err, "encode data")
	}
	return nil
}

// MarshalJSON ...
// func (o LuaString) MarshalJSON() ([]byte, error) {
// 	return ffjson.Marshal([]interface{}{
// 		LuaTypeString,
// 		o,
// 	})
// }

// LuaBool ... bool eg:[3, {"v": true}]
type LuaBool struct {
	V bool `json:"v"`
}

// Marshal ...
func (o LuaBool) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(LuaTypeBool); err != nil {
		return errors.Annotate(err, "encode type")
	}

	if err := enc.Encode(o.V); err != nil {
		return errors.Annotate(err, "encode data")
	}
	return nil
}

// MarshalJSON ...
// func (o LuaBool) MarshalJSON() ([]byte, error) {
// 	return ffjson.Marshal([]interface{}{
// 		LuaTypeBool,
// 		o,
// 	})
// }

// LuaType ... [type_index, {"v":type_obj}]
type LuaType []interface{}

// LuaKey ... eg: {"key": [2, {"v": "lua string test"}]}
type LuaKey struct {
	Key LuaType `json:"key"`
}

// LuaMap ...
type LuaMap []LuaTypeItem

// LuaTable ...  eg:[4,{"v":[[{"key":[2,{"v":"name"}]},[2,{"v":"Jack"}]],[{"key":[2,{"v":"age"}]},[0,{"v":22}]],[{"key":[2,{"v":"sex"}]},[2,{"v":"man"}]],[{"key":[2,{"v":"height"}]},[0,{"v":170}]],[{"key":[2,{"v":"weight"}]},[1,{"v":65.3}]],[{"key":[2,{"v":"address"}]},[2,{"v":"CN"}]],[{"key":[2,{"v":"is_student"}]},[3,{"v":false}]]]}]
type LuaTable struct {
	Value LuaMap `json:"v"`
}

// Marshal ...
func (o LuaTable) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(LuaTypeTable); err != nil {
		return errors.Annotate(err, "encode type")
	}

	if err := enc.Encode(o.Value); err != nil {
		return errors.Annotate(err, "encode data")
	}
	return nil
}

// MarshalJSON ...
// func (o LuaTable) MarshalJSON() ([]byte, error) {
// 	return ffjson.Marshal([]interface{}{
// 		LuaTypeTable,
// 		o,
// 	})
// }

// LuaFunction ... eg: [5, {"v": {"is_var_arg": true, "arglist": ["arg1", "arg2", "arg3"]}}]
type LuaFunction struct {
	IsVarArg bool     `json:"is_var_arg"`
	ArgList  []string `json:"arglist"`
}

// Marshal ...
func (o LuaFunction) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(o.IsVarArg); err != nil {
		return errors.Annotate(err, "encode IsVarArg")
	}

	if err := enc.Encode(o.ArgList); err != nil {
		return errors.Annotate(err, "encode ArgList")
	}
	return nil
}

// MarshalJSON ...
// func (o LuaFunction) MarshalJSON() ([]byte, error) {
// 	return ffjson.Marshal([]interface{}{
// 		LuaTypeFunction,
// 		o,
// 	})
// }

// LuaTypeItem ... [LuaKey, LuaType]
type LuaTypeItem []interface{}

// Marshal ...
func (o LuaTypeItem) Marshal(enc *util.TypeEncoder) error {
	if o == nil {
		return nil
	}

	if len(o) != 2 {
		return errors.Errorf("lua map item type length error. size: %v", len(o))
	}

	if err := enc.Encode(o[0]); err != nil {
		return errors.Annotate(err, "encode item key")
	}

	if err := enc.Encode(o[1]); err != nil {
		return errors.Annotate(err, "encode item value")
	}

	return nil
}

// MarshalJSON ...
func (o LuaTypeItem) MarshalJSON() ([]byte, error) {
	return ffjson.Marshal(o)
}

// UnmarshalJSON ...
func (o *LuaTypeItem) UnmarshalJSON(data []byte) error {
	raw := make([]json.RawMessage, 2)
	if err := ffjson.Unmarshal(data, &raw); err != nil {
		return errors.Annotate(err, "unmarshal raw object")
	}

	var key LuaKey
	if err := ffjson.Unmarshal(raw[0], &key); err != nil {
		return errors.Annotate(err, "unmarshal key")
	}

	var value LuaType
	if err := ffjson.Unmarshal(raw[1], &value); err != nil {
		return errors.Annotate(err, "unmarshal value")
	}

	// keyResult, err := parseLuaKey(key)
	// if err != nil {
	// 	return err
	// }
	var keyResult []interface{}
	err := parseLuaType(key.Key, &keyResult)
	if err != nil {
		return err
	}

	var result []interface{}
	err = parseLuaType(value, &result)
	if err != nil {
		return err
	}

	itemResult := make(LuaTypeItem, 2)
	itemResult[0] = keyResult
	switch len(result) {
	case 1:
		itemResult[1] = result[0]
	default:
		itemResult[1] = result
	}
	*o = itemResult

	return nil
}

// LuaKey ... int number string bool
func parseLuaKey(keyObject LuaKey) (interface{}, error) {
	valueType := keyObject.Key[0].(float64)
	value := keyObject.Key[1].(map[string]interface{})
	valuev := value["v"] // basic type

	var result interface{}
	switch valueType {
	case LuaTypeInt, LuaTypeNumber:
		switch valuev.(type) {
		case float64:
			result = valuev.(float64)
		case string:
			strf, err := strconv.ParseFloat(valuev.(string), 64)
			if err != nil {
				fmt.Printf("string ParseFloat64 failed. src: %v\n", valuev)
				return result, err
			}
			result = strf
		}
	case LuaTypeString:
		result = valuev.(string)
	case LuaTypeBool:
		result = valuev.(bool)
	default:
		fmt.Println("UnKnow LuaKey type. ", fmt.Sprintf("%T", valueType))
		return nil, errors.Errorf("UnKnow LuaKey type, type index: %v", valueType)
	}
	return result, nil
}

// LuaType ... basic type: int number string bool, complex: LuaMap, LuaTable, LuaFunction, maybe include basic and complex type
func parseLuaType(pair []interface{}, dst *[]interface{}) error {
	var result interface{}

	valueType := pair[0].(float64)
	value := pair[1].(map[string]interface{})
	valuev := value["v"] // basic type

	switch valueType {
	case LuaTypeInt, LuaTypeNumber:
		switch valuev.(type) {
		case float64:
			result = valuev.(float64)
		case string:
			// result = valuev.(string) // 强转成float64
			strf, err := strconv.ParseFloat(valuev.(string), 64)
			if err != nil {
				fmt.Printf("string ParseFloat64 failed. src: %v\n", valuev)
				return err
			}
			result = strf
		}
	case LuaTypeString:
		result = valuev.(string)
	case LuaTypeBool:
		result = valuev.(bool)
	case LuaTypeFunction:
		isVarArg := value["is_var_arg"].(bool)
		argList := value["arglist"].([]interface{})
		args := []string{}
		for _, arg := range argList {
			args = append(args, arg.(string))
		}
		result = LuaFunction{IsVarArg: isVarArg, ArgList: args}
	case LuaTypeTable:
		vMap := pair[1].(map[string]interface{})
		for _, vMapList := range vMap {
			for _, eList := range vMapList.([]interface{}) {
				for _, element := range eList.([]interface{}) {
					switch element.(type) {
					case []interface{}:
						err := parseLuaType(element.([]interface{}), dst)
						if err != nil {
							return err
						}
					case map[string]interface{}:
						ePair := eList.([]interface{}) // [0] -- map[string]interface {} | [1] -- []interface {}
						first := ePair[0]              // map[key:[2 map[v:PARTNER_INCENTIE]]]
						second := ePair[1]

						firstMap := first.(map[string]interface{})
						value := firstMap["key"] // pair type, parse by type
						err := parseLuaType(value.([]interface{}), dst)
						if err != nil {
							fmt.Printf("parseLuaType error. data: %v, error: %v\n", value, err)
							return err
						}

						err = parseLuaType(second.([]interface{}), dst)
						if err != nil {
							fmt.Printf("parseLuaType error. data: %v, error: %v\n", second, err)
							return err
						}
					}
				}
			}
		}
	default:
		fmt.Println("UnKnow LuaType type. ", fmt.Sprintf("%T", valueType))
		return errors.Errorf("UnKnow LuaType type, type index: %v", valueType)
	}
	*dst = append(*dst, result)
	return nil
}

// type LuaTypeType

// func (o LuaTypeType) MarshalJSON() ([]byte, error) {
// 	switch o.(type) {
// 	case LuaInt:
// 		return ffjson.Marshal([]interface{}{
// 			LuaTypeInt,
// 			o,
// 		})
// 	case LuaNumber:
// 		return ffjson.Marshal([]interface{}{
// 			LuaTypeNumber,
// 			o,
// 		})
// 	case LuaString:
// 		return ffjson.Marshal([]interface{}{
// 			LuaTypeString,
// 			o,
// 		})
// 	case LuaBool:
// 		return ffjson.Marshal([]interface{}{
// 			LuaTypeBool,
// 			o,
// 		})
// 	case LuaFunction:
// 		return ffjson.Marshal([]interface{}{
// 			LuaTypeFunction,
// 			o,
// 		})
// 	case LuaTable:
// 		return ffjson.Marshal([]interface{}{
// 			LuaTypeTable,
// 			o,
// 		})
// 	}

// 	return nil, errors.Annotate(err, "Unknow lua type")
// }

// func (o LuaTypeType) Marshal(enc *util.TypeEncoder) error {
// 	if err := enc.Encode(o); err != nil {
// 		return errors.Annotate(err, "encode luatype type")
// 	}

// 	return nil
// }

// ValueList ...
// type LuaTypesVec []LuaTypeType

// func (o LuaTypesVec) Marshal(enc *util.TypeEncoder) error {
// 	if error := enc.EncodeUVarint(uint64(len(o))); error != nil {
// 		return errors.Annotate(err, "encode length")
// 	}

// 	for _, item := range o {
// 		if err := enc.Encode(item); err != nil {
// 			return errors.Annotate(err, "encode LuaType item")
// 		}
// 	}

// 	return nil
// }
