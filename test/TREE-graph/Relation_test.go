package treegraph_test

import (
	"testing"
	treegraph "tree-document-topology-visualizer/TREE-graph"
)

func TestTheOperatorOfARelationShouldBeConvertable(t *testing.T) {
	aNode := treegraph.Node{Url: ""}
	rawLiteral := "foo"
	destination := aNode
	for k, v := range *treegraph.OperatorMapping() {
		relation := treegraph.Relation{
			RawOperator: k,
			Literal:     rawLiteral,
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
		Literal:     rawLiteral,
		Destination: destination,
	}
	relation.Operator()
}

func TestARelationShouldReturnAnEquationOfTheRightFormat(t *testing.T) {
	aNode := treegraph.Node{Url: ""}
	rawLiteral := "2"
	destination := aNode
	correctEquation := "m > 2"

	relation := treegraph.Relation{
		RawOperator: "https://w3id.org/tree#GreaterThanRelation",
		Literal:     rawLiteral,
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
		Literal:     rawLiteral,
		Destination: destination,
	}

	if relation.Equation() != correctEquation {
		t.Fatalf("The equation should be {%v} but it is {%v}", correctEquation, relation.Equation())
	}
}
