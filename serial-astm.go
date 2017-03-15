package main

import (
	//"fmt"
	"bytes"
	"encoding/hex"
	"github.com/tarm/serial"
	"log"
	"reflect"
	"strings"
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
	ETB = 0x17
)

func main() {
	c := &serial.Config{Name: "COM2", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	//for {
	for !waitForResp(s, []byte{ENQ}) {
		send(s, []byte{NAK})
	}
	send(s, []byte{ACK})

	for !waitForValidData(s) {
		send(s, []byte{NAK})
	}

	//}
}

func send(s *serial.Port, msg []byte) {
	_, err := s.Write(msg)
	if err != nil {
		log.Fatal(err)
	}
	//printASTMMessage(msg)
}

func waitForResp(s *serial.Port, resp []byte) bool {
	buf := make([]byte, 128)
	n, err := s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	if reflect.DeepEqual(buf[:n], resp) {
		return true
	} else {
		return false
	}
}

func waitForValidData(s *serial.Port) bool {
	buf := make([]byte, 128)
	n, err := s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	msg := buf[:n]
	datamsg := searchBetween(msg, []byte{STX}, []byte{CR})
	// Se completa con el mensaje para el checksum
	data4cs := append(datamsg, []byte{CR, ETX}...)
	datacs := searchBetween(msg, []byte{ETX}, []byte{CR, LF})
	cs := checkSumASCII(checkSum8Mod256(data4cs))
	if bytes.Equal(datacs, cs) {
		return true
	} else {
		return false
	}
}

func searchBetween(text []byte, ini []byte, end []byte) []byte {
	iniI := bytes.Index(text, ini) + 1
	endI := bytes.Index(text, end)
	out := text[iniI:endI]
	return out
}

//http://www.asciitohex.com/
func checkSumASCII(sum byte) []byte {
	hexString := strings.ToUpper(hex.EncodeToString([]byte{sum}))
	hexBytes := []byte(hexString)
	return hexBytes
}

//http://www.scadacore.com/field-applications/programming-calculators/online-checksum-calculator/
//http://www.hendricksongroup.com/code_003.aspx
//http://foro6x.velneo.es/viewtopic.php?t=12299
func checkSum8Mod256(data []byte) byte {
	var sum byte = 0x00
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	return sum
}
