package main

import (
	"fmt"
	"log"
	"tree-document-topology-visualizer/communication"
)

func main() {
	fmt.Println("Hello, World!")
	r, err := communication.GetTreeRelation("http://localhost:3000/ldes/test")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(r)
}
