package serverlogger

import (
	"log"
	"net/http"
	"os"
)

// append and format log according to api endpoint
func Appendtoserverlog(endpoint string) {
	f, err := os.OpenFile("serverlog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, endpoint, log.LstdFlags)
	logger.Println(http.StatusOK)
}
