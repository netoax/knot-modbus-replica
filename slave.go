package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/tbrandon/mbserver"
)

func uint16ToBytes(i uint16) []byte {
	var buf = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, i)
	return buf
}

func main() {
	fmt.Println("Running Modbus simulator for KNoT platform")
	serv := mbserver.NewServer()

	serv.RegisterFunctionHandler(3,
		func(s *mbserver.Server, frame mbserver.Framer) ([]byte, *mbserver.Exception) {
			data := frame.GetData()
			register := int(binary.BigEndian.Uint16(data[0:2]))
			numRegs := int(binary.BigEndian.Uint16(data[2:4]))
			endRegister := register + numRegs

			if endRegister > 65536 {
				return []byte{}, &mbserver.IllegalDataAddress
			}

			rand.Seed(time.Now().UnixNano())
			min := 17
			max := 25
			temp := rand.Intn(max-min+1) + min

			return append([]byte{byte(numRegs * 2)}, uint16ToBytes(uint16(temp))...), &mbserver.Success
		})

	err := serv.ListenTCP(":1502")
	if err != nil {
		log.Printf("%v\n", err)
	}
	defer serv.Close()

	// Wait forever
	for {
		time.Sleep(1 * time.Second)
	}
}
