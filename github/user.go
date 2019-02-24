package github

import "time"

// User represents a user with limited data.
type User struct {
	Login           string       `json:"login"`
	ID              int          `json:"id"`
	Followers       []User       `json:"followers,omitempty"`
	Repositories    []Repository `json:"repositories,omitempty"`
	LoadedTimestamp time.Time    `json:"-"` // having issues copying it so lets just not include it.
}
