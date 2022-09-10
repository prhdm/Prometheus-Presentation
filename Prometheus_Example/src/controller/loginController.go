package controller

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Infoln("received a request to login")
	NumberOfRequestsCounter.WithLabelValues("login").Inc()
	payload := Login{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Fatal(err)
	}

	if DB.Get("users", payload.Username) == nil {
		NumberOfRequestsGauge.WithLabelValues("login").Dec()
		w.WriteHeader(404)
		_, err = w.Write([]byte("User not found"))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if DB.Get("users", payload.Username) != payload.Password {
		NumberOfUnAuthorizedRequestsCounter.WithLabelValues("login").Inc()
		NumberOfRequestsGauge.WithLabelValues("login").Dec()
		w.WriteHeader(401)
		_, err = w.Write([]byte("Unauthorized"))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	NumberOfRequestsGauge.WithLabelValues("login").Inc()
	NumberOfRequestsPerUserSummary.WithLabelValues("login", payload.Username).Observe(1)
	w.WriteHeader(200)
	_, err = w.Write([]byte("User logged in"))
	if err != nil {
		log.Fatal(err)
	}

}
