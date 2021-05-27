package db

import (
	"context"
	"log"
	"time"

	"github.com/RoundofThree/nyxeon/config"
	"github.com/gomodule/redigo/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database

var db *mongo.Database

func Init() {
	c := config.GetConfig()

	host := c.GetString("mongo.host")
	dbName := c.GetString("mongo.db_name")
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(host))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err = mongoClient.Connect(ctx)
	if err != nil {
		log.Fatal("Could not connect to the MongoDB client")
	}
	defer cancel()
	db = mongoClient.Database(dbName)
}

func GetDB() *mongo.Database {
	return db
}

// Cache

var redisClient redis.Conn

func InitRedis() {
	c := config.GetConfig()
	redis_url := c.GetString("redis.url")
	conn, err := redis.DialURL(redis_url)
	if err != nil {
		panic(err)
	}
	redisClient = conn
}

func GetCache() redis.Conn {
	return redisClient
}
