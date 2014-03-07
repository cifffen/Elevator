package main

import (
  "net"
  "fmt"
  "log"
  "encoding/json"
  "strings"
)
var sock *net.UDPConn
type ActionType int
const NetworkIp = "192.168.1.255"
const NetworkPort = ":2224"
const(
	NewOrder ActionType = iota
	DeleteOrder
	Tender
)

type ButtonMessage struct {
	Floor int
	Button int
	Action ActionType
}

func send(msg ButtonMessage)(){
	fmt.Printf("her \n")
	var err error
	var addr net.Addr
	for try :=0; try <2; try++ {
		fmt.Printf("try=",try)
		if try ==0{
			addr, err= net.ResolveUDPAddr("udpp", NetworkIp+NetworkPort)
		}else if try ==1 {
			addr, err= net.ResolveUDPAddr("udp", NetworkIp+NetworkPort)
		}
		if err != nil{
			log.Printf("Error: %v. Runing without network connetion", err)
			
		} else{
			fmt.Printf("jepp",try, addr)
			break
		}
	}
	buf,err := json.Marshal(msg)
	//rAddr, _ := net.ResolveUDPAddr("udp", "192.168.1.255:2224")
	_,err =sock.WriteTo(buf,addr)
	fmt.Printf("naaa\n")
	fmt.Println(sock.LocalAddr())
	if err != nil{
		fmt.Println(err)
		return
	}
}

func getSelfIP() string {
	conn,err :=net.Dial("udp", "google.com:80") 
	if err !=nil {
		log.Printf("Error: %v. Runing without self address checking.", err)
		return "localhost"

	} else{
		return strings.Split(string(conn.LocalAddr().String()), ":" )[0] 

	} 
}
	 
func main() {
	xere := getSelfIP()
	//xere :=strings.Split(string(connn.LocalAddr().String()), ":" )[0] 
	fmt.Println(xere)
	var m ButtonMessage
	addr, err := net.ResolveUDPAddr("udp", ":2224")
	//mine:=net.LookupHost(sock.LocalAddr())
	
	if err != nil {
		log.Println(err)
	}
	sock, _ = net.ListenUDP("udp", addr)
	//fmt.Println((net.*UDPConn).LocalAddr())
	fmt.Println("Listening on", addr)
	m1 :=ButtonMessage{
		Floor: 5,
		Button: 2,
		Action: Tender,
	}
	go send(m1)
	fmt.Printf("nope \n")
	i:=0
	for {
		if i >=2{
			break
		}
		i++
		buf := make([]byte, 1024)
		rlen, address, err := sock.ReadFromUDP(buf)
		sAddr, _ := net.ResolveUDPAddr("udp", "localhost:2224")
		if address != sAddr{
			
			err = json.Unmarshal(buf[0:rlen], &m)
			fmt.Println(address)
			fmt.Println(sAddr)
			if err != nil {
			  fmt.Println(err)
			}
			send(m)
			//_, err = sock.WriteToUDP(buf,address)
			if err != nil{
				fmt.Println(err)
			}
			//fmt.Printf("Floor: %d, Button: %d\n, Action: %d", m.Floor, m.Button, m.Action)
			//go handlePacket(buf, rlen)
		}
	}
}