// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/spidernet-io/spiderpool/api/v1beta/spiderpool-controller/server/restapi/controller"
	runtimeops "github.com/spidernet-io/spiderpool/api/v1beta/spiderpool-controller/server/restapi/runtime"
)

// NewSpiderpoolControllerAPIAPI creates a new SpiderpoolControllerAPI instance
func NewSpiderpoolControllerAPIAPI(spec *loads.Document) *SpiderpoolControllerAPIAPI {
	return &SpiderpoolControllerAPIAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		ControllerGetIpamStatusHandler: controller.GetIpamStatusHandlerFunc(func(params controller.GetIpamStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation controller.GetIpamStatus has not yet been implemented")
		}),
		RuntimeGetRuntimeLivenessHandler: runtimeops.GetRuntimeLivenessHandlerFunc(func(params runtimeops.GetRuntimeLivenessParams) middleware.Responder {
			return middleware.NotImplemented("operation runtime.GetRuntimeLiveness has not yet been implemented")
		}),
		RuntimeGetRuntimeReadinessHandler: runtimeops.GetRuntimeReadinessHandlerFunc(func(params runtimeops.GetRuntimeReadinessParams) middleware.Responder {
			return middleware.NotImplemented("operation runtime.GetRuntimeReadiness has not yet been implemented")
		}),
		RuntimeGetRuntimeStartupHandler: runtimeops.GetRuntimeStartupHandlerFunc(func(params runtimeops.GetRuntimeStartupParams) middleware.Responder {
			return middleware.NotImplemented("operation runtime.GetRuntimeStartup has not yet been implemented")
		}),
		ControllerPostIpamGcIpsHandler: controller.PostIpamGcIpsHandlerFunc(func(params controller.PostIpamGcIpsParams) middleware.Responder {
			return middleware.NotImplemented("operation controller.PostIpamGcIps has not yet been implemented")
		}),
		ControllerPutIpamIPHandler: controller.PutIpamIPHandlerFunc(func(params controller.PutIpamIPParams) middleware.Responder {
			return middleware.NotImplemented("operation controller.PutIpamIP has not yet been implemented")
		}),
	}
}

/*SpiderpoolControllerAPIAPI Spiderpool Controller */
type SpiderpoolControllerAPIAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// ControllerGetIpamStatusHandler sets the operation handler for the get ipam status operation
	ControllerGetIpamStatusHandler controller.GetIpamStatusHandler
	// RuntimeGetRuntimeLivenessHandler sets the operation handler for the get runtime liveness operation
	RuntimeGetRuntimeLivenessHandler runtimeops.GetRuntimeLivenessHandler
	// RuntimeGetRuntimeReadinessHandler sets the operation handler for the get runtime readiness operation
	RuntimeGetRuntimeReadinessHandler runtimeops.GetRuntimeReadinessHandler
	// RuntimeGetRuntimeStartupHandler sets the operation handler for the get runtime startup operation
	RuntimeGetRuntimeStartupHandler runtimeops.GetRuntimeStartupHandler
	// ControllerPostIpamGcIpsHandler sets the operation handler for the post ipam gc ips operation
	ControllerPostIpamGcIpsHandler controller.PostIpamGcIpsHandler
	// ControllerPutIpamIPHandler sets the operation handler for the put ipam IP operation
	ControllerPutIpamIPHandler controller.PutIpamIPHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *SpiderpoolControllerAPIAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *SpiderpoolControllerAPIAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *SpiderpoolControllerAPIAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *SpiderpoolControllerAPIAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *SpiderpoolControllerAPIAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *SpiderpoolControllerAPIAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *SpiderpoolControllerAPIAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *SpiderpoolControllerAPIAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *SpiderpoolControllerAPIAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the SpiderpoolControllerAPIAPI
func (o *SpiderpoolControllerAPIAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.ControllerGetIpamStatusHandler == nil {
		unregistered = append(unregistered, "controller.GetIpamStatusHandler")
	}
	if o.RuntimeGetRuntimeLivenessHandler == nil {
		unregistered = append(unregistered, "runtime.GetRuntimeLivenessHandler")
	}
	if o.RuntimeGetRuntimeReadinessHandler == nil {
		unregistered = append(unregistered, "runtime.GetRuntimeReadinessHandler")
	}
	if o.RuntimeGetRuntimeStartupHandler == nil {
		unregistered = append(unregistered, "runtime.GetRuntimeStartupHandler")
	}
	if o.ControllerPostIpamGcIpsHandler == nil {
		unregistered = append(unregistered, "controller.PostIpamGcIpsHandler")
	}
	if o.ControllerPutIpamIPHandler == nil {
		unregistered = append(unregistered, "controller.PutIpamIPHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *SpiderpoolControllerAPIAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *SpiderpoolControllerAPIAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *SpiderpoolControllerAPIAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *SpiderpoolControllerAPIAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *SpiderpoolControllerAPIAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *SpiderpoolControllerAPIAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the spiderpool controller API API
func (o *SpiderpoolControllerAPIAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *SpiderpoolControllerAPIAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/ipam/status"] = controller.NewGetIpamStatus(o.context, o.ControllerGetIpamStatusHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/runtime/liveness"] = runtimeops.NewGetRuntimeLiveness(o.context, o.RuntimeGetRuntimeLivenessHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/runtime/readiness"] = runtimeops.NewGetRuntimeReadiness(o.context, o.RuntimeGetRuntimeReadinessHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/runtime/startup"] = runtimeops.NewGetRuntimeStartup(o.context, o.RuntimeGetRuntimeStartupHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/ipam/gc_ips"] = controller.NewPostIpamGcIps(o.context, o.ControllerPostIpamGcIpsHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/ipam/ip"] = controller.NewPutIpamIP(o.context, o.ControllerPutIpamIPHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *SpiderpoolControllerAPIAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *SpiderpoolControllerAPIAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *SpiderpoolControllerAPIAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *SpiderpoolControllerAPIAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *SpiderpoolControllerAPIAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
