package coinbase

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
)

// ServiceOAuthAuthentication Struct implements the Authentication interface
// and takes care of authenticating OAuth RPC requests on behalf of the service
// (i.e GetTokens())
type ServiceOAuthAuthentication struct {
	BaseUrl string
	Client  http.Client
}

// ServiceOAuth instantiates ServiceOAuthAuthentication with the coinbase certificate file
func ServiceOAuth(certFilePath string) (*ServiceOAuthAuthentication, error) {
	// First we read the cert
	certs := x509.NewCertPool()
	pemData, err := ioutil.ReadFile(certFilePath)
	if err != nil {
		return nil, err
	}
	certs.AppendCertsFromPEM(pemData)
	mTLSConfig := &tls.Config{
		RootCAs: certs, //Add the cert as a TLS config
	}
	a := ServiceOAuthAuthentication{
		BaseUrl: "https://coinbase.com/",
		Client: http.Client{
			Transport: &http.Transport{
				Dial:            dialTimeout,
				TLSClientConfig: mTLSConfig,
			},
		},
	}
	return &a, nil
}

// Service OAuth authentication requires no additional headers to be sent. The
// Coinbase Public Certificate is set as a TLS config in the http.Client
func (a ServiceOAuthAuthentication) Authenticate(req *http.Request, endpoint string, params []byte) error {
	return nil // No additional headers needed for service OAuth requests
}

func (a ServiceOAuthAuthentication) GetBaseUrl() string {
	return a.BaseUrl
}

func (a ServiceOAuthAuthentication) GetClient() *http.Client {
	return &a.Client
}
