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

// NoSQL database

var db *mongo.Database

// Connect to the database.
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

// Return a pointer to the database.
func GetDB() *mongo.Database {
	return db
}

// Close the connection to the database,
func Close() error {
	return db.Client().Disconnect(context.Background())
}

// Cache

var redisClient redis.Conn

// Connect to Redis cache.
func InitRedis() {
	c := config.GetConfig()
	redis_url := c.GetString("redis.url")
	conn, err := redis.DialURL(redis_url)
	if err != nil {
		panic(err)
	}
	redisClient = conn
}

// Return the connection to the Redis cache.
func GetCache() redis.Conn {
	return redisClient
}

// Close the connection to the Redis cache.
func CloseCache() error {
	return redisClient.Close()
}
