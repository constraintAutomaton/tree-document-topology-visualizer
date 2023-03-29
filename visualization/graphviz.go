package visualization

import (
	"bytes"
	"log"
	"os"
	"path"
	"strings"
	treegraph "tree-document-topology-visualizer/TREE-graph"

	graphviz "github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

func GenerateGraphvizGraph(treeGraph treegraph.Graph, graphPath string) {
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
	if extension := path.Ext(graphPath); extension != "" {
		GenerateFile(g, graph, graphviz.Format(extension[1:]), graphPath)
	} else {
		log.Fatalln("The path of the output graph as no extension")
	}
}

func GenerateFile(instance *graphviz.Graphviz, graph *cgraph.Graph, format graphviz.Format, path string) {
	var buf bytes.Buffer
	if err := instance.Render(graph, format, &buf); err != nil {
		log.Fatal(err)
	}
	if format == graphviz.XDOT {
		f, err := os.Create(path) // creates a file at current directory
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		_, err = f.Write(buf.Bytes())
		if err != nil {
			log.Fatal(err)
		}
	} else {
		if err := instance.RenderFilename(graph, format, path); err != nil {
			log.Fatal(err)
		}
	}

}

func createNode(node treegraph.Node, graph *cgraph.Graph) *cgraph.Node {
	node.BuildId()
	n, err := graph.CreateNode(node.Url)
	if err != nil {
		log.Fatal(err)
	}
	n = n.SetLabel(node.Id())
	formateUrl := strings.ReplaceAll(node.Url, "&", "&amp;")
	n = n.SetURL(formateUrl)
	n = n.SetShape(cgraph.BoxShape)
	n = n.SetTarget("_blank")
	return n
}
