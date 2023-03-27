package main

import (
	"fmt"
	"log"
	treegraph "tree-document-topology-visualizer/TREE-graph"
	"tree-document-topology-visualizer/communication"
	"tree-document-topology-visualizer/visualization"
)

func main() {
	treeDocument := "https://treecg.github.io/demo_data/vtmk.ttl" //"http://localhost:3000/ldes/test"
	fmt.Println("Starting")
	queryOutput, err := communication.GetTreeRelation(treeDocument)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Query executed")
	graph := treegraph.NewGraphFromSparlRelationOutputs(queryOutput)
	fmt.Printf("%+v", len(queryOutput))
	fmt.Println("Graph of the TREE document constituted")
	visualization.GenerateDotFileFromTreeGraph(graph)
	fmt.Println("Visualization generated")

}
