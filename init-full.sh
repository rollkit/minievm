#!/bin/sh

export L1_CHAIN_ID=minitia-test
export L2_CHAIN_ID=minitia-test
export DENOM=unova
export OPERATOR=operator
export BRIDGE_EXECUTOR=bridge-executor
export RELAYER=relayer

CHAIN_ID=gm
BASE_DIR="$HOME/.minitia_fn"

rm -rf $BASE_DIR

minitiad --home "$BASE_DIR" init FullNode --chain-id $L2_CHAIN_ID --denom $DENOM

cp -R "$HOME/.minitia/config/genesis.json" "$BASE_DIR/config/genesis.json"

# DA height at which the rollkit sequencer produced the first rollup block
DA_BLOCK_HEIGHT=your-block-height

# P2p id of the rollkit sequencer
# you can find this from the sequencer logs 12:33PM INF listening on address=/ip4/10.0.0.154/tcp/36656/p2p/12D3KooWHMHQVm6xJp3CptTpbydvmzpVHj72w7yjVzKpyEqd8Zoc module=p2p
# e.g., `12D3KooWHMHQVm6xJp3CptTpbydvmzpVHj72w7yjVzKpyEqd8Zoc` from above log
P2P_ID="your-p2p-id"


# query the DA Layer start height, in this case we are querying
# our local devnet at port 26657, the RPC. The RPC endpoint is
# to allow users to interact with Celestia's nodes by querying
# the node's state and broadcasting transactions on the Celestia
# network. The default port is 26657.

DA_BLOCK_HEIGHT=11
#$(curl http://0.0.0.0:26657/block | jq -r '.result.block.header.height')

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

# start the chain
minitiad start --rollkit.da_auth_token=$AUTH_TOKEN --rollkit.da_namespace 00000000000000000000000000000000000000000008e5f679bf7116cb --rollkit.da_start_height $DA_BLOCK_HEIGHT --rpc.laddr tcp://127.0.0.1:46657 --grpc.address 127.0.0.1:9390 --p2p.seeds $P2P_ID@127.0.0.1:36656 --p2p.laddr "0.0.0.0:46656" --api.address tcp://localhost:1417 --minimum-gas-prices="0.025unova" --home $BASE_DIR

