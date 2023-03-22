package treegraph_test

import (
	"fmt"
	"testing"
	treegraph "tree-document-topology-visualizer/TREE-graph"
)

func TestCreatingNNodeShouldHaveDifferentId(t *testing.T) {
	treegraph.ResetNodeCounter()
	url := ""
	for i := uint(0); i < 100; i++ {
		if treegraph.NewNode(url).Id() != fmt.Sprintf("n%v", i) {
			t.Fatalf("the node should have always a new id.\n The node has id {%v} and the id should be {n%v}", treegraph.NewNode(url).Id(), i)
		}
	}
}

func TestAUrlShouldReturnTheRightUrl(t *testing.T) {
	url := "foo"
	node := treegraph.NewNode(url)

	if node.Url() != url {
		t.Fatalf("The URL should be the same. The node has {%v} but it should be {%v}", node.Url(), url)
	}
}

func TestTheIdOfANodeShouldBeFormated(t *testing.T) {
	treegraph.ResetNodeCounter()
	url := "foo"
	correctId := "n0"
	node := treegraph.NewNode(url)

	if node.Id() != correctId {
		t.Fatalf("The id the of the node should be {%v} but it is {%v}", correctId, node.Id())
	}
}

func TestTheOperatorOfARelationShouldBeConvertable(t *testing.T) {
	aNode := treegraph.NewNode("")
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

	aNode := treegraph.NewNode("")
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
	aNode := treegraph.NewNode("")
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

	aNode := treegraph.NewNode("")
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

	aNode := treegraph.NewNode("")
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
	aNode := treegraph.NewNode("")
	rawLiteral := "\"2\""
	destination := aNode
	correctEquation := "x > 2"

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

	aNode := treegraph.NewNode("")
	rawLiteral := "\"2\""
	destination := aNode
	correctEquation := "x > 2"

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

	aNode := treegraph.NewNode("")
	rawLiteral := "\"2"
	destination := aNode
	correctEquation := "x > 2"

	relation := treegraph.Relation{
		RawOperator: "https://w3id.org/tree#GreaterThanRelation",
		RawLiteral:  rawLiteral,
		Destination: destination,
	}

	if relation.Equation() != correctEquation {
		t.Fatalf("The equation should be {%v} but it is {%v}", correctEquation, relation.Equation())
	}
}
