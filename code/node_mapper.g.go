// Code generated by tool/code_generator. DO NOT EDIT.

package code

// Node is the interface that everything in this package implements
type Node interface {
	// isNode is only a type-guard to limit what can be used as a Node.
	isNode()
}

func (f *Function) isNode() {}

func (f *FunctionProperty) isNode() {}

func (l *LiteralInt) isNode() {}

func (l *LiteralRune) isNode() {}

func (l *LiteralString) isNode() {}

func (l *LiteralBool) isNode() {}

func (l *LiteralList) isNode() {}

func (k *KeyValue) isNode() {}

func (l *LiteralMap) isNode() {}

func (l *LiteralSet) isNode() {}

func (e *EmptyList) isNode() {}

func (e *EmptySet) isNode() {}

func (f *FieldDef) isNode() {}

func (a *ArgumentDef) isNode() {}

func (m *MethodDef) isNode() {}

func (m *ModelDef) isNode() {}

func (v *Variable) isNode() {}

func (p *Property) isNode() {}

func (m *Module) isNode() {}

func (c *ConstantDef) isNode() {}

func (f *FunctionDef) isNode() {}

func (u UnaryOperator) isNode() {}

func (u *UnaryOperation) isNode() {}

func (b BinaryOperator) isNode() {}

func (b *BinaryOperation) isNode() {}

func (b *Block) isNode() {}

func (a *Assignment) isNode() {}

func (i *If) isNode() {}

func (e *ElseIf) isNode() {}

func (e *Else) isNode() {}

func (c *Conditional) isNode() {}

func (r *Return) isNode() {}

func (d *Declare) isNode() {}

func (f *For) isNode() {}

func (f *ForIn) isNode() {}

func (a *AddToSet) isNode() {}

func (p *Push) isNode() {}

func (m *Model) isNode() {}

func (p Primitive) isNode() {}

func (l *List) isNode() {}

func (m *Map) isNode() {}

func (s *Set) isNode() {}

func (c *Call) isNode() {}

func (l *Lookup) isNode() {}

func (n *New) isNode() {}

func (l *Length) isNode() {}

func (s *SetContains) isNode() {}

func (p *Pop) isNode() {}

type NodeMapper[T any] interface {
	MapFunction(original *Function) (T, error)

	MapFunctionProperty(original *FunctionProperty) (T, error)

	MapLiteralInt(original *LiteralInt) (T, error)

	MapLiteralRune(original *LiteralRune) (T, error)

	MapLiteralString(original *LiteralString) (T, error)

	MapLiteralBool(original *LiteralBool) (T, error)

	MapLiteralList(original *LiteralList) (T, error)

	MapKeyValue(original *KeyValue) (T, error)

	MapLiteralMap(original *LiteralMap) (T, error)

	MapLiteralSet(original *LiteralSet) (T, error)

	MapEmptyList(original *EmptyList) (T, error)

	MapEmptySet(original *EmptySet) (T, error)

	MapFieldDef(original *FieldDef) (T, error)

	MapArgumentDef(original *ArgumentDef) (T, error)

	MapMethodDef(original *MethodDef) (T, error)

	MapModelDef(original *ModelDef) (T, error)

	MapVariable(original *Variable) (T, error)

	MapProperty(original *Property) (T, error)

	MapModule(original *Module) (T, error)

	MapConstantDef(original *ConstantDef) (T, error)

	MapFunctionDef(original *FunctionDef) (T, error)

	MapUnaryOperator(original UnaryOperator) (T, error)

	MapUnaryOperation(original *UnaryOperation) (T, error)

	MapBinaryOperator(original BinaryOperator) (T, error)

	MapBinaryOperation(original *BinaryOperation) (T, error)

	MapBlock(original *Block) (T, error)

	MapAssignment(original *Assignment) (T, error)

	MapIf(original *If) (T, error)

	MapElseIf(original *ElseIf) (T, error)

	MapElse(original *Else) (T, error)

	MapConditional(original *Conditional) (T, error)

	MapReturn(original *Return) (T, error)

	MapDeclare(original *Declare) (T, error)

	MapFor(original *For) (T, error)

	MapForIn(original *ForIn) (T, error)

	MapAddToSet(original *AddToSet) (T, error)

	MapPush(original *Push) (T, error)

	MapModel(original *Model) (T, error)

	MapPrimitive(original Primitive) (T, error)

	MapList(original *List) (T, error)

	MapMap(original *Map) (T, error)

	MapSet(original *Set) (T, error)

	MapCall(original *Call) (T, error)

	MapLookup(original *Lookup) (T, error)

	MapNew(original *New) (T, error)

	MapLength(original *Length) (T, error)

	MapSetContains(original *SetContains) (T, error)

	MapPop(original *Pop) (T, error)
}

func MapNode[T any](node Node, mapper NodeMapper[T]) (T, error) {
	switch value := node.(type) {

	case *Function:
		return mapper.MapFunction(value)

	case *FunctionProperty:
		return mapper.MapFunctionProperty(value)

	case *LiteralInt:
		return mapper.MapLiteralInt(value)

	case *LiteralRune:
		return mapper.MapLiteralRune(value)

	case *LiteralString:
		return mapper.MapLiteralString(value)

	case *LiteralBool:
		return mapper.MapLiteralBool(value)

	case *LiteralList:
		return mapper.MapLiteralList(value)

	case *KeyValue:
		return mapper.MapKeyValue(value)

	case *LiteralMap:
		return mapper.MapLiteralMap(value)

	case *LiteralSet:
		return mapper.MapLiteralSet(value)

	case *EmptyList:
		return mapper.MapEmptyList(value)

	case *EmptySet:
		return mapper.MapEmptySet(value)

	case *FieldDef:
		return mapper.MapFieldDef(value)

	case *ArgumentDef:
		return mapper.MapArgumentDef(value)

	case *MethodDef:
		return mapper.MapMethodDef(value)

	case *ModelDef:
		return mapper.MapModelDef(value)

	case *Variable:
		return mapper.MapVariable(value)

	case *Property:
		return mapper.MapProperty(value)

	case *Module:
		return mapper.MapModule(value)

	case *ConstantDef:
		return mapper.MapConstantDef(value)

	case *FunctionDef:
		return mapper.MapFunctionDef(value)

	case UnaryOperator:
		return mapper.MapUnaryOperator(value)

	case *UnaryOperation:
		return mapper.MapUnaryOperation(value)

	case BinaryOperator:
		return mapper.MapBinaryOperator(value)

	case *BinaryOperation:
		return mapper.MapBinaryOperation(value)

	case *Block:
		return mapper.MapBlock(value)

	case *Assignment:
		return mapper.MapAssignment(value)

	case *If:
		return mapper.MapIf(value)

	case *ElseIf:
		return mapper.MapElseIf(value)

	case *Else:
		return mapper.MapElse(value)

	case *Conditional:
		return mapper.MapConditional(value)

	case *Return:
		return mapper.MapReturn(value)

	case *Declare:
		return mapper.MapDeclare(value)

	case *For:
		return mapper.MapFor(value)

	case *ForIn:
		return mapper.MapForIn(value)

	case *AddToSet:
		return mapper.MapAddToSet(value)

	case *Push:
		return mapper.MapPush(value)

	case *Model:
		return mapper.MapModel(value)

	case Primitive:
		return mapper.MapPrimitive(value)

	case *List:
		return mapper.MapList(value)

	case *Map:
		return mapper.MapMap(value)

	case *Set:
		return mapper.MapSet(value)

	case *Call:
		return mapper.MapCall(value)

	case *Lookup:
		return mapper.MapLookup(value)

	case *New:
		return mapper.MapNew(value)

	case *Length:
		return mapper.MapLength(value)

	case *SetContains:
		return mapper.MapSetContains(value)

	case *Pop:
		return mapper.MapPop(value)

	default:
		panic("unreachable")
	}
}

func MapNodes[T any, V Node](nodes []V, mapper NodeMapper[T]) ([]T, error) {
	var resultNodes []T
	for _, node := range nodes {
		resultNode, err := MapNode(node, mapper)
		if err != nil {
			return nil, err
		}

		resultNodes = append(resultNodes, resultNode)
	}

	return resultNodes, nil
}

type NodeMapperNoError[T any] interface {
	MapFunction(original *Function) T

	MapFunctionProperty(original *FunctionProperty) T

	MapLiteralInt(original *LiteralInt) T

	MapLiteralRune(original *LiteralRune) T

	MapLiteralString(original *LiteralString) T

	MapLiteralBool(original *LiteralBool) T

	MapLiteralList(original *LiteralList) T

	MapKeyValue(original *KeyValue) T

	MapLiteralMap(original *LiteralMap) T

	MapLiteralSet(original *LiteralSet) T

	MapEmptyList(original *EmptyList) T

	MapEmptySet(original *EmptySet) T

	MapFieldDef(original *FieldDef) T

	MapArgumentDef(original *ArgumentDef) T

	MapMethodDef(original *MethodDef) T

	MapModelDef(original *ModelDef) T

	MapVariable(original *Variable) T

	MapProperty(original *Property) T

	MapModule(original *Module) T

	MapConstantDef(original *ConstantDef) T

	MapFunctionDef(original *FunctionDef) T

	MapUnaryOperator(original UnaryOperator) T

	MapUnaryOperation(original *UnaryOperation) T

	MapBinaryOperator(original BinaryOperator) T

	MapBinaryOperation(original *BinaryOperation) T

	MapBlock(original *Block) T

	MapAssignment(original *Assignment) T

	MapIf(original *If) T

	MapElseIf(original *ElseIf) T

	MapElse(original *Else) T

	MapConditional(original *Conditional) T

	MapReturn(original *Return) T

	MapDeclare(original *Declare) T

	MapFor(original *For) T

	MapForIn(original *ForIn) T

	MapAddToSet(original *AddToSet) T

	MapPush(original *Push) T

	MapModel(original *Model) T

	MapPrimitive(original Primitive) T

	MapList(original *List) T

	MapMap(original *Map) T

	MapSet(original *Set) T

	MapCall(original *Call) T

	MapLookup(original *Lookup) T

	MapNew(original *New) T

	MapLength(original *Length) T

	MapSetContains(original *SetContains) T

	MapPop(original *Pop) T
}

func MapNodeNoError[T any](node Node, mapper NodeMapperNoError[T]) T {
	switch value := node.(type) {

	case *Function:
		return mapper.MapFunction(value)

	case *FunctionProperty:
		return mapper.MapFunctionProperty(value)

	case *LiteralInt:
		return mapper.MapLiteralInt(value)

	case *LiteralRune:
		return mapper.MapLiteralRune(value)

	case *LiteralString:
		return mapper.MapLiteralString(value)

	case *LiteralBool:
		return mapper.MapLiteralBool(value)

	case *LiteralList:
		return mapper.MapLiteralList(value)

	case *KeyValue:
		return mapper.MapKeyValue(value)

	case *LiteralMap:
		return mapper.MapLiteralMap(value)

	case *LiteralSet:
		return mapper.MapLiteralSet(value)

	case *EmptyList:
		return mapper.MapEmptyList(value)

	case *EmptySet:
		return mapper.MapEmptySet(value)

	case *FieldDef:
		return mapper.MapFieldDef(value)

	case *ArgumentDef:
		return mapper.MapArgumentDef(value)

	case *MethodDef:
		return mapper.MapMethodDef(value)

	case *ModelDef:
		return mapper.MapModelDef(value)

	case *Variable:
		return mapper.MapVariable(value)

	case *Property:
		return mapper.MapProperty(value)

	case *Module:
		return mapper.MapModule(value)

	case *ConstantDef:
		return mapper.MapConstantDef(value)

	case *FunctionDef:
		return mapper.MapFunctionDef(value)

	case UnaryOperator:
		return mapper.MapUnaryOperator(value)

	case *UnaryOperation:
		return mapper.MapUnaryOperation(value)

	case BinaryOperator:
		return mapper.MapBinaryOperator(value)

	case *BinaryOperation:
		return mapper.MapBinaryOperation(value)

	case *Block:
		return mapper.MapBlock(value)

	case *Assignment:
		return mapper.MapAssignment(value)

	case *If:
		return mapper.MapIf(value)

	case *ElseIf:
		return mapper.MapElseIf(value)

	case *Else:
		return mapper.MapElse(value)

	case *Conditional:
		return mapper.MapConditional(value)

	case *Return:
		return mapper.MapReturn(value)

	case *Declare:
		return mapper.MapDeclare(value)

	case *For:
		return mapper.MapFor(value)

	case *ForIn:
		return mapper.MapForIn(value)

	case *AddToSet:
		return mapper.MapAddToSet(value)

	case *Push:
		return mapper.MapPush(value)

	case *Model:
		return mapper.MapModel(value)

	case Primitive:
		return mapper.MapPrimitive(value)

	case *List:
		return mapper.MapList(value)

	case *Map:
		return mapper.MapMap(value)

	case *Set:
		return mapper.MapSet(value)

	case *Call:
		return mapper.MapCall(value)

	case *Lookup:
		return mapper.MapLookup(value)

	case *New:
		return mapper.MapNew(value)

	case *Length:
		return mapper.MapLength(value)

	case *SetContains:
		return mapper.MapSetContains(value)

	case *Pop:
		return mapper.MapPop(value)

	default:
		panic("unreachable")
	}
}

func MapNodesNoError[T any, V Node](nodes []V, mapper NodeMapperNoError[T]) []T {
	var resultNodes []T
	for _, node := range nodes {
		resultNodes = append(resultNodes, MapNodeNoError(node, mapper))
	}

	return resultNodes
}

type CallableMapper[T any] interface {
	MapFunction(original *Function) (T, error)

	MapFunctionProperty(original *FunctionProperty) (T, error)
}

func MapCallable[T any](nodeType Callable, mapper CallableMapper[T]) (T, error) {
	switch value := nodeType.(type) {

	case *Function:
		return mapper.MapFunction(value)

	case *FunctionProperty:
		return mapper.MapFunctionProperty(value)

	default:
		panic("unreachable")
	}
}

func MapCallables[T any](nodes []Callable, mapper CallableMapper[T]) ([]T, error) {
	var resultNodes []T
	for _, node := range nodes {
		resultNode, err := MapCallable(node, mapper)
		if err != nil {
			return nil, err
		}

		resultNodes = append(resultNodes, resultNode)
	}

	return resultNodes, nil
}

type CallableMapperNoError[T any] interface {
	MapFunction(original *Function) T

	MapFunctionProperty(original *FunctionProperty) T
}

func MapCallableNoError[T any](node Callable, mapper CallableMapperNoError[T]) T {
	switch value := node.(type) {

	case *Function:
		return mapper.MapFunction(value)

	case *FunctionProperty:
		return mapper.MapFunctionProperty(value)

	default:
		panic("unreachable")
	}
}

func MapCallablesNoError[T any](nodes []Callable, mapper CallableMapperNoError[T]) []T {
	var resultNodes []T
	for _, node := range nodes {
		resultNodes = append(resultNodes, MapCallableNoError(node, mapper))
	}

	return resultNodes
}

type ConstantValueMapper[T any] interface {
	MapLiteralInt(original *LiteralInt) (T, error)

	MapLiteralRune(original *LiteralRune) (T, error)

	MapLiteralString(original *LiteralString) (T, error)

	MapLiteralBool(original *LiteralBool) (T, error)

	MapLiteralList(original *LiteralList) (T, error)

	MapLiteralMap(original *LiteralMap) (T, error)

	MapLiteralSet(original *LiteralSet) (T, error)

	MapEmptyList(original *EmptyList) (T, error)

	MapEmptySet(original *EmptySet) (T, error)
}

