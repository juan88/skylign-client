package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"mime/multipart"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"bytes"
	//"net/http/httputil"
)

const url = "http://skylign.org/"

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


func GenerateLogo(file string, params map[string]string) (err error) {
	var b bytes.Buffer
	var response UploadedAlignFileResponse;
	err = UploadData(file, params, &response)
	
	if err == nil && response.Error == nil {
		req, err := http.NewRequest("GET", response.Url, &b)
	    if err != nil {
	    	fmt.Errorf("Error while setting GET request")
	    	log.Fatal(err)
	        return err
	    }

	    // Don't forget to set the content type, this will contain the boundary.
	    req.Header.Set("Accept", "image/png")
	    req.Header.Set("Accept-Encoding", "gzip")
	    
	    //debug(httputil.DumpRequestOut(req, true))
	    

	    // Submit the request
	    client := &http.Client{}
	    res, err := client.Do(req)
	    if err != nil {
	    	fmt.Println("Error making request")
	    	log.Fatal(err)
	        return err
	    }

	    //debug(httputil.DumpResponse(res, true))

	    // Check the response
	    if res.StatusCode == http.StatusOK {
	        data, err := ioutil.ReadAll(res.Body)
	        res.Body.Close()
	        ioutil.WriteFile(strings.Replace(file, ".sto", ".png", 1), data, 0666)

	        return err
	    }
	}

	return err
}


func UploadData(file string, params map[string]string, response *UploadedAlignFileResponse) (err error) {
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