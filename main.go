package main

import (
	"encoding/json"
	"fmt"
	"github.com/superky13/ky13BuildBoxApi/serverlogger"
	"log"
	"net/http"
)

type serverinfo struct {
	ID         string `json="id"`
	Servername string `json="servername"`
	Ipaddress  string `json="ipaddress"`
}

var servers []serverinfo

func returnServerInfo(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("/api/v1/serverinfo", time.Now(), http.StatusOK)
	serverlogger.Appendtoserverlog("/api/v1/serverinfo ")
	json.NewEncoder(w).Encode(servers)
}
func apiWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ky13 buildbox api server")
	serverlogger.Appendtoserverlog("/api/v1 ")
	// log.Output(1, http.StatusText(200))
}
func handleRequests() {
	http.HandleFunc("/api/v1", apiWelcome)
	http.HandleFunc("/api/v1/serverinfo", returnServerInfo)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
func main() {
	servers = []serverinfo{
		serverinfo{ID: "1", Servername: "dns.ky13buildbox.com", Ipaddress: "192.168.122.120"},
		serverinfo{ID: "2", Servername: "git.ky13buildbox.com", Ipaddress: "192.168.122.189"},
	}
	handleRequests()
}
