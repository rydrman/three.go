package objects

import "github.com/golang/glog"

// Object acts as a common base to many more
// specific types in the marshmallow library
type Object struct {
	parent   Node
	children []Node
}

// NewObject creates a new Object instace with default values
func NewObject() *Object {
	return &Object{}
}

// Add adds the given node as a child of this object
func (o *Object) Add(n Node) {

	if n == Node(o) {
		glog.Error("cannot add object as a child of itself")
		return
	}

	o.children = append(o.children, n)
	n.setParent(o)

}

// Remove remvoves the given node from this nodes children,
// returns true if the node was found and removed successfully
func (o *Object) Remove(n Node) bool {

	for i, child := range o.children {

		if child == n {

			o.children = append(o.children[:i], o.children[i+1:]...)
			child.setParent(nil)
			return true

		}

	}

	return false

}

// GetParent returns the current parent of this object
func (o *Object) GetParent() Node {

	return o.parent

}

// GetChildren returns the current slice of child nodes for this object
func (o *Object) GetChildren() []Node {

	return o.children

}

// setParent sets the parent of this object to the given node
// this is a hidden function because it needs to be kept up to date
// through the Add and Remove functions
func (o *Object) setParent(n Node) {

	if o.parent != nil {

		o.parent.Remove(o)

	}

	o.parent = n

}
