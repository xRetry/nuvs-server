package listener

import (
	"net"
	"time"
)

type Record struct {
	ip string
	header string
	active_since time.Time
}

func newRecord(ip net.Addr, header []byte) Record {
	return Record{ip.String(), string(header), time.Now()}
}

func listen_to_broadcast(record_chan chan map[string]Record) {
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
	record_map := <- record_chan	
	record_map[addr.String()] = newRecord(addr, buf[:n])
	record_chan <- record_map
}

func run_listen_routine(record_chan chan map[string]Record) {
	for true {
		listen_to_broadcast(record_chan)
	}
}
