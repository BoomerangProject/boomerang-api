package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/BoomerangProject/boomerang-api/parameters"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getEthClient(endpointAddress string) *ethclient.Client {

	client, err := ethclient.Dial(endpointAddress)
	if err != nil {
		panic(err)
	}

	return client
}

type JsonObject struct {
	Abi json.RawMessage `json:"abi"`
}

func getAbiFromString(contractJson string) abi.ABI {

	parsed, err := abi.JSON(strings.NewReader(contractJson))

	if err != nil {
		fmt.Println(err)
		panic("error parsing contract json")
	}

	return parsed
}

func BoomerangContract() {

	endpointAddress := parameters.GetRinkebyEndpointAddress()
	client := getEthClient(endpointAddress)
	_ = client

	payerAddress := parameters.GetPayerAddress()
	_ = payerAddress
}

