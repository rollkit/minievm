// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

address constant INFERENCE_ADDRESS = 0x0000000000000000000000000000000000000011;

IInference constant INFERENCE_CONTRACT = IInference(INFERENCE_ADDRESS);

interface IInference {

    function infer(
        bytes memory model, bytes memory input
    ) external returns (bytes memory);

}

contract Test {

    function run() public {
        INFERENCE_CONTRACT.infer(
            abi.encodePacked("QmbbzDwqSxZSgkz1EbsNHp2mb67rYeUYHYWJ4wECE24S7A"),
            abi.encodePacked("[[0.07286679744720459, 0.4486280083656311]]")
        );
    }
}