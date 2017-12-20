package main

import (
	"github.com/brycelol/learning-golang/chitchat/app/data"
	"net/http"
)

// Authenticates a user and returns a cookie in the response header which
// has already logged in
func authenticate(w http.ResponseWriter, r *http.Request) {

	// Parses the request's raw form and populates r.Form
	r.ParseForm()

	// Querying data for user via the email form value
	user, _ := data.UserByEmail(r.PostFormValue("email"))

	// Grab the user password, then check if the posted password is matching
	// after encryption of it
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		// If so, we create a session for the user
		session := user.CreateSession()

		// Then a corresponding cookie
		cookie := http.Cookie{
			Name:  "chitchat-session-id",
			Value: session.Uuid,
			// HttpOnly means only http requests can access the cookie (not JS etc)
			HttpOnly: true,
		}

		// We then set the cookie and redirect to the index
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		// Else redirect back to the login page
		http.Redirect(w, r, "/login", 302)
	}
}
