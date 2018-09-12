package services

import (
	"fmt"
	"log"

	"github.com/BoomerangProject/boomerang-api/parameters"
	"github.com/BoomerangProject/boomerang-api/contracts"
)

func BoomerangContract() {

	endpointAddress := parameters.GetRinkebyEndpointAddress()
	client := getEthClient(endpointAddress)

	payerAddress := parameters.GetPayerAddress()

	instance, err := contracts.NewBoomerangToken(payerAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)
}
