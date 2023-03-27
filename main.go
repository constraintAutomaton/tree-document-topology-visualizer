package main

import (
	"fmt"
	"log"
	treegraph "tree-document-topology-visualizer/TREE-graph"
	"tree-document-topology-visualizer/communication"
	"tree-document-topology-visualizer/visualization"
)

func main() {
	fmt.Println("Hello, World!")
	queryOutput, err := communication.GetTreeRelation("http://localhost:3000/ldes/test")
	if err != nil {
		log.Fatal(err.Error())
	}
	graph := treegraph.NewGraphFromSparlRelationOutputs(queryOutput)
	visualization.GenerateDotFileFromTreeGraph(graph)

}
