package parameters

const infuraToken = "2wx4womFMFUEyRBJKbKq"

func GetRopstenEndpointAddress() string {
	return "https://ropsten.infura.io/" + infuraToken
}

func GetRinkebyEndpointAddress() string {
	return "https://rinkeby.infura.io/" + infuraToken
}

func GetKovanEndpointAddress() string {
	return "https://kovan.infura.io/" + infuraToken
}
