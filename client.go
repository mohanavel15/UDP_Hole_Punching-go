package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func client(ip string) {
	ServerAddr, err := net.ResolveUDPAddr("udp", ip)
	HandleError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:0")
	HandleError(err)

	Conn, err := net.ListenUDP("udp", LocalAddr)
	HandleError(err)

	Conn.WriteToUDP([]byte("Hello server!"), ServerAddr)

	buffer := make([]byte, 1024)

	for {
		n, addr, err := Conn.ReadFromUDP(buffer)
		HandleError(err)
		bin_data := buffer[0:n]

		var msg Message
		err = json.Unmarshal(bin_data, &msg)
		HandleError(err)

		data, err := json.Marshal(msg.Data)
		HandleError(err)

		if msg.Type == "new_node" {
			var node Node
			err := json.Unmarshal([]byte(data), &node)
			HandleError(err)

			fmt.Println("New node : ", node.IP)
			MessageNode(node.IP, Conn)
		}

		if msg.Type == "node_list" {
			var nodes Nodes
			err := json.Unmarshal([]byte(data), &nodes)
			HandleError(err)

			fmt.Println("Nodes list : ", nodes.Nodes)
			for _, node := range nodes.Nodes {
				MessageNode(node.IP, Conn)
			}
		}

		if msg.Type == "from_node" {
			fmt.Println(addr.String(), " : ", string(data))
		}
	}
}

func MessageNode(ip string, Conn *net.UDPConn) {
	node_addr, err := net.ResolveUDPAddr("udp", ip)
	HandleError(err)

	new_msg := Message{
		Type: "from_node",
		Data: "Hello from another node",
	}

	new_msg_data, err := json.Marshal(new_msg)
	HandleError(err)

	int_, err := Conn.WriteToUDP(new_msg_data, node_addr)
	HandleError(err)
	fmt.Println("Sent ", int_, " bytes to ", node_addr.String())
}
