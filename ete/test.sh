export FILE_NAME=Test
export FILE_PATH=./$FILE_NAME.sol
export KEY_NAME=operator
export CHAIN_ID=minitia-test
export NODE_RPC=tcp://127.0.0.1:36657

# compile solidity
mkdir build
solc $FILE_PATH --bin --abi -o build --overwrite

# Run the first command and capture the output
DEPLOY_OUTPUT=$(minitiad tx evm create ./build/$FILE_NAME.bin --from $KEY_NAME --gas auto --node $NODE_RPC --yes --chain-id $CHAIN_ID)

# Extract the transaction hash from the output
export DEPLOY_TX_HASH=$(echo "$DEPLOY_OUTPUT" | grep txhash | sed 's/.*: //')
echo "Deploy tx hash: '$DEPLOY_TX_HASH'"

# Wait for txn processing
sleep 1

# Get the deploy contract tx
CONTRACT_OUTPUT=$(minitiad q tx --node $NODE_RPC $DEPLOY_TX_HASH)

# Extract the contract address from the output
export CONTRACT_ADDRESS=$(echo "$CONTRACT_OUTPUT" | grep -A 1 "key: contract" | grep "value:" | sed 's/.*value: //')
echo "Contract address: $CONTRACT_ADDRESS"