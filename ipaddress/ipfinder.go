package main

import (
    "fmt"
    "net"
)

// A program that will print the IP Address for a given host name
// Returns error if the host name does not exist
func main() {

    var host string // var to hold given address
    fmt.Print("Enter a host name:\t")
    fmt.Scanln(&host)

    // output address is assigned
    if address, err := net.LookupIP(host); err != nil {
        fmt.Printf("hostname %s does not exist", host)
    } else {
        fmt.Printf("The IP addresses for %s are:\n", host)
        for _, ipAddress := range address{
            fmt.Println(ipAddress)
        }
    }
}

