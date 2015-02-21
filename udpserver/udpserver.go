package main

import (
    "net"
    "fmt"
    "os"
)

const PORT string = "127.0.0.1:1234"

func main() {

    address, err := net.ResolveUDPAddr("udp", PORT)
    if err != nil {
        fmt.Println(err.Error())
    }

    listen, err := net.ListenUDP("udp", address)
    if err != nil {
        // exit if connection fails
        fmt.Println("Error listening: ", err.Error())
        os.Exit(1)
    }

    fmt.Println("Listening...")
    defer listen.Close()
    for {

        // listen and wait
        _, conn, err := listen.ReadFromUDP([]byte("\n"))
        if err != nil {
            fmt.Println(err.Error())
        }

        // log message
        fmt.Printf("Recieved message from %s \n",
            conn.String())

        // handle connection
        go handleConnection(listen, conn)
    }
}

func handleConnection(listen *net.UDPConn, conn *net.UDPAddr ){
    _, err := listen.WriteToUDP([]byte("Message recieved"), conn)
    if err != nil {
        fmt.Println(err.Error())
    }
}