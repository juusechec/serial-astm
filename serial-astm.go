// You can edit this code!
// Click here and start typing.
package main

import "fmt"

const (
	ENQ = 0x05
	ACK = 0x06
	NAK = 0x15
	STX = 0x02
	ETX = 0x03
	CR  = 0x0D
	LF  = 0x0A
	EOT = 0x04
)

func main() {
	var entradaEjemplo = "\xbd\xbd\xbd\xbd\xbd\x05\xbd\xbd\xbd\xbd\xbd\xb2\x3d\xbc\x20\xe2\x8c\x0D\x0A\x98\x98\x98\x98\x98\x98"

	buf := make([]byte, 40)
	bufUlimos2 := [2]uint8{0x00, 0x00}
	lecturaIniciada := false
	for i := 0; i < len(entradaEjemplo); i++ {
		var entrada = entradaEjemplo[i]
		if entrada == ENQ || lecturaIniciada == true {
			lecturaIniciada = true
			bufUlimos2[0] = bufUlimos2[1]
			bufUlimos2[1] = entrada
			fmt.Println(bufUlimos2)
			finDeLinea := [2]uint8{CR, LF}
			if bufUlimos2 == finDeLinea {
				fmt.Println("Terminó linea")
				break
			}
			buf[i] = entrada
		}
	}
	fmt.Println(buf)
}
