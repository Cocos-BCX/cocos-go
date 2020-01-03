package types

import (
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

//go:generate ffjson $GOFILE
type Witness struct {
	ID                    WitnessID         `json:"id"`
	WitnessAccount        AccountID         `json:"witness_account"`
	LastAslot             UInt64            `json:"last_aslot"`
	SigningKey            PublicKey         `json:"signing_key"`
	PayVestingBalance     *VestingBalanceID `json:"pay_vb,omitempty"`
	VoteID                VoteID            `json:"vote_id"`
	TotalVotes            UInt64            `json:"total_votes"`
	URL                   String            `json:"url"`
	TotalMissed           UInt64            `json:"total_missed"`
	LastConfirmedBlockNum UInt32            `json:"last_confirmed_block_num"`
	WorkStatus            bool              `json:"work_status"`
	NextMaintenanceTime   Time              `json:"next_maintenance_time"`
	Supporters            SupporterType     `json:"supporters"`
}

func (p Witness) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(p.WitnessAccount); err != nil {
		return errors.Annotate(err, "encode WitnessAccount")
	}

	if err := enc.Encode(p.LastAslot); err != nil {
		return errors.Annotate(err, "encode LastAslot")
	}

	if err := enc.Encode(p.SigningKey); err != nil {
		return errors.Annotate(err, "encode SigningKey")
	}

	if err := enc.Encode(p.PayVestingBalance == nil); err != nil {
		return errors.Annotate(err, "encode have PayVestingBalance")
	}

	if err := enc.Encode(p.PayVestingBalance); err != nil {
		return errors.Annotate(err, "encode PayVestingBalance")
	}

	if err := enc.Encode(p.VoteID); err != nil {
		return errors.Annotate(err, "encode VoteID")
	}

	if err := enc.Encode(p.TotalVotes); err != nil {
		return errors.Annotate(err, "encode TotalVotes")
	}

	if err := enc.Encode(p.URL); err != nil {
		return errors.Annotate(err, "encode URL")
	}

	if err := enc.Encode(p.TotalMissed); err != nil {
		return errors.Annotate(err, "encode TotalMissed")
	}

	if err := enc.Encode(p.LastConfirmedBlockNum); err != nil {
		return errors.Annotate(err, "encode LastConfirmedBlockNum")
	}

	if err := enc.Encode(p.WorkStatus); err != nil {
		return errors.Annotate(err, "encode WorkStatus")
	}

	if err := enc.Encode(p.NextMaintenanceTime); err != nil {
		return errors.Annotate(err, "encode NextMaintenanceTime")
	}

	if err := enc.Encode(p.Supporters); err != nil {
		return errors.Annotate(err, "encode Supporters")
	}

	return nil
}
