package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	//server address address
	fullAddress := "localhost:5000"

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
	connection.Write([]byte("Message received"))

	//connect to secondary client
	strEcho := received
	servAddr := "localhost:5800"
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

	println("sent to client #2 = ", strEcho)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)

	if err != nil {
		println("Read from server failed:", err.Error())
		os.Exit(1)
	}
	time.Sleep(2 * time.Second)
	println("reply from server = ", string(reply))

	connection.Close()
}
