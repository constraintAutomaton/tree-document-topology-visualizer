package treegraph

import (
	"fmt"
	"log"
	"strings"
)

var nNode uint = 0

type Node struct {
	url string
	id  uint
}

func (n Node) Url() string {
	return n.url
}

func (n Node) Id() uint {
	return n.id
}

func NewNode(url string) Node {
	return Node{
		url: url,
		id:  nNode,
	}
}

type Relation struct {
	RawOperator string
	RawLiteral  string
	Origin      Node
	Destination Node
}

func (r Relation) Operator() string {
	operator, ok := operatorMapping[r.RawOperator]
	if !ok {
		log.Panic("the TREE document contain an unsupported relation operator")
	}
	return operator
}

func (r Relation) Literal() string {
	rawLiteralStrip := r.RawLiteral[1:]
	indexLastGuillemet := strings.Index(rawLiteralStrip, "\"")
	return rawLiteralStrip[:indexLastGuillemet]
}

func (r Relation) Equation() string {
	return fmt.Sprintf("x %v %v", r.Operator(), r.Literal())
}

type Graph map[Node][]Relation
