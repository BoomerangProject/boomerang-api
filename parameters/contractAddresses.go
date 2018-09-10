package parameters

import "github.com/ethereum/go-ethereum/common"

func GetPayerAddress() common.Address {
	return common.HexToAddress("0xeb37ebf8ee37a3aabde4de006cf47a40837e029e")
}

func GetBoomerangTokenAddress() common.Address {
	return common.HexToAddress("0x4ba8fe72892481827d4d666f4ab17c130809cfeb")
}
