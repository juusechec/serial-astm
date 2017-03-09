package main

import (
	//"fmt"
	"github.com/tarm/serial"
	"log"
)

const (
	ENQ = 0x05
	ACK = 0x06
	NAK = 0x15
	STX = 0x02
	ETX = 0x03
	CR  = 0x0D
	LF  = 0x0A
	EOT = 0x04
	AAA = 0x41
)

func main() {
	c := &serial.Config{Name: "COM2", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	for {
    buf := make([]byte, 128)
  	n, err := s.Read(buf)
  	if err != nil {
  		log.Fatal(err)
  	}
  	log.Printf("%q", buf[:n])
	}
}
