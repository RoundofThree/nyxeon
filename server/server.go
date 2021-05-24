package server

import "github.com/RoundofThree/nyxeon/config"

func Init() {
	c := config.GetConfig()
	r := NewRouter()
	r.Run(c.GetString("server.port"))
}
