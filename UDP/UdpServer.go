package main

import (
  "net"
  "fmt"
  "encoding/json"
)
var sock *net.UDPConn
type ActionType int

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
	addr, _ := net.ResolveUDPAddr("udp", "192.168.1.255:2224")
	buf,err1 := json.Marshal(msg)
	if err1 != nil{
		fmt.Println(err1)
	}
	//rAddr, _ := net.ResolveUDPAddr("udp", "192.168.1.255:2224")
	_,err1 =sock.WriteTo(buf,addr)
	if err1 != nil{
		fmt.Println(err1)
	}
	fmt.Printf("her \n")
}
func main() {
	var m ButtonMessage
	addr, _ := net.ResolveUDPAddr("udp", ":2224")
	sock, _ = net.ListenUDP("udp", addr)
	fmt.Println("Connected")
	m1 :=ButtonMessage{
		Floor: 5,
		Button: 2,
		Action: Tender,
	}
	send(m1)
	i:=0
	for {
		if i >=10{
			break
		}
		i++
		buf := make([]byte, 1024)
		rlen, address, err := sock.ReadFromUDP(buf)
		sAddr, _ := net.ResolveUDPAddr("udp", "localhost:2224")
		if address != sAddr{
			err = json.Unmarshal(buf[0:rlen], &m)
			fmt.Println(address)
			if err != nil {
			  fmt.Println(err)
			}
			fmt.Println(buf[0:rlen])
			//send(m)
			//_, err = sock.WriteToUDP(buf,address)
			if err != nil{
				fmt.Println(err)
			}
			fmt.Printf("Floor: %d, Button: %d\n, Action: %d", m.Floor, m.Button, m.Action)
			//go handlePacket(buf, rlen)
		}
	}
}