package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	treegraph "tree-document-topology-visualizer/TREE-graph"
	"tree-document-topology-visualizer/communication"
	"tree-document-topology-visualizer/visualization"
)

var treeDocumentUrl string
var graphPath string
var limit uint

func main() {
	cliParseCliArgs()
	if limit == math.MaxUint {
		fmt.Printf("Starting to get the relations in the data source {%v}\n", treeDocumentUrl)
	} else {
		fmt.Printf("Starting to get the relations with a limit of %v in the data source {%v}\n", limit, treeDocumentUrl)
	}
	queryOutput, err := communication.GetTreeRelation(treeDocumentUrl, limit)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Relation %v fetched\n", len(queryOutput))
	fmt.Println("Starting to generate a graph from the TREE relation")
	graph := treegraph.NewGraphFromSparlRelationOutputs(queryOutput)
	fmt.Println("Graph of the TREE document constituted")
	fmt.Printf("Starting to generate the graph that will be outputed at path {%v}\n", graphPath)
	visualization.GenerateGraphvizGraph(graph, graphPath)
	fmt.Println("Visualization generated")
	fmt.Println("Closing of the program, keep on living.")

}

func cliParseCliArgs() {
	flag.StringVar(&treeDocumentUrl, "t", "http://localhost:3000/ldes/test", "URL of the TREE document")

	flag.StringVar(&graphPath, "p", "./generated/graph.svg", "Resulting path of the graph")

	flag.UintVar(&limit, "l", math.MaxUint, "The maximum number of relation")
	flag.Parse()

}
