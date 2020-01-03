package types

//go:generate ffjson $GOFILE

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Cocos-BCX/cocos-go/config"
	"github.com/Cocos-BCX/cocos-go/logging"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/pquerna/ffjson/ffjson"

	"github.com/juju/errors"
)

const (
	TxExpirationDefault = 30 * time.Second
)

type SignedTransactions []SignedTransaction
type AgreedTaskPair []string

type SignedTransaction struct {
	Transaction
	AgreedTask AgreedTaskPair `json:"agreed_task,omitempty"`
	Signatures Signatures     `json:"signatures"`
}

func (p SignedTransaction) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(p.Transaction); err != nil {
		return errors.Annotate(err, "encode Transaction")
	}

	if err := enc.Encode(p.AgreedTask); err != nil {
		return errors.Annotate(err, "encode AgreedTask")
	}

	if err := enc.Encode(p.Signatures); err != nil {
		return errors.Annotate(err, "encode Signatures")
	}

	return nil
}

//SerializeTrx serializes the transaction wihout signatures.
func (p SignedTransaction) SerializeTrx() ([]byte, error) {
	var b bytes.Buffer
	enc := util.NewTypeEncoder(&b)
	if err := enc.Encode(p.Transaction); err != nil {
		return nil, errors.Annotate(err, "encode Transaction")
	}

	return b.Bytes(), nil
}

//ToHex returns th hex representation of the underlying transaction + signatures.
func (p SignedTransaction) ToHex() (string, error) {
	var b bytes.Buffer
	enc := util.NewTypeEncoder(&b)
	if err := enc.Encode(p); err != nil {
		return "", errors.Annotate(err, "encode SignedTransaction")
	}

	return hex.EncodeToString(b.Bytes()), nil
}

//Digest calculates ths sha256 hash of the transaction.
func (tx SignedTransaction) Digest(chain *config.ChainConfig) ([]byte, error) {
	if chain == nil {
		return nil, ErrChainConfigIsUndefined
	}

	// fmt.Println("-------------> Digest: ")
	// if txJSON, err := tx.MarshalJSON(); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("tx: %s\n", string(txJSON))
	// }

	writer := sha256.New()
	rawChainID, err := hex.DecodeString(chain.ID)
	if err != nil {
		return nil, errors.Annotatef(err, "failed to decode chain ID: %v", chain.ID)
	}

	//	digestChainID := sha256.Sum256(rawChainID)
	//	util.Dump("digest chainID", hex.EncodeToString(digestChainID[:]))

	if _, err := writer.Write(rawChainID); err != nil {
		return nil, errors.Annotate(err, "Write [chainID]")
	}

	rawTrx, err := tx.SerializeTrx()
	if err != nil {
		return nil, errors.Annotatef(err, "Serialize")
	}
	// fmt.Printf("rawTrx - Hex: %v\n", hex.EncodeToString(rawTrx[:]))

	//	digestTrx := sha256.Sum256(rawTrx)
	//	util.Dump("digest trx", hex.EncodeToString(digestTrx[:]))

	if _, err := writer.Write(rawTrx); err != nil {
		return nil, errors.Annotate(err, "Write [trx]")
	}

	digest := writer.Sum(nil)
	//	util.Dump("digest trx all", hex.EncodeToString(digest[:]))
	// fmt.Printf("digest: %v, hex: %s\n", digest[:], hex.EncodeToString(digest[:]))
	
	return digest[:], nil
}

//NewSignedTransactionWithBlockData creates a new SignedTransaction and initializes
//relevant Blockdata fields and expiration.
func NewSignedTransactionWithBlockData(props *DynamicGlobalProperties) (*SignedTransaction, error) {
	prefix, err := props.RefBlockPrefix()
	if err != nil {
		return nil, errors.Annotate(err, "RefBlockPrefix")
	}

	tx := SignedTransaction{
		Transaction: Transaction{
			Extensions:     Extensions{},
			RefBlockNum:    props.RefBlockNum(),
			Expiration:     props.Time.Add(TxExpirationDefault),
			RefBlockPrefix: prefix,
		},
		AgreedTask: AgreedTaskPair{},
		Signatures: Signatures{},
	}

	return &tx, nil
}

//NewSignedTransaction creates an new SignedTransaction
func NewSignedTransaction() *SignedTransaction {
	tm := time.Now().UTC().Add(TxExpirationDefault)
	tx := SignedTransaction{
		Transaction: Transaction{
			Extensions: Extensions{},
			Expiration: Time{tm},
		},
		Signatures: Signatures{},
	}

	return &tx
}

type ProcessedTransaction struct {
	SignedTransaction
	Operationresults OperationResults `json: "operation_results"`
}

func (p ProcessedTransaction) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(p.SignedTransaction); err != nil {
		return errors.Annotate(err, "encode Transaction")
	}

	if err := enc.Encode(p.Operationresults); err != nil {
		return errors.Annotate(err, "encode Signatures")
	}

	return nil
}

type SignedTransactionsWithTrxID []SignedTransactionWithTransactionId

type SignedTransactionWithTransactionId struct {
	TransactionId     string
	SignedTransaction ProcessedTransaction
}

func (p SignedTransactionWithTransactionId) Marshal(enc *util.TypeEncoder) error {
	// type is marshaled by operation
	if err := enc.Encode(p); err != nil {
		return errors.Annotate(err, "Encode SignedTransactionWithTransactionId")
	}

	return nil
}

func (p SignedTransactionWithTransactionId) MarshalJSON() ([]byte, error) {
	return ffjson.Marshal([]interface{}{
		p.TransactionId,
		p.SignedTransaction,
	})
}

func (p *SignedTransactionWithTransactionId) UnmarshalJSON(data []byte) error {
	raw := make([]json.RawMessage, 2)
	if err := ffjson.Unmarshal(data, &raw); err != nil {
		return errors.Annotate(err, "Unmarshal [raw]")
	}

	if len(raw) != 2 {
		return ErrInvalidInputLength
	}

	if err := ffjson.Unmarshal(raw[0], &p.TransactionId); err != nil {
		return errors.Annotate(err, "Unmarshal [TransactionId]")
	}

	if err := ffjson.Unmarshal(raw[1], &p.SignedTransaction); err != nil {
		logging.DDumpUnmarshaled(
			fmt.Sprintf("TransactionId %s", p.TransactionId),
			raw[1],
		)
		return errors.Annotatef(err, "Unmarshal SignedTransaction %v", p.SignedTransaction)
	}

	return nil
}
