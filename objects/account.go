package objects

//go:generate ffjson -force-regenerate $GOFILE

type Account struct {
	ID                            GrapheneID     `json:"id"`
	Name                          string         `json:"name"`
	Statistics                    GrapheneID     `json:"statistics"`
	MembershipExpirationDate      RFC3339Time    `json:"membership_expiration_date"`
	NetworkFeePercentage          int64          `json:"network_fee_percentage"`
	LifetimeReferrerFeePercentage int64          `json:"lifetime_referrer_fee_percentage"`
	ReferrerRewardsPercentage     int64          `json:"referrer_rewards_percentage"`
	TopNControlFlags              int64          `json:"top_n_control_flags"`
	WhitelistingAccounts          []string       `json:"whitelisting_accounts"`
	BlacklistingAccounts          []string       `json:"blacklisting_accounts"`
	WhitelistedAccounts           []string       `json:"whitelisted_accounts"`
	BlacklistedAccounts           []string       `json:"blacklisted_accounts"`
	Options                       AccountOptions `json:"options"`
	Registrar                     GrapheneID     `json:"registrar"`
	Referrer                      GrapheneID     `json:"referrer"`
	LifetimeReferrer              GrapheneID     `json:"lifetime_referrer"`
	CashbackVB                    GrapheneID     `json:"cashback_vb"`
	Owner                         Authority      `json:"owner"`
	Active                        Authority      `json:"active"`
	OwnerSpecialAuthority         []interface{}  `json:"owner_special_authority"`
	ActiveSpecialAuthority        []interface{}  `json:"active_special_authority"`
}

//NewAccount creates a new Account object
func NewAccount(id ObjectID) *Account {
	acc := Account{
		ID: *NewGrapheneID(id),
	}

	return &acc
}
