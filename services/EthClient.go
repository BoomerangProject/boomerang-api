package services

import "github.com/ethereum/go-ethereum/ethclient"

func getEthClient(endpointAddress string) *ethclient.Client {

	client, err := ethclient.Dial(endpointAddress)
	if err != nil {
		panic(err)
	}

	return client
}