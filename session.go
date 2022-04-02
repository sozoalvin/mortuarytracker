package main

import (
	"fmt"
	"net/http"
	"time"
	"wecare/dbMemorySessions"
)

func alreadyLoggedIn(w http.ResponseWriter, req *http.Request) (string, bool) { //checks if a user is logged or by returning equivalent bool value

	c, err := req.Cookie("session")
	if err != nil {
		fmt.Println("unable to get session cookie information")
		return "", false
	}
	username, ok := dbMemorySessions.RetrieverUsername(c.Value)

	if ok {
		fmt.Println("username can be found")
		return username, true
	} else {
		return "", false
	}

}

func cleanSessions() {
	fmt.Println("(before) db session cleaned") // for demonstration purposes
	showSessions()                             // for demonstration purposes
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("(after) db session cleaned") // for demonstration purposes
	showSessions()                            // for demonstration purposes
}

func showSessions() {
	fmt.Println("********")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}
