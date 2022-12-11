package udp_service

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"net"
	"time"
	"sync"
)


var RecordsMtx sync.RWMutex
var Records map[string]Record


type Record struct {
	Ip string
	Header string
	ActiveSince time.Time
}


func newRecord(ip net.Addr, header []byte) Record {
	return Record{ip.String(), string(header), time.Now()}
}


func checkLocalhost() (string, error) {
	fmt.Println("checking")
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


func sendBroadcast(conn *net.PacketConn) {
	addr,err := net.ResolveUDPAddr("udp4", "10.21.255.255:2010")
	if err != nil {
		panic(err)
	}

	for {
		message, err := checkLocalhost()
		if err == nil {
			_,err = (*conn).WriteTo([]byte(message), addr)
			fmt.Println("broadcasting")
			if err != nil {
				panic(err)
			}
		}

		time.Sleep(time.Second * 60)
	}
}


func listenToBroadcast(conn *net.PacketConn) {
	fmt.Println("enter listening")

	for {
		buf := make([]byte, 1024)
		n,addr,err := (*conn).ReadFrom(buf)
		if err != nil {
			panic(err)
		}

		go addToMap(newRecord(addr, buf[:n]))
	}
}


func addToMap(record Record) {
	RecordsMtx.Lock()
	fmt.Println("adding to map")
	Records[record.Ip] = record
	RecordsMtx.Unlock()
}


func RunUdpService() {
	conn, err := net.ListenPacket("udp4", ":2010")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	go listenToBroadcast(&conn)
	go sendBroadcast(&conn)

	waitGroup.Wait()
}
