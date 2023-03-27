package treegraph

import (
	"fmt"
	"log"
	"strings"
)

type Relation struct {
	RawOperator string
	RawLiteral  string
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
	if indexLastGuillemet == -1 || indexLastGuillemet == 0 || strings.Count(r.RawLiteral, "\"") != 2 {
		log.Panic("the literal should be between guillement")
	}
	return rawLiteralStrip[:indexLastGuillemet]
}

func (r Relation) Equation() string {
	return fmt.Sprintf("m %v %v", r.Operator(), r.Literal())
}
