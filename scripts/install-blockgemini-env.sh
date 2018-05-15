#!/usr/bin/env bash
CORE_PEER_LOCALMSPID="OrgBlockGeminiMSP"
CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org_blockgemini.example.com/peers/peer0.org_blockgemini.example.com/tls/ca.crt
CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org_blockgemini.example.com/users/Admin@org_blockgemini.example.com/msp
CORE_PEER_ADDRESS=peer0.org_blockgemini.example.com:7051