package treegraph

import "fmt"

var nNode uint = 0

func ResetNodeCounter() {
	nNode = 0
}

type Node struct {
	Url string
	id  uint
}

func (n Node) Id() string {
	return fmt.Sprintf("n%v", n.id)
}

func (n *Node) BuildId() {
	defer func() { nNode++ }()
	if n.id != 0 {
		n.id = nNode
	}

}
