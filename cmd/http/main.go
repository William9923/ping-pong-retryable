package main

import "log"

const ()

type serverConfig struct {
	MaxSimultaneousUser    int
	ReqDurationMinInSecond int
	ReqDurationMaxInSecond int
}

func main() {

	config := serverConfig{
		MaxSimultaneousUser:    5,
		ReqDurationMinInSecond: 1,
		ReqDurationMaxInSecond: 5,
	}

	log.Println("starting http server...")
	serve(config)
}
