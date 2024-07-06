package main

import (
    "fmt"
    "net"
)

func main() {
    // Define the port on which the server will listen
    port := "8080"

    // Start listening on the specified port
    listener, err := net.Listen("tcp", ":"+port)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        return
    }
    defer listener.Close()
    fmt.Println("Server is listening on port", port)

    for {
        // Wait for a connection request from a client
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err.Error())
            return
        }

        // Handle connections concurrently using goroutines
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    // Close the connection when the function exits
    defer conn.Close()

    // Buffer to hold incoming data
    buffer := make([]byte, 1024)

    for {
        // Read data from the connection
        bytesRead, err := conn.Read(buffer)
        if err != nil {
            fmt.Println("Error reading:", err.Error())
            return
        }

        // Echo back to client
        _, err = conn.Write(buffer[:bytesRead])
        if err != nil {
            fmt.Println("Error writing:", err.Error())
            return
        }
    }
}
