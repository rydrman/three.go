package objects

// Node represents an item that can exist in an object hierearchy.
// A node is the most basic of objects in the marshmallow library
type Node interface {
	Add(Node)
	Remove(Node) bool

	GetParent() Node
	GetChildren() []Node

	setParent(Node)
}
