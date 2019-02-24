package github

import "time"

// Should probably increase this nonsense.
const bufferTimeout = 5 * time.Minute * -1

var token = ""

// SetToken will set our token.
func SetToken(githubToken string) {
	token = githubToken
}

var buffer bool

// SetBuffering will set whether github will buffer the results.
func SetBuffering(buff bool) {
	buffer = buff
}
