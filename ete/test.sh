export FILE_NAME=Test
export FILE_PATH=./$FILE_NAME.sol
export KEY_NAME=operator
export CHAIN_ID=minitiad-test
export NODE_RPC=tcp://127.0.0.1:36657

# compile solidity
mkdir build
solc $FILE_PATH --bin --abi -o build --overwrite

export DEPLOY_TX_HASH=$(     \
    minitiad tx evm create   \
    ./build/$FILE_NAME.bin   \
    --from $KEY_NAME         \
    --gas auto               \
    --node $NODE_RPC         \
    --yes                    \
    --chain-id $CHAIN_ID     \
    | grep txhash            \
    | sed 's/.*: //'         \
)
echo "Deploy tx hash: '$DEPLOY_TX_HASH'"

# find contract address from the tx
export CONTRACT_ADDRESS=$(
    minitiad q tx $DEPLOY_TX_HASH \
    --node $NODE_RPC              \
    | grep contract -A 1
)
echo "Contract address: $CONTRACT_ADDRESS"