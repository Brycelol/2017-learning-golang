package main

import (
	"errors"
	"fmt"
	"github.com/brycelol/learning-golang/chitchat/data"
	"html/template"
	"net/http"
)

// generateHTML can be used to generate templates based upon supplied
// file names and a set of data
//
// Note that data is an empty interface type - meaning it can take any type.
// Go gets around being statically typed by accepting different types using interfaces
// Interfaces in Go are constructs that are a set fo methods and are also TYPES.
// An empty interface is then an empty set, meaning ANY TYPE can be an empty interface.
// This allows us to pass any type into this function as data
//
// This function is a variadic function, meaning it can take 0..n parameters in its
// last variadic parameter
//
// In variadic functions, the variadic parameter must always be last
func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {

	// Declaring a slice of files
	var files []string

	// Iterate over the file names and append their formatted template html
	// file to the slice of files
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	// Here we are passing a slice of templates from the passed in file names
	//
	// The ... after passing in the slice is known as a variadic parameter
	// This means you can pass any number of args to the function
	// It might sound counter intuitive but we add the ... to the slice as
	// it will pass the slice contents as "a", "b", "c" rather than
	// ["a","b","c"]
	//
	// template.Must is a helper wrapper which wraps creating the templates
	// It basically means these MUST template or it panics (avoids us error
	// handling ourselves).
	templates := template.Must(template.ParseFiles(files...))

	// This executes the templating of the html
	// "layout" maps to the top level layout template - although this could possibly be passed
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {

	// Grab the cookie
	cookie, err := r.Cookie("chitchat-session-id")

	// Expect an error if the cookie doesn't exist
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
