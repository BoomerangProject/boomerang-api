package handlers

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func AddFile2() {

	remoteUrl := "http://ec2-54-172-136-192.compute-1.amazonaws.com:5001/api/v0/add?pin=true"

	//prepare the reader instances to encode
	values := map[string]io.Reader{
		// "file":  mustOpen("file.json"), // lets assume its this file
		"other": strings.NewReader("hello world2!"),
	}
	err := Upload2(http.DefaultClient, remoteUrl, values)
	if err != nil {
		panic(err)
	}
}

func Upload2(client *http.Client, url string, values map[string]io.Reader) (err interface{}) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}

	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}

	fmt.Println("the request is:")
	fmt.Println(req)
	fmt.Println("--------------------------------")

	defer req.Body.Close()

	//block forever at the next line
	reqBody, _ := ioutil.ReadAll(req.Body)

	fmt.Println(string(reqBody))

	fmt.Println("*****")

	fmt.Println("the response is:")
	fmt.Println(res)
	fmt.Println("--------------------------------")

	defer res.Body.Close()

	//block forever at the next line
	content, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(content))

	return
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return r
}
