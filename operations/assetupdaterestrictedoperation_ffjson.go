// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: assetupdaterestrictedoperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Cocos-BCX/cocos-go/types"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *AssetUpdateRestrictedOperation) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *AssetUpdateRestrictedOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "payer":`)

	{

		obj, err = j.Payer.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"target_asset":`)

	{

		obj, err = j.TargetAsset.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	if j.IsAdd {
		buf.WriteString(`,"isadd":true`)
	} else {
		buf.WriteString(`,"isadd":false`)
	}
	buf.WriteString(`,"restricted_type":`)
	fflib.FormatBits2(buf, uint64(j.RestrictedType), 10, false)
	buf.WriteString(`,"restricted_List":`)
	if j.RestrictedList != nil {
		buf.WriteString(`[`)
		for i, v := range j.RestrictedList {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				obj, err = v.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"extensions":`)

	{

		obj, err = j.Extensions.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteByte(',')
	if j.Fee != nil {
		if true {
			/* Struct fall back. type=types.AssetAmount kind=struct */
			buf.WriteString(`"fee":`)
			err = buf.Encode(j.Fee)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtAssetUpdateRestrictedOperationbase = iota
	ffjtAssetUpdateRestrictedOperationnosuchkey

	ffjtAssetUpdateRestrictedOperationPayer

	ffjtAssetUpdateRestrictedOperationTargetAsset

	ffjtAssetUpdateRestrictedOperationIsAdd

	ffjtAssetUpdateRestrictedOperationRestrictedType

	ffjtAssetUpdateRestrictedOperationRestrictedList

	ffjtAssetUpdateRestrictedOperationExtensions

	ffjtAssetUpdateRestrictedOperationFee
)

var ffjKeyAssetUpdateRestrictedOperationPayer = []byte("payer")

var ffjKeyAssetUpdateRestrictedOperationTargetAsset = []byte("target_asset")

var ffjKeyAssetUpdateRestrictedOperationIsAdd = []byte("isadd")

var ffjKeyAssetUpdateRestrictedOperationRestrictedType = []byte("restricted_type")

var ffjKeyAssetUpdateRestrictedOperationRestrictedList = []byte("restricted_List")

var ffjKeyAssetUpdateRestrictedOperationExtensions = []byte("extensions")

var ffjKeyAssetUpdateRestrictedOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *AssetUpdateRestrictedOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *AssetUpdateRestrictedOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtAssetUpdateRestrictedOperationbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtAssetUpdateRestrictedOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'e':

					if bytes.Equal(ffjKeyAssetUpdateRestrictedOperationExtensions, kn) {
						currentKey = ffjtAssetUpdateRestrictedOperationExtensions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyAssetUpdateRestrictedOperationFee, kn) {
						currentKey = ffjtAssetUpdateRestrictedOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyAssetUpdateRestrictedOperationIsAdd, kn) {
						currentKey = ffjtAssetUpdateRestrictedOperationIsAdd
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffjKeyAssetUpdateRestrictedOperationPayer, kn) {
						currentKey = ffjtAssetUpdateRestrictedOperationPayer
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffjKeyAssetUpdateRestrictedOperationRestrictedType, kn) {
						currentKey = ffjtAssetUpdateRestrictedOperationRestrictedType
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyAssetUpdateRestrictedOperationRestrictedList, kn) {
						currentKey = ffjtAssetUpdateRestrictedOperationRestrictedList
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffjKeyAssetUpdateRestrictedOperationTargetAsset, kn) {
						currentKey = ffjtAssetUpdateRestrictedOperationTargetAsset
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyAssetUpdateRestrictedOperationFee, kn) {
					currentKey = ffjtAssetUpdateRestrictedOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetUpdateRestrictedOperationExtensions, kn) {
					currentKey = ffjtAssetUpdateRestrictedOperationExtensions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetUpdateRestrictedOperationRestrictedList, kn) {
					currentKey = ffjtAssetUpdateRestrictedOperationRestrictedList
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetUpdateRestrictedOperationRestrictedType, kn) {
					currentKey = ffjtAssetUpdateRestrictedOperationRestrictedType
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetUpdateRestrictedOperationIsAdd, kn) {
					currentKey = ffjtAssetUpdateRestrictedOperationIsAdd
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyAssetUpdateRestrictedOperationTargetAsset, kn) {
					currentKey = ffjtAssetUpdateRestrictedOperationTargetAsset
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyAssetUpdateRestrictedOperationPayer, kn) {
					currentKey = ffjtAssetUpdateRestrictedOperationPayer
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtAssetUpdateRestrictedOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtAssetUpdateRestrictedOperationPayer:
					goto handle_Payer

				case ffjtAssetUpdateRestrictedOperationTargetAsset:
					goto handle_TargetAsset

				case ffjtAssetUpdateRestrictedOperationIsAdd:
					goto handle_IsAdd

				case ffjtAssetUpdateRestrictedOperationRestrictedType:
					goto handle_RestrictedType

				case ffjtAssetUpdateRestrictedOperationRestrictedList:
					goto handle_RestrictedList

				case ffjtAssetUpdateRestrictedOperationExtensions:
					goto handle_Extensions

				case ffjtAssetUpdateRestrictedOperationFee:
					goto handle_Fee

				case ffjtAssetUpdateRestrictedOperationnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Payer:

	/* handler: j.Payer type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Payer.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_TargetAsset:

	/* handler: j.TargetAsset type=types.AssetID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.TargetAsset.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_IsAdd:

	/* handler: j.IsAdd type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				j.IsAdd = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				j.IsAdd = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_RestrictedType:

	/* handler: j.RestrictedType type=types.UInt8 kind=uint8 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.RestrictedType.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_RestrictedList:

	/* handler: j.RestrictedList type=operations.ObjectIDs kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ObjectIDs", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.RestrictedList = nil
		} else {

			j.RestrictedList = []types.ObjectID{}

			wantVal := true

			for {

				var tmpJRestrictedList types.ObjectID

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJRestrictedList type=types.ObjectID kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJRestrictedList.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.RestrictedList = append(j.RestrictedList, tmpJRestrictedList)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Extensions:

	/* handler: j.Extensions type=types.Extensions kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Extensions.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Fee:

	/* handler: j.Fee type=types.AssetAmount kind=struct quoted=false*/

	{
		/* Falling back. type=types.AssetAmount kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Fee)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
