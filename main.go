package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/RoundofThree/nyxeon/config"
	"github.com/RoundofThree/nyxeon/db"
	"github.com/RoundofThree/nyxeon/server"
)

func main() {
	environment := flag.String("m", "production", "Production or development")
	flag.Usage = func() {
		fmt.Println("Usage: server -m {development | production}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	db.InitRedis()
	server.Init()
	// defer disconnect db and redis
	defer db.Close()
	defer db.CloseCache()
}
