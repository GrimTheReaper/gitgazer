package network

import (
	"net/http"
	"regexp"

	"github.com/grimthereaper/gitgazer/github"
)

func getFollowersForUser(rw http.ResponseWriter, r *http.Request, p *regexp.Regexp) {
	username := getFirstCaptureGroup(r, p)

	user, err := github.GetFollowerRecursive(username)
	if err != nil {
		rw.Write([]byte("Failed to retrieve response from github"))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	serveFormatted(rw, user.Followers)
}
