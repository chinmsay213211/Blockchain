#Installation process
CHANNEL_NAME=mainchannel TIMEOUT=50000 docker-compose -f docker-compose-cli.yaml up -d --force-recreate

#companies_chaincode
CHAINCODE=companies_chaincode
peer chaincode install -n $CHAINCODE -v 1.0 -p github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example01
sleep 2
peer chaincode instantiate -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n $CHAINCODE -v 1.0 -c '{"Args":["init"]}' -P "OR('OrgTristarMSP.member','OrgAgilityMSP.member', 'OrgBlockGeminiMSP.member')"
sleep 10

peer chaincode invoke -o orderer.example.com:7050  --tls $CORE_PEER_TLS_ENABLED --cafile $ORDERER_CA -C $CHANNEL_NAME -n $CHAINCODE -c '{"Args":["create", "{\"Hash\":\"hash_\",\"Name\":\"Tristar\",\"PhoneNumber\":\"7777777\",\"Email\":\"company@mail.com\",\"Active\":true,\"Deleted\":true,\"Timestamp\":111}"]}'
sleep 10

peer chaincode query -C $CHANNEL_NAME -n $CHAINCODE -c '{"Args":["get","hash_"]}'

org.acme.sample.SampleParticipant
