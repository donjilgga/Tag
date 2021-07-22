#!/bin/bash

cd ../java
sudo mvn install

sleep 5

cd target
sudo cp blockchain-java-sdk-0.0.1-SNAPSHOT-jar-with-dependencies.jar blockchain-client.jar

sleep 5

sudo cp blockchain-client.jar ../../network_resources


cd ../../network_resources
sudo java -cp blockchain-client.jar org.example.network.CreateChannel

sudo java -cp blockchain-client.jar org.example.network.DeployInstantiateChaincode

sudo java -cp blockchain-client.jar org.example.user.RegisterEnrollUser

sudo java -cp blockchain-client.jar org.example.chaincode.invocation.InvokeChaincode

sudo java -cp blockchain-client.jar org.example.chaincode.invocation.QueryChaincode