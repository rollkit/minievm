// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

/// @dev The ICosmos contract's address.
address constant INFERENCE_ADDRESS = 0x0000000000000000000000000000000000000011;

/// @dev The ICosmos contract's instance.
IInference constant INFERENCE_CONTRACT = IInference(INFERENCE_ADDRESS);

interface IInference {

    function infer(
        bytes memory model, bytes memory input
    ) external returns (bytes memory);

}
