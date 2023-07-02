package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// HandsOn Method
type HandsOn struct {
	Time     time.Time `json:"time"`
	HostName string    `json:"hostname"`
}

func index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	resp := HandsOn{
		Time:     time.Now(),
		HostName: os.Getenv("HOSTNAME"),
	}
	jsonResp, err := json.Marshal(&resp)
	if err != nil {
		w.Write([]byte("Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(jsonResp)
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Preparing to listen on port 3000...")
	http.ListenAndServe(":3000", nil)

}
