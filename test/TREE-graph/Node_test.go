package treegraph_test

import (
	"fmt"
	"testing"
	treegraph "tree-document-topology-visualizer/TREE-graph"
)

func TestCreatedNodeShouldHaveDifferentId(t *testing.T) {
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

func TestNodeShouldReturnTheRightUrl(t *testing.T) {
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
