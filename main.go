package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	director := func(req *http.Request) {
		origin, _ := url.Parse("http://www.fanfiction.net/")
		req.Header.Add("X-Forwarded-Host", req.Host)
		req.Header.Add("X-Origin-Host", origin.Host)
		req.URL.Scheme = "http"
		req.URL.Host = origin.Host
	}
	responseModder := func(res *http.Response) error {
		res.Header.Add("Access-Control-Allow-Origin", "*")
		res.Header.Add("Access-Control-Allow-Headers", "Content-Type")
		return nil
	}
	proxy := &httputil.ReverseProxy{
		Director: director,
		ModifyResponse: responseModder,
	}

	// Install handlers
	http.Handle("/", proxy)

	// Run server
	http.ListenAndServe(":" + port, nil)
}
