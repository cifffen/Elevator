package main
import (
  "net"
  "fmt"
  "encoding/json"

)
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

func main() {

	x:=5
	m :=ButtonMessage{
		Floor: x,
		Button: 2,
		Action: Tender,
	}
	buf,err1 := json.Marshal(m)
	if err1 != nil{
		fmt.Println(err1)
	}
	/*
	aaaaaaaaa
	*/
	addr, _ := net.ResolveUDPAddr("udp", "192.168.1.255:2224")
	addr2, _ := net.ResolveUDPAddr("udp", "192.168.1.4:2422")
	con, _ := net.DialUDP("udp", addr2, addr)
	//buf := []byte("bla bla bla I am the packet\n")
	_, err := con.Write(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("1") 
	var m1 ButtonMessage
	buf2 := make([]byte, 1024)
	var rlen int
	rlen,_, err =con.ReadFromUDP(buf2)
	_ = json.Unmarshal(buf2[0:rlen],&m1)
	fmt.Printf("1")
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Printf("Floor: %d, Button: %d", m1.Floor, m1.Button)
}