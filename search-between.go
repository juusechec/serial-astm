package main

import (
   "bytes"
   "fmt"
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
  //https://www.socketloop.com/references/golang-bytes-trimleft-function-example
  //http://stackoverflow.com/questions/26916952/go-retrieve-a-string-from-between-two-characters-or-other-strings
  str := []byte("<h1>Hello World!</h2>")
  out := bytes.TrimLeft(bytes.TrimRight(str,"</h2>"),"<h1>")
  //fmt.Println(out)
  //fmt.Println(string(out))

  str = []byte{STX, 0x31, 0x32, 0x33, CR, ETX, 0x35, 0x36, 0x37, CR, LF}
  ini := bytes.Index(str, []byte{STX}) + 1
  end := bytes.Index(str, []byte{CR})
  out = str[ini:end]
  // fmt.Println(str)
  // fmt.Println(out)
  // fmt.Println(string(out))

  ini = bytes.Index(str, []byte{ETX}) + 1
  end = bytes.Index(str, []byte{CR, LF})
  out = str[ini:end]
  fmt.Println(str)
  fmt.Println(out)
  fmt.Println(string(out))
}