func MapConstantValue[T any](nodeType ConstantValue, mapper ConstantValueMapper[T]) (T, error) {
	switch value := nodeType.(type) {

	case *LiteralInt:
		return mapper.MapLiteralInt(value)

	case *LiteralRune:
		return mapper.MapLiteralRune(value)

	case *LiteralString:
		return mapper.MapLiteralString(value)

	case *LiteralBool:
		return mapper.MapLiteralBool(value)

	case *LiteralList:
		return mapper.MapLiteralList(value)

	case *LiteralMap:
		return mapper.MapLiteralMap(value)

	case *LiteralSet:
		return mapper.MapLiteralSet(value)

	case *EmptyList:
		return mapper.MapEmptyList(value)

	case *EmptySet:
		return mapper.MapEmptySet(value)

	default:
		panic("unreachable")
	}
}

func MapConstantValues[T any](nodes []ConstantValue, mapper ConstantValueMapper[T]) ([]T, error) {
	var resultNodes []T
	for _, node := range nodes {
		resultNode, err := MapConstantValue(node, mapper)
		if err != nil {
			return nil, err
		}

		resultNodes = append(resultNodes, resultNode)
	}

	return resultNodes, nil
}

type ConstantValueMapperNoError[T any] interface {
	MapLiteralInt(original *LiteralInt) T

	MapLiteralRune(original *LiteralRune) T

	MapLiteralString(original *LiteralString) T

	MapLiteralBool(original *LiteralBool) T

	MapLiteralList(original *LiteralList) T

	MapLiteralMap(original *LiteralMap) T

	MapLiteralSet(original *LiteralSet) T

	MapEmptyList(original *EmptyList) T

	MapEmptySet(original *EmptySet) T
}

func MapConstantValueNoError[T any](node ConstantValue, mapper ConstantValueMapperNoError[T]) T {
	switch value := node.(type) {

	case *LiteralInt:
		return mapper.MapLiteralInt(value)

	case *LiteralRune:
		return mapper.MapLiteralRune(value)

	case *LiteralString:
		return mapper.MapLiteralString(value)

	case *LiteralBool:
		return mapper.MapLiteralBool(value)

	case *LiteralList:
		return mapper.MapLiteralList(value)

	case *LiteralMap:
		return mapper.MapLiteralMap(value)

	case *LiteralSet:
		return mapper.MapLiteralSet(value)

	case *EmptyList:
		return mapper.MapEmptyList(value)

	case *EmptySet:
		return mapper.MapEmptySet(value)

	default:
		panic("unreachable")
	}
}

func MapConstantValuesNoError[T any](nodes []ConstantValue, mapper ConstantValueMapperNoError[T]) []T {
	var resultNodes []T
	for _, node := range nodes {
		resultNodes = append(resultNodes, MapConstantValueNoError(node, mapper))
	}

	return resultNodes
}

type DefinitionMapper[T any] interface {
	MapFieldDef(original *FieldDef) (T, error)

	MapArgumentDef(original *ArgumentDef) (T, error)

	MapConstantDef(original *ConstantDef) (T, error)

	MapDeclare(original *Declare) (T, error)

	MapForIn(original *ForIn) (T, error)
}

func MapDefinition[T any](nodeType Definition, mapper DefinitionMapper[T]) (T, error) {
	switch value := nodeType.(type) {

	case *FieldDef:
		return mapper.MapFieldDef(value)

	case *ArgumentDef:
		return mapper.MapArgumentDef(value)

	case *ConstantDef:
		return mapper.MapConstantDef(value)

	case *Declare:
		return mapper.MapDeclare(value)

	case *ForIn:
		return mapper.MapForIn(value)

	default:
		panic("unreachable")
	}
}

