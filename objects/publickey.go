package objects

import "github.com/denkhaus/bitshares/util"

type PublicKey struct {
	key string
}

func NewPublicKey(key string) *PublicKey {
	k := PublicKey{
		key: key,
	}

	return &k
}

func (p *PublicKey) UnmarshalJSON(data []byte) error {
	if data == nil || len(data) == 0 {
		return nil
	}

	str := string(data)
	if len(str) > 0 && str != "null" {
		key, err := util.SafeUnquote(str)
		if err != nil {
			return err
		}
		p.key = key
	}

	return nil
}
