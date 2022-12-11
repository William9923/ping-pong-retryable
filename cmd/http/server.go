package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

var currUser int = 0

type pongServer struct {
	config serverConfig
}

func (s *pongServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if currUser >= s.config.MaxSimultaneousUser {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("server overloaded"))
		return
	}

	currUser++
	rand := rand.New(rand.NewSource(time.Now().UnixMilli()))
	duration := time.Duration(rand.Intn(s.config.ReqDurationMaxInSecond-s.config.ReqDurationMinInSecond) + s.config.ReqDurationMaxInSecond)
	time.Sleep(duration * time.Second)
	currUser--

	w.Write([]byte("pong"))
	w.WriteHeader(http.StatusOK)
}

func serve(config serverConfig) {

	http.Handle("/ping", &pongServer{config})
	log.Fatal(http.ListenAndServe(":8080", nil))

}
