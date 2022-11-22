package contact

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	"net"
)


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


func run_contact_routine() {
	body, err := connect_to_localhost()
	if err == nil {
		broadcast_message(body)
	}
}
