package main

import (
	"encoding/hex"
	"fmt"
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
	c := &serial.Config{Name: "COM1", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

  // Inicio
	msg := []byte{ENQ}
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
  // Fin

  // Inicio
	msg = []byte{'1', 0x52, 0x7c, 0x32, 0x7c, 0x5e, 0x5e, 0x5e, 0x31, 0x2e, 0x30, 0x30, 0x30, 0x30, 0x2b, 0x39, 0x35, 0x30, 0x2b, 0x31, 0x2e, 0x30, 0x7c, 0x31, 0x35, 0x7c, 0x7c, 0x7c, 0x5e, 0x35, 0x5e, 0x7c, 0x7c, 0x56, 0x7c, 0x7c, 0x33, 0x34, 0x30, 0x30, 0x31, 0x36, 0x33, 0x37, 0x7c, 0x32, 0x30, 0x30, 0x38, 0x30, 0x35, 0x31, 0x36, 0x31, 0x35, 0x33, 0x35, 0x34, 0x30, 0x7c, 0x32, 0x30, 0x30, 0x38, 0x30, 0x35, 0x31, 0x36, 0x31, 0x35, 0x33, 0x36, 0x30, 0x32, 0x7c, 0x33, 0x34, 0x30, 0x30, 0x31, 0x36, 0x33, 0x37}
	msg = createMessage(msg)
	send(s, msg)

  for !waitForResp(s) {
		send(s, msg)
	}
  // Fin

  // Inicio
  msg = []byte("1H|\\^&|||SAT||||||P|E 1394-97|20080731103023")
  msg = createMessage(msg)
	send(s, msg)

  for !waitForResp(s) {
		send(s, msg)
	}
  // Fin

  // Inicio
	msg = []byte{EOT}
	send(s, msg)

  // Fin
	for !waitForResp(s) {
		send(s, msg)
	}
}

func createMessage(data []byte) []byte {
	//<STX><Frame Data><CR><ETX><CHECKSUM 1><CHECKSUM 2><CR><LF>
  //<Frame Data> = <Frame Number><Data>
	msg := []byte{}
	msg = append(msg, data...)
	msg = append(msg, []byte{CR, ETX}...)
	cs := checkSumASCII(checkSum8Mod256(msg))
	msg = append(msg, []byte{cs[0], cs[1], CR, LF}...)
	msg = append([]byte{STX}, msg...)
	return msg
}

func send(s *serial.Port, msg []byte) {
	_, err := s.Write(msg)
	if err != nil {
		log.Fatal(err)
	}
	printASTMMessage(msg)
}

func printASTMMessage(msg []byte) {
	for i := 0; i < len(msg); i++ {
		hexString := ""
		switch {
		case msg[i] == ENQ:
			hexString = "ENQ"
		case msg[i] == ACK:
			hexString = "ACK"
		case msg[i] == NAK:
			hexString = "NAK"
		case msg[i] == STX:
			hexString = "STX"
		case msg[i] == ETX:
			hexString = "ETX"
		case msg[i] == CR:
			hexString = "CR"
		case msg[i] == LF:
			hexString = "LF"
		case msg[i] == EOT:
			hexString = "EOT"
		case '0' <= msg[i] && msg[i] <= '9' || 'a' <= msg[i] && msg[i] <= 'z' || 'A' <= msg[i] && msg[i] <= 'Z':
			hexString = string(msg[i])
		default:
			hexString = "0x" + strings.ToUpper(hex.EncodeToString([]byte{msg[i]}))
		}
		fmt.Print(hexString + " ")
	}
	fmt.Println("")
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

func waitForResp(s *serial.Port) bool {
	buf := make([]byte, 128)
	n, err := s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	if reflect.DeepEqual(buf[:n], []byte{ACK}) {
		return true
	} else if reflect.DeepEqual(buf[:n], []byte{NAK}) {
		return false
	} else {
		return false
	}
}
