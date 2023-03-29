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

	fmt.Printf("Starting to get the relation of {%v}\n", treeDocumentUrl)
	queryOutput, err := communication.GetTreeRelation(treeDocumentUrl, limit)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Relation fetched")
	fmt.Println("Starting to generate a graph from the TREE relation")
	graph := treegraph.NewGraphFromSparlRelationOutputs(queryOutput)
	fmt.Println("Graph of the TREE document constituted")
	fmt.Printf("Starting to generate the graph that will be outputed at path {%v}\n", graphPath)
	visualization.GenerateGraphvizGraph(graph, graphPath)
	fmt.Println("Visualization generated")
	fmt.Println("Closing of the program, keep on living!")

}

func cliParseCliArgs() {
	flag.StringVar(&treeDocumentUrl, "t", "http://localhost:3000/ldes/test", "URL of the TREE document")
	flag.StringVar(&treeDocumentUrl, "tree-document-url", "http://localhost:3000/ldes/test", "URL of the TREE document")

	flag.StringVar(&graphPath, "p", "./generated/graph.svg", "Resulting path of the graph")
	flag.StringVar(&graphPath, "path", "./generated/graph.svg", "Resulting path of the graph")

	flag.UintVar(&limit, "l", math.MaxUint, "The maximum number of relation")
	flag.UintVar(&limit, "limit", math.MaxUint, "The maximum number of relation")
	flag.Parse()

}
