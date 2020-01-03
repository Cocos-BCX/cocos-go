package types

import (
	"fmt"

	"github.com/pquerna/ffjson/ffjson"

	"strconv"
	"strings"

	sort "github.com/emirpasic/gods/utils"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

type Votes []VoteID

//TODO: define this
func (p Votes) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	//TODO: remove duplicates
	//copy votes and sort
	votes := make([]interface{}, len(p))
	for idx, id := range p {
		votes[idx] = id
	}

	sort.Sort(votes, voteIDComparator)
	for _, v := range votes {
		if err := enc.Encode(v); err != nil {
			return errors.Annotate(err, "encode VoteID")
		}
	}

	return nil
}

type VoteID struct {
	Typ      int
	Instance int
}

func (p *VoteID) UnmarshalJSON(data []byte) error {
	var str string

	if err := ffjson.Unmarshal(data, &str); err != nil {
		return errors.Annotate(err, "Unmarshal")
	}

	tk := strings.Split(str, ":")
	if len(tk) != 2 {
		return errors.Errorf("unable to unmarshal Vote from %s", str)
	}

	t, err := strconv.Atoi(tk[0])
	if err != nil {
		return errors.Annotate(err, "Atoi VoteID [type]")
	}
	p.Typ = t

	in, err := strconv.Atoi(tk[1])
	if err != nil {
		return errors.Annotate(err, "Atoi VoteID [instance]")
	}
	p.Instance = in

	return nil
}

func (p VoteID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%d:%d"`, p.Typ, p.Instance)), nil
}

//TODO: define this
func (p VoteID) Marshal(enc *util.TypeEncoder) error {
	bin := (p.Typ & 0xff) | (p.Instance << 8)
	if err := enc.Encode(uint32(bin)); err != nil {
		return errors.Annotate(err, "encode ID")
	}

	return nil
}

func NewVoteID(id string) *VoteID {
	v := VoteID{}
	if err := v.UnmarshalJSON([]byte(id)); err != nil {
		panic(errors.Annotatef(err, "unmarshal VoteID from %v", id))
	}

	return &v
}

func voteIDComparator(a, b interface{}) int {
	aID := a.(VoteID)
	bID := b.(VoteID)

	switch {
	case aID.Instance > bID.Instance:
		return 1
	case aID.Instance < bID.Instance:
		return -1
	default:
		return 0
	}
}

const (
	VoteTypeCommittee = 0
	VoteTypeWitness   = 1
	VoteTypeVoteNoone = 2
)
