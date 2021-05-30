package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//UserActivity request activity
type UserActivity struct {
	Header       http.Header
	Method       string
	Body         interface{}
	Status       int
	ResponseTime time.Time
	IP           string
	UserAgent    string
}

// DebugMiddleware : Request debug
func DebugMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		addIfExists(c, c.Request.Header, "debug")
		addIfExists(c, c.Request.Header, "trace_depth")
		logReq(c)
	}
}

//logReq will log request if debug is passed
func logReq(ctx *gin.Context) {
	var body interface{}
	ctx.Request.Body, body = copyReqBody(ctx.Request.Body)
	var ua = &UserActivity{
		Header: ctx.Request.Header,
		Method: ctx.Request.Method,
		Body:   body,
	}

	ctx.Next()
	log.Println(ua)
}

func copyReqBody(reqBody io.ReadCloser) (oBody io.ReadCloser, dBody interface{}) {
	bodyByte, _ := ioutil.ReadAll(reqBody)
	if len(bodyByte) > 0 {
		var out bytes.Buffer
		if err := json.Indent(&out, bodyByte, "", "  "); err == nil {
			dBody = string(out.String())
		}
	}
	oBody = ioutil.NopCloser(bytes.NewBuffer(bodyByte))
	return
}

func addIfExists(c *gin.Context, header http.Header, key string) {
	if v := header.Get(key); v != "" {
		c.Set(key, v)
	}
}
