// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// unnamed API: Client
//
// Command:
// $ goagen
// --design=github.com/Microkubes/authorization-server/design
// --out=$(GOPATH)/src/github.com/Microkubes/authorization-server
// --version=v1.2.0-dirty

package client

import (
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
)

// Client is the  service client.
type Client struct {
	*goaclient.Client
	Oauth2ClientBasicAuthSigner goaclient.Signer
	OAuth2Signer                goaclient.Signer
	Encoder                     *goa.HTTPEncoder
	Decoder                     *goa.HTTPDecoder
}

// New instantiates the client.
func New(c goaclient.Doer) *Client {
	client := &Client{
		Client:  goaclient.New(c),
		Encoder: goa.NewHTTPEncoder(),
		Decoder: goa.NewHTTPDecoder(),
	}

	// Setup encoders and decoders
	client.Encoder.Register(goa.NewJSONEncoder, "application/json")
	client.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	client.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	client.Decoder.Register(goa.NewJSONDecoder, "application/json")
	client.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	client.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	client.Encoder.Register(goa.NewJSONEncoder, "*/*")
	client.Decoder.Register(goa.NewJSONDecoder, "*/*")

	return client
}

// SetOauth2ClientBasicAuthSigner sets the request signer for the oauth2_client_basic_auth security scheme.
func (c *Client) SetOauth2ClientBasicAuthSigner(signer goaclient.Signer) {
	c.Oauth2ClientBasicAuthSigner = signer
}

// SetOAuth2Signer sets the request signer for the OAuth2 security scheme.
func (c *Client) SetOAuth2Signer(signer goaclient.Signer) {
	c.OAuth2Signer = signer
}
