package main

import (
	"fmt"
	"log"
	"strings"

	"go.bug.st/serial"
)

func main() {
	test()
}
func test() {

	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}

	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}
	port, err := serial.Open(ports[0], mode)
	if err != nil {
		log.Fatal(err)
	}

	n, err := port.Write([]byte("\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sent %v bytes\n", n)

	buf := make([]byte, 4096)
	for {

		n, err := port.Read(buf)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n == 0 {
			fmt.Println("\nEOF")
			break
		}

		resp := string(buf[:n])
		fmt.Printf("%v", resp)

		if strings.Index(resp, "login:") != -1 &&
			strings.Index(resp, "Last login:") == -1 {
			port.Write([]byte("pi\n"))
		} else if strings.Index(resp, "Password:") != -1 {
			port.Write([]byte("raspberry\n"))
		} else if strings.Index(resp, ":~$ ") != -1 {
			port.Write([]byte("exit\n"))
		}
	}

}
