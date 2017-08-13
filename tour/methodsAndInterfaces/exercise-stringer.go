package main

import (
	"fmt"
)

type IpAddr [4]byte

func (self IpAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", self[0], self[1], self[2], self[3])
}

func main() {
	hosts := map[string]IpAddr{
		"loop":      {127, 0, 0, 1},
		"googleDns": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v:%v\n", name, ip)
	}
}
