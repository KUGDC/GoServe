package main

import ("net"
        "fmt"
)

func handleConnection(connection net.Conn) {
    var output []byte = make([]byte, 1024)
    for {
        i, err := connection.Read(output)
        if err != nil {
            fmt.Print(err)
            break
        }
        if i > 0 {
            fmt.Printf("Read %i bytes\n", i)
            fmt.Println(output)
        }
        connection.Close()
    }
}

func main() {
    ln, err := net.Listen("tcp", ":8956")
    if err != nil {
        // handle error
        fmt.Print(err)
        return
    }
    for {
        fmt.Print("...")
        conn, err := ln.Accept()
        if err != nil {
            // handle error
            fmt.Print(err)
            continue
        }
        go handleConnection(conn)
    }
}
