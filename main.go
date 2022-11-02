package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {

	tlsConfig := &tls.Config{
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	tlsConfig.PreferServerCipherSuites = true
	tlsConfig.MinVersion = tls.VersionTLS10
	tlsConfig.MaxVersion = tls.VersionTLS10

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := &http.Client{Transport: transport}

	req, err := http.NewRequest(http.MethodGet, "https://tls-v1-0.badssl.com:1010/", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Print(err)
		}
	}()
}
