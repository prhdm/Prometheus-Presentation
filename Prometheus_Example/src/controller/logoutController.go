package controller

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	log.Infoln("received a request to logout")
	NumberOfRequestsCounter.WithLabelValues("logOut").Inc()
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
		NumberOfRequestsGauge.WithLabelValues("logOut").Dec()
		w.WriteHeader(404)
		_, err = w.Write([]byte("User not found"))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if DB.Get("users", payload.Username) != payload.Password {
		NumberOfUnAuthorizedRequestsCounter.WithLabelValues("logout").Inc()
		NumberOfRequestsGauge.WithLabelValues("logout").Dec()
		w.WriteHeader(401)
		_, err = w.Write([]byte("Unauthorized"))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	NumberOfRequestsGauge.WithLabelValues("logout").Inc()
	NumberOfRequestsPerUserSummary.WithLabelValues("logout", payload.Username).Observe(1)
	DB.Delete("users", payload.Username)
	w.WriteHeader(200)
	_, err = w.Write([]byte("User logged out"))
	if err != nil {
		log.Fatal(err)
	}

}
