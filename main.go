package main

import (
	"fmt"
	"net"
	"net/http"
	"sort"
)

func getClientIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.Header.Get("X-Real-IP")
	}
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return ip
}

func handler(w http.ResponseWriter, r *http.Request) {
	clientIP := getClientIP(r)
	fmt.Fprintf(w, "Client IP: %s\n\n", clientIP)

	fmt.Fprintf(w, "Request Method: %s\n", r.Method)

	fmt.Fprintf(w, "Request Path: %s\n", r.URL.Path)

	fmt.Fprintf(w, "Protocol Version: %s\n\n", r.Proto)

	fmt.Fprintf(w, "Request Headers:\n")
	fmt.Fprintf(w, "Host: %s\n", r.Host) // 从 r.Host 获取 Host
	var headers []string
	for name := range r.Header {
		headers = append(headers, name)
	}
	sort.Strings(headers)

	for _, name := range headers {
		values := r.Header[name]
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)

	port := ":8080"
	fmt.Printf("Server is starting, listening on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
