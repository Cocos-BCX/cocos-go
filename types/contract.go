package types

//go:generate ffjson $GOFILE

type Contract struct {
	ID                     ContractID        `json:"id"`
	CreationDate           Time              `json:"creation_date"`
	Owner                  AccountID         `json:"owner"`
	Name                   string            `json:"name"`
	UserInvokeSharePercent uint32            `json:"user_invoke_share_percent"`
	IsRelease              bool              `json:"is_release"`
	CurrentVersion         string            `json:"current_version"`
	CheckContractAuthority bool              `json:"check_contract_authority"`
	ContractAuthority      PublicKey         `json:"contract_authority"`
	ContractData           LuaMap            `json:"contract_data"`
	ContractABI            LuaMap            `json:"contract_ABI"`
	LuaCodeBID             ContractBinCodeID `json:"lua_code_b_id"`
}
