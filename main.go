package main

import (
	"BaseWebProject/config"
	"BaseWebProject/env"
	"BaseWebProject/routes"
)

func main() {
	config.Init()
	r := routes.NewRouter()
	_ = r.Run(env.HttpPort)
}


