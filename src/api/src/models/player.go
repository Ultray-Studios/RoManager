package models

import "github.com/gocql/gocql"

// Player model for database and API
type Player struct {
	ID    gocql.UUID `json:"id"`
	Name  string     `json:"name"`
	Score int        `json:"score"`
}
