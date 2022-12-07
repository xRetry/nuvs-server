package http_service

import (
    "fmt"
    "net/http"
	"github.com/xRetry/nuvs-server/internal/udp-service"
	"encoding/json"
	"time"
)


type DiffRecord struct {
	Ip string `json:"ip"`
	Header string `json:"header"`
	Diff int `json:"valid"`
}


func serve_json(w http.ResponseWriter, req *http.Request, active bool) {

	udp_service.RecordsMtx.RLock()

	fmt.Fprintf(w, "[\n")
	num_valid := 0
	for _, record := range udp_service.Records {
		seconds := time.Now().Sub(record.ActiveSince).Seconds()

		if active && seconds > 60 { continue } 
		
		if num_valid > 0 { fmt.Fprintf(w, ",\n") }
		num_valid += 1

		diff_record := DiffRecord{record.Ip, record.Header, int(seconds)}

		b, err := json.MarshalIndent(diff_record, "\t", "\t")
		if err != nil {
			continue
		}
		fmt.Fprintf(w, "\t")
		fmt.Fprintf(w, string(b)+"")

	}
	fmt.Fprintf(w, "\n]")

	udp_service.RecordsMtx.RUnlock()
}


func RunHttpRoutine() {

    http.HandleFunc("/api/v1.0/active-http-services", 
		func(w http.ResponseWriter, req *http.Request) {
			serve_json(w, req, true)
	})

    http.HandleFunc("/api/v1.0/all-http-services", 
		func(w http.ResponseWriter, req *http.Request) {
			serve_json(w, req, false)
	})

    http.ListenAndServe(":2020", nil)
}
