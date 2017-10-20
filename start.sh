#!/usr/bin/env bash
#/bin/bash
#rm -rf $GOPATH/src/lib
#rm -rf $GOROOT/src/lib
#cp -rf chaincode/lib $GOPATH/src/
#cp -rf chaincode/lib $GOROOT/src/

echo " ==================== Cleaning up containers ==================== "
docker ps | grep 'dev-peer0\|hyperledger' | awk '{print $1}' | xargs docker rm -f
echo " ==================== Cleaning up chaincode images ==================== "
##docker images | grep 'dev-peer0' | awk '{print $1}' | xargs docker rmi -f
CHANNEL_NAME=mainchannel TIMEOUT=50000 docker-compose -f docker-compose-cli.yaml up -d
docker logs -f cli
