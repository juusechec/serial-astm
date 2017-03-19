package main

import (
	//"fmt"
	"bytes"
	"encoding/hex"
	"github.com/tarm/serial"
	"log"
	"os"
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

var (
	Transferencia bytes.Buffer
)

func main() {
	c := &serial.Config{Name: "COM2", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	// for {
	// Espera por la iniciacion del envio ENQ
	for !waitForResp(s, []byte{ENQ}) {
		send(s, []byte{NAK})
	}
	send(s, []byte{ACK})

	// Espera por la el envio de datos o por un EOT
	// break if is EOT
	for {
		isEOT := false
    // Si no es valido retorna un NAK
    // for init; condition; post { }
		for v, e := waitForValidData(s); !v; v, e = waitForValidData(s){
			isEOT = e
      // Si no es valido pero es un EOT sale del for
			if e {
				break
			}
			send(s, []byte{NAK})
		}
    // Si es un EOT termina el mensaje
		if isEOT {
			break
		}
		send(s, []byte{ACK})
	}
	send(s, []byte{ACK}) // ACK of EOT
	Transferencia.WriteTo(os.Stdout)
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
	buf := make([]byte, 256)
	n, err := s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	if reflect.DeepEqual(buf[:n], resp) {
    log.Println(buf[:n])
		return true
	} else {
		return false
	}
}

func waitForValidData(s *serial.Port) (valid bool, eot bool) {
  log.Println("Llegue")
	buf := make([]byte, 128)
	n, err := s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	msg := buf[:n]
  log.Println(string(msg))
	if bytes.Equal(msg, []byte{EOT}) {
		return false, true
	}
	datamsg := searchBetween(msg, []byte{STX}, []byte{CR})
	datacs := searchBetween(msg, []byte{ETX}, []byte{CR, LF})
  // Se completa con el mensaje para el checksum
	data4cs := append(datamsg, []byte{CR, ETX}...)
	cs := checkSumASCII(checkSum8Mod256(data4cs))
  log.Println(datacs, cs)
  // Es nulo cuando no encontro el mensaje entre los limites especificados
  // El mensaje esta corrupto
	if datamsg == nil || datacs == nil {
		return false, false
	} else if bytes.Equal(datacs, cs) {
		Transferencia.Write(datamsg)
		return true, false
	} else {
		return false, false
	}
}

func searchBetween(text []byte, ini []byte, end []byte) []byte {
	iniI := bytes.Index(text, ini) + 1
	endI := bytes.Index(text, end)
  //log.Println(text, ini, end)
	if iniI == 0 || endI == -1 {
		return nil
	}
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
