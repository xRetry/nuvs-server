package main

import (
	"github.com/xRetry/nuvs-server/internal/contact"
	"github.com/xRetry/nuvs-server/internal/listener"
	"github.com/xRetry/nuvs-server/internal/response"
)


func main() {

	ch := make(chan map[string]listener.Record)

	go contact.run_contact_routine()
	go listener.run_listen_routine(ch)
	go response.run_response_routine(ch)

}
