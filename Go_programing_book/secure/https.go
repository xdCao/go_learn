package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

const SERVER_PORT = 8888
const SERVER_DOMAIN = "localhost"
const RESPONSE_TEMPLATE = "hello"

func rootHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html")
	w.Header().Set("Content-length", fmt.Sprint(len(RESPONSE_TEMPLATE)))
	w.Write([]byte(RESPONSE_TEMPLATE))
}

func main() {
	http.HandleFunc(fmt.Sprint("%s:%d/", SERVER_DOMAIN, SERVER_PORT), rootHandler)
	err := http.ListenAndServeTLS(fmt.Sprintf(":%d", SERVER_PORT), "rui.crt", "rui.key", nil)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func YourListenAndServeTLS(addr string, certFile string, keyFile string, handler http.Handler) error {
	config := &tls.Config{
		Rand:       rand.Reader,
		Time:       time.Now,
		NextProtos: []string{"http/1.1"},
	}

	var err error
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0], err = YourLoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return err
	}
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	tlsListener := tls.NewListener(conn, config)
	return http.Serve(tlsListener, handler)
}

func YourLoadX509KeyPair(certFile, keyFile string) (cert tls.Certificate, err error) {
	certPEMBlock, err := os.ReadFile(certFile)
	if err != nil {
		return
	}
	certDERBlock, restPEMBlock := pem.Decode(certPEMBlock)
	if certDERBlock == nil {
		err = errors.New("failed to parse PEM data")
		return
	}
	certDERBlockCHain, _ := pem.Decode(restPEMBlock)
	if certDERBlockCHain == nil {
		cert.Certificate = [][]byte{certDERBlock.Bytes}
	} else {
		cert.Certificate = [][]byte{certDERBlock.Bytes, certDERBlockCHain.Bytes}
	}
	keyPEMBlock, err := os.ReadFile(keyFile)
	if err != nil {
		return
	}
	keyDERBlock, _ := pem.Decode(keyPEMBlock)
	if keyDERBlock == nil {
		err = errors.New("crypto/tls: failed to parse key PEM data")
		return
	}
	key, err := x509.ParsePKCS1PrivateKey(keyDERBlock.Bytes)
	if err != nil {
		err = errors.New("crypto/tls: failed to parse key")
		return
	}
	cert.PrivateKey = key
	x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	if err != nil {
		return
	}
	if x509Cert.PublicKeyAlgorithm != x509.RSA || x509Cert.PublicKey.(*rsa.PublicKey).N.Cmp(key.PublicKey.N) != 0 {
		err = errors.New("crypto/tls: private key does not match public key")
		return
	}
	return
}
