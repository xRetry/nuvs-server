package response

import (
    "fmt"
    "net/http"
	"github.com/xRetry/nuvs-server/internal/listener"
)

func active_services_response(w http.ResponseWriter, req *http.Request) {
	// TODO: Create response
    fmt.Fprintf(w, "V1.0 m01602279\naddNumbers(x,y)\nisPrime(val)\n")
}


func run_response_routine(record_chan chan map[string]listener.Record) {

    http.HandleFunc("/api/v1.0/active-http-services", active_services_response)

    http.ListenAndServe(":2020", nil)
}
