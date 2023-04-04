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

var graphvizValidFormat map[graphviz.Format]bool = map[graphviz.Format]bool{
	graphviz.JPG:  true,
	graphviz.PNG:  true,
	graphviz.SVG:  true,
	graphviz.XDOT: true,
}

type GraphvizTreeVisualizer struct {
	instance *graphviz.Graphviz
	graph    *cgraph.Graph
}

func NewGraphvizTreeVisualizer(treeGraph treegraph.Graph) (Visualizer, error) {
	g := graphviz.New()
	nodeRegistry := map[treegraph.Node]*cgraph.Node{}
	graph, err := g.Graph()
	graph = graph.SetRankDir(cgraph.LRRank)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()

	for node, relations := range treeGraph {
		if _, exist := nodeRegistry[node]; !exist {
			n, err := createNode(node, graph)
			if err != nil {
				return nil, err
			}
			nodeRegistry[node] = n

		}

		for _, relation := range relations {
			var destinationNode *cgraph.Node
			if n, exist := nodeRegistry[relation.Destination]; exist {
				destinationNode = n

			} else {
				n, err := createNode(relation.Destination, graph)
				if err != nil {
					return nil, err
				}
				nodeRegistry[relation.Destination] = n
				destinationNode = n
			}
			e, err := graph.CreateEdge(relation.Equation(), nodeRegistry[node], destinationNode)
			if err != nil {
				return nil, err
			}
			e.SetLabel(relation.Equation())
		}
	}
	return GraphvizTreeVisualizer{
		instance: g,
		graph:    graph,
	}, nil
}

func (g GraphvizTreeVisualizer) GenerateFile(graphPath string) error {
	var buf bytes.Buffer
	var format graphviz.Format
	if extension := path.Ext(graphPath); extension != "" {
		format = graphviz.Format(extension[1:])
		if !isValidGraphvizFileFormat(format) {
			return GraphFilePathInvalidFormat{format: string(format)}
		}
	} else {
		return GraphFilePathNoExtension{Path: graphPath}
	}

	if err := g.instance.Render(g.graph, format, &buf); err != nil {
		return err
	}
	if format == graphviz.XDOT {
		f, err := os.Create(graphPath)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = f.Write(buf.Bytes())
		if err != nil {
			return err
		}
	} else {
		if err := g.instance.RenderFilename(g.graph, format, graphPath); err != nil {
			return err
		}
	}
	return nil
}

func isValidGraphvizFileFormat(format graphviz.Format) bool {
	_, exist := graphvizValidFormat[format]
	return exist
}

func createNode(node treegraph.Node, graph *cgraph.Graph) (*cgraph.Node, error) {
	node.BuildId()
	n, err := graph.CreateNode(node.Url)
	if err != nil {
		return nil, err
	}
	n = n.SetLabel(node.Id())
	formateUrl := strings.ReplaceAll(node.Url, "&", "&amp;")
	n = n.SetURL(formateUrl)
	n = n.SetShape(cgraph.BoxShape)
	n = n.SetTarget("_blank")
	return n, nil
}
