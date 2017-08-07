package client

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOkResponse(t *testing.T) {
	var jsonBlob = []byte(`{
		"url":"http://skylign.org:8000/logo/6BBFEB96-E7E0-11E2-A243-DF86A4A34227",
		"uuid":"6BBFEB96-E7E0-11E2-A243-DF86A4A34227",
		"message":"Logo generated successfully"
	}`)

	var response UploadedAlignFileResponse
	json.Unmarshal(jsonBlob, &response)

	assert.Equal(t, response.Url, "http://skylign.org:8000/logo/6BBFEB96-E7E0-11E2-A243-DF86A4A34227", "Fields should be parsed ok")
	assert.Equal(t, response.Uuid, "6BBFEB96-E7E0-11E2-A243-DF86A4A34227", "Fields should be parsed ok")
	assert.Equal(t, response.Message, "Logo generated successfully", "Fields should be parsed ok")
	assert.Nil(t, response.Error)
}

func TestErrorFile(t *testing.T) {
	var jsonBlob = []byte(`{
	  "error" : {
	    "upload" : "Please choose an alignment or HMM file to upload."
	  }
	}`)

	var response UploadedAlignFileResponse
	json.Unmarshal(jsonBlob, &response)

	errors := map[string]string{
		"upload": "Please choose an alignment or HMM file to upload.",
	}

	assert.Equal(t, response.Error.ErrorMsg, errors)
}

func TestErrorMalformedStockholm(t *testing.T) {
	var jsonBlob = []byte(`{
    "error": {
        "hmmbuild": "There was a problem parsing your multiple sequence alignment. It looked like you were trying to upload using STOCKHOLM format, but we couldn't parse it because: missing // terminator after MSA at or near line 4.\n"
    	}
	}`)

	var response UploadedAlignFileResponse
	json.Unmarshal(jsonBlob, &response)

	errors := map[string]string{
		"hmmbuild": "There was a problem parsing your multiple sequence alignment. It looked like you were trying to upload using STOCKHOLM format, but we couldn't parse it because: missing // terminator after MSA at or near line 4.\n",
	}

	assert.Equal(t, response.Error.ErrorMsg, errors)
}

func TestMultipleErrors(t *testing.T) {
	var jsonBlob = []byte(`{
    "error": {
        "hmmbuild": "There was a problem parsing your multiple sequence alignment. It looked like you were trying to upload using STOCKHOLM format, but we couldn't parse it because: missing // terminator after MSA at or near line 4.\n",
        "upload": "Please choose an alignment or HMM file to upload."
    	}
	}`)

	var response UploadedAlignFileResponse
	json.Unmarshal(jsonBlob, &response)

	errors := map[string]string{
		"hmmbuild": "There was a problem parsing your multiple sequence alignment. It looked like you were trying to upload using STOCKHOLM format, but we couldn't parse it because: missing // terminator after MSA at or near line 4.\n",
		"upload":   "Please choose an alignment or HMM file to upload.",
	}

	assert.Equal(t, response.Error.ErrorMsg, errors)
}
