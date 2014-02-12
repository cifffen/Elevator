package main

import (
  "net"
  "fmt"
)

func main() {
  addr, _ := net.ResolveUDPAddr("udp", ":2224")
  sock, _ := net.ListenUDP("udp", addr)
  fmt.Println("Connected")
  for {
    buf := make([]byte, 1024)
    rlen, address, err := sock.ReadFromUDP(buf)
	fmt.Println(address)
    if err != nil {
      fmt.Println(err)
    }
	_, err = sock.WriteToUDP(buf,address)
    fmt.Println(string(buf[0:rlen]))
    //go handlePacket(buf, rlen)
  }
}