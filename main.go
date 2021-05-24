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
	environment := flag.String("m", "development", "Production or development")
	flag.Usage = func() {
		fmt.Println("Usage: server -m {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()
}