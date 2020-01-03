package types

import (
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

//go:generate ffjson $GOFILE
type CommitteeMember struct {
	ID                     CommitteeMemberID `json:"id"`
	CommitteeMemberAccount AccountID         `json:"committee_member_account"`
	VoteID                 VoteID            `json:"vote_id"`
	TotalVotes             UInt64            `json:"total_votes"`
	URL                    String            `json:"url"`
	WorkStatus             bool              `json:"work_status"`
	NextMaintenanceTime    Time              `json:"next_maintenance_time"`
	Supporters             SupporterType     `json:"supporters"`
}

func (p CommitteeMember) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(p.CommitteeMemberAccount); err != nil {
		return errors.Annotate(err, "encode CommitteeMemberAccount")
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
