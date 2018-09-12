package services

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

const url = "http://ec2-54-172-136-192.compute-1.amazonaws.com:5001/api/v0/add?pin=true"

func AddJsonToIpfs(jsonString string) string {

	var bytesBuffer bytes.Buffer
	multipartWriter := multipart.NewWriter(&bytesBuffer)

	fieldWriter, _ := multipartWriter.CreateFormField("other")
	io.Copy(fieldWriter, strings.NewReader(jsonString))
	multipartWriter.Close()

	request, _ := http.NewRequest("POST", url, &bytesBuffer)

	// setting the content type will contain the boundary
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	response, _ := http.DefaultClient.Do(request)

	if response.StatusCode != http.StatusOK {
		panic("error when adding json to ipfs node: " + response.Status)
	}

	return getHash(&response.Body)
}

func getHash(responseBody *io.ReadCloser) string {

	type Params struct {
		Name string `json:"Name"`
		Hash string `json:"Hash"`
		Size string `json:"Size"`
	}

	decoder := json.NewDecoder(*responseBody)
	var params Params
	decoder.Decode(&params)

	return params.Hash
}
