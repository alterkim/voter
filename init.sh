~!/usr/bin/env bash

rm -rf ~/.voterd
rm -rf ~/.votercli

voterd init test --chain-id=voterchain

votercli config output json
votercli config indent true
votercli config trust-node true
votercli config chain-id voterchain
votercli config keyring-backend test

votercli keys add user1
votercli keys add user2

voterd add-genesis-account $(nameservicecli keys show user1 -a) 1000token,100000000stake
voterd add-genesis-account $(nameservicecli keys show user2 -a) 1000token,100000000stake

nameserviced gentx --name user1 --keyring-backend test

echo "Collecting genesis txs..."
nameserviced collect-gentxs

echo "Validating genesis file..."
nameserviced validate-genesis