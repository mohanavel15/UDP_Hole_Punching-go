package main

import "fmt"

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type Node struct {
	IP string `json:"ip"`
}

type Nodes struct {
	Nodes []Node `json:"nodes"`
}

func HandleError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
}
