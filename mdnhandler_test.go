package mdnshandler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMDNSHandler(t *testing.T) {
	// Create a new MDNSHandler instance with the test parameters.
	handler := MDNSHandler{
		Name:    "example.local",
		Service: "_http._tcp.local",
	}

	// Create a new HTTP request to the test URL.
	request := httptest.NewRequest(http.MethodGet, "http://example.local/test", nil)

	// Create a new HTTP response recorder to capture the response.
	recorder := httptest.NewRecorder()

	// Call the handler's ServeHTTP method with the test request and response.
	if err := handler.ServeHTTP(recorder, request, nil); err != nil {
		t.Fatalf("failed to serve request: %v", err)
	}

	// Check that the response status code is 502.
	if recorder.Code != http.StatusBadGateway {
		t.Fatalf("unexpected response status code: %d", recorder.Code)
	}

	// Check that the response body is empty.
	responseBody, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	if len(responseBody) != 0 {
		t.Fatalf("unexpected response body: %s", responseBody)
	}
}
