package service

import (
	configs "PrometheusExample/config"
	"context"
	"errors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type PrometheusMetricsServer struct {
	Config configs.Monitor
}

func NewPrometheusMetricsServer() *PrometheusMetricsServer {
	config, err := configs.GetConfigFromYaml()
	if err != nil {
		log.Fatal(err)
	}
	return &PrometheusMetricsServer{Config: config.Monitor}
}

func (s PrometheusMetricsServer) Register(collector prometheus.Collector) {
	prometheus.MustRegister(collector)
}

func (s PrometheusMetricsServer) Start() {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	addr := fmt.Sprintf(":%s", s.Config.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)
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
