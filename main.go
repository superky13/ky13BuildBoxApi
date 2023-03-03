package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/superky13/ky13BuildBoxApi/serverlogger"
)

// swagger:parameters admin serverinfo
type serverinfo struct {
	// name: id
	// in: path
	// type: string
	// required: true
	ID string `json="id"`
	// name: servername
	// in: path
	// type: string
	// required: true
	Servername string `json="servername"`
	// name: ipaddress
	// in: path
	// type: string
	// required: true
	Ipaddress string `json="ipaddress"`
}

var servers []serverinfo

func returnServerInfo(w http.ResponseWriter, r *http.Request) {
	serverlogger.Appendtoserverlog("/api/v1/serverinfo ")
	json.NewEncoder(w).Encode(servers)
}
func apiWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ky13 buildbox api server")
	serverlogger.Appendtoserverlog("/api/v1 ")
}

// swagger:route GET /api/v1 admin handleApiRequests
//
// security:
// - apiKey: []
// responses:
// 200:
// 400:
func handleApiRequests() {
	http.HandleFunc("/api/v1", apiWelcome)
	http.HandleFunc("/api/v1/serverinfo", returnServerInfo)
	// developer swagger docs
	http.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	http.Handle("/docs", sh)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
func main() {
	servers = []serverinfo{
		serverinfo{ID: "1", Servername: "dns.ky13buildbox.com", Ipaddress: "192.168.122.120"},
		serverinfo{ID: "2", Servername: "git.ky13buildbox.com", Ipaddress: "192.168.122.189"},
	}
	handleApiRequests()
}
