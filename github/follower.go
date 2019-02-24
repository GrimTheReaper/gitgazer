package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// I _think_ I'm reading these instructions correctly and I need to limit data..
const (
	followerLimit = 5
	followerdepth = 3
	followerURI   = "https://api.github.com/users/%v/followers"
)

var followerBuffer = make(map[string]User)

// GetFollowers will return the followers. Non-recursive.
func GetFollowers(login string) (users []User, err error) {
	// Github is non case insensitive, why shouldn't we be?
	// login = strings.ToLower(login)

	if buffer {
		if follower, ok := followerBuffer[login]; ok {
			// if the follower not was loaded now minus the timeout
			if !follower.LoadedTimestamp.Before(time.Now().Add(bufferTimeout)) {
				fmt.Printf("%v was loaded from the buffer with time %v\n", login, follower.LoadedTimestamp)
				users = make([]User, len(follower.Followers))
				copy(users, follower.Followers)
				return
			}
			delete(followerBuffer, login)
		}
	}

	request, err := http.NewRequest("GET", fmt.Sprintf(followerURI, login), nil)
	if err != nil {
		return
	}
	if token != "" {
		request.Header.Add("Authorization", fmt.Sprintf("token %v", token))
	}

	// Add a timeout to this nonsense...
	http.DefaultClient.Timeout = 10 * time.Second
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	decoder.Decode(&users)

	if len(users) > followerLimit {
		users = users[:followerLimit]
	}

	if buffer {
		followerBuffer[login] = User{
			Followers:       users,
			LoadedTimestamp: time.Now(),
		}
	}

	return
}

// GetFollowerRecursive is written for this challenge.
func GetFollowerRecursive(login string) (user User, err error) {
	user.Login = login

	user.Followers, err = recurseGetFollowers(user, 0)

	return
}

func recurseGetFollowers(user User, depth int) (followers []User, err error) {
	if depth > followerdepth {
		return nil, err
	}

	followers, err = GetFollowers(user.Login)
	if err != nil {
		return
	}

	for ii := range followers {
		followers[ii].Followers, err = recurseGetFollowers(followers[ii], depth+1)
		if err != nil {
			return
		}
	}

	return
}

// Man if only we didnt have that pesky rate limiter we could use this lovely.
func recurseAsync(user *User, depth int) (err error) {
	if depth > followerdepth {
		return nil
	}

	user.Followers, err = GetFollowers(user.Login)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for i := 0; i < len(user.Followers); i++ {
		wg.Add(1)
		follower := &user.Followers[i]
		go func(user *User) {
			defer wg.Done()
			recurseAsync(user, depth)
		}(follower)
	}

	wg.Wait()
	return
}
