package main

import (
	"net/http"
	"net/http/httputil"
)

// CustomRoundTripper is a custom implementation of http.RoundTripper
type CustomRoundTripper struct {
	Transport http.RoundTripper
}

// RoundTrip is the method to handle the HTTP round trip
func (c *CustomRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Remove the "X-Forwarded-For" header from the request before sending it to the backend server
	req.Header.Del("X-Forwarded-For")

	// Perform the request using the original transport
	return c.Transport.RoundTrip(req)
}

func TestCustomRoundTripper() {
	// Create a custom RoundTripper and set it as the transport for the http.Client
	customTransport := &CustomRoundTripper{
		Transport: http.DefaultTransport,
	}

	client := &http.Client{
		Transport: customTransport,
	}
	_ = client
	// Create a reverse proxy with the custom http.Client
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			// Modify the request as needed before forwarding it
			req.URL.Scheme = "http"
			req.URL.Host = "backend-server:8080" // Replace with your backend server address
		},
		Transport: customTransport, // Use the same custom transport for the reverse proxy
	}

	// Start the proxy server
	http.HandleFunc("/", proxy.ServeHTTP)
	http.ListenAndServe(":8080", nil)
}
