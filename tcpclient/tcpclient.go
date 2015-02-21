package main
import (
    "net"
    "fmt"
    "os"
    "bufio"
)

const PORT string = ":1234"

func main() {
    conn, err := net.Dial("tcp", PORT)
    if err != nil {
        fmt.Println("Error dialing: ", err.Error())
        os.Exit(1)
    }

    defer conn.Close()
    fmt.Fprintf(conn, "I'm a client")
    status, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        fmt.Println("Error: ", err.Error())
    }

    fmt.Println(status)
}