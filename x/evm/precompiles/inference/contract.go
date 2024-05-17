package inferenceprecompile

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/vm"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/initia-labs/minievm/x/evm/contracts/i_inference"
	"github.com/initia-labs/minievm/x/evm/types"
)

var _ vm.ExtendedPrecompiledContract = InferencePrecompile{}
var _ vm.PrecompiledContract = InferencePrecompile{}
var _ types.WithContext = InferencePrecompile{}

type InferencePrecompile struct {
	*abi.ABI
	ctx context.Context
}

func NewInferencePrecompile() (InferencePrecompile, error) {
	abi, err := i_inference.IInferenceMetaData.GetAbi()
	if err != nil {
		return InferencePrecompile{}, err
	}

	return InferencePrecompile{ABI: abi}, nil
}

func (i InferencePrecompile) WithContext(ctx context.Context) vm.PrecompiledContract {
	i.ctx = ctx
	return i
}

const (
	METHOD_INFER = "infer"
)

// ExtendedRun implements vm.ExtendedPrecompiledContract.
func (i InferencePrecompile) ExtendedRun(caller vm.ContractRef, input []byte, suppliedGas uint64, readOnly bool) (resBz []byte, usedGas uint64, err error) {
	method, err := i.ABI.MethodById(input)
	if err != nil {
		return nil, 0, types.ErrPrecompileFailed.Wrap(err.Error())
	}

	args, err := method.Inputs.Unpack(input[4:])
	if err != nil {
		return nil, 0, types.ErrPrecompileFailed.Wrap(err.Error())
	}

	ctx := sdk.UnwrapSDKContext(i.ctx).WithGasMeter(storetypes.NewGasMeter(suppliedGas))

	// TODO: calculate gas
	ctx.GasMeter().ConsumeGas(storetypes.Gas(10), "input bytes")

	switch method.Name {
	case METHOD_INFER:
		if readOnly {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrNonReadOnlyMethod.Wrap(method.Name)
		}

		var inferArguments InferenceRequestArguments
		if err := method.Inputs.Copy(&inferArguments, args); err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}

		input := inferArguments.Input

		resBz, err = method.Outputs.Pack(input)
		if err != nil {
			return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrPrecompileFailed.Wrap(err.Error())
		}
	default:
		return nil, ctx.GasMeter().GasConsumedToLimit(), types.ErrUnknownPrecompileMethod.Wrap(method.Name)
	}

	usedGas = ctx.GasMeter().GasConsumedToLimit()
	return resBz, usedGas, nil
}

// RequiredGas implements vm.PrecompiledContract.
func (i InferencePrecompile) RequiredGas(input []byte) uint64 {
	return 0
}

// Run implements vm.PrecompiledContract.
func (i InferencePrecompile) Run(input []byte) ([]byte, error) {
	return nil, errors.New("the InferencePrecompile works exclusively with ExtendedRun")
}
