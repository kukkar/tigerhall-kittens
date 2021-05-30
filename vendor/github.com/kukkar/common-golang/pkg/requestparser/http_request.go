package requestparser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kukkar/common-golang/globalconst"
	"github.com/kukkar/common-golang/pkg/logger"
)

// // Get Body Data without draining body.
func getBody(c *gin.Context) ([]byte, error) {

	var apiKey string
	if uniqueID, ok := c.Get(globalconst.UniqueAPIKey); ok {
		apiKey = uniqueID.(string)
	} else {
		return nil, fmt.Errorf("unable to reterive api key")
	}
	var output []byte
	recievedContentType, _ := c.Request.Header[ContentType]
	var contentType string
	if len(recievedContentType) > 0 {
		contentType = recievedContentType[0]
	} else {
		contentType = "application/json"
	}
	contentTypeSplitData := strings.Split(contentType, ";")
	contentType = contentTypeSplitData[0]
	buildData := make(map[string]interface{})
	switch contentType {
	case ContentFormData:
		var err error
		c.Request.ParseForm()
		values := c.Request.Form
		for key, eachValue := range values {
			if len(eachValue) == 1 {
				if eachValue[0] == "" && key == "verifyotp" || key == "requestotp" {
					eachValue[0] = "empty"
				}
				buildData[key] = eachValue[0]
			} else {
				buildData[key] = eachValue
			}
		}
		// Marshal the map into a JSON string.
		output, err = json.Marshal(buildData)
		if err != nil {
			return nil, err
		}
	default:
		b := c.Request.Body
		var buf bytes.Buffer
		if _, err := buf.ReadFrom(b); err != nil {
			c.Request.Body = b
			return nil, err
		}
		if err := b.Close(); err != nil {
			c.Request.Body = b
			return nil, err
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(buf.Bytes()))
		output = buf.Bytes()
		//c.Set(REQUEST_CONTAINER,buf.Bytes())
	}
	logger.Logger.Info(fmt.Sprintf("%s Api Request %s", apiKey, output))
	return output, nil
}

// // Load Body into supplied struct.
func LoadBody(c *gin.Context, payload interface{}) error {
	dataBytes, err := getBody(c)
	if err != nil {
		return fmt.Errorf("*Request#LoadBody -> %s", err.Error())
	}
	return json.Unmarshal(dataBytes, payload)
}
