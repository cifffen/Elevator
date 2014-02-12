package main
import (
  "net"
  "fmt"
)

func main() {
  con, _ := net.Dial("udp", "192.168.1.255:2224")
  buf := []byte("bla bla bla I am the packet\n")
  _, err := con.Write(buf)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(1) 
  buf2 := make([]byte, 1024)
  rlen, err :=con.Read(buf2)
  fmt.Println(1)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(string(buf[0:rlen]))
}