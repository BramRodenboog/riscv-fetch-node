package main

import (
	"encoding/binary"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:4000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var instruction uint32 = 0x00510113

	buf := make([]byte, 4)

	binary.LittleEndian.PutUint32(buf, instruction)

	result, err := conn.Write(buf)

	if err != nil {
		log.Println(err)
	}
	if result != 4 {
		log.Printf("Warning: only %d bytes sent (expected 4)", result)
	}

	log.Println("Send instruction 0x%08x (%d btyes)\n: ", instruction, result)

}