func MapDefinitions[T any](nodes []Definition, mapper DefinitionMapper[T]) ([]T, error) {
	var resultNodes []T
	for _, node := range nodes {
		resultNode, err := MapDefinition(node, mapper)
		if err != nil {
			return nil, err
		}

		resultNodes = append(resultNodes, resultNode)
	}

	return resultNodes, nil
}

type DefinitionMapperNoError[T any] interface {
	MapFieldDef(original *FieldDef) T

	MapArgumentDef(original *ArgumentDef) T

	MapConstantDef(original *ConstantDef) T

	MapDeclare(original *Declare) T

	MapForIn(original *ForIn) T
}

func MapDefinitionNoError[T any](node Definition, mapper DefinitionMapperNoError[T]) T {
	switch value := node.(type) {

	case *FieldDef:
		return mapper.MapFieldDef(value)

	case *ArgumentDef:
		return mapper.MapArgumentDef(value)

	case *ConstantDef:
		return mapper.MapConstantDef(value)

	case *Declare:
		return mapper.MapDeclare(value)

	case *ForIn:
		return mapper.MapForIn(value)

	default:
		panic("unreachable")
	}
}

func MapDefinitionsNoError[T any](nodes []Definition, mapper DefinitionMapperNoError[T]) []T {
	var resultNodes []T
	for _, node := range nodes {
		resultNodes = append(resultNodes, MapDefinitionNoError(node, mapper))
	}

	return resultNodes
}

type StatementMapper[T any] interface {
	MapAssignment(original *Assignment) (T, error)

	MapConditional(original *Conditional) (T, error)

	MapReturn(original *Return) (T, error)

	MapDeclare(original *Declare) (T, error)

	MapFor(original *For) (T, error)

	MapForIn(original *ForIn) (T, error)

	MapAddToSet(original *AddToSet) (T, error)

	MapPush(original *Push) (T, error)

	MapCall(original *Call) (T, error)

	MapPop(original *Pop) (T, error)
}

func MapStatement[T any](nodeType Statement, mapper StatementMapper[T]) (T, error) {
	switch value := nodeType.(type) {

	case *Assignment:
		return mapper.MapAssignment(value)

	case *Conditional:
		return mapper.MapConditional(value)

	case *Return:
		return mapper.MapReturn(value)

	case *Declare:
		return mapper.MapDeclare(value)

	case *For:
		return mapper.MapFor(value)

	case *ForIn:
		return mapper.MapForIn(value)

	case *AddToSet:
		return mapper.MapAddToSet(value)

	case *Push:
		return mapper.MapPush(value)

	case *Call:
		return mapper.MapCall(value)

	case *Pop:
		return mapper.MapPop(value)

	default:
		panic("unreachable")
	}
}

func MapStatements[T any](nodes []Statement, mapper StatementMapper[T]) ([]T, error) {
	var resultNodes []T
	for _, node := range nodes {
		resultNode, err := MapStatement(node, mapper)
		if err != nil {
			return nil, err
		}

		resultNodes = append(resultNodes, resultNode)
	}

	return resultNodes, nil
}

type StatementMapperNoError[T any] interface {
	MapAssignment(original *Assignment) T

	MapConditional(original *Conditional) T

	MapReturn(original *Return) T

	MapDeclare(original *Declare) T

	MapFor(original *For) T

	MapForIn(original *ForIn) T

	MapAddToSet(original *AddToSet) T

	MapPush(original *Push) T

	MapCall(original *Call) T

	MapPop(original *Pop) T
}

func MapStatementNoError[T any](node Statement, mapper StatementMapperNoError[T]) T {
	switch value := node.(type) {

	case *Assignment:
		return mapper.MapAssignment(value)

	case *Conditional:
		return mapper.MapConditional(value)

	case *Return:
		return mapper.MapReturn(value)

	case *Declare:
		return mapper.MapDeclare(value)

	case *For:
		return mapper.MapFor(value)

	case *ForIn:
		return mapper.MapForIn(value)

	case *AddToSet:
		return mapper.MapAddToSet(value)

	case *Push:
		return mapper.MapPush(value)

	case *Call:
		return mapper.MapCall(value)

	case *Pop:
		return mapper.MapPop(value)

	default:
		panic("unreachable")
	}
}

