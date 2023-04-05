package treegraph_test

import (
	"reflect"
	"testing"
	treegraph "tree-document-topology-visualizer/TREE-graph"
	"tree-document-topology-visualizer/communication"
)

func TestCreateNewGraph(t *testing.T) {
	rawOperator := "https://w3id.org/tree#GreaterThanRelation"
	value := "\"2022-08-07T08:08:21.000Z\"^^<http://www.w3.org/2001/XMLSchema#dateTime>"

	outputs := []communication.SparqlRelationOutput{
		{
			Operator: rawOperator,
			Value:    value,
			NextNode: "foo",
			Node:     "bar",
		},
		{
			Operator: rawOperator,
			Value:    value,
			NextNode: "foo2",
			Node:     "bar2",
		},
	}

	expectedGraph := treegraph.Graph{
		treegraph.Node{Url: "foo"}:  []treegraph.Relation{},
		treegraph.Node{Url: "bar"}:  []treegraph.Relation{{RawOperator: rawOperator, Literal: value, Destination: treegraph.Node{Url: "foo"}}},
		treegraph.Node{Url: "foo2"}: []treegraph.Relation{},
		treegraph.Node{Url: "bar2"}: []treegraph.Relation{{RawOperator: rawOperator, Literal: value, Destination: treegraph.Node{Url: "foo2"}}},
	}

	resp := treegraph.NewGraphFromSparlRelationOutputs(outputs)

	if !reflect.DeepEqual(resp, expectedGraph) {
		t.Fatalf("The two graph should be the same.\nThe expected graph is {%v}\nThe returned graph is {%v}", expectedGraph, resp)
	}

}

func TestCreateNewGraphWithDestinationNodeWithRelation(t *testing.T) {
	rawOperator := "https://w3id.org/tree#GreaterThanRelation"
	value := "\"2022-08-07T08:08:21.000Z\"^^<http://www.w3.org/2001/XMLSchema#dateTime>"

	outputs := []communication.SparqlRelationOutput{
		{
			Operator: rawOperator,
			Value:    value,
			NextNode: "foo",
			Node:     "bar",
		},
		{
			Operator: rawOperator,
			Value:    value,
			NextNode: "foo2",
			Node:     "bar2",
		},
		{
			Operator: rawOperator,
			Value:    value,
			NextNode: "bar3",
			Node:     "foo2",
		},
	}

	expectedGraph := treegraph.Graph{
		treegraph.Node{Url: "foo"}:  []treegraph.Relation{},
		treegraph.Node{Url: "bar3"}: []treegraph.Relation{},
		treegraph.Node{Url: "bar"}:  []treegraph.Relation{{RawOperator: rawOperator, Literal: value, Destination: treegraph.Node{Url: "foo"}}},
		treegraph.Node{Url: "foo2"}: []treegraph.Relation{{RawOperator: rawOperator, Literal: value, Destination: treegraph.Node{Url: "bar3"}}},
		treegraph.Node{Url: "bar2"}: []treegraph.Relation{{RawOperator: rawOperator, Literal: value, Destination: treegraph.Node{Url: "foo2"}}},
	}

	resp := treegraph.NewGraphFromSparlRelationOutputs(outputs)

	if !reflect.DeepEqual(resp, expectedGraph) {
		t.Fatalf("The two graph should be the same.\nThe expected graph is {%v}\nThe returned graph is {%v}", expectedGraph, resp)
	}

}
