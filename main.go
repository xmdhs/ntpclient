package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/beevik/ntp"
)

var ntpServer string

func init() {
	flag.StringVar(&ntpServer, "n", "time.apple.com", "")
	flag.Parse()
}

func main() {
	r, err := ntp.Query(ntpServer)
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

	fmt.Printf("offset: %v\n", r.ClockOffset)
	fmt.Printf("rtt: %v\n", r.RTT)
}
