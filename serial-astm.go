package main

import (
	"fmt"
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
	// Buffer de la transmision/trasferencia (entre ENQ y EOT)
	Transferencia bytes.Buffer
)

func main() {
	// Configuración de lectura del serial
	c := &serial.Config{Name: "COM2", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	// for {
	// Espera por la iniciacion del envio ENQ
	// Inicio
	for !waitForResp(s, []byte{ENQ}) {
		send(s, []byte{NAK})
	}
	send(s, []byte{ACK})
	// Fin

	// Espera por la el envio de datos o por un EOT
	// break if is EOT
  loopmsg:
	for {
		// Si no es valido retorna un NAK
		// for init; condition; post { }
		for v, e := waitForValidData(s); !v; v, e = waitForValidData(s) {
			// Si no es valido pero es un EOT sale del for
			if e {
				break loopmsg
			}
			// Si no es un EOT (fin de la transmisión) se envia un NAK
			send(s, []byte{NAK})
		}
		send(s, []byte{ACK}) // ACK of Valid Data
	}
	send(s, []byte{ACK}) // ACK of EOT
  fmt.Println("Transferencia terminada: ")
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
		//log.Println(buf[:n])
		return true
	} else {
		return false
	}
}

func read(s *serial.Port) []byte {
	buf := make([]byte, 128)
	n, err := s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	msg := buf[:n]
	return msg
}

func waitForValidData(s *serial.Port) (valid bool, eot bool) {
	//<STX><Frame Data><CR><ETX><CHECKSUM 1><CHECKSUM 2><CR><LF>
	//<Frame Data> = <Frame Number><Data>
	// Lee el dato
	msg := read(s)
	// Valida si es un EOT
	if bytes.Equal(msg, []byte{EOT}) {
		return false, true
	}
	// Valida si es un mensaje cortado (puede pasar de a un byte o varios bytes)
	// Un mensaje cortado comienza con STX pero no termina con LFs
	if msg[0] == STX {
		//fmt.Println("msg[0] == STX")
		// Si el ultimo caracter del slide es LF
		if msg[len(msg)-1] == LF {
			//fmt.Println("msg[len(msg)-1] == LF")
			// Continua, el mensaje es correcto
		} else {
			//fmt.Println("msg[len(msg)-1] != LF")
			// Buffer para la linea (entre STX y LF), make([]T, len, cap)
			tempBuffer := make([]byte, 0, 256)
			// Se inicializa el buffer con el contenido de msg
			tempBuffer = append(tempBuffer, msg...)
			// Mientras que el mensaje no termine en LF se guarda en tempBuffer
			for {
				newmsg := read(s)
				tempBuffer = append(tempBuffer, newmsg...)
				if(newmsg[len(newmsg)-1] == LF) {
					break
				}
			}
			msg = nil // clear
			msg = tempBuffer
		}
	}
	fmt.Println("Mensaje Valido: ")
	printASTMMessage(msg)
	//panic("EXIT")
	datamsg := searchBetween(msg, []byte{STX}, []byte{CR})
	datacs := searchBetween(msg, []byte{ETX}, []byte{CR, LF})
	// Se completa con el mensaje para el checksum
	data4cs := append(datamsg, []byte{CR, ETX}...)
	cs := checkSumASCII(checkSum8Mod256(data4cs))
	//log.Println(datacs, cs)
	// Es nulo cuando no encontro el mensaje entre los limites especificados
	// El mensaje esta corrupto
	if datamsg == nil || datacs == nil {
		return false, false
	} else if bytes.Equal(datacs, cs) {
		Transferencia.Write(datamsg)
    // Salto de linea CRLF despues de cada mensaje
    Transferencia.Write([]byte{CR, LF})
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
