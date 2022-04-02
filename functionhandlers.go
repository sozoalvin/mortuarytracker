package main

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"regexp"
	"strconv"
	"wecare/dbMemoryCasualty"
	"wecare/dbMemorySessions"
	"wecare/entities"
	mortLogin "wecare/login"

	uuid "github.com/satori/go.uuid"
)

func index(w http.ResponseWriter, req *http.Request) {

	// u := getUser(w, req)

	username, ok := alreadyLoggedIn(w, req)

	if !ok {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		// tpl.ExecuteTemplate(w, "login.gohtml", nil)
		return
	}

	parseData := struct {
		U    string
		Data string
	}{
		username, "",
	}

	// showSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "index.gohtml", parseData)
	// tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	_, ok := alreadyLoggedIn(w, req)
	if ok {
		http.Redirect(w, req, "/", http.StatusSeeOther) //if alreadyLoggedIn == true -> returns them to see what they're supposed to see etc.
		return
	}

	var me = make(map[string]string) //make map for error

	if req.Method == http.MethodPost {

		unUnsanitized := req.FormValue("username") //the 'name' of the field
		p := req.FormValue("password")             //the 'name' of the field
		un := html.EscapeString(unUnsanitized)

		if !mortLogin.CheckUserNameValid(un) {
			me["Username1"] = "username is not a registered mort usernmame"
			tpl.ExecuteTemplate(w, "login.gohtml", me)
			return
		}

		if p != un {
			me["Password"] = "Invaid Password Entered. Please try again"
			tpl.ExecuteTemplate(w, "login.gohtml", me)
			return
		}
		// if password same as username, then it is a match

		sID, err := uuid.NewV4()
		////err handling
		if err != nil {
			fmt.Printf("Something went wrong when assigning cookies: %s", err)
		}

		c := &http.Cookie{
			Name:     "session",
			Value:    sID.String(),
			HttpOnly: true,
		}

		http.SetCookie(w, c)
		dbMemorySessions.AddNewSessions(c.Value, un)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	// showSessions() // for demonstration purposes
	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}
func SearchCasualty(w http.ResponseWriter, req *http.Request) {
	username, ok := alreadyLoggedIn(w, req)

	if !ok {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		casualtyCaseID := req.FormValue("casualtyCaseID") //pid is also known as the productID
		fmt.Println("req method post for the form detected")
		buildParam := `/CasualtySearchResultPage` + `?` + "casualtyCaseID" + `=` + casualtyCaseID
		http.Redirect(w, req, buildParam, http.StatusSeeOther)
	}

	parseData := struct {
		U    string
		Data string
	}{
		username, "",
	}

	tpl.ExecuteTemplate(w, "searchCasualty.gohtml", parseData)

}

func SearchCasualtyResult(w http.ResponseWriter, req *http.Request) {
	username, ok := alreadyLoggedIn(w, req)
	var err error

	if !ok {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	casualtyCaseID := req.URL.Query().Get("casualtyCaseID")

	casualtyCaseIDInt, _ := strconv.Atoi(casualtyCaseID)

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	data := dbMemoryCasualty.SearchCasualty(casualtyCaseIDInt)
	scanData := dbMemoryCasualty.SearchCasualtyScanLogs(casualtyCaseIDInt)

	if data == nil {
		data = nil

		parseData := struct {
			U       string
			Results string
		}{
			username,
			"empty",
		}
		tpl.ExecuteTemplate(w, "searchCasualtyResult.gohtml", parseData)
		return

	}

	parseData := struct {
		U        string
		Data     entities.Casualty
		ScanData []entities.ScanLogs
		CaseID   string
	}{
		username, *data, scanData, casualtyCaseID,
	}

	tpl.ExecuteTemplate(w, "searchCasualtyResult.gohtml", parseData)

}

func AllCasualtyLogs(w http.ResponseWriter, req *http.Request) {
	username, ok := alreadyLoggedIn(w, req)

	if !ok {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	data := dbMemoryCasualty.AllCasualtyLogs()

	if data == nil {
		data = nil

		parseData := struct {
			U       string
			Results string
		}{
			username,
			"empty",
		}
		tpl.ExecuteTemplate(w, "allCasualtyLogs.gohtml", parseData)
		return

	}

	parseData := struct {
		U    string
		Data map[int]entities.Casualty
	}{
		username, data,
	}

	tpl.ExecuteTemplate(w, "allCasualtyLogs.gohtml", parseData)

}
func users(w http.ResponseWriter, req *http.Request) {

	username, ok := alreadyLoggedIn(w, req)

	if !ok {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	parseData := struct {
		U    string
		Data string
	}{
		username, "",
	}
	tpl.ExecuteTemplate(w, "users_2.gohtml", parseData)
}

func DeleteAllCasualtyLogs(w http.ResponseWriter, req *http.Request) {

	dbMemoryCasualty.DeleteAllCasualtyLogs()
}

func process(w http.ResponseWriter, req *http.Request) {
	fmt.Println("process function handler got called")

	if req.Method != "POST" {
		fmt.Println("wrong type of method used for should log error")
		http.Error(w, "Bad request - Go away!", 400)
		return

	} else {
		fmt.Println("welcome post method")

		defer req.Body.Close()

		var c entities.CasualtyProcessing

		decoder := json.NewDecoder(req.Body)
		// decoder.DisallowUnknownFields()
		err := decoder.Decode(&c)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Json Unmarshalling errored out. Check Server", 500)
			return

		} else {
			http.Error(w, "Accepted request!", 200)
			fmt.Printf("this is our casualty %+v\n", c)
			dbMemoryCasualty.UpsertCasualty(&c)
			return
		}
	}
}

func logout(w http.ResponseWriter, req *http.Request) {
	_, ok := alreadyLoggedIn(w, req)
	if !ok { //if you are not logged in, there's nothing you need to do. whatever UUID cookie value, belongs to a non-logged in visitor
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")

	dbMemorySessions.DeleteSessions(c.Value)

	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	http.Redirect(w, req, "/", http.StatusSeeOther) //goes back to home page after logging out
}

func PERSCOM_501(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "assets/PERSCOM-501.png")
}

func qrCodeJS(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "src/html5-qrcode.min.js")
}
func beepSuccesful(w http.ResponseWriter, req *http.Request) {
	fmt.Println("successfuly handler served")
	http.ServeFile(w, req, "assets/beepSuccessful.mp3")
}
func beepFailed(w http.ResponseWriter, req *http.Request) {
	fmt.Println("failed handler served")
	http.ServeFile(w, req, "assets/beepFailed.mp3")
}
func rs2(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "assets/rs2.png")
}

func rs3(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "assets/rs3.png")
}

func rs4(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "assets/rs4.png")
}

func rs5(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "assets/rs5.png")
}

func rs6(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "assets/rs6.png")
}

func rs7(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "assets/rs7.png")
}
func rs8(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "assets/rs8.png")
}

func loginbutton(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "assets/login.png")
}

func signupbutton(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "assets/signup.png")
}

func validateInputs(un string, p string, f string, l string, me map[string]string) (bool, map[string]string) {

	rx := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(un) == 0 || len(un) > 40 {
		me["Username"] = "Username is not valid. Please Enter again"
	} else if !rx.MatchString(un) {
		me["Username"] = "Username is not valid email address. Please Enter again"
	}
	if len(p) == 0 {
		me["Password"] = "Password is not valid. Please Enter again"
	}
	if len(f) == 0 || len(f) > 45 {
		me["FirstName"] = "First Name is not valid. Please Enter again"
	}
	if len(l) == 0 || len(l) > 45 {
		me["LastName"] = "Last Name is not valid. Please Enter again"
	}
	if len(un) != 0 && len(p) != 0 && len(un) != 0 && len(l) != 0 && rx.MatchString(un) && len(un) < 40 && len(f) < 45 && len(l) < 45 {
		return true, me
	}
	return false, me
}
