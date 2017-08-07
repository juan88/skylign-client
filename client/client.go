package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"net/url"
	"mime/multipart"
	"io"
	"io/ioutil"
	"os"
	"bytes"
	//"net/http/httputil"
)

type UploadedAlignFileResponse struct {
	Url   string   `json:"url,omitempty"`
	Uuid     string `json:"uuid,omitempty"`
	Message   string `json:"message,omitempty"`
	Error *UploadResponseError
}

type UploadResponseError struct {
	ErrorMsg interface{} `json:"error,omitempty"`
}

func (e *UploadResponseError) UnmarshalJSON(data []byte) (err error) {
	dec := json.NewDecoder(bytes.NewReader(data))
    //var msg json.RawMessage
    var msg map[string]string

    if err := dec.Decode(&msg); err != nil {
        return fmt.Errorf("Error while decoding error response: %v", err)
    }

    e.ErrorMsg = msg
    return err
}



func Hello() {

	url := "http://skylign.org/"
	
	// Build the request
	//req, err := http.NewRequest("GET", url, nil)
	//if err != nil {
	//	log.Fatal("New request failed: ", err)
	//	return
	//}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	//client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	//resp, err := client.Do(req)
	//if err != nil {
	//	log.Fatal("Error while doing request: ", err)
	//	return
	//}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	//defer resp.Body.Close()

	var response UploadedAlignFileResponse;
	params := map[string]string{
		"processing": "observed",
		"letter_height": "info_content_all",
		"frag":	"full",
	}
	UploadData(url, "ADIN0.sto", params, &response)

}

func UploadData(url, file string, params map[string]string, response *UploadedAlignFileResponse) (err error) {
    // Prepare a form that you will submit to that URL.
    var b bytes.Buffer
    w := multipart.NewWriter(&b)
    // Add your image file
    f, err := os.Open(file)
    if err != nil {
        //log.Fatal(err)
        return 
    }
    defer f.Close()

    fw, err := w.CreateFormFile("file", file)
    if err != nil {
        //log.Fatal(err)
        return 
    }
    if _, err = io.Copy(fw, f); err != nil {
        //log.Fatal(err)
        return
    }
    
    for key, val := range params {
		_ = w.WriteField(key, val)
	}

    // Don't forget to close the multipart writer.
    // If you don't close it, your request will be missing the terminating boundary.
    err = w.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

    // Now that you have a form, you can submit it to your handler.
    req, err := http.NewRequest("POST", url, &b)
    if err != nil {
    	fmt.Errorf("Error while setting POST request")
    	log.Fatal(err)
        return 
    }

    // Don't forget to set the content type, this will contain the boundary.
    req.Header.Add("Content-Type", w.FormDataContentType())
    req.Header.Set("Accept", "application/json")
    req.Header.Set("User-Agent", "skylign-go")
    req.Header.Set("Host", "skylign.org")
    req.Header.Add("Accept-Encoding", "gzip")
    
    //debug(httputil.DumpRequestOut(req, true))
    

    // Submit the request
    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
    	fmt.Println("Error making request")
    	log.Fatal(err)
        return 
    }

    // Check the response
    if res.StatusCode != http.StatusOK {
        err = fmt.Errorf("bad status: %s", res.Status)
    }


    //fmt.Println("Code: %s", res.Status)
    //fmt.Println("Parseando respuesta")

    //debug(httputil.DumpResponse(res, true))

    // Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		fmt.Errorf("Decoding error!")
		log.Println(err)
	}

	bodyBytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bodyBytes))

	defer res.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(response.Url)
	//fmt.Println(response.Message)

	return
}

func debug(data []byte, err error) {
    if err == nil {
        fmt.Printf("%s\n\n", data)
    } else {
        log.Fatalf("%s\n\n", err)
    }
}