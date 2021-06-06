/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

type Product struct {
	Tagid       string    `json:"tagid"`
	Productcode string    `json:"productcode"`
	Modelname   string    `json:"modelname"`
	Brand       string    `json:"brand"`
	Color       string    `json:"color"`
	Photo       string    `json:"photo"`
	Date        time.Time `json:"date"`
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryProduct" {
		return s.queryProduct(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createProduct" {
		return s.createProduct(APIstub, args)
	} else if function == "changeProduct" {
		return s.changeProduct(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	products := []Product{
		Product{
			Tagid:       "0000",
			Productcode: "0000",
			Brand:       "ZERO",
			Color:       "BLACK",
			Modelname:   "ZERO000",
			Photo:       "https://file.mk.co.kr/meet/neds/2017/01/image_readtop_2017_51822_14850815582756240.jpg",
			Date:        time.Now(),
		},
	}

	i := 0
	for i < len(products) {
		productAsBytes, _ := json.Marshal(products[i])
		APIstub.PutState(products[i].Tagid, productAsBytes)
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) queryProduct(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	productsAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(productsAsBytes)
}

func (s *SmartContract) createProduct(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var product = Product{
		Tagid:       args[0],
		Productcode: args[1],
		Modelname:   args[2],
		Brand:       args[3],
		Photo:       args[4],
		Color:       args[5],
		Date:        time.Now(),
	}

	productAsBytes, _ := json.Marshal(product)
	APIstub.PutState(args[0], productAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) changeProduct(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	productAsBytes, _ := APIstub.GetState(args[0])
	product := Product{}

	json.Unmarshal(productAsBytes, &product)
	product.Tagid = args[0]
	product.Productcode = args[1]
	product.Modelname = args[2]
	product.Brand = args[3]
	product.Photo = args[4]
	product.Color = args[5]

	productAsBytes, _ = json.Marshal(product)
	APIstub.PutState(args[0], productAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
