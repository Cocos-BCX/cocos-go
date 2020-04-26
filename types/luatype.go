package types

// import (
// 	"encoding/json"
// 	"fmt"
// 

// 	"github.com/Cocos-BCX/cocos-go/logging"
// 	"github.com/Cocos-BCX/cocos-go/util"
// 	"github.com/juju/errors"
// 	"github.com/pquerna/ffjson/ffjson"
// )

// type LuaTypeType int8

// // LuaType ... lua type index number
// const (
// 	LuaTypeInt      = iota // 0
// 	LuaTypeNumber          // 1
// 	LuaTypeString          // 2
// 	LuaTypeBool            // 3
// 	LuaTypeTable           // 4
// 	LuaTypeFunction        // 5
// 	// LuaTypeUserData
// )

// // LuaKey ... lua key index number
// const (
// 	LuaKeyInt = iota
// 	LuaKeyNumber
// 	LuaKeyString
// 	LuaKeyBool
// )

// type GetLuaTypeFunc func() LuaType

// var (
// 	LuaTypeMMap = make(map[LuaTypeType]GetLuaTypeFunc)
// )

// type LuaType interface {
// 	util.TypeMarshaler
// 	Type() LuaTypeType
// }

// type LuaTypeEnvelopeHolder struct {
// 	Op LuaTypeEnvelope `json:"key"`
// }

// type LuaTypeEnvelopeHolders []LuaTypeEnvelopeHolder

// func (p LuaTypeEnvelopeHolders) Marshal(enc *util.TypeEncoder) error {
// 	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
// 		return errors.Annotate(err, "encode length")
// 	}

// 	for _, op := range p {
// 		if err := enc.Encode(op.Op); err != nil {
// 			return errors.Annotate(err, "encode Op")
// 		}
// 	}

// 	return nil
// }

// type LuaTypeEnvelope struct {
// 	Type         LuaTypeType
// 	LuaTypeValue LuaType
// }

// func (p LuaTypeEnvelope) Marshal(enc *util.TypeEncoder) error {
// 	if err := enc.Encode(p.LuaTypeValue); err != nil {
// 		return errors.Annotate(err, "encode LuaTypeValue")
// 	}

// 	return nil
// }

// func (p LuaTypeEnvelope) MarshalJSON() ([]byte, error) {
// 	return ffjson.Marshal([]interface{}{
// 		p.Type,
// 		p.LuaTypeValue,
// 	})
// }

// func (p *LuaTypeEnvelope) UnmarshalJSON(data []byte) error {
// 	raw := make([]json.RawMessage, 2)
// 	if err := ffjson.Unmarshal(data, &raw); err != nil {
// 		return errors.Annotate(err, "unmarshal raw object")
// 	}

// 	if err := ffjson.Unmarshal(raw[0], &p.Type); err != nil {
// 		return errors.Annotate(err, "unmarshal LuaTypeType")
// 	}

// 	descr := fmt.Sprintf("LuaTypeType %s", p.Type)

// 	if getOp, ok := LuaTypeMMap[p.Type]; ok {

// 		p.LuaTypeValue = getOp()
// 		if err := ffjson.Unmarshal(raw[1], p.LuaTypeValue); err != nil {
// 			logging.DDumpUnmarshaled(descr, raw[1])
// 			return errors.Annotatef(err, "unmarshal LuaType %s", p.Type)
// 		}
// 	} else {
// 		fmt.Printf("LuaType type %s not yet supported\n", p.Type)
// 		logging.DDumpUnmarshaled(descr, raw[1])
// 	}

// 	return nil
// }

// type LuaTypes []LuaType

// func (p LuaTypes) Marshal(enc *util.TypeEncoder) error {
// 	fmt.Println("LuaTypes Marshal")
// 	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
// 		return errors.Annotate(err, "encode LuaTypes length")
// 	}

// 	for _, op := range p {
// 		if err := enc.Encode(op); err != nil {
// 			return errors.Annotate(err, "encode LuaType")
// 		}
// 	}

// 	return nil
// }

// func (p LuaTypes) MarshalJSON() ([]byte, error) {
// 	env := make([]LuaTypeEnvelope, len(p))
// 	for idx, op := range p {
// 		env[idx] = LuaTypeEnvelope{
// 			Type:         op.Type(),
// 			LuaTypeValue: op,
// 		}
// 	}

// 	return ffjson.Marshal(env)
// }

// func (p *LuaTypes) UnmarshalJSON(data []byte) error {
// 	var envs []LuaTypeEnvelope
// 	if err := ffjson.Unmarshal(data, &envs); err != nil {
// 		return err
// 	}

// 	ops := make(LuaTypes, len(envs))
// 	for idx, env := range envs {
// 		if env.LuaTypeValue != nil {
// 			ops[idx] = env.LuaTypeValue.(LuaType)
// 		}
// 	}

// 	*p = ops
// 	return nil
// }

// func (p LuaTypes) Envelopes() []LuaTypeEnvelope {
// 	ret := make([]LuaTypeEnvelope, len(p))
// 	for idx, op := range p {
// 		ret[idx] = LuaTypeEnvelope{
// 			Type:         op.Type(),
// 			LuaTypeValue: op,
// 		}
// 	}

// 	return ret
// }
