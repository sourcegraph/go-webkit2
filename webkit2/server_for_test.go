package webkit2

import (
	"net/http"
	"net/http/httptest"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// webView is the WebView being tested.
	webView *WebView

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

// setup sets up a test HTTP server and a WebView. Tests should register
// handlers on mux which provide mock responses for the method being tested.
func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	webView = NewWebView()
}

// teardown closes the test HTTP server and webengine.View.
func teardown() {
	server.Close()
	webView.Destroy()
}
