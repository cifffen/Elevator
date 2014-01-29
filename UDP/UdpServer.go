package main

import (
  "net"
  "fmt"
)

func main() {
  addr, _ := net.ResolveUDPAddr("udp", ":2000")
  sock, _ := net.ListenUDP("udp", addr)
  fmt.Println("Connected")
  for {
    buf := make([]byte, 1024)
    rlen, address, err := sock.ReadFromUDP(buf)
	fmt.Println(1)
    if err != nil {
      fmt.Println(err)
    }
	_, err = sock.WriteToUDP(buf[0:rlen],address)
    fmt.Println(string(buf[0:rlen]))
    //go handlePacket(buf, rlen)
  }
}