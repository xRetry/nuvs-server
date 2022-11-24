package main

import (
	"sync"
	"github.com/xRetry/nuvs-server/internal/udp-service"
	"github.com/xRetry/nuvs-server/internal/response"
	"time"
)

var waitGroup sync.WaitGroup

func main() {

	ch := make(chan map[string]udp_service.Record, 1)

	m := make(map[string]udp_service.Record) 
	m["sfds"] = udp_service.Record{"23.42", "dfsfsfd", time.Now()}
	m["dfse"] = udp_service.Record{"23.45.2", "dfsfsfd", time.Now()}
	ch <- m

	waitGroup.Add(3)

	go udp_service.RunUdpService(ch)
	go response.RunResponseRoutine(ch)

	waitGroup.Wait()

}
