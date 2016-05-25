// Package httputil provides various convenience functions for request/response handling
package httputil

import (
	"net/http"
)

func NewResponseUtil(response http.ResponseWriter) ResponseUtil {
	ru := ResponseUtil{}
	ru.response = response
	return ru
}

type ResponseUtil struct {
	response http.ResponseWriter
}

func (ru ResponseUtil) WriteString(content string) {
	ru.response.Write([]byte(content))
}
