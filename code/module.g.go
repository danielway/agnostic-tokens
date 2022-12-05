// Code generated by tool/code_generator. DO NOT EDIT.

package code

type Module struct {
	Name		string
	Models		[]*ModelDef
	Functions	[]*FunctionDef
	Constants	[]*ConstantDef
	ModuleMetadata
}

type ConstantDef struct {
	Name	string
	Value	ConstantValue
	ConstantDefMetadata
}

func (c *ConstantDef) isDefinition()	{}

type FunctionDef struct {
	Name		string
	Arguments	[]*ArgumentDef
	Block		*Block
	ReturnType	Type
	FunctionDefMetadata
}
