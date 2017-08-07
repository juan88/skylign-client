package main

import (
	"github.com/juan88/skylign-client/client"
)

func main() {

	url := "http://skylign.org/"
	
	var response client.UploadedAlignFileResponse;
	params := map[string]string{
		"processing": "observed",
		"letter_height": "info_content_all",
		"frag":	"full",
	}
	client.UploadData(url, "ADIN0.sto", params, &response)
}