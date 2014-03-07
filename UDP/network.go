package main

import(
	"net"
	"fmt"
	"encoding/json"
)

var sock *net.UDPConn
type ActionType int
const NetworkIp = "192.168.1.255"
const NetworkPort = ":2224"

const (
	InvalidMsg ActionType  = iota  //	Only used to check if the message recieved is of type ButtonMsg.
	NewOrder		 //
	DeleteOrder
	Tender
	AddOrder
)
type OrderType struct{
	Button 	int			// Holds the button on the floor, Up or Down
	Floor 	int			// Holds the floor
}

type ButtonMsg struct {
	Action    	ActionType   	// Holds what the information of what to do with the message
	Order 		OrderType 		// Holds the floor and button of the order
	TenderVal 	int				// If the action is a Tender, this will hold the cost from the sender, that is, the value from the cost function for this order
}
func BroadcastOnNet(msg ButtonMessage)(){
	addr, _ := net.ResolveUDPAddr("udp", NetworkIp+NetworkPort)
	buf,err1 := json.Marshal(msg)
	if err1 != nil{
		fmt.Println(err1)
	}
	//rAddr, _ := net.ResolveUDPAddr("udp", "192.168.1.255:2224")
	_,err1 =sock.WriteTo(buf,addr)
	if err1 != nil{
		log.Println(err1)
	}
}
//func listenOnNetwork()(){
func main()(){
	var m ButtonMessage
	addr, _ := net.ResolveUDPAddr("udp", NetworkPort)
	sock, _ = net.ListenUDP("udp", addr)
	fmt.Println("Connected")
	/*
	m1 :=ButtonMessage{
		Floor: 5,
		Button: 2,
		Action: Tender,
	}
	BroadcastOnNet(m1)
	*/
	for {
		buf := make([]byte, 1024)
		rlen, address, err := sock.ReadFromUDP(buf)
		//sAddr, _ := net.ResolveUDPAddr("udp", "localhost:2224")

		err = json.Unmarshal(buf[0:rlen], &m)
		fmt.Println(address)
		//BroadcastOnNet(m)
		if err != nil {
		  fmt.Println(err)
		}
		//fmt.Println(buf[0:rlen])
		//send(m)
		//_, err = sock.WriteToUDP(buf,address)
		fmt.Printf("Floor: %d, Button: %d\n, Action: %d", m.Order.Floor, m.Order.Button, m.Action)
	}
}