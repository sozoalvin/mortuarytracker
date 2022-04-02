package login

import "wecare/dbMemoryUsers"

func CheckUserNameValid(username string) bool {
	for _, v := range dbMemoryUsers.DBLogInUsers {
		if v == username {
			return true
		}
	}
	// if after ranging through the entire user list and still can't find, then it really doesn't exist. have to add the user in
	return false
}
