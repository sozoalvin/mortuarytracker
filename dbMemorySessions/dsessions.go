package dbMemorySessions

import (
	"errors"
)

var DBsessions map[string]string

// func AddNewSessions(w http.ResponseWriter, req *http.Request, username string) {
func AddNewSessions(cValue, username string) {
	// initializes DBsessions
	if DBsessions == nil {
		DBsessions = map[string]string{}
	}

	_, ok := DBsessions[cValue]
	if !ok {
		// since key doesn't exist we'll write add it into the map
		DBsessions[cValue] = username
	}
}

// func RetrieverUsername(w http.ResponseWriter, req *http.Request, username string) string {
// 	c, err := req.Cookie("session")
// 	if err != nil {
// 		fmt.Println("unable to get session cookie information")
// 	}
// 	username, ok := DBsessions[c.Value]
// 	if ok {
// 		return username
// 	} else {
// 		return "username not found in system"
// 	}
// }

func RetrieverUsername(cvalue string) (string, bool) {
	// c, err := req.Cookie("session")
	// if err != nil {
	// 	fmt.Println("unable to get session cookie information")
	// }
	username, ok := DBsessions[cvalue]
	if ok {
		return username, true
	} else {
		return "username not found in system", false
	}
}

func DeleteSessions(cvalue string) error {
	_, ok := DBsessions[cvalue]

	if ok {
		delete(DBsessions, cvalue)
		return nil
	} else {
		return errors.New("session value is not found")
	}
}
