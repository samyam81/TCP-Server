# TCP-Server 

This repository contains a simple TCP server implemented in Go (`echoserver.go`). The server listens on a specified port (`8080` by default) and echoes back any data it receives from clients.

## Usage

1. Clone the repository:

   ```bash
   git clone https://github.com/samyam81/TCP-Server.git
   cd TCP-Server
   ```

2. Run the server:

   ```bash
   go run echoserver.go
   ```

   By default, the server will listen on port `8080`.

3. Connect to the server using a TCP client (e.g., Telnet or netcat):

   ```bash
   telnet localhost 8080
   ```

4. Send data from the client, and the server will echo it back.

5. To stop the server, press `Ctrl + C`.

## Implementation Details

- The server uses Go's `net` package to listen for incoming TCP connections.
- Each incoming connection is handled concurrently using goroutines.
- The server reads data from clients and immediately writes it back (echoes) to the client.

## Example

Assuming the server is running on `localhost` with the default port (`8080`), you can test it using Telnet:

```bash
telnet localhost 8080
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
Hello, server!
Hello, server!
```

In this example, the server echoes back the message "Hello, server!" that was sent from the Telnet client.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
