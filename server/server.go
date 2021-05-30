package server

import "github.com/RoundofThree/nyxeon/config"

func Init() {
	c := config.GetConfig()
	r := NewRouter()
	if c.GetString("server") == "" {
		r.Run()
	} else {
		r.Run(c.GetString("server"))
	}
}
