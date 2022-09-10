package controller

import (
	"PrometheusExample/src/models"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func joinGameHandler(w http.ResponseWriter, r *http.Request) {
	log.Infoln("received a request to join a game")
	NumberOfRequestsCounter.WithLabelValues("join_game").Inc()
	payload := CreateGame{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Fatal(err)
	}

	if DB.Get("users", payload.Username) == nil {
		NumberOfRequestsGauge.WithLabelValues("join_game").Dec()
		w.WriteHeader(404)
		_, err = w.Write([]byte("User not found"))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if DB.Get("users", payload.Username) != payload.Password {
		NumberOfUnAuthorizedRequestsCounter.WithLabelValues("join_game").Inc()
		NumberOfRequestsGauge.WithLabelValues("join_game").Dec()
		w.WriteHeader(401)
		_, err = w.Write([]byte("Unauthorized"))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	game := DB.Get("games", payload.GameID)
	if game != nil {
		NumberOfRequestsGauge.WithLabelValues("join_game").Dec()
		w.WriteHeader(409)
		_, err = w.Write([]byte("Game already exists"))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	g := game.(models.Game)
	g.AddPlayer(payload.Username)
	NumberOfGamePlayersGauge.WithLabelValues(payload.GameID).Set(float64(len(g.Players)))
	NumberOfRequestsPerUserSummary.WithLabelValues("join_game", payload.Username).Observe(float64(1))
	NumberOfRequestsGauge.WithLabelValues("join_game").Inc()
	w.WriteHeader(200)
	_, err = w.Write([]byte("Game created"))
	if err != nil {
		log.Fatal(err)
	}

}
