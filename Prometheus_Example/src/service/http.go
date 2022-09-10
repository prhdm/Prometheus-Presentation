package service

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type HttpServer struct {
	server *http.Server
	mux    *http.ServeMux
}

func (h *HttpServer) getMux() *http.ServeMux {
	return h.mux
}

func NewHttpServer() *HttpServer {
	mux := http.NewServeMux()
	addr := ":8080"
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	return &HttpServer{server: server, mux: mux}
}

func (h *HttpServer) Start() {
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)
	log.Info("Starting example service")
	server := h.server
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Monitor PrometheusMetricsServer Error: %v", err)
		}
	}()

	<-interrupt

	if err := server.Shutdown(context.Background()); err != nil {
		log.Error(err)
	}
}

func (h *HttpServer) Register(url string, handler func(http.ResponseWriter, *http.Request)) {
	h.mux.HandleFunc(url, handler)
}
