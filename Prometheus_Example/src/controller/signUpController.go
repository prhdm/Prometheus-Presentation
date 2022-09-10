package controller

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {
	log.Infoln("Signup")
	NumberOfRequestsCounter.WithLabelValues("signup").Inc()
	payload := Login{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Fatal(err)
	}

	if DB.Get("users", payload.Username) != nil {
		NumberOfRequestsGauge.WithLabelValues("signup").Dec()
		w.WriteHeader(409)
		_, err = w.Write([]byte("Username already exists"))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	DB.Set("users", payload.Username, payload.Password)
	NumberOfRequestsGauge.WithLabelValues("signup").Inc()
	//NumberOfRequestsPerUserSummary.WithLabelValues("signup", payload.Username).Observe(1)
	w.WriteHeader(200)
	_, err = w.Write([]byte("User created"))
	if err != nil {
		log.Fatal(err)
	}
}
