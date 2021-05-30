package models

import (
	"strconv"

	"github.com/RoundofThree/nyxeon/db"
)

type SessionManager struct{}

// Get the user id from the session token.
func (m *SessionManager) FetchSession(sessionToken string) (string, error) {
	cache := db.GetCache()
	userID, err := cache.Do("GET", sessionToken)
	return string(userID.([]byte)), err
}

// Add session token mapping to user id in cache. If the session token
// already exists in cache, its expiry time resets to 24*60*60.
func (m *SessionManager) UpdateSession(sessionToken, userID string) error {
	cache := db.GetCache()
	_, err := cache.Do("SETEX", sessionToken, strconv.Itoa(24*60*60), userID)
	return err
}

// Delete the session token in cache.
func (m *SessionManager) DeleteSession(sessionToken string) error {
	cache := db.GetCache()
	_, err := cache.Do("DEL", sessionToken)
	return err
}
