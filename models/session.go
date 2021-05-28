package models

import (
	"fmt"
	"strconv"

	"github.com/RoundofThree/nyxeon/db"
)

type SessionManager struct {
}

// var cache redis.Conn = db.GetCache()

// type Session string -> email

func (m *SessionManager) FetchSession(sessionToken string) (string, error) {
	// lookup sessionToken in Redis
	cache := db.GetCache()
	userID, err := cache.Do("GET", sessionToken)
	fmt.Println("Got userID from cache", userID)
	return string(userID.([]byte)), err
}

func (m *SessionManager) UpdateSession(sessionToken, userID string) error {
	cache := db.GetCache()
	fmt.Println("Setting sessiontoken in cache", sessionToken)
	_, err := cache.Do("SETEX", sessionToken, strconv.Itoa(24*60*60), userID)
	fmt.Println("Is it stored?")
	test, _ := cache.Do("GET", sessionToken)
	fmt.Println(test)
	return err
}
