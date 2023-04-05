package treegraph

import (
	"fmt"
	"log"
)

type Relation struct {
	RawOperator string
	Literal     string
	Destination Node
}

func (r Relation) Operator() string {
	operator, ok := operatorMapping[r.RawOperator]
	if !ok {
		log.Panic("the TREE document contain an unsupported relation operator")
	}
	return operator
}

func (r Relation) Equation() string {
	return fmt.Sprintf("m %v %v", r.Operator(), r.Literal)
}
