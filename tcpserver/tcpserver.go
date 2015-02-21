package main

import (
    "net"
    "fmt"
    "os"
)

const PORT string = ":1234"

func main() {

    listen, err := net.Listen("tcp", PORT)
    if err != nil {
        // exit if connection fails
        fmt.Println("Error listening: ", err.Error())
        os.Exit(1)
    }

    fmt.Println("Listening...")
    defer listen.Close()
    for {

        // listen and wait...
        conn, err := listen.Accept()
        if err != nil {
           fmt.Print("Error connecting: ", err.Error())
        }

        // log the message
        fmt.Printf("Received message %s -> %s \n",
            conn.RemoteAddr(),
            conn.LocalAddr())

        // handle incoming connections
        go handleConnection(conn)
    }
}

// processes the connection
// Receive byte slice, send strings
func handleConnection(conn net.Conn)  {

    // buffer is the input from the connection
    buffer := make([]byte, 1024)
    defer conn.Close()
    _, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error recieving", err.Error())
    }

    // convert incoming message to string
    clientMsg := []byte("Message recieved\n")

    conn.Write(clientMsg)
    fmt.Printf("%s says %s", conn.RemoteAddr(), string(buffer))
}
