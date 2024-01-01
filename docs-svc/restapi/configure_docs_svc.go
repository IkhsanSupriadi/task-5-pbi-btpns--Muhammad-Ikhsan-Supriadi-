// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/yusrilsabir22/orderfaz/docs-svc/restapi/operations"
	"github.com/yusrilsabir22/orderfaz/docs-svc/restapi/operations/auth"
	"github.com/yusrilsabir22/orderfaz/docs-svc/restapi/operations/logistic"
)

//go:generate swagger generate server --target ../../docs-svc --name DocsSvc --spec ../docs/swagger.json --principal interface{}

func configureFlags(api *operations.DocsSvcAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.DocsSvcAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.AuthMixin0Handler == nil {
		api.AuthMixin0Handler = auth.Mixin0HandlerFunc(func(params auth.Mixin0Params) middleware.Responder {
			return middleware.NotImplemented("operation auth.Mixin0 has not yet been implemented")
		})
	}
	if api.AuthMixin1Handler == nil {
		api.AuthMixin1Handler = auth.Mixin1HandlerFunc(func(params auth.Mixin1Params) middleware.Responder {
			return middleware.NotImplemented("operation auth.Mixin1 has not yet been implemented")
		})
	}
	if api.LogisticMixin3Handler == nil {
		api.LogisticMixin3Handler = logistic.Mixin3HandlerFunc(func(params logistic.Mixin3Params) middleware.Responder {
			return middleware.NotImplemented("operation logistic.Mixin3 has not yet been implemented")
		})
	}
	if api.LogisticMixin4Handler == nil {
		api.LogisticMixin4Handler = logistic.Mixin4HandlerFunc(func(params logistic.Mixin4Params) middleware.Responder {
			return middleware.NotImplemented("operation logistic.Mixin4 has not yet been implemented")
		})
	}
	if api.AuthPostAuthLoginHandler == nil {
		api.AuthPostAuthLoginHandler = auth.PostAuthLoginHandlerFunc(func(params auth.PostAuthLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostAuthLogin has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
