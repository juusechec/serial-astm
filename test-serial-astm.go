package main

import (
	"fmt"
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
	c := &serial.Config{Name: "COM3", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 512)
	buf := make([]byte, 512)
	tamano := 0
	bufUlimos2 := [2]uint8{0x00, 0x00}
	lecturaIniciada := false
	for {
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		//log.Printf("%q", buf[:n])
		//fmt.Println(buf[:n])
		j := tamano
		for i := 0; i < n; i++ {
			entrada := buf[i]
			if lecturaIniciada == true {
				bufUlimos2[0] = bufUlimos2[1]
				bufUlimos2[1] = entrada
				finDeLinea := [2]uint8{CR, LF}
				//finDeLinea := [2]uint8{AAA, CR}
				//fmt.Println(bufUlimos2, finDeLinea)
				if bufUlimos2 == finDeLinea {
					fmt.Println("Terminó linea")
					log.Printf("%q", buffer[:tamano-1])
					lecturaIniciada = false
					procesarCadena(buffer[:tamano])
					break
				}
				buffer[j] = entrada
			}
			j++
			if entrada == ENQ {
				fmt.Println("Inició linea")
				_, err := s.Write([]byte{ACK})
				if err != nil {
					log.Fatal(err)
				}
				lecturaIniciada = true
			}
		}

		if lecturaIniciada == true {
			tamano = tamano + n
		}

		//log.Printf("%q", buffer[:tamano])
		//log.Println("indice", indice)
		//fmt.Println()
	}
}

func procesarCadena(cadena []byte) {
	strCadena := string(cadena[:])
	log.Println(strCadena)
}
