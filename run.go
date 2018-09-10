package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/BoomerangProject/boomerang-api/services"
	"github.com/ethereum/go-ethereum/accounts/abi"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func main() {

	services.BoomerangContract()
	// client, err := ethclient.Dial("https://rinkeby.infura.io/2wx4womFMFUEyRBJKbKq")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(client)

	// address := parameters.GetPayerAddress()
	// fmt.Println(address)

	// jsonFile := getContractJsonAsIoReader()
	// fmt.Println(jsonFile)

	// // someObject := jsonFile.Unmarshal(raw, &tentacles)
	// // fmt.Println(someObject)

	// // contractAbi := getAbiFromString(jsonFile)
	// contractAbi := getAbiFromIoReader(jsonFile)

	// fmt.Println(contractAbi)

	// instance, err := store.NewStore(address, client)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

// // NewStore creates a new instance of Store, bound to a specific deployed contract.
// func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
// 	contract, err := bindStore(address, backend, backend, backend)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
// }

// bindStore binds a generic wrapper to an already deployed contract.
// func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {

// 	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
// }

func getAbiFromString(contractJson string) abi.ABI {

	parsed, err := abi.JSON(strings.NewReader(contractJson))

	if err != nil {
		fmt.Println(err)
		panic("error parsing contract json")
	}

	return parsed
}

func getAbiFromIoReader(contractJson io.Reader) abi.ABI {

	parsed, err := abi.JSON(contractJson)

	if err != nil {
		fmt.Println(err)
		panic("error parsing contract json")
	}

	return parsed
}
