package main
import (
    "net"
    "fmt"
    "os"
    "bufio"
)

const PORT string = "127.0.0.1:1234"

func main() {

    buffer := make([]byte, 2048)
    conn, err := net.Dial("udp", PORT)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    defer conn.Close()
    fmt.Fprintf(conn, "I'm a client")
    status, err := bufio.NewReader(conn).Read(buffer)
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println(status)
    }

    fmt.Println(string(buffer))
}
