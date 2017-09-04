// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// unnamed API: Application Controllers
//
// Command:
// $ goagen
// --design=github.com/JormungandrK/authorization-server/design
// --out=$(GOPATH)/src/github.com/JormungandrK/authorization-server
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// AuthUIController is the controller interface for the AuthUI actions.
type AuthUIController interface {
	goa.Muxer
	ConfirmAuthorization(*ConfirmAuthorizationAuthUIContext) error
	PromptAuthorization(*PromptAuthorizationAuthUIContext) error
}

// MountAuthUIController "mounts" a AuthUI resource controller on the given service.
func MountAuthUIController(service *goa.Service, ctrl AuthUIController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewConfirmAuthorizationAuthUIContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.ConfirmAuthorization(rctx)
	}
	service.Mux.Handle("POST", "/auth/confirm-authorization", ctrl.MuxHandler("confirmAuthorization", h, nil))
	service.LogInfo("mount", "ctrl", "AuthUI", "action", "ConfirmAuthorization", "route", "POST /auth/confirm-authorization")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewPromptAuthorizationAuthUIContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.PromptAuthorization(rctx)
	}
	service.Mux.Handle("GET", "/auth/authorize-client", ctrl.MuxHandler("promptAuthorization", h, nil))
	service.LogInfo("mount", "ctrl", "AuthUI", "action", "PromptAuthorization", "route", "GET /auth/authorize-client")
}

// Oauth2ProviderController is the controller interface for the Oauth2Provider actions.
type Oauth2ProviderController interface {
	goa.Muxer
	Authorize(*AuthorizeOauth2ProviderContext) error
	GetToken(*GetTokenOauth2ProviderContext) error
}

// MountOauth2ProviderController "mounts" a Oauth2Provider resource controller on the given service.
func MountOauth2ProviderController(service *goa.Service, ctrl Oauth2ProviderController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAuthorizeOauth2ProviderContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Authorize(rctx)
	}
	service.Mux.Handle("GET", "/oauth2/authorize", ctrl.MuxHandler("authorize", h, nil))
	service.LogInfo("mount", "ctrl", "Oauth2Provider", "action", "Authorize", "route", "GET /oauth2/authorize")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetTokenOauth2ProviderContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*TokenPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.GetToken(rctx)
	}
	h = handleSecurity("oauth2_client_basic_auth", h)
	service.Mux.Handle("POST", "/oauth2/token", ctrl.MuxHandler("get_token", h, unmarshalGetTokenOauth2ProviderPayload))
	service.LogInfo("mount", "ctrl", "Oauth2Provider", "action", "GetToken", "route", "POST /oauth2/token", "security", "oauth2_client_basic_auth")
}

// unmarshalGetTokenOauth2ProviderPayload unmarshals the request body into the context request data Payload field.
func unmarshalGetTokenOauth2ProviderPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &tokenPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// PublicController is the controller interface for the Public actions.
type PublicController interface {
	goa.Muxer
	goa.FileServer
}

// MountPublicController "mounts" a Public resource controller on the given service.
func MountPublicController(service *goa.Service, ctrl PublicController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/login", ctrl.MuxHandler("preflight", handlePublicOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/login", "public/login/login-form.html")
	h = handlePublicOrigin(h)
	service.Mux.Handle("GET", "/login", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Public", "files", "public/login/login-form.html", "route", "GET /login")
}

// handlePublicOrigin applies the CORS response headers corresponding to the origin.
func handlePublicOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
