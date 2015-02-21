package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    address, err := net.InterfaceAddrs()
    if err != nil {
        fmt.Println("Oops: ", err.Error())
        os.Exit(1)
    }

    for _, addr := range address {
        if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                fmt.Println("Your Local IP Address is ", ipnet.IP.String())
            }
        }
    }
}
