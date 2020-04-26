package types

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"

	"github.com/juju/errors"
)

//go:generate ffjson $GOFILE

type VariantInt interface{}

type DynamicGlobalProperties struct {
	ID                             DynamicGlobalPropertyID `json:"id"`
	HeadBlockNumber                UInt32                  `json:"head_block_number"`
	HeadBlockID                    string                  `json:"head_block_id"`
	Time                           Time                    `json:"time"`
	CurrentWitness                 WitnessID               `json:"current_witness"`
	CurrentTransactionCount        UInt32                  `json:"current_transaction_count"`
	NextMaintenanceTime            Time                    `json:"next_maintenance_time"`
	LastBudgetTime                 Time                    `json:"last_budget_time"`
	WitnessBudget                  VariantInt              `json:"witness_budget"`
	AccountsRegisteredThisInterval int                     `json:"accounts_registered_this_interval"`
	RecentlyMissedCount            int64                   `json:"recently_missed_count"`
	CurrentAslot                   int64                   `json:"current_aslot"`
	RecentSlotsFilled              string                  `json:"recent_slots_filled"`
	DynamicFlags                   int                     `json:"dynamic_flags"`
	LastIrreversibleBlockNum       UInt32                  `json:"last_irreversible_block_num"`
}

func (p DynamicGlobalProperties) RefBlockNum() UInt16 {
	return UInt16(p.HeadBlockNumber)
}

func (p DynamicGlobalProperties) RefBlockPrefix() (UInt32, error) {
	rawBlockID, err := hex.DecodeString(p.HeadBlockID)
	if err != nil {
		return 0, errors.Annotatef(err, "DecodeString HeadBlockID: %v", p.HeadBlockID)
	}

	if len(rawBlockID) < 8 {
		return 0, errors.Errorf("invalid HeadBlockID: %v", p.HeadBlockID)
	}

	rawPrefix := rawBlockID[4:8]

	var prefix uint32
	if err := binary.Read(bytes.NewReader(rawPrefix), binary.LittleEndian, &prefix); err != nil {
		return 0, errors.Annotatef(err, "failed to read block prefix: %v", rawPrefix)
	}

	return UInt32(prefix), nil
}
