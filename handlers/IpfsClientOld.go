package handlers

// import (
// 	"bytes"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"

// 	"mime/multipart"

// 	// "github.com/valyala/fasthttp"
// )

func AddFile3() {

	// url := "http://ec2-54-172-136-192.compute-1.amazonaws.com:5001/api/v0/add?pin=true"

	// filename := "/Users/panda/Desktop/file.json"

	// request, response := postFile2(filename, url)

	// bodyBytes := response.Body()
	// println(string(bodyBytes))
	// println(string(request.Body()))
	// response, err := postFile(filename, url)

	// _ = response
	// _ = err
	// fmt.Println(response)
	// fmt.Println(err)

}















// func postFile2(filename string, url string) (*fasthttp.Request, *fasthttp.Response) {

// 	bodyBuffer := &bytes.Buffer{}
// 	bodyWriter := multipart.NewWriter(bodyBuffer)

// 	// use the bodyWriter to write the Part headers to the buffer
// 	bodyWriter.CreateFormFile("upfile", filename)

// 	// the file data will be the second part of the body
// 	fh, _ := os.Open(filename)

// 	boundary := bodyWriter.Boundary()
// 	close_buf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))

// 	// use multi-reader to defer the reading of the file data until writing to the socket buffer.
// 	_ = io.MultiReader(bodyBuffer, fh, close_buf)
// 	fi, _ := fh.Stat()

// 	request := fasthttp.AcquireRequest()
// 	request.SetRequestURI(url)

// 	request.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
// 	request.Header.SetContentLength(int(fi.Size() + int64(bodyBuffer.Len()) + int64(close_buf.Len())))

// 	resp := fasthttp.AcquireResponse()
// 	client := &fasthttp.Client{}
// 	client.Do(request, resp)
// 	return request, resp
// }

// func postFile(filename string, target_url string) (*http.Response, interface{}) {

// 	bodyBuffer := bytes.NewBufferString("")
// 	bodyWriter := multipart.NewWriter(bodyBuffer)

// 	// use the bodyWriter to write the Part headers to the buffer
// 	_, err := bodyWriter.CreateFormFile("upfile", filename)
// 	if err != nil {
// 		fmt.Println("error writing to buffer")
// 		return nil, err
// 	}

// 	// the file data will be the second part of the body
// 	fh, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Println("error opening file")
// 		return nil, err
// 	}
// 	// need to know the boundary to properly close the part myself.
// 	boundary := bodyWriter.Boundary()
// 	fmt.Printf("\r\n--%s--\r\n", boundary)
// 	close_buf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))

// 	// use multi-reader to defer the reading of the file data until writing to the socket buffer.
// 	request_reader := io.MultiReader(bodyBuffer, fh, close_buf)
// 	fi, err := fh.Stat()
// 	if err != nil {
// 		fmt.Printf("Error Stating file: %s", filename)
// 		return nil, err
// 	}
// 	req, err := http.NewRequest("POST", target_url, request_reader)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Set headers for multipart, and Content Length
// 	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
// 	req.ContentLength = fi.Size() + int64(bodyBuffer.Len()) + int64(close_buf.Len())

// 	fmt.Println("the request is:")
// 	fmt.Println(req)
// 	fmt.Println("--------------------------------")

// 	return http.DefaultClient.Do(req)
// }
