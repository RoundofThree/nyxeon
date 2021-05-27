package models

import "github.com/RoundofThree/nyxeon/db"

type SessionManager struct {
}

// type Session string

func (m *SessionManager) FetchSession(sessionToken string) (string, error) {
	// lookup sessionToken in Redis
	userID, err := db.GetCache().Do("GET", sessionToken)
	return userID.(string), err
}
