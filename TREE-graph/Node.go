package treegraph

import "fmt"

var nNode uint = 0

// ResetNodeCounter reset the counter of the id of the Node(s)
func ResetNodeCounter() {
	nNode = 0
}

// Node describes a tree:Node.
// It has also an optional unique id for more intuitive topology visualization.
type Node struct {
	Url string
	id  uint
}

func (n Node) Id() string {
	return fmt.Sprintf("n%v", n.id)
}

func (n *Node) BuildId() {
	defer func() { nNode++ }()
	n.id = nNode
}
