package response

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

func active_services_response(w http.ResponseWriter, req *http.Request, record_chan chan map[string]udp_service.Record) {
	record_map := <- record_chan

	fmt.Fprintf(w, "[\n")
	num_valid := 0
	for _, record := range record_map {
		seconds := time.Now().Sub(record.ActiveSince).Seconds()
		if seconds > 60 { continue } 
		
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

	record_chan <- record_map
}


func RunResponseRoutine(record_chan chan map[string]udp_service.Record) {

    http.HandleFunc("/api/v1.0/active-http-services", 
		func(w http.ResponseWriter, req *http.Request) {
			active_services_response(w, req, record_chan)
	})

    http.ListenAndServe(":2020", nil)
}
