package main

import (
	"kilikili/config"
	v1 "kilikili/internal/controller/http/v1"
)

func main() {
	config.Init()
	r := v1.NewRouter()
	_ = r.Run(config.HttpPort)
}
