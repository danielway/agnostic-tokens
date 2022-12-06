// Code generated by tool/code_generator. DO NOT EDIT.

package code

type Model struct {
	Name	string
	ModelMetadata
}

func (m *Model) isType()	{}

type Primitive int

const (
	Boolean	Primitive	= iota + 1
	Int
	Rune
	String
	Void
)

func (p Primitive) isType()	{}

type List struct {
	Base	Type
	ListMetadata
}

func (l *List) isType()	{}

type Map struct {
	Key, Value	Type
	MapMetadata
}

func (m *Map) isType()	{}

type Set struct {
	Base	Type
	SetMetadata
}

func (s *Set) isType()	{}
