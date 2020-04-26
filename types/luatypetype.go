package types

// //go:generate ffjson $GOFILE

// import (
// 	"fmt"
// 

// 	"github.com/Cocos-BCX/cocos-go/util"
// 	"github.com/juju/errors"
// )

// func init() {
// 	LuaTypeMMap[LuaTypeInt] = func() LuaType {
// 		obj := &LuaInt{}
// 		return obj
// 	}
// }

// // LuaInt ... type LuaInt uint64
// type LuaInt struct {
// 	V uint64 `json:"v"`
// }

// func (p LuaInt) Type() LuaTypeType {
// 	return LuaTypeInt
// }

// func (p LuaInt) Marshal(enc *util.TypeEncoder) error {
// 	fmt.Println("LuaInt Marshal")
// 	fmt.Println("  ->type: ", p.Type())
// 	if err := enc.Encode(int8(p.Type())); err != nil {
// 		return errors.Annotate(err, "encode OperationType")
// 	}

// 	fmt.Println("  ->Fee: ", p.V)
// 	if err := enc.Encode(p.V); err != nil {
// 		return errors.Annotate(err, "encode fee")
// 	}

// 	return nil
// }

// /////////////////////////

// func init() {
// 	LuaTypeMMap[LuaTypeNumber] = func() LuaType {
// 		obj := &LuaNumber{}
// 		return obj
// 	}
// }

// // LuaNumber ... type LuaInt uint64
// type LuaNumber struct {
// 	V float64 `json:"v"`
// }

// func (p LuaNumber) Type() LuaTypeType {
// 	return LuaTypeNumber
// }

// func (p LuaNumber) Marshal(enc *util.TypeEncoder) error {
// 	fmt.Println("LuaNumber Marshal")
// 	fmt.Println("  ->type: ", p.Type())
// 	if err := enc.Encode(int8(p.Type())); err != nil {
// 		return errors.Annotate(err, "encode OperationType")
// 	}

// 	fmt.Println("  ->Fee: ", p.V)
// 	if err := enc.Encode(p.V); err != nil {
// 		return errors.Annotate(err, "encode fee")
// 	}

// 	return nil
// }

// /////////////////////////

// func init() {
// 	LuaTypeMMap[LuaTypeBool] = func() LuaType {
// 		obj := &LuaBool{}
// 		return obj
// 	}
// }

// // LuaNumber ... type LuaInt uint64
// type LuaBool struct {
// 	V bool `json:"v"`
// }

// func (p LuaBool) Type() LuaTypeType {
// 	return LuaTypeBool
// }

// func (p LuaBool) Marshal(enc *util.TypeEncoder) error {
// 	fmt.Println("LuaBool Marshal")
// 	fmt.Println("  ->type: ", p.Type())
// 	if err := enc.Encode(int8(p.Type())); err != nil {
// 		return errors.Annotate(err, "encode OperationType")
// 	}

// 	fmt.Println("  ->Fee: ", p.V)
// 	if err := enc.Encode(p.V); err != nil {
// 		return errors.Annotate(err, "encode fee")
// 	}

// 	return nil
// }

// /////////////////////////

// func init() {
// 	LuaTypeMMap[LuaTypeString] = func() LuaType {
// 		obj := &LuaString{}
// 		return obj
// 	}
// }

// // LuaNumber ... type LuaInt uint64
// type LuaString struct {
// 	V string `json:"v"`
// }

// func (p LuaString) Type() LuaTypeType {
// 	return LuaTypeString
// }

// func (p LuaString) Marshal(enc *util.TypeEncoder) error {
// 	fmt.Println("LuaInt Marshal")
// 	fmt.Println("  ->type: ", p.Type())
// 	if err := enc.Encode(int8(p.Type())); err != nil {
// 		return errors.Annotate(err, "encode OperationType")
// 	}

// 	fmt.Println("  ->Fee: ", p.V)
// 	if err := enc.Encode(p.V); err != nil {
// 		return errors.Annotate(err, "encode fee")
// 	}

// 	return nil
// }

// /////////////////////////

// func init() {
// 	LuaTypeMMap[LuaTypeFunction] = func() LuaType {
// 		obj := &LuaFunction{}
// 		return obj
// 	}
// }

// // LuaNumber ... type LuaInt uint64
// type LuaFunction struct {
// 	IsVarArg bool     `json:"is_var_arg"`
// 	ArgList  []string `json:"arglist"`
// }

// func (p LuaFunction) Type() LuaTypeType {
// 	return LuaTypeFunction
// }

// func (p LuaFunction) Marshal(enc *util.TypeEncoder) error {
// 	fmt.Println("LuaFunction Marshal")
// 	fmt.Println("  ->type: ", p.Type())
// 	if err := enc.Encode(int8(p.Type())); err != nil {
// 		return errors.Annotate(err, "encode OperationType")
// 	}

// 	if err := enc.Encode(p.IsVarArg); err != nil {
// 		return errors.Annotate(err, "encode IsVarArg")
// 	}

// 	if err := enc.Encode(p.ArgList); err != nil {
// 		return errors.Annotate(err, "encode ArgList")
// 	}
// 	return nil
// }

// /////////////////////////

// func init() {
// 	LuaTypeMMap[LuaTypeTable] = func() LuaType {
// 		obj := &LuaTable{}
// 		return obj
// 	}
// }

// // LuaTable ...
// type LuaTable struct {
// 	Value LuaTypeMap `json:"v"`
// }

// func (p LuaTable) Type() LuaTypeType {
// 	return LuaTypeTable
// }

// func (p LuaTable) Marshal(enc *util.TypeEncoder) error {
// 	fmt.Println("LuaFunction Marshal")
// 	fmt.Println("  ->type: ", p.Type())
// 	if err := enc.Encode(int8(p.Type())); err != nil {
// 		return errors.Annotate(err, "encode OperationType")
// 	}

// 	if err := enc.Encode(p.Value); err != nil {
// 		return errors.Annotate(err, "encode IsVarArg")
// 	}

// 	return nil
// }

// type LuaTypeMap map[LuaTypeType]LuaType

// // type LuaMapItem struct {
// // }
