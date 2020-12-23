package common

import "fmt"

type node struct {
	item     string
	count    int
	parent   *node
	children []*node
}

func NewNode() *node {
	return &node{
		children: make([]*node, 0),
	}
}

func NewNodeWithItem(item string) *node {
	return &node{
		item:     item,
		children: make([]*node, 0),
	}
}

func (n *node) AddCount() {
	n.count += 1
}

func (n *node) SetCount(c int) {
	n.count = c
}

func (n *node) AddChild(child *node) *node {
	child.parent = n
	n.children = append(n.children, child)
	return child
}

func (n *node) FindChild(item string) *node {
	for _, c := range n.children {
		if c.item == item {
			return c
		}
	}
	return nil
}

func (n *node) GetSelf() *node {
	return n
}

func (n *node) Scan(hp map[string]int, ancestors []string) {
	//对当前节点，将其和所有祖先的对加上当前的 supp
	for _, ancestor := range ancestors {
		pair := ancestor + " " + n.item
		if _, ok := hp[pair]; ok {
			hp[pair] += n.count
		} else {
			hp[pair] = n.count
		}
	}
	for _, c := range n.children {

		if len(n.item) > 0 {
			newAncestors := make([]string, len(ancestors)+1)
			copy(newAncestors, ancestors)
			newAncestors[len(ancestors)] = n.item
			c.Scan(hp, newAncestors)
		} else {
			c.Scan(hp, ancestors)
		}
	}
}

func (n *node) String() string {
	return fmt.Sprintf("%s:%d at %p", n.item, n.count, n)
}
