// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// unnamed API: Application Media Types
//
// Command:
// $ goagen
// --design=github.com/Microkubes/authorization-server/design
// --out=$(GOPATH)/src/github.com/Microkubes/authorization-server
// --version=v1.2.0-dirty

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// OAuth2 error response, see https://tools.ietf.org/html/rfc6749#section-5.2 (default view)
//
// Identifier: application/vnd.goa.example.oauth2.error+json; view=default
type OAuth2ErrorMedia struct {
	// Error returned by authorization server
	Error string `form:"error" json:"error" xml:"error"`
	// Human readable ASCII text providing additional information
	ErrorDescription *string `form:"error_description,omitempty" json:"error_description,omitempty" xml:"error_description,omitempty"`
	// A URI identifying a human-readable web page with information about the error
	ErrorURI *string `form:"error_uri,omitempty" json:"error_uri,omitempty" xml:"error_uri,omitempty"`
}

// Validate validates the OAuth2ErrorMedia media type instance.
func (mt *OAuth2ErrorMedia) Validate() (err error) {
	if mt.Error == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "error"))
	}
	if !(mt.Error == "invalid_request" || mt.Error == "invalid_client" || mt.Error == "invalid_grant" || mt.Error == "unauthorized_client" || mt.Error == "unsupported_grant_type") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.error`, mt.Error, []interface{}{"invalid_request", "invalid_client", "invalid_grant", "unauthorized_client", "unsupported_grant_type"}))
	}
	return
}

// DecodeOAuth2ErrorMedia decodes the OAuth2ErrorMedia instance encoded in resp body.
func (c *Client) DecodeOAuth2ErrorMedia(resp *http.Response) (*OAuth2ErrorMedia, error) {
	var decoded OAuth2ErrorMedia
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// OAuth2 access token request successful response, see https://tools.ietf.org/html/rfc6749#section-5.1 (default view)
//
// Identifier: application/vnd.goa.example.oauth2.token+json; view=default
type TokenMedia struct {
	// The access token issued by the authorization server
	AccessToken string `form:"access_token" json:"access_token" xml:"access_token"`
	// The lifetime in seconds of the access token
	ExpiresIn *int `form:"expires_in,omitempty" json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	// The refresh token
	RefreshToken *string `form:"refresh_token,omitempty" json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	// The scope of the access token
	Scope *string `form:"scope,omitempty" json:"scope,omitempty" xml:"scope,omitempty"`
	// The type of the token issued, e.g. "bearer" or "mac"
	TokenType string `form:"token_type" json:"token_type" xml:"token_type"`
}

// Validate validates the TokenMedia media type instance.
func (mt *TokenMedia) Validate() (err error) {
	if mt.AccessToken == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "access_token"))
	}
	if mt.TokenType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token_type"))
	}
	return
}

// DecodeTokenMedia decodes the TokenMedia instance encoded in resp body.
func (c *Client) DecodeTokenMedia(resp *http.Response) (*TokenMedia, error) {
	var decoded TokenMedia
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
