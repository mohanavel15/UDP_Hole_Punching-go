# UDP hole punching in Go
UDP hole punching example in  go

## How it works
- server waits for clients to connect
- client 1 connects to server
- when client 2 joins 
- server sends client 1 info to client 2
- server sends client 2 info to client 1
- now client 1 & 2 can communicate without the server

### run server
```bash
$ ./UPU_Hole_Punching server
```
- Server runs on port 5000
- Server needs port forwarding

### run client
```bash
$ ./UPU_Hole_Punching client [server's public ip]:5000
```