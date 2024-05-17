package inferenceprecompile

type InferenceRequestArguments struct {
	ModelId []byte `abi:"model"`
	Input   []byte `abi:"input"`
}
