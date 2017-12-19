package main

import (
	"errors"
	"github.com/gclendinning02/learning-go/web-programming/chapter2/chitchat/data"
	"net/http"
)

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {

	// Grab the cookie
	cookie, err := r.Cookie("_cookie")

	if err == nil {

		// Create a session struct with the cookie as the uuid
		sess = data.Session{Uuid: cookie.Value}

		// Check the session. If it's not ok then set a new error
		if ok, _ := sess.Check(); !ok {
			err = errors.New("invalid session")
		}
	}

	return
}
