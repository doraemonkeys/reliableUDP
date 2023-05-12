package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Doraemonkeys/reliableUDP"
)

func main() {
	udpconn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: 12346})
	if err != nil {
		log.Fatal(err)
	}
	rudp := reliableUDP.New(udpconn)
	defer rudp.Close()
	go func() {
		for {
			d, addr, err := rudp.ReceiveAll(0)
			if err != nil {
				log.Println(err)
			}
			fmt.Println("receive", string(d), addr.String())
		}
	}()
	i := 0
	rudp.SetGlobalReceive()
	for {
		msg := fmt.Sprintf("hello %d", i)
		err := rudp.Send(&net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 12345}, []byte(msg), 0)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second * 2)
		i++
		if i > 5 {
			break
		}
	}
}