func MapStatementsNoError[T any](nodes []Statement, mapper StatementMapperNoError[T]) []T {
	var resultNodes []T
	for _, node := range nodes {
		resultNodes = append(resultNodes, MapStatementNoError(node, mapper))
	}

	return resultNodes
}

type TypeMapper[T any] interface {
	MapModel(original *Model) (T, error)

	MapPrimitive(original Primitive) (T, error)

	MapList(original *List) (T, error)

	MapMap(original *Map) (T, error)

	MapSet(original *Set) (T, error)
}

func MapType[T any](nodeType Type, mapper TypeMapper[T]) (T, error) {
	switch value := nodeType.(type) {

	case *Model:
		return mapper.MapModel(value)

	case Primitive:
		return mapper.MapPrimitive(value)

	case *List:
		return mapper.MapList(value)

	case *Map:
		return mapper.MapMap(value)

	case *Set:
		return mapper.MapSet(value)

	default:
		panic("unreachable")
	}
}

func MapTypes[T any](nodes []Type, mapper TypeMapper[T]) ([]T, error) {
	var resultNodes []T
	for _, node := range nodes {
		resultNode, err := MapType(node, mapper)
		if err != nil {
			return nil, err
		}

		resultNodes = append(resultNodes, resultNode)
	}

	return resultNodes, nil
}

type TypeMapperNoError[T any] interface {
	MapModel(original *Model) T

	MapPrimitive(original Primitive) T

	MapList(original *List) T

	MapMap(original *Map) T

	MapSet(original *Set) T
}

func MapTypeNoError[T any](node Type, mapper TypeMapperNoError[T]) T {
	switch value := node.(type) {

	case *Model:
		return mapper.MapModel(value)

	case Primitive:
		return mapper.MapPrimitive(value)

	case *List:
		return mapper.MapList(value)

	case *Map:
		return mapper.MapMap(value)

	case *Set:
		return mapper.MapSet(value)

	default:
		panic("unreachable")
	}
}

func MapTypesNoError[T any](nodes []Type, mapper TypeMapperNoError[T]) []T {
	var resultNodes []T
	for _, node := range nodes {
		resultNodes = append(resultNodes, MapTypeNoError(node, mapper))
	}

	return resultNodes
}

type ValueMapper[T any] interface {
	MapLiteralInt(original *LiteralInt) (T, error)

	MapLiteralRune(original *LiteralRune) (T, error)

	MapLiteralString(original *LiteralString) (T, error)

	MapLiteralBool(original *LiteralBool) (T, error)

	MapLiteralList(original *LiteralList) (T, error)

	MapLiteralMap(original *LiteralMap) (T, error)

	MapLiteralSet(original *LiteralSet) (T, error)

	MapEmptyList(original *EmptyList) (T, error)

	MapEmptySet(original *EmptySet) (T, error)

	MapVariable(original *Variable) (T, error)

	MapProperty(original *Property) (T, error)

	MapUnaryOperation(original *UnaryOperation) (T, error)

	MapBinaryOperation(original *BinaryOperation) (T, error)

	MapCall(original *Call) (T, error)

	MapLookup(original *Lookup) (T, error)

	MapNew(original *New) (T, error)

	MapLength(original *Length) (T, error)

	MapSetContains(original *SetContains) (T, error)

	MapPop(original *Pop) (T, error)
}

