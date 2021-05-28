package main

/*
import (
	"fmt"

	"github.com/RoundofThree/nyxeon/config"
	"github.com/RoundofThree/nyxeon/db"
)
*/

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
	db.InitRedis()
	fmt.Println("Listening...")
	server.Init()
	fmt.Println("Closing connection...")
	// defer disconnect db and redis
	defer db.Close()
	defer db.CloseCache()
}

/*
func main() {
	TestDBConnection()
}

func TestDBConnection() {
	config.Init("development")
	db.Init()
	database := db.GetDB()
	if database == nil {
		fmt.Println("Database is nil")
	} else {
		fmt.Println("Database is ", database)
		fmt.Println("Collection is ", database.Collection("quests"))
	}

}
*/
