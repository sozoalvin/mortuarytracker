package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"text/template"
	"time"
	"wecare/dbMemoryCasualty"
	"wecare/dbMemoryUsers"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/acme/autocert"
)

var (
	domain        string
	productionFlg bool = false
)

type session struct {
	un           string
	lastActivity time.Time
}

var tpl *template.Template

var dbSessions = map[string]session{}
var dbSessionsCleaned time.Time

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	// dbSessionsCleaned = time.Now()
	go dbMemoryUsers.Init()
}

func main() {

	flag.StringVar(&domain, "domain", "", "domain name to request your certificate")
	flag.BoolVar(&productionFlg, "productionFlg", false, "if true, we start HTTPS server")
	flag.Parse()

	currentTime := time.Now()

	fmt.Println("System Message : System is Ready", currentTime.Format("2006-01-02 15:04:05"))

	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/login", login)
	router.HandleFunc("/logout", logout)
	router.HandleFunc("/process", process).Methods("POST")
	router.HandleFunc("/searchCasualty", SearchCasualty).Methods("GET", "POST")
	router.HandleFunc("/CasualtySearchResultPage", SearchCasualtyResult)
	router.HandleFunc("/AllCasualtyLogs", AllCasualtyLogs)
	router.HandleFunc("/users", users)
	router.HandleFunc("/deleteAll", DeleteAllCasualtyLogs)
	router.HandleFunc("/assets/PERSCOM-501.png", PERSCOM_501)
	router.HandleFunc("/assets/rs2.png", rs2)
	router.HandleFunc("/assets/beepSuccessful.mp3", beepSuccesful)
	router.HandleFunc("/assets/beepFailed.mp3", beepFailed)
	router.HandleFunc("/src/html5-qrcode.min.js", qrCodeJS)
	router.Handle("/favicon.ico", http.NotFoundHandler()) //NotFoundHandler returns a simple request handler that replies to each request with a “404 page not found” reply.

	if productionFlg {
		certManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(domain),
			Cache:      autocert.DirCache("certs"),
		}
		tlsConfig := certManager.TLSConfig()
		tlsConfig.GetCertificate = getSelfSignedOrLetsEncryptCert(&certManager)
		server := http.Server{
			Addr:      ":443",
			Handler:   router,
			TLSConfig: tlsConfig,
		}

		go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
		fmt.Println("Server listening on", server.Addr)
		fmt.Println("production flag is activated")
		if err := server.ListenAndServeTLS("", ""); err != nil {
			fmt.Println(err)
		}
	} else {
		port := os.Getenv("PORT")
		fmt.Println("ProductionFlg is not activated\nServer listening on :80")
		http.ListenAndServe(port, router)
	}
	go func() {
		for {
			for _, v := range dbMemoryCasualty.CasualtyDB.Data {
				fmt.Printf("this is the casualty data every 5 secs %+v\n", v)
			}
			time.Sleep(5 * time.Second)
		}
	}()
} // close main functioin

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
} //end function check
