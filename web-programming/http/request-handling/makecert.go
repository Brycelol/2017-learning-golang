// The purpose of this go file is to generate a tls cert pair
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {
	// Generate a max int for the randomizer
	max := new(big.Int).Lsh(big.NewInt(1), 128)

	// rand.Int returns a uniform random value in [0, max).
	// rand.Reader is a global, shared instance of a cryptographically
	// strong pseudo-random generator.
	serialNumber, _ := rand.Int(rand.Reader, max)

	// Subject for the certificate
	subject := pkix.Name{
		Organization:       []string{"Test Organization"},
		OrganizationalUnit: []string{"Test Organizational Unit"},
		CommonName:         "Test Certificate",
	}

	// Template with which the certificate will be generated
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}

	// Generate a private key of size 2048 bits using the rand Reader
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	// Generate a certificate using the template defined above. Note that a pointer
	// to the template is passed twice. The 2nd parameter is the parent cert and indicates that
	// this cert will be self signed
	// The resulting certificate is DER encoded, hence the variable name
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)

	// Create file to write the cert to
	certOut, _ := os.Create("cert.pem")

	// Pem file encoded with the public certificate
	// We pass in the cert file we created, and a new pointer to a CERTIFICATE block
	// that takes our DER encoded certificate as its bytes
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	// Close off the cert file we created
	certOut.Close()

	// Create file to write the private key to
	keyOut, _ := os.Create("key.pem")

	// Pem file encoded with the private key
	// MarshalPKCS1PrivateKey takes a private key and converts it to a DER encoded form
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})

	// Close the key
	keyOut.Close()
}
