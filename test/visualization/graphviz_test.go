package visualization_test

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	treegraph "tree-document-topology-visualizer/TREE-graph"
	"tree-document-topology-visualizer/visualization"

	"github.com/goccy/go-graphviz"
)

const FILE_PATH = "./temp"
const (
	rawOperator = "https://w3id.org/tree#GreaterThanRelation"
	value       = "\"%v-08-07T08:08:21.000Z\"^^<http://www.w3.org/2001/XMLSchema#dateTime>"
)

func aGraph() *treegraph.Graph {
	return &treegraph.Graph{
		treegraph.Node{Url: "foo"}:  []treegraph.Relation{},
		treegraph.Node{Url: "bar3"}: []treegraph.Relation{},
		treegraph.Node{Url: "bar"}:  []treegraph.Relation{{RawOperator: rawOperator, Literal: fmt.Sprintf(value, "2022"), Destination: treegraph.Node{Url: "foo"}}},
		treegraph.Node{Url: "foo2"}: []treegraph.Relation{{RawOperator: rawOperator, Literal: fmt.Sprintf(value, "2023"), Destination: treegraph.Node{Url: "bar3"}}},
		treegraph.Node{Url: "bar2"}: []treegraph.Relation{{RawOperator: rawOperator, Literal: fmt.Sprintf(value, "1993"), Destination: treegraph.Node{Url: "foo2"}}},
	}
}

func tearDown() {

	if err := os.RemoveAll(FILE_PATH); err != nil {
		log.Fatal(err)
	}
}

func setup() {
	if err := os.Mkdir(FILE_PATH, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

// / For the moment we are not able to get the edges, but we can validate their count.
func TestNewGraphvizTreeVisualizerWithANonEmptyGraph(t *testing.T) {
	visualizer, err := visualization.NewGraphvizTreeVisualizer(*aGraph())
	if err != nil {
		t.Fatal(err)
	}

	graph := visualizer.(visualization.GraphvizTreeVisualizer).Graph()
	defer graph.Close()

	if nNode := graph.NumberNodes(); nNode != len(*aGraph()) {
		t.Errorf("The graphViz graph has {%v} nodes whereas it should have {%v} nodes", nNode, len(*aGraph()))
	}
	nRelations := 0

	for node, relations := range *aGraph() {
		cNode, err := graph.Node(node.Url)
		nRelations += len(relations)

		if err != nil {
			t.Error(err)
		}

		if name := cNode.Name(); name != node.Url {
			t.Errorf("The name of the cNode doesn't correspond to the graph node.\nThe cNode has the name {%v} whereas the graph node has {%v}", name, node.Id())
		}

		if url := cNode.Get("URL"); url != node.Url {
			t.Errorf("The url of the cNode doesn't correspond to the graph node.\nThe cNode has the url {%v} whereas the graph node has {%v}", url, node.Url)
		}
	}

	if nEdges := graph.NumberEdges(); nEdges != nRelations {
		t.Errorf("The graphViz graph has {%v} edges whereas it should have {%v} edges", nEdges, nRelations)
	}
}

func TestNewGraphvizTreeVisualizerWithAnEmptyGraph(t *testing.T) {
	visualizer, err := visualization.NewGraphvizTreeVisualizer(treegraph.Graph{})
	if err != nil {
		t.Fatal(err)
	}

	graph := visualizer.(visualization.GraphvizTreeVisualizer).Graph()
	defer graph.Close()

	if nNode := graph.NumberNodes(); nNode != 0 {
		t.Errorf("The graphViz graph has {%v} nodes whereas it should have {%v} nodes", nNode, 0)
	}

	if nEdges := graph.NumberEdges(); nEdges != 0 {
		t.Errorf("The graphViz graph has {%v} edges whereas it should have {%v} edges", nEdges, 0)
	}
}

func TestGetGenerateFile(t *testing.T) {
	baseName := "graph"
	fileExtensions := []graphviz.Format{
		graphviz.JPG,
		graphviz.PNG,
		graphviz.SVG,
		graphviz.XDOT,
	}
	for _, fileExtension := range fileExtensions {
		visualizer, err := visualization.NewGraphvizTreeVisualizer(*aGraph())
		if err != nil {
			t.Fatal(err)
		}
		graphPath := filepath.Join("./", FILE_PATH, fmt.Sprintf("%v.%v", baseName, fileExtension))

		if err := visualizer.GenerateFile(graphPath); err != nil {
			t.Fatal(err)
		}

		if _, err := os.Stat(graphPath); err == nil {

		} else if errors.Is(err, os.ErrNotExist) {
			t.Error(err)
		} else {
			t.Error(err)
		}
	}

}
