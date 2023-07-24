package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net"

	"github.com/beevik/ntp"
)

var ntpServer string

func init() {
	flag.StringVar(&ntpServer, "n", "time.apple.com", "")
	flag.Parse()
}

func main() {
	ips, err := net.DefaultResolver.LookupIP(context.TODO(), "ip", ntpServer)
	if err != nil {
		panic(err)
	}
	if len(ips) == 0 {
		panic("not host")
	}
	selectIp := ips[0].String()
	r, err := ntp.Query(net.JoinHostPort(selectIp, "123"))
	if err != nil {
		panic(err)
	}
	bf := &bytes.Buffer{}
	e := json.NewEncoder(bf)
	e.SetEscapeHTML(false)
	e.SetIndent("", "    ")
	err = e.Encode(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(bf.String())

	fmt.Printf("server: %v (%v)\n", ntpServer, selectIp)
	fmt.Printf("offset: %v\n", r.ClockOffset)
	fmt.Printf("rtt: %v\n", r.RTT)
}
