package controller

import "github.com/prometheus/client_golang/prometheus"

var (
	NumberOfRequestsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "example",
			Name:      "requests",
			Help:      "count requests",
		}, []string{"topic"},
	)
	NumberOfRequestsGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "example",
			Name:      "requests",
			Help:      "count requests",
		}, []string{"topic"},
	)
	NumberOfUnAuthorizedRequestsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "example",
			Name:      "unauthorized_requests",
			Help:      "count unauthorized requests",
		}, []string{"topic"},
	)
	NumberOfGamesCreatedCounter = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "example",
			Name:      "games_created",
			Help:      "count games created",
		}, []string{"topic"},
	)
	NumberOfGamePlayersGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "example",
			Name:      "game_players",
			Help:      "count game players",
		}, []string{"topic"},
	)
)
