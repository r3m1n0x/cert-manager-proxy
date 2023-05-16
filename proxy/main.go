package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
)

func main() {
	// Parse command-line flags
	backend := flag.String("backend", "http://localhost:8080", "backend server URL")
	frontend := flag.String("frontend", ":80", "frontend server address")
	flag.Parse()

	// Create the reverse proxy with ProxyProtocol support
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			// Set the backend URL as the request destination
			req.URL.Scheme = "http"
			req.URL.Host = *backend

			// Modify the request to add Proxy Protocol header
			req.Header.Set("Proxy-Protocol", req.RemoteAddr)
		},
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				// Dial the backend server
				conn, err := net.Dial(network, addr)
				if err != nil {
					return nil, err
				}

				// Check if the connection is using TCP
				tcpConn, ok := conn.(*net.TCPConn)
				if !ok {
					return nil, nil // Skip Proxy Protocol for non-TCP connections
				}

				// Write the Proxy Protocol header before returning the connection
				proxyProtoHeader := []byte("PROXY " + tcpConn.RemoteAddr().String() + " " + tcpConn.LocalAddr().String() + "\r\n")
				_, err = tcpConn.Write(proxyProtoHeader)
				if err != nil {
					return nil, err
				}

				return conn, nil
			},
		},
	}

	// Start the frontend server
	log.Printf("Starting frontend server on %s", *frontend)
	err := http.ListenAndServe(*frontend, proxy)
	if err != nil {
		log.Fatal(err)
	}
}
