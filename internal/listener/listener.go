package listener

import (
	"net"
	"time"
	"fmt"
)

type Record struct {
	Ip string
	Header string
	ActiveSince time.Time
}

func newRecord(ip net.Addr, header []byte) Record {
	return Record{ip.String(), string(header), time.Now()}
}

func listen_to_broadcast() Record {
	fmt.Print("enter listening\n")
	defer fmt.Print("leaving listening\n")
	pc,err := net.ListenPacket("udp4", ":2010")
	if err != nil {
		panic(err)
	}
	defer pc.Close()

	buf := make([]byte, 1024)
	n,addr,err := pc.ReadFrom(buf)
	if err != nil {
		panic(err)
	}

	return newRecord(addr, buf[:n])
}

func add_to_map(record_chan chan map[string]Record, record Record) {
	record_map := <- record_chan	
	record_map[record.Ip] = record
	record_chan <- record_map
}

func RunListenRoutine(record_chan chan map[string]Record) {
	for true {
		record := listen_to_broadcast()
		add_to_map(record_chan, record)
	}
}
