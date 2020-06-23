package operations

//go:generate ffjson $GOFILE

import (
	"github.com/Cocos-BCX/cocos-go/types"
	"github.com/Cocos-BCX/cocos-go/util"
	"github.com/juju/errors"
)

func init() {
	types.OperationMap[types.OperationsTypeUpdateGlobalPropertyExtensions] = func() types.Operation {
		op := &UpdateGlobalPropertyExtensionsOperation{}
		return op
	}
}

type UpdateGlobalPropertyExtensionsOperation struct {
	types.OperationFee
	WitnessMaxVotes         types.UInt16 `json:"witness_max_votes"`
	CommitteeMaxVotes       types.UInt16 `json:"committee_max_votes"`
	ContractPrivateDataSize types.UInt64 `json:"contract_private_data_size"`
	ContractTotalDataSize   types.UInt64 `json:"contract_total_data_size"`
	ContractMaxDataSize     types.UInt64 `json:"contract_max_data_size"`
}

func (p UpdateGlobalPropertyExtensionsOperation) Type() types.OperationType {
	return types.OperationsTypeUpdateGlobalPropertyExtensions
}

func (p UpdateGlobalPropertyExtensionsOperation) MarshalFeeScheduleParams(params types.M, enc *util.TypeEncoder) error {
	if fee, ok := params["fee"]; ok {
		if err := enc.Encode(types.UInt64(fee.(float64))); err != nil {
			return errors.Annotate(err, "encode Fee")
		}
	}

	return nil
}

func (p UpdateGlobalPropertyExtensionsOperation) Marshal(enc *util.TypeEncoder) error {
	if err := enc.Encode(int8(p.Type())); err != nil {
		return errors.Annotate(err, "encode OperationType")
	}

	if err := enc.Encode(p.Fee); err != nil {
		return errors.Annotate(err, "encode fee")
	}

	if err := enc.Encode(p.WitnessMaxVotes); err != nil {
		return errors.Annotate(err, "encode WitnessMaxVotes")
	}

	if err := enc.Encode(p.ContractPrivateDataSize); err != nil {
		return errors.Annotate(err, "encode ContractPrivateDataSize")
	}

	if err := enc.Encode(p.ContractTotalDataSize); err != nil {
		return errors.Annotate(err, "encode ContractTotalDataSize")
	}

	if err := enc.Encode(p.ContractMaxDataSize); err != nil {
		return errors.Annotate(err, "encode ContractMaxDataSize")
	}

	return nil
}
