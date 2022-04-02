package dbMemoryUsers

var DBLogInUsers []string

// init for log in Users
func Init() {
	// everything here will be added by default
	DBLogInUsers = []string{
		"123",
		"BodyCleanIn",
		"BodyCleanOut",
		"BodyDressIn",
		"BodyDressOut",
		"EcoffinIn",
		"EcoffinOut",
		"InPro",
		"OutPro",
		"OutproInPECheck",
		"PeGoStorage",
		"PeInPro",
		"PeIntoStorage",
		"PeOutPro",
		"PeOutofStorage",
		"ProcessedCleanIn",
		"ProcessedCleanOut",
		"TransportIn",
		"TransportOut",
		"UnprocessedIn",
		"UnprocessedOut",
		"MortCommand",
	}
}

// this adds extra user namme to your program
func AddUsers(username string) {
	DBLogInUsers = append(DBLogInUsers, username)
}

// TODO @ use slice method to add and delete users
func DeleteUsers(username string) {
}
