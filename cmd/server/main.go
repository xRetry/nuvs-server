package main

import (
	"sync"
	"github.com/xRetry/nuvs-server/internal/udp-service"
	"github.com/xRetry/nuvs-server/internal/http-service"
)

func main() {
	udp_service.Records = make(map[string]udp_service.Record) 
	udp_service.RecordsMtx = sync.RWMutex{}

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	go udp_service.RunUdpService()
	go http_service.RunHttpRoutine()

	waitGroup.Wait()

}
