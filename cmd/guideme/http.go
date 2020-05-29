package main

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
	organizationsvr "guide.me/gen/http/organization/server"
	stepsvr "guide.me/gen/http/step/server"
	swaggersvr "guide.me/gen/http/swagger/server"
	walkthroughsvr "guide.me/gen/http/walkthrough/server"
	organization "guide.me/gen/organization"
	step "guide.me/gen/step"
	walkthrough "guide.me/gen/walkthrough"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, organizationEndpoints *organization.Endpoints, walkthroughEndpoints *walkthrough.Endpoints, stepEndpoints *step.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		organizationServer *organizationsvr.Server
		walkthroughServer  *walkthroughsvr.Server
		stepServer         *stepsvr.Server
		swaggerServer      *swaggersvr.Server
	)
	{
		eh := errorHandler(logger)
		organizationServer = organizationsvr.New(organizationEndpoints, mux, dec, enc, eh, nil)
		walkthroughServer = walkthroughsvr.New(walkthroughEndpoints, mux, dec, enc, eh, nil)
		stepServer = stepsvr.New(stepEndpoints, mux, dec, enc, eh, nil)
		swaggerServer = swaggersvr.New(nil, mux, dec, enc, eh, nil)
		if debug {
			servers := goahttp.Servers{
				organizationServer,
				walkthroughServer,
				stepServer,
				swaggerServer,
			}
			servers.Use(httpmdlwr.Debug(mux, os.Stdout))
		}
	}
	// Configure the mux.
	organizationsvr.Mount(mux, organizationServer)
	walkthroughsvr.Mount(mux, walkthroughServer)
	stepsvr.Mount(mux, stepServer)
	swaggersvr.Mount(mux)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(adapter)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler}
	for _, m := range organizationServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range walkthroughServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range stepServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}
	for _, m := range swaggerServer.Mounts {
		logger.Printf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Printf("HTTP server listening on %q", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Printf("shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger *log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Printf("[%s] ERROR: %s", id, err.Error())
	}
}
