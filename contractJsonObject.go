package main

import (
	"io"
	"io/ioutil"
	"os"
)

func getContractJsonAsString() string {

	jsonFile, err := ioutil.ReadFile("boomerangTokenContract.json")

	if err != nil {
		panic("error opening contract json file")
	}

	return string(jsonFile)
}

func getContractJsonAsIoReader() io.Reader {

	jsonFile, err := os.Open("boomerangTokenContract.json")

	if err != nil {
		panic("error opening contract json file")
	}

	return jsonFile
}
