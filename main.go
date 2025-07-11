package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:4000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	instructions := []uint32{
		0x00510113,
		0x00100093,
		0x00208133,
	}

	for _, instruction := range instructions {

		buf := binary.LittleEndian.AppendUint32(nil, instruction)
		fmt.Println(buf)
		result, err := conn.Write(buf)

		if err != nil {
			log.Println(err)
		}
		if result != 4 {
			log.Printf("Warning: only %d bytes sent (expected 4)", result)
		}

		log.Printf("Send instruction 0x%08x (%d bytes):\n", instruction, result)

	}

}
