# reliableUDP
a simple implementation of reliable udp





## example

```go
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
	defer udpconn.Close()
	rudp := reliableUDP.NewReliableUDP(udpconn)
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
	for {
		msg := fmt.Sprintf("hello %d", i)
		err := rudp.Send([]byte(msg), &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: 12345})
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
```

```go
package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Doraemonkeys/reliableUDP"
)

func main() {
	udpconn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: 12345})
	if err != nil {
		log.Fatal(err)
	}
	defer udpconn.Close()
	rudp := reliableUDP.NewReliableUDP(udpconn)
	defer rudp.Close()
	ch := make(chan string)
	go func() {
		addr := <-ch
		raddr, err := net.ResolveUDPAddr("udp", addr)
		if err != nil {
			log.Println(err)
		}
		i := 0
		for {
			msg := fmt.Sprintf("hello %d", i)
			err := rudp.Send([]byte(msg), raddr)
			if err != nil {
				log.Println(err)
			}
			time.Sleep(1 * time.Second)
			i++
		}
	}()
	i := 0
	for {
		d, addr, err := rudp.ReceiveAll(0)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("receive", string(d), addr.String())
		if i == 0 {
			go func() {
				ch <- addr.String()
			}()
		}
		i++
	}
}
```

