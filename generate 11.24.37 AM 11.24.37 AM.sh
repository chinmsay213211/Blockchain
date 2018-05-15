#!/bin/bash
export FABRIC_CFG_PATH=$PWD
export CHANNEL_NAME=mainchannel

./bin/cryptogen generate --config=./crypto-config.yaml
./bin/configtxgen -profile MainOrdererGenesis -outputBlock ./channel-artifacts/genesis.block
./bin/configtxgen -profile MainChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID $CHANNEL_NAME
./bin/configtxgen -profile MainChannel -outputAnchorPeersUpdate ./channel-artifacts/OrgTristarMSPanchors.tx -channelID $CHANNEL_NAME -asOrg OrgTristarMSP
./bin/configtxgen -profile MainChannel -outputAnchorPeersUpdate ./channel-artifacts/OrgAgilityMSPanchors.tx -channelID $CHANNEL_NAME -asOrg OrgAgilityMSP
./bin/configtxgen -profile MainChannel -outputAnchorPeersUpdate ./channel-artifacts/OrgBlockGeminiMSPanchors.tx -channelID $CHANNEL_NAME -asOrg OrgBlockGeminiMSP
