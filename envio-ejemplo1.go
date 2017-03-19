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
	msg = []byte("1H|\\^&|||SAT||||||P|E 1394-97|20080731103023")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("2P|1||Paid127||Name_127^F-Name_127||19800127|M|||||Phy_127||||||||||||Dep_127|")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("3O|1|AUTOSID127||^^^LMG|||080731103023|080731103023||||||||||||||||||F||||")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("4C|1|alarm^^^G1|I")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("5R|1|^^^MPV^776-5|7,6|1||||F||||20080731103023|")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("6R|2|^^^PDW^X-PDW|12,9|1||||F||||20080731103023|")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("7R|3|^^^PLT^777-3|230|1||||F||||20080731103023|")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("0C|1|I|curve^PLT^0^63^00000000030A0F151B21262D3135363636333333322E2F2D2B292726221F1F1E1B1B1A17171613120F0F0E0E0D0B0A0A0A090707060505050505030303030303|G")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("1C|2|I|curve^PLT^64^127^03030303030302020202020201010101010101010101010101010001010100010101000001000000010101010101010101010100000000000000000000010101|G")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("2C|3|I|threshold^PLT^69|G")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("3R|4|^^^THT^X-PCT|0,175|1||||F||||20080731103023|")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("4R|5|^^^HCT^4544-3|43,6|1||||F||||20080731103023|")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("5R|6|^^^HGB^717-9|14,4|1||||F||||20080731103023|")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("6R|7|^^^MCH^785-6|32,8|1||H||F||||20080731103023|")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("7R|8|^^^MCHC^786-4|33,0|1||||F||||20080731103023|")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("0R|9|^^^MCV^787-2|99|1||||F||||20080731103023|")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("1R|10|^^^RBC^789-9|4,40|1||||F||||20080731103023|")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("2C|1|I|curve^RBC^0^63^00000000000000000000000001030102010102010201020305060B121A2338464F758996B1BAD8DFCFDCD2C1BEAA9B9F788170575C4B403A36231F1A1F181612|G")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("3C|2|I|curve^RBC^64^127^130F110D0E0B0F0A090908080A0A0A0C080707070605060704050405030304030203030100010101010100000000010000000000000000000000000000000006|G")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("4R|11|^^^RDW^788-0|13,5|1||||F||||20080731103023|")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("5R|12|^^^GRA#^20482-6|8,60|1||||F||||20080731103023|")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("6R|13|^^^GRA%^14773-6|91,9|1||H||F||||20080731103023|")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("7R|14|^^^LYM#^731-0|0,40|1||L||F||||20080731103023|")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("0R|15|^^^LYM%^736-9|5,3|1||L||F||||20080731103023|")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("1R|16|^^^MON#^742-7|0,20|1||||F||||20080731103023|")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("2R|17|^^^MON%^744-3|2,8|1||||F||||20080731103023|")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("3R|18|^^^WBC^804-5|9,2|1||||F||||20080731103023|")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("4C|1|I|curve^WBC^0^63^000000000000000214324441342B241B16120F0F0F12120B0B09090B0D0F1626323F4B5F768191A6B8BCBCBAB6A89D8A7D64514B443F342D282D3434343B4856|G")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("5C|2|I|curve^WBC^64^127^667883838F96B1C1D3D5DFD8D8D5DCD8D3C8C1B6BCB3AA9A938D83746856544841342B1F1B191916120D0D090909060404040202000204040402020202020228|G")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("6C|3|I|threshold^WBC^00^00^00^19^22|G")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte("7L|1|N")
	msg = createMessage(msg)
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
	// Fin

	// Inicio
	msg = []byte{EOT}
	send(s, msg)

	for !waitForResp(s) {
		send(s, msg)
	}
  // Fin

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
