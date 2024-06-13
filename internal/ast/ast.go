package ast

type Node interface {
	Offset() int
	End() int
}
