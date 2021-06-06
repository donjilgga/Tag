#!/bin/bash

cd ../java
mvn install

sleep 5

cd target
cp blockchain-java-sdk-0.0.1-SNAPSHOT-jar-with-dependencies.jar blockchain-client.jar

sleep 5

cp blockchain-client.jar ../../network_resources


cd ../../network_resources
java -cp blockchain-client.jar org.example.network.CreateChannel


java -cp blockchain-client.jar org.example.network.DeployInstantiateChaincode


java -cp blockchain-client.jar org.example.user.RegisterEnrollUser

java -cp blockchain-client.jar org.example.chaincode.invocation.InvokeChaincode

java -cp blockchain-client.jar org.example.chaincode.invocation.QueryChaincode