package main

import (
	"fmt"
	"net"
	"os"
	"proxyclient/Cypher"
	"strings"
	"time"
)

func main() {
	// Cypher start
	HEMMELIGHET := "Møte i Ålesund 1. juni kl. 2"
	HEMMChar := []rune(strings.ToLower(HEMMELIGHET))
	fmt.Println("Message set to: " + HEMMELIGHET)
	// test for how many bytes the message consists of BEFORE encryption
	bytes := []byte(HEMMELIGHET)
	fmt.Println(len(bytes))

	for index := 0; index < len(HEMMChar); index++ {
		HEMMChar[index] = Cypher.EncryptLetter(HEMMChar[index])
	}

	fmt.Println("Message encrypted with +4, new message: \n" + string(HEMMChar))
	// test for how many bytes the message consists of AFTER encryption
	encBytes := []byte(string(HEMMChar))
	fmt.Println(len(encBytes))

	//Cypher end

	strEcho := string(HEMMChar)
	servAddr := "localhost:5000"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write([]byte(strEcho))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("write to server = ", strEcho)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)

	if err != nil {
		println("Read from server failed:", err.Error())
		os.Exit(1)
	}
	time.Sleep(2 * time.Second)
	println("reply from server=", string(reply))

	conn.Close()
}
