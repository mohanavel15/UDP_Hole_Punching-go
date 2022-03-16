package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func server() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":5000")
	HandleError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	HandleError(err)

	var clients []*net.UDPAddr
	var nodes Nodes

	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		HandleError(err)
		fmt.Println(addr.String(), " : ", string(buffer[0:n]))

		is_in_clients := false
		for _, client := range clients {
			if client.String() == addr.String() {
				is_in_clients = true
				break
			}
		}

		if !is_in_clients {
			new_node := Node{
				IP: addr.String(),
			}

			if len(clients) > 0 {
				ncli_msg := Message{
					Type: "new_node",
					Data: new_node,
				}

				ncli_data, err := json.Marshal(ncli_msg)
				HandleError(err)

				for _, client := range clients {
					conn.WriteToUDP(ncli_data, client)
				}
			}

			if len(nodes.Nodes) > 0 {
				nl_msg := Message{
					Type: "node_list",
					Data: nodes,
				}

				nl_data, err := json.Marshal(nl_msg)
				HandleError(err)

				conn.WriteToUDP(nl_data, addr)
			}

			clients = append(clients, addr)
			nodes.Nodes = append(nodes.Nodes, new_node)
			fmt.Println("New client: ", addr.String())
		}
	}
}
