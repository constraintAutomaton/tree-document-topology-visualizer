package main

import (
	"fmt"
	"log"
	treegraph "tree-document-topology-visualizer/TREE-graph"
	"tree-document-topology-visualizer/communication"
	"tree-document-topology-visualizer/visualization"
)

func main() {
	treeDocument := "https://demo.netwerkdigitaalerfgoed.nl/fragments/wo2/" //"http://localhost:3000/ldes/test"
	graphPath := "./generated/graph.svg"
	fmt.Println("Starting")
	queryOutput, err := communication.GetTreeRelation(treeDocument)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Query executed")
	graph := treegraph.NewGraphFromSparlRelationOutputs(queryOutput)
	fmt.Println("Graph of the TREE document constituted")
	visualization.GenerateGraphvizGraph(graph, graphPath)
	fmt.Println("Visualization generated")

}
