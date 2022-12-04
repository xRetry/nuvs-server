package main

import (
	"sync"
	"github.com/xRetry/nuvs-server/internal/udp-service"
	"github.com/xRetry/nuvs-server/internal/response"
)

var waitGroup sync.WaitGroup

func main() {

	ch := make(chan map[string]udp_service.Record, 1)

	m := make(map[string]udp_service.Record) 
	ch <- m

	waitGroup.Add(2)

	go udp_service.RunUdpService(ch)
	go response.RunResponseRoutine(ch)

	waitGroup.Wait()

}
