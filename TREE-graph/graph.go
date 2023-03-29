package treegraph

import (
	"tree-document-topology-visualizer/communication"
)

type Graph map[Node][]Relation

func NewGraphFromSparlRelationOutputs(outputs []communication.SparqlRelationOutput) Graph {
	graph := Graph{}
	// we create all the nodes
	for _, output := range outputs {
		graph[Node{Url: output.Node}] = []Relation{}

		graph[Node{Url: output.NextNode}] = []Relation{}

	}

	// we create all the edges
	for _, output := range outputs {
		relation := Relation{
			RawOperator: output.Operator,
			Literal:     output.Value,
			Destination: Node{Url: output.NextNode},
		}
		graph[Node{Url: output.Node}] = append(graph[Node{Url: output.Node}], relation)
	}

	return graph
}