func MapValue[T any](nodeType Value, mapper ValueMapper[T]) (T, error) {
	switch value := nodeType.(type) {

	case *LiteralInt:
		return mapper.MapLiteralInt(value)

	case *LiteralRune:
		return mapper.MapLiteralRune(value)

	case *LiteralString:
		return mapper.MapLiteralString(value)

	case *LiteralBool:
		return mapper.MapLiteralBool(value)

	case *LiteralList:
		return mapper.MapLiteralList(value)

	case *LiteralMap:
		return mapper.MapLiteralMap(value)

	case *LiteralSet:
		return mapper.MapLiteralSet(value)

	case *EmptyList:
		return mapper.MapEmptyList(value)

	case *EmptySet:
		return mapper.MapEmptySet(value)

	case *Variable:
		return mapper.MapVariable(value)

	case *Property:
		return mapper.MapProperty(value)

	case *UnaryOperation:
		return mapper.MapUnaryOperation(value)

	case *BinaryOperation:
		return mapper.MapBinaryOperation(value)

	case *Call:
		return mapper.MapCall(value)

	case *Lookup:
		return mapper.MapLookup(value)

	case *New:
		return mapper.MapNew(value)

	case *Length:
		return mapper.MapLength(value)

	case *SetContains:
		return mapper.MapSetContains(value)

	case *Pop:
		return mapper.MapPop(value)

	default:
		panic("unreachable")
	}
}

func MapValues[T any](nodes []Value, mapper ValueMapper[T]) ([]T, error) {
	var resultNodes []T
	for _, node := range nodes {
		resultNode, err := MapValue(node, mapper)
		if err != nil {
			return nil, err
		}

		resultNodes = append(resultNodes, resultNode)
	}

	return resultNodes, nil
}

type ValueMapperNoError[T any] interface {
	MapLiteralInt(original *LiteralInt) T

	MapLiteralRune(original *LiteralRune) T

	MapLiteralString(original *LiteralString) T

	MapLiteralBool(original *LiteralBool) T

	MapLiteralList(original *LiteralList) T

	MapLiteralMap(original *LiteralMap) T

	MapLiteralSet(original *LiteralSet) T

	MapEmptyList(original *EmptyList) T

	MapEmptySet(original *EmptySet) T

	MapVariable(original *Variable) T

	MapProperty(original *Property) T

	MapUnaryOperation(original *UnaryOperation) T

	MapBinaryOperation(original *BinaryOperation) T

	MapCall(original *Call) T

	MapLookup(original *Lookup) T

	MapNew(original *New) T

	MapLength(original *Length) T

	MapSetContains(original *SetContains) T

	MapPop(original *Pop) T
}

func MapValueNoError[T any](node Value, mapper ValueMapperNoError[T]) T {
	switch value := node.(type) {

	case *LiteralInt:
		return mapper.MapLiteralInt(value)

	case *LiteralRune:
		return mapper.MapLiteralRune(value)

	case *LiteralString:
		return mapper.MapLiteralString(value)

	case *LiteralBool:
		return mapper.MapLiteralBool(value)

	case *LiteralList:
		return mapper.MapLiteralList(value)

	case *LiteralMap:
		return mapper.MapLiteralMap(value)

	case *LiteralSet:
		return mapper.MapLiteralSet(value)

	case *EmptyList:
		return mapper.MapEmptyList(value)

	case *EmptySet:
		return mapper.MapEmptySet(value)

	case *Variable:
		return mapper.MapVariable(value)

	case *Property:
		return mapper.MapProperty(value)

	case *UnaryOperation:
		return mapper.MapUnaryOperation(value)

	case *BinaryOperation:
		return mapper.MapBinaryOperation(value)

	case *Call:
		return mapper.MapCall(value)

	case *Lookup:
		return mapper.MapLookup(value)

	case *New:
		return mapper.MapNew(value)

	case *Length:
		return mapper.MapLength(value)

	case *SetContains:
		return mapper.MapSetContains(value)

	case *Pop:
		return mapper.MapPop(value)

	default:
		panic("unreachable")
	}
}

func MapValuesNoError[T any](nodes []Value, mapper ValueMapperNoError[T]) []T {
	var resultNodes []T
	for _, node := range nodes {
		resultNodes = append(resultNodes, MapValueNoError(node, mapper))
	}

	return resultNodes
}
