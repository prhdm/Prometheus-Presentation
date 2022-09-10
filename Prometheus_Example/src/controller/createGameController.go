package controller

import (
	"PrometheusExample/src/models"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type CreateGame struct {
	Username string `json:"username"`
	Password string `json:"password"`
	GameID   string `json:"gameID"`
}

func createGameHandler(w http.ResponseWriter, r *http.Request) {
	log.Infoln("received a request to create a game")
	NumberOfRequestsCounter.WithLabelValues("create_game").Inc()
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
		NumberOfRequestsGauge.WithLabelValues("create_game").Dec()
		w.WriteHeader(404)
		_, err = w.Write([]byte("User not found"))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if DB.Get("users", payload.Username) != payload.Password {
		NumberOfUnAuthorizedRequestsCounter.WithLabelValues("create_game").Inc()
		NumberOfRequestsGauge.WithLabelValues("create_game").Dec()
		w.WriteHeader(401)
		_, err = w.Write([]byte("Unauthorized"))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if DB.Get("games", payload.GameID) != nil {
		NumberOfRequestsGauge.WithLabelValues("create_game").Dec()
		w.WriteHeader(409)
		_, err = w.Write([]byte("Game already exists"))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	game := models.NewGame(payload.GameID)
	game.AddPlayer(payload.Username)
	NumberOfGamePlayersGauge.WithLabelValues(payload.GameID).Set(float64(len(game.Players)))
	DB.Set("games", payload.GameID, game)
	NumberOfRequestsGauge.WithLabelValues("create_game").Inc()
	NumberOfRequestsPerUserSummary.WithLabelValues("create_game", payload.Username).Observe(1)
	w.WriteHeader(200)
	_, err = w.Write([]byte("Game created"))
	if err != nil {
		log.Fatal(err)
	}

}
