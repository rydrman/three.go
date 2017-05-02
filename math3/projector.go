package math3

type Projector interface {
	Positioner
	GetProjectionMatrix() *Matrix4
}
