package main

import (
	"fmt"
	"net"
	"proxyclient2/Cypher"
)

func main() {

	//proxy client #2 address
	fullAddress := "localhost:5800"

	// Listen for incoming connections.
	l, err := net.Listen("tcp", fullAddress)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + fullAddress)
	for {
		// Listen for an incoming connection.
		connection, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
		}
		// Handle connections in a new function.
		go handleRequest(connection)
	}
}

// Handles incoming requests.
func handleRequest(connection net.Conn) {
	// Make a buffer to hold incoming data.
	buffer := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	n, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	received := string(buffer[:n])

	// print out read contents
	fmt.Println("Received" + " -> " + received)
	// Send a response back to person contacting us.
	connection.Write([]byte("received message"))

	// message decryption

	receivedChar := []rune(received)

	for index := 0; index < len(receivedChar); index++ {
		receivedChar[index] = Cypher.DecryptLetter(receivedChar[index])
	}

	fmt.Println("Message decrypted with +4, new message: \n" + string(receivedChar))

	connection.Close()
}
