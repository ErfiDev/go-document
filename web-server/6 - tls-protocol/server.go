package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	srv := NewServer()
	// install openssl library
	// run the command `openssl genrsa -out server.key 2048` in the root of your app
	// then run this command `openssl ecparam -genkey -name secp384r1 -out server.key`
	/*
		if the server return warning
		called bad tls handshake error
		don't worry about that
		becuase this tls we generated in localy
		and browser check global tls keys
		and return the warning
	*/
	err := srv.ListenAndServeTLS("server.crt", "server.key")
	log.Println("Server run on 5000 port")
	if err != nil {
		log.Fatalf("we have error when running server: %v", err)
	}
}

func NewServer() http.Server {
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Goood!"))
	})

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
		},
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	server := http.Server{
		Addr:      ":5000",
		Handler:   router,
		TLSConfig: tlsConfig,
	}

	return server
}
