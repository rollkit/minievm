#!/bin/sh

export L1_CHAIN_ID=minitia-test
export L2_CHAIN_ID=minitia-test
export DENOM=unova
export OPERATOR=operator
export BRIDGE_EXECUTOR=bridge-executor
export RELAYER=relayer

# query the DA Layer start height, in this case we are querying
# our local devnet at port 26657, the RPC. The RPC endpoint is
# to allow users to interact with Celestia's nodes by querying
# the node's state and broadcasting transactions on the Celestia
# network. The default port is 26657.
DA_BLOCK_HEIGHT=$(curl http://0.0.0.0:26657/block | jq -r '.result.block.header.height')

AUTH_TOKEN=$(docker exec $(docker ps -q) celestia bridge auth admin --node.store /home/celestia/bridge)
# AUTH_TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJwdWJsaWMiLCJyZWFkIiwid3JpdGUiLCJhZG1pbiJdfQ.NiEBetac-cQ_gTRioOY2ZMBitRlzQyScs1claK6sfmE"

# rollkit logo
cat <<'EOF'

                 :=+++=.                
              -++-    .-++:             
          .=+=.           :++-.         
       -++-                  .=+=: .    
   .=+=:                        -%@@@*  
  +%-                       .=#@@@@@@*  
    -++-                 -*%@@@@@@%+:   
       .=*=.         .=#@@@@@@@%=.      
      -++-.-++:    =*#@@@@@%+:.-++-=-   
  .=+=.       :=+=.-: @@#=.   .-*@@@@%  
  =*=:           .-==+-    :+#@@@@@@%-  
     :++-               -*@@@@@@@#=:    
        =%+=.       .=#@@@@@@@#%:       
     -++:   -++-   *+=@@@@%+:   =#*##-  
  =*=.         :=+=---@*=.   .=*@@@@@%  
  .-+=:            :-:    :+%@@@@@@%+.  
      :=+-             -*@@@@@@@#=.     
         .=+=:     .=#@@@@@@%*-         
             -++-  *=.@@@#+:            
                .====+*-.  

   ______         _  _  _     _  _   
   | ___ \       | || || |   (_)| |  
   | |_/ /  ___  | || || | __ _ | |_ 
   |    /  / _ \ | || || |/ /| || __|
   | |\ \ | (_) || || ||   < | || |_ 
   \_| \_| \___/ |_||_||_|\_\|_| \__|
EOF

# echo variables for the chain
echo -e "\n Your DA_BLOCK_HEIGHT is $DA_BLOCK_HEIGHT \n"
echo -e "\n Your AUTH_TOKEN is $AUTH_TOKEN \n"

# minitiad init --chain-id=<chain-id> --denom <denom> <moniker>
minitiad init --chain-id=$L2_CHAIN_ID --denom $DENOM $OPERATOR

# minitiad keys add <key-name>
minitiad keys add $OPERATOR
minitiad keys add $BRIDGE_EXECUTOR
echo "milk verify alley price trust come maple will suit hood clay exotic" | minitiad keys add $RELAYER --recover

# minitiad genesis add-genesis-validator <key-name>
minitiad genesis add-genesis-validator $OPERATOR

sed -i -e \
    "s/\"bridge_executor\": \"[a-zA-Z0-9]*\",$/\"bridge_executor\": \"$(minitiad keys show $BRIDGE_EXECUTOR -a)\",/g" \
    ~/.minitia/config/genesis.json

# minitiad genesis add-genesis-account <key-or-address> <amount>
minitiad genesis add-genesis-account $OPERATOR 1000000000000000$DENOM
minitiad genesis add-genesis-account $BRIDGE_EXECUTOR 1000000000000000$DENOM
minitiad genesis add-genesis-account $RELAYER 1000000000000000$DENOM

# minitiad genesis add-fee-whitelist <key-or-address> 
minitiad genesis add-fee-whitelist $OPERATOR 
minitiad genesis add-fee-whitelist $BRIDGE_EXECUTOR 
minitiad genesis add-fee-whitelist $RELAYER 

#minitiad start
# copy centralized sequencer address into genesis.json
# Note: validator and sequencer are used interchangeably here
ADDRESS=$(jq -r '.address' ~/.minitia/config/priv_validator_key.json)
PUB_KEY=$(jq -r '.pub_key' ~/.minitia/config/priv_validator_key.json)
jq --argjson pubKey "$PUB_KEY" '.consensus["validators"]=[{"address": "'$ADDRESS'", "pub_key": $pubKey, "power": "1", "name": "Rollkit Sequencer"}]' ~/.minitia/config/genesis.json > temp.json && mv temp.json ~/.minitia/config/genesis.json

# create a restart-local.sh file to restart the chain later
[ -f restart-local.sh ] && rm restart-local.sh
echo "DA_BLOCK_HEIGHT=$DA_BLOCK_HEIGHT" >> restart-local.sh
echo "AUTH_TOKEN=$AUTH_TOKEN" >> restart-local.sh

echo "minitiad start --rollkit.aggregator --rollkit.aggregator --rollkit.da_auth_token=\$AUTH_TOKEN --rollkit.da_namespace 00000000000000000000000000000000000000000008e5f679bf7116cb --rollkit.da_start_height \$DA_BLOCK_HEIGHT --rpc.laddr tcp://127.0.0.1:36657 --grpc.address 127.0.0.1:9290 --p2p.laddr \"0.0.0.0:36656\" --minimum-gas-prices="0.025stake"" >> restart-local.sh

# start the chain
minitiad start --rollkit.aggregator --rollkit.da_auth_token=$AUTH_TOKEN --rollkit.da_namespace 00000000000000000000000000000000000000000008e5f679bf7116cb --rollkit.da_start_height $DA_BLOCK_HEIGHT --rpc.laddr tcp://127.0.0.1:36657 --grpc.address 127.0.0.1:9290 --p2p.laddr "0.0.0.0:36656" --minimum-gas-prices="0.025unova"

