package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	//"github.com/aws/aws-sdk-go/service/s3"
)

type Params struct {
	BusinessAddress     string `json:"businessAddress"`
	BusinessName        string `json:"businessName"`
	BusinessDescription string `json:"businessDescription"`
}

func error(responseWriter http.ResponseWriter, errorMessage string) {
	responseWriter.WriteHeader(http.StatusBadRequest)
	io.WriteString(responseWriter, errorMessage)
}

func RegisterAsBusiness(responseWriter http.ResponseWriter, httpRequest *http.Request) {

	decoder := json.NewDecoder(httpRequest.Body)
	var params Params
	decoder.Decode(&params)

	if len(params.BusinessAddress) < 1 {
		error(responseWriter, "businessAddress is required")
		return
	}

	if len(params.BusinessName) < 1 {
		error(responseWriter, "businessName is required")
		return
	}

	if len(params.BusinessDescription) < 1 {
		error(responseWriter, "businessDescription is required")
		return
	}

	AddFile2()

	// add to ipfs

	// do aws stuff with s3

	// create the transaction
	// sign the transaction
	// broadcast it to infura

	// send the response back to the user

	responseWriter.WriteHeader(http.StatusOK)
	io.WriteString(responseWriter, "okay")
}
