package main

import (
	"net"
	"os"
	"fmt"
	"time"
	"os/exec"
	)

func startSlave()() {

	cmd := exec.Command("cmd", "/C", "start", "processpairs.exe", "slave")
	fmt.Printf("Slave started\n")	
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}	
}	

func udpListenToMaster(number chan<- int, quit <-chan int)() {	
	addr, err1 := net.ResolveUDPAddr("udp", ":2224")
	if err1 != nil {
				time.Sleep(time.Second*4)
		fmt.Println(err1)
	}
	sock, err2:= net.ListenUDP("udp", addr)
	if err2 != nil {
		fmt.Println(err2)
		time.Sleep(time.Second*4)
	}
	fmt.Println("Connected")
	l:
		for {
			buf :=make([]byte,1024)
			_, _, err := sock.ReadFromUDP(buf)
			if err != nil {
				time.Sleep(time.Second*4)
				fmt.Println(err)	
			}
			select{
				case <-quit:
					sock.Close()
					break l
				default:
					number<-int(buf[0])
			}
	}
}

func udpBroadcast(number int)(){
	con,_ := net.Dial("udp", "localhost:2224")
	buf :=[]byte(string(number))
	_, err := con.Write(buf)
	if err != nil {
		fmt.Println(err)
	}
}

func startCount(number int)() {
	for{
		fmt.Printf("%d\n",number)	
		number++
		udpBroadcast(number)
		time.Sleep(time.Second*1)
	}
}
	

func main() {
	if len(os.Args)==1{
		startSlave()
		startCount(2)
	} else {
		
		ticker := time.NewTicker(time.Second*8)
		numberChan := make(chan int)
		quitChan := make(chan int)
		go udpListenToMaster(numberChan, quitChan)
		var number int			
		for{
			select{
				case number = <-numberChan:
					ticker.Stop()
					ticker = time.NewTicker(time.Second*8)
				case <-ticker.C:
					close(quitChan)
					time.Sleep(time.Second*1)
					startSlave()		
					startCount(number)
					
			}
		}
	}
}

