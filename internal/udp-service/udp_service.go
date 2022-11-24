package udp_service

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"net"
	"time"
)


type Record struct {
	Ip string
	Header string
	ActiveSince time.Time
}


func newRecord(ip net.Addr, header []byte) Record {
	return Record{ip.String(), string(header), time.Now()}
}


func connect_to_localhost() (string, error) {
	resp, err := http.Get("http://127.0.0.1:2000/")
	if err != nil {
		return "", fmt.Errorf("Unable to connect to server")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Unable to read body") 
	}
	return strings.Split(string(body), "\n")[0], nil
}

func broadcast_message(message string) {
	fmt.Print("enter broadcast\n")
	defer fmt.Print("leaving broadcast\n")
	pc, err := net.ListenPacket("udp4", ":2010")
	if err != nil {
		panic(err)
	}
	defer pc.Close()

	addr,err := net.ResolveUDPAddr("udp4", "10.21.0.255:2010")
	if err != nil {
		panic(err)
	}

	_,err = pc.WriteTo([]byte(message), addr)
	if err != nil {
		panic(err)
	}
}

func listen_to_broadcast(record_chan chan map[string]Record) {
	fmt.Print("enter listening\n")

	pc,err := net.ListenPacket("udp4", ":2010")
	if err != nil {
		panic(err)
	}
	pc.SetReadDeadline(time.Now().Add(60 * time.Second))
	defer pc.Close()

	for {
		buf := make([]byte, 1024)
		n,addr,err := pc.ReadFrom(buf)
		if err != nil {
			if e, ok := err.(net.Error); !ok || !e.Timeout() {
				panic(err)
			}
			break
		}

		fmt.Print("adding to map\n")
		go add_to_map(record_chan, newRecord(addr, buf[:n]))
	}

	fmt.Print("leaving listening\n")
}

func add_to_map(record_chan chan map[string]Record, record Record) {
	record_map := <- record_chan	
	record_map[record.Ip] = record
	record_chan <- record_map
}

func RunUdpService(record_chan chan map[string]Record) {
	for true {
		listen_to_broadcast(record_chan)
		body, err := connect_to_localhost()
		if err == nil {
			broadcast_message(body)
		}
	}
}
