package treegraph_test

import (
	"fmt"
	"reflect"
	"testing"
	treegraph "tree-document-topology-visualizer/TREE-graph"
	"tree-document-topology-visualizer/communication"
)

func TestCreatingNNodeShouldHaveDifferentId(t *testing.T) {
	treegraph.ResetNodeCounter()
	url := ""
	for i := uint(0); i < 100; i++ {
		node := treegraph.Node{Url: url}
		node.BuildId()
		if node.Id() != fmt.Sprintf("n%v", i) {
			t.Fatalf("the node should have always a new id.\n The node has id {%v} and the id should be {n%v}", treegraph.Node{Url: url}.Id(), i)
		}
	}
}

func TestAUrlShouldReturnTheRightUrl(t *testing.T) {
	url := "foo"
	node := treegraph.Node{Url: url}

	if node.Url != url {
		t.Fatalf("The URL should be the same. The node has {%v} but it should be {%v}", node.Url, url)
	}
}

func TestTheIdOfANodeShouldBeFormated(t *testing.T) {
	treegraph.ResetNodeCounter()
	url := "foo"
	correctId := "n0"
	node := treegraph.Node{Url: url}

	if node.Id() != correctId {
		t.Fatalf("The id the of the node should be {%v} but it is {%v}", correctId, node.Id())
	}
}

func TestTheOperatorOfARelationShouldBeConvertable(t *testing.T) {
	aNode := treegraph.Node{Url: ""}
	rawLiteral := "foo"
	destination := aNode
	for k, v := range *treegraph.OperatorMapping() {
		relation := treegraph.Relation{
			RawOperator: k,
			RawLiteral:  rawLiteral,
			Destination: destination,
		}
		if relation.Operator() != v {
			t.Fatalf("The operator should be {%v} but it was {%v}", v, relation.Operator())
		}
	}
}

func TestTheRelationShouldPanicIfItWasGivenAFaltyOperator(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	aNode := treegraph.Node{Url: ""}
	rawLiteral := "foo"
	destination := aNode

	relation := treegraph.Relation{
		RawOperator: "",
		RawLiteral:  rawLiteral,
		Destination: destination,
	}
	relation.Operator()
}

func TestARelationShouldBeAbleToConvertLiteralBetweenGuillemet(t *testing.T) {
	aNode := treegraph.Node{Url: ""}
	rawLiteral := "\"abc\""
	destination := aNode

	relation := treegraph.Relation{
		RawOperator: "",
		RawLiteral:  rawLiteral,
		Destination: destination,
	}
	relation.Literal()
}

func TestARelationShouldPanicWhenTheLiteralIsHaveOnlyAGuillemetToTheRight(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	aNode := treegraph.Node{Url: ""}
	rawLiteral := "\"abc"
	destination := aNode

	relation := treegraph.Relation{
		RawOperator: "",
		RawLiteral:  rawLiteral,
		Destination: destination,
	}
	relation.Literal()
}

func TestARelationShouldPanicWhenTheLiteralIsHaveOnlyAGuillemetToTheLeft(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	aNode := treegraph.Node{Url: ""}
	rawLiteral := "abc\""
	destination := aNode

	relation := treegraph.Relation{
		RawOperator: "",
		RawLiteral:  rawLiteral,
		Destination: destination,
	}
	relation.Literal()
}

func TestARelationShouldReturnAnEquationOfTheRightFormat(t *testing.T) {
	aNode := treegraph.Node{Url: ""}
	rawLiteral := "\"2\""
	destination := aNode
	correctEquation := "m > 2"

	relation := treegraph.Relation{
		RawOperator: "https://w3id.org/tree#GreaterThanRelation",
		RawLiteral:  rawLiteral,
		Destination: destination,
	}

	if relation.Equation() != correctEquation {
		t.Fatalf("The equation should be {%v} but it is {%v}", correctEquation, relation.Equation())
	}
}

func TestARelationShouldPanicIfTheOperatorIsNotValid(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	aNode := treegraph.Node{Url: ""}
	rawLiteral := "\"2\""
	destination := aNode
	correctEquation := "m > 2"

	relation := treegraph.Relation{
		RawOperator: "https://w3id./tree#GreaterThanRelation",
		RawLiteral:  rawLiteral,
		Destination: destination,
	}

	if relation.Equation() != correctEquation {
		t.Fatalf("The equation should be {%v} but it is {%v}", correctEquation, relation.Equation())
	}
}

func TestARelationShouldPanicIfTheLiteralIsNotValid(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	aNode := treegraph.Node{Url: ""}
	rawLiteral := "\"2"
	destination := aNode
	correctEquation := "m > 2"

	relation := treegraph.Relation{
		RawOperator: "https://w3id.org/tree#GreaterThanRelation",
		RawLiteral:  rawLiteral,
		Destination: destination,
	}

	if relation.Equation() != correctEquation {
		t.Fatalf("The equation should be {%v} but it is {%v}", correctEquation, relation.Equation())
	}
}

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
		treegraph.Node{Url: "bar"}:  []treegraph.Relation{{RawOperator: rawOperator, RawLiteral: value, Destination: treegraph.Node{Url: "foo"}}},
		treegraph.Node{Url: "foo2"}: []treegraph.Relation{},
		treegraph.Node{Url: "bar2"}: []treegraph.Relation{{RawOperator: rawOperator, RawLiteral: value, Destination: treegraph.Node{Url: "foo2"}}},
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
		treegraph.Node{Url: "bar"}:  []treegraph.Relation{{RawOperator: rawOperator, RawLiteral: value, Destination: treegraph.Node{Url: "foo"}}},
		treegraph.Node{Url: "foo2"}: []treegraph.Relation{{RawOperator: rawOperator, RawLiteral: value, Destination: treegraph.Node{Url: "bar3"}}},
		treegraph.Node{Url: "bar2"}: []treegraph.Relation{{RawOperator: rawOperator, RawLiteral: value, Destination: treegraph.Node{Url: "foo2"}}},
	}

	resp := treegraph.NewGraphFromSparlRelationOutputs(outputs)

	if !reflect.DeepEqual(resp, expectedGraph) {
		t.Fatalf("The two graph should be the same.\nThe expected graph is {%v}\nThe returned graph is {%v}", expectedGraph, resp)
	}

}
