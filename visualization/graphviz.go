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
	graph = graph.SetRankDir(cgraph.LRRank)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()

	for node, relations := range treeGraph {
		if _, exist := nodeRegistry[node]; !exist {
			n := createNode(node, graph)
			nodeRegistry[node] = n

		}

		for _, relation := range relations {
			var destinationNode *cgraph.Node
			if n, exist := nodeRegistry[relation.Destination]; exist {
				destinationNode = n

			} else {
				n := createNode(relation.Destination, graph)
				nodeRegistry[relation.Destination] = n
				destinationNode = n
			}
			e, err := graph.CreateEdge(relation.Equation(), nodeRegistry[node], destinationNode)
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

func createNode(node treegraph.Node, graph *cgraph.Graph) *cgraph.Node {
	node.BuildId()
	n, err := graph.CreateNode(node.Url)
	if err != nil {
		log.Fatal(err)
	}
	n = n.SetLabel(node.Id())
	n = n.SetURL(node.Url)
	n = n.SetShape(cgraph.BoxShape)
	//n = n.SetTarget()
	return n
}
