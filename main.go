package telehash

import (
	"flag"
	"fmt"
	"log"
	"net"
	"telehash/exchange"
	"telehash/telex"
)

var (
	port = flag.Int("port", 4242, "Specify the UDP port to listen on")
)

func init() {
	flag.Parse()
}

func main() {

	exchange, err := listener.New((*port))
	if err != nil {
		log.Fatal(err)
	}

	defer exchange.Close()

	for {
		msg, err := exchange.Read()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(msg)
	}
}
