package views

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeGraphiQLReturnsHTML(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(ServeGraphiQL))
	defer testServer.Close()

	expectedContent := GraphiQLContent

	res, _ := http.Get(testServer.URL)
	actualContent, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	bytes.Equal(expectedContent, actualContent)
}
