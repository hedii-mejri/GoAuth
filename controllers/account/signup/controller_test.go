package signup

import (
	"authsys/middlewares"
	"bytes"
	//"fmt"
	"github.com/codegangsta/negroni"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createHttpServerTest(handler http.Handler) *httptest.Server {
	return httptest.NewServer(handler)
}

func createMiddlewareHandler(middlewares []negroni.Handler, handler http.Handler) *negroni.Negroni {

	n := negroni.New(middlewares...)
	n.UseHandler(handler)
	return n

}

func TestInputFields(t *testing.T) {

	ts := createHttpServerTest(createMiddlewareHandler(middlewares.New(), New()))
	defer ts.Close()

	client := &http.Client{}
	response, err := client.Get(ts.URL)
	assert.NoError(t, err, "Should not contain any error.")

	buffer := new(bytes.Buffer)
	io.Copy(buffer, response.Body)
	assert.Contains(t, buffer.String(), "Sign Up", "Should contain sign up title.")

}

func TestCaptchaAvailable(t *testing.T) {
		ts := createHttpServerTest(createMiddlewareHandler(middlewares.New(), New()))
	defer ts.Close()

	client := &http.Client{}
	response, err := client.Get(ts.URL)
	assert.NoError(t, err, "Should not contain any error.")

	buffer := new(bytes.Buffer)
	io.Copy(buffer, response.Body)
	assert.Contains(t, buffer.String(), "<img", "Should contain image tag.")
}
