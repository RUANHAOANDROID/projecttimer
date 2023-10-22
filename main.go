package main

import (
	"projecttimer/api"
	"projecttimer/config"
	"projecttimer/db"
)

func main() {
	conf, err := config.Load("config.yml")
	if err != nil {
		panic("config.yml read error")
	}
	db.Create(conf)
	api.Start(conf)
}
