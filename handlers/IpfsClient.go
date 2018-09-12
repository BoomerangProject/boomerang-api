package handlers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/valyala/fasthttp"
)

func AddFile() {

	remoteUrl := "http://ec2-54-172-136-192.compute-1.amazonaws.com:5001/api/v0/add?pin=true"

	//prepare the reader instances to encode
	values := map[string]io.Reader{
		"other": strings.NewReader("blaaamm"),
	}
	err := Upload(remoteUrl, values)
	if err != nil {
		panic(err)
	}
}

func Upload(url string, values map[string]io.Reader) (err interface{}) {
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

	

	request := fasthttp.AcquireRequest()
	request.Header.SetMethod("POST")
	request.Header.Set("Content-Type", w.FormDataContentType())
	request.SetRequestURI(url)
	// fasthttp.WriteMultipartForm(request.BodyWriter(), w, w.Boundary())
	request.SetBodyStream(&b, -1)

	fmt.Println("the request is:")
	fmt.Println(request)
	fmt.Println("--------------------------------")

	res := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	client.Do(request, res)

	fmt.Println("the response is:")
	fmt.Println(res)
	fmt.Println("--------------------------------")

	fmt.Println(string(res.Body()))

	return
}
