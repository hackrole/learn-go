package main

import (
	"fmt"
)

type IPAddr [4]byte

func (o IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", o[0], o[1], o[2], o[3])

}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
