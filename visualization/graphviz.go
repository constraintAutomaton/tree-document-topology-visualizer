package visualization

import (
	"bytes"
	"log"
	"os"
	treegraph "tree-document-topology-visualizer/TREE-graph"

	graphviz "github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

func GenerateDotFileFromTreeGraph(treeGraph treegraph.Graph) {
	g := graphviz.New()
	nodeRegistry := map[treegraph.Node]*cgraph.Node{}
	graph, err := g.Graph()
	graph.SetRankDir(cgraph.RLRank)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()

	for node := range treeGraph {
		n, err := graph.CreateNode(node.Url)
		if err != nil {
			log.Fatal(err)

		}
		nodeRegistry[node] = n
	}

	for node, relations := range treeGraph {
		for _, relation := range relations {
			e, err := graph.CreateEdge(relation.Equation(), nodeRegistry[node], nodeRegistry[relation.Destination])
			if err != nil {
				log.Fatal(err)
			}
			e.SetLabel(relation.Equation())
		}
	}

	var buf bytes.Buffer
	if err := g.Render(graph, "dot", &buf); err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("./graph.gv") // creates a file at current directory
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

}
