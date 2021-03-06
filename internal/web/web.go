package web

import (
	"context"
	"errors"
	"net"
	"net/http"
	"sync"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"goauthentik.io/internal/config"
)

type WebServer struct {
	Bind    string
	BindTLS bool

	LegacyProxy bool

	stop chan struct{} // channel for waiting shutdown

	m   *mux.Router
	lh  *mux.Router
	log *log.Entry
}

func NewWebServer() *WebServer {
	mainHandler := mux.NewRouter()
	if config.G.ErrorReporting.Enabled {
		mainHandler.Use(recoveryMiddleware())
	}
	mainHandler.Use(handlers.ProxyHeaders)
	mainHandler.Use(handlers.CompressHandler)
	logginRouter := mainHandler.NewRoute().Subrouter()
	logginRouter.Use(loggingMiddleware)

	ws := &WebServer{
		LegacyProxy: true,

		m:   mainHandler,
		lh:  logginRouter,
		log: log.WithField("logger", "authentik.g.web"),
	}
	ws.configureStatic()
	ws.configureProxy()
	return ws
}

func (ws *WebServer) Run() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		ws.listenPlain()
	}()
	go func() {
		defer wg.Done()
		ws.listenTLS()
	}()
	wg.Done()
}

func (ws *WebServer) listenPlain() {
	ln, err := net.Listen("tcp", config.G.Web.Listen)
	if err != nil {
		ws.log.WithError(err).Fatalf("failed to listen")
	}
	ws.log.WithField("addr", config.G.Web.Listen).Info("Running")

	ws.serve(ln)

	ws.log.WithField("addr", config.G.Web.Listen).Info("Running")
	http.ListenAndServe(config.G.Web.Listen, ws.m)
}

func (ws *WebServer) serve(listener net.Listener) {
	srv := &http.Server{
		Handler: ws.m,
	}

	// See https://golang.org/pkg/net/http/#Server.Shutdown
	idleConnsClosed := make(chan struct{})
	go func() {
		<-ws.stop // wait notification for stopping server

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			ws.log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	err := srv.Serve(listener)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		ws.log.Errorf("ERROR: http.Serve() - %s", err)
	}
	<-idleConnsClosed
}
