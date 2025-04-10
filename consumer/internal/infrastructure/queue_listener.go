package infrastructure

import (
    "encoding/json"
    "fmt"
    "go-api-ddd/consumer/internal/domain"
    "go-api-ddd/consumer/internal/application"
    "net"
)

func StartListener() {
    ln, err := net.Listen("tcp", ":5672")
    if err != nil {
        panic(err)
    }
    fmt.Println("Consumer listening on :5672")

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Accept error:", err)
            continue
        }
        go handleConn(conn)
    }
}

func handleConn(conn net.Conn) {
    defer conn.Close()
    buf := make([]byte, 1024)
    n, err := conn.Read(buf)
    if err != nil {
        fmt.Println("Read error:", err)
        return
    }

    var msg domain.Message
    if err := json.Unmarshal(buf[:n], &msg); err != nil {
        fmt.Println("Unmarshal error:", err)
        return
    }

    if err := application.ProcessMessage(msg); err != nil {
        fmt.Println("DB insert error:", err)
    }
}
