package main

import (
	"sync"
	"github.com/xRetry/nuvs-server/internal/contact"
	"github.com/xRetry/nuvs-server/internal/listener"
	"github.com/xRetry/nuvs-server/internal/response"
	"time"
)

var waitGroup sync.WaitGroup

func main() {

	ch := make(chan map[string]listener.Record, 1)

	m := make(map[string]listener.Record) 
	m["sfds"] = listener.Record{"23.42", "dfsfsfd", time.Now()}
	m["dfse"] = listener.Record{"23.45.2", "dfsfsfd", time.Now()}
	ch <- m

	waitGroup.Add(3)

	go contact.RunContactRoutine()
	go listener.RunListenRoutine(ch)
	go response.RunResponseRoutine(ch)

	waitGroup.Wait()

}
