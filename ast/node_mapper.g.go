// Code generated by tool/code_generator. DO NOT EDIT.

package ast

// Node is the interface that everything in this package implements
type Node interface {
	// isNode is only a type-guard to limit what can be used as a Node.
	isNode()
}

func (f Function) isNode() {}

func (f FunctionProperty) isNode() {}

func (l LiteralInt) isNode() {}

func (l LiteralRune) isNode() {}

func (l LiteralString) isNode() {}

func (l LiteralList) isNode() {}

func (k KeyValue) isNode() {}

func (l LiteralMap) isNode() {}

func (f FieldDef) isNode() {}

func (a ArgumentDef) isNode() {}

func (m MethodDef) isNode() {}

func (m ModelDef) isNode() {}

func (v Variable) isNode() {}

func (p Property) isNode() {}

func (m Module) isNode() {}

func (c ConstantDef) isNode() {}

func (f FunctionDef) isNode() {}

func (u UnaryOperator) isNode() {}

func (u UnaryOperation) isNode() {}

func (b BinaryOperator) isNode() {}

func (b BinaryOperation) isNode() {}

func (b Block) isNode() {}

func (a Assignment) isNode() {}

func (i If) isNode() {}

func (e ElseIf) isNode() {}

func (e Else) isNode() {}

func (c Conditional) isNode() {}

func (r Return) isNode() {}

func (d Declare) isNode() {}

func (f For) isNode() {}

func (f ForIn) isNode() {}

func (m Model) isNode() {}

func (p Primitive) isNode() {}

func (l List) isNode() {}

func (m Map) isNode() {}

func (c Call) isNode() {}

func (l Lookup) isNode() {}

func (n New) isNode() {}

func (l Length) isNode() {}

type NodeMapper[T any] interface {
	MapFunction(original Function) (T, error)

	MapFunctionProperty(original FunctionProperty) (T, error)

	MapLiteralInt(original LiteralInt) (T, error)

	MapLiteralRune(original LiteralRune) (T, error)

	MapLiteralString(original LiteralString) (T, error)

	MapLiteralList(original LiteralList) (T, error)

	MapKeyValue(original KeyValue) (T, error)

	MapLiteralMap(original LiteralMap) (T, error)

	MapFieldDef(original FieldDef) (T, error)

	MapArgumentDef(original ArgumentDef) (T, error)

	MapMethodDef(original MethodDef) (T, error)

	MapModelDef(original ModelDef) (T, error)

	MapVariable(original Variable) (T, error)

	MapProperty(original Property) (T, error)

	MapModule(original Module) (T, error)

	MapConstantDef(original ConstantDef) (T, error)

	MapFunctionDef(original FunctionDef) (T, error)

	MapUnaryOperator(original UnaryOperator) (T, error)

	MapUnaryOperation(original UnaryOperation) (T, error)

	MapBinaryOperator(original BinaryOperator) (T, error)

	MapBinaryOperation(original BinaryOperation) (T, error)

	MapBlock(original Block) (T, error)

	MapAssignment(original Assignment) (T, error)

	MapIf(original If) (T, error)

	MapElseIf(original ElseIf) (T, error)

	MapElse(original Else) (T, error)

	MapConditional(original Conditional) (T, error)

	MapReturn(original Return) (T, error)

	MapDeclare(original Declare) (T, error)

	MapFor(original For) (T, error)

	MapForIn(original ForIn) (T, error)

	MapModel(original Model) (T, error)

	MapPrimitive(original Primitive) (T, error)

	MapList(original List) (T, error)

	MapMap(original Map) (T, error)

	MapCall(original Call) (T, error)

	MapLookup(original Lookup) (T, error)

	MapNew(original New) (T, error)

	MapLength(original Length) (T, error)
}

func MapNode[T any](node Node, mapper NodeMapper[T]) (T, error) {
	switch value := node.(type) {

	case Function:
		return mapper.MapFunction(value)

	case FunctionProperty:
		return mapper.MapFunctionProperty(value)

	case LiteralInt:
		return mapper.MapLiteralInt(value)

	case LiteralRune:
		return mapper.MapLiteralRune(value)

	case LiteralString:
		return mapper.MapLiteralString(value)

	case LiteralList:
		return mapper.MapLiteralList(value)

	case KeyValue:
		return mapper.MapKeyValue(value)

	case LiteralMap:
		return mapper.MapLiteralMap(value)

	case FieldDef:
		return mapper.MapFieldDef(value)

	case ArgumentDef:
		return mapper.MapArgumentDef(value)

	case MethodDef:
		return mapper.MapMethodDef(value)

	case ModelDef:
		return mapper.MapModelDef(value)

	case Variable:
		return mapper.MapVariable(value)

	case Property:
		return mapper.MapProperty(value)

	case Module:
		return mapper.MapModule(value)

	case ConstantDef:
		return mapper.MapConstantDef(value)

	case FunctionDef:
		return mapper.MapFunctionDef(value)

	case UnaryOperator:
		return mapper.MapUnaryOperator(value)

	case UnaryOperation:
		return mapper.MapUnaryOperation(value)

	case BinaryOperator:
		return mapper.MapBinaryOperator(value)

	case BinaryOperation:
		return mapper.MapBinaryOperation(value)

	case Block:
		return mapper.MapBlock(value)

	case Assignment:
		return mapper.MapAssignment(value)

	case If:
		return mapper.MapIf(value)

	case ElseIf:
		return mapper.MapElseIf(value)

	case Else:
		return mapper.MapElse(value)

	case Conditional:
		return mapper.MapConditional(value)

	case Return:
		return mapper.MapReturn(value)

	case Declare:
		return mapper.MapDeclare(value)

	case For:
		return mapper.MapFor(value)

	case ForIn:
		return mapper.MapForIn(value)

	case Model:
		return mapper.MapModel(value)

	case Primitive:
		return mapper.MapPrimitive(value)

	case List:
		return mapper.MapList(value)

	case Map:
		return mapper.MapMap(value)

	case Call:
		return mapper.MapCall(value)

	case Lookup:
		return mapper.MapLookup(value)

	case New:
		return mapper.MapNew(value)

	case Length:
		return mapper.MapLength(value)

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
	MapFunctionNoError(original Function) T

	MapFunctionPropertyNoError(original FunctionProperty) T

	MapLiteralIntNoError(original LiteralInt) T

	MapLiteralRuneNoError(original LiteralRune) T

	MapLiteralStringNoError(original LiteralString) T

	MapLiteralListNoError(original LiteralList) T

	MapKeyValueNoError(original KeyValue) T

	MapLiteralMapNoError(original LiteralMap) T

	MapFieldDefNoError(original FieldDef) T

	MapArgumentDefNoError(original ArgumentDef) T

	MapMethodDefNoError(original MethodDef) T

	MapModelDefNoError(original ModelDef) T

	MapVariableNoError(original Variable) T

	MapPropertyNoError(original Property) T

	MapModuleNoError(original Module) T

	MapConstantDefNoError(original ConstantDef) T

	MapFunctionDefNoError(original FunctionDef) T

	MapUnaryOperatorNoError(original UnaryOperator) T

	MapUnaryOperationNoError(original UnaryOperation) T

	MapBinaryOperatorNoError(original BinaryOperator) T

	MapBinaryOperationNoError(original BinaryOperation) T

	MapBlockNoError(original Block) T

	MapAssignmentNoError(original Assignment) T

	MapIfNoError(original If) T

	MapElseIfNoError(original ElseIf) T

	MapElseNoError(original Else) T

	MapConditionalNoError(original Conditional) T

	MapReturnNoError(original Return) T

	MapDeclareNoError(original Declare) T

	MapForNoError(original For) T

	MapForInNoError(original ForIn) T

	MapModelNoError(original Model) T

	MapPrimitiveNoError(original Primitive) T

	MapListNoError(original List) T

	MapMapNoError(original Map) T

	MapCallNoError(original Call) T

	MapLookupNoError(original Lookup) T

	MapNewNoError(original New) T

	MapLengthNoError(original Length) T
}

func MapNodeNoError[T any](node Node, mapper NodeMapperNoError[T]) T {
	switch value := node.(type) {

	case Function:
		return mapper.MapFunctionNoError(value)

	case FunctionProperty:
		return mapper.MapFunctionPropertyNoError(value)

	case LiteralInt:
		return mapper.MapLiteralIntNoError(value)

	case LiteralRune:
		return mapper.MapLiteralRuneNoError(value)

	case LiteralString:
		return mapper.MapLiteralStringNoError(value)

	case LiteralList:
		return mapper.MapLiteralListNoError(value)

	case KeyValue:
		return mapper.MapKeyValueNoError(value)

	case LiteralMap:
		return mapper.MapLiteralMapNoError(value)

	case FieldDef:
		return mapper.MapFieldDefNoError(value)

	case ArgumentDef:
		return mapper.MapArgumentDefNoError(value)

	case MethodDef:
		return mapper.MapMethodDefNoError(value)

	case ModelDef:
		return mapper.MapModelDefNoError(value)

	case Variable:
		return mapper.MapVariableNoError(value)

	case Property:
		return mapper.MapPropertyNoError(value)

	case Module:
		return mapper.MapModuleNoError(value)

	case ConstantDef:
		return mapper.MapConstantDefNoError(value)

	case FunctionDef:
		return mapper.MapFunctionDefNoError(value)

	case UnaryOperator:
		return mapper.MapUnaryOperatorNoError(value)

	case UnaryOperation:
		return mapper.MapUnaryOperationNoError(value)

	case BinaryOperator:
		return mapper.MapBinaryOperatorNoError(value)

	case BinaryOperation:
		return mapper.MapBinaryOperationNoError(value)

	case Block:
		return mapper.MapBlockNoError(value)

	case Assignment:
		return mapper.MapAssignmentNoError(value)

	case If:
		return mapper.MapIfNoError(value)

	case ElseIf:
		return mapper.MapElseIfNoError(value)

	case Else:
		return mapper.MapElseNoError(value)

	case Conditional:
		return mapper.MapConditionalNoError(value)

	case Return:
		return mapper.MapReturnNoError(value)

	case Declare:
		return mapper.MapDeclareNoError(value)

	case For:
		return mapper.MapForNoError(value)

	case ForIn:
		return mapper.MapForInNoError(value)

	case Model:
		return mapper.MapModelNoError(value)

	case Primitive:
		return mapper.MapPrimitiveNoError(value)

	case List:
		return mapper.MapListNoError(value)

	case Map:
		return mapper.MapMapNoError(value)

	case Call:
		return mapper.MapCallNoError(value)

	case Lookup:
		return mapper.MapLookupNoError(value)

	case New:
		return mapper.MapNewNoError(value)

	case Length:
		return mapper.MapLengthNoError(value)

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
	MapFunction(original Function) (T, error)

	MapFunctionProperty(original FunctionProperty) (T, error)
}

func MapCallable[T any](nodeType Callable, mapper CallableMapper[T]) (T, error) {
	switch value := nodeType.(type) {

	case Function:
		return mapper.MapFunction(value)

	case FunctionProperty:
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
	MapFunctionNoError(original Function) T

	MapFunctionPropertyNoError(original FunctionProperty) T
}

func MapCallableNoError[T any](node Callable, mapper CallableMapperNoError[T]) T {
	switch value := node.(type) {

	case Function:
		return mapper.MapFunctionNoError(value)

	case FunctionProperty:
		return mapper.MapFunctionPropertyNoError(value)

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
	MapLiteralInt(original LiteralInt) (T, error)

	MapLiteralRune(original LiteralRune) (T, error)

	MapLiteralString(original LiteralString) (T, error)

	MapLiteralList(original LiteralList) (T, error)

	MapLiteralMap(original LiteralMap) (T, error)
}

func MapConstantValue[T any](nodeType ConstantValue, mapper ConstantValueMapper[T]) (T, error) {
	switch value := nodeType.(type) {

	case LiteralInt:
		return mapper.MapLiteralInt(value)

	case LiteralRune:
		return mapper.MapLiteralRune(value)

	case LiteralString:
		return mapper.MapLiteralString(value)

	case LiteralList:
		return mapper.MapLiteralList(value)

	case LiteralMap:
		return mapper.MapLiteralMap(value)

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
	MapLiteralIntNoError(original LiteralInt) T

	MapLiteralRuneNoError(original LiteralRune) T

	MapLiteralStringNoError(original LiteralString) T

	MapLiteralListNoError(original LiteralList) T

	MapLiteralMapNoError(original LiteralMap) T
}

func MapConstantValueNoError[T any](node ConstantValue, mapper ConstantValueMapperNoError[T]) T {
	switch value := node.(type) {

	case LiteralInt:
		return mapper.MapLiteralIntNoError(value)

	case LiteralRune:
		return mapper.MapLiteralRuneNoError(value)

	case LiteralString:
		return mapper.MapLiteralStringNoError(value)

	case LiteralList:
		return mapper.MapLiteralListNoError(value)

	case LiteralMap:
		return mapper.MapLiteralMapNoError(value)

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
	MapFieldDef(original FieldDef) (T, error)

	MapArgumentDef(original ArgumentDef) (T, error)

	MapConstantDef(original ConstantDef) (T, error)

	MapDeclare(original Declare) (T, error)

	MapForIn(original ForIn) (T, error)
}

func MapDefinition[T any](nodeType Definition, mapper DefinitionMapper[T]) (T, error) {
	switch value := nodeType.(type) {

	case FieldDef:
		return mapper.MapFieldDef(value)

	case ArgumentDef:
		return mapper.MapArgumentDef(value)

	case ConstantDef:
		return mapper.MapConstantDef(value)

	case Declare:
		return mapper.MapDeclare(value)

	case ForIn:
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
	MapFieldDefNoError(original FieldDef) T

	MapArgumentDefNoError(original ArgumentDef) T

	MapConstantDefNoError(original ConstantDef) T

	MapDeclareNoError(original Declare) T

	MapForInNoError(original ForIn) T
}

func MapDefinitionNoError[T any](node Definition, mapper DefinitionMapperNoError[T]) T {
	switch value := node.(type) {

	case FieldDef:
		return mapper.MapFieldDefNoError(value)

	case ArgumentDef:
		return mapper.MapArgumentDefNoError(value)

	case ConstantDef:
		return mapper.MapConstantDefNoError(value)

	case Declare:
		return mapper.MapDeclareNoError(value)

	case ForIn:
		return mapper.MapForInNoError(value)

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
	MapAssignment(original Assignment) (T, error)

	MapConditional(original Conditional) (T, error)

	MapReturn(original Return) (T, error)

	MapDeclare(original Declare) (T, error)

	MapFor(original For) (T, error)

	MapForIn(original ForIn) (T, error)
}

func MapStatement[T any](nodeType Statement, mapper StatementMapper[T]) (T, error) {
	switch value := nodeType.(type) {

	case Assignment:
		return mapper.MapAssignment(value)

	case Conditional:
		return mapper.MapConditional(value)

	case Return:
		return mapper.MapReturn(value)

	case Declare:
		return mapper.MapDeclare(value)

	case For:
		return mapper.MapFor(value)

	case ForIn:
		return mapper.MapForIn(value)

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
	MapAssignmentNoError(original Assignment) T

	MapConditionalNoError(original Conditional) T

	MapReturnNoError(original Return) T

	MapDeclareNoError(original Declare) T

	MapForNoError(original For) T

	MapForInNoError(original ForIn) T
}

func MapStatementNoError[T any](node Statement, mapper StatementMapperNoError[T]) T {
	switch value := node.(type) {

	case Assignment:
		return mapper.MapAssignmentNoError(value)

	case Conditional:
		return mapper.MapConditionalNoError(value)

	case Return:
		return mapper.MapReturnNoError(value)

	case Declare:
		return mapper.MapDeclareNoError(value)

	case For:
		return mapper.MapForNoError(value)

	case ForIn:
		return mapper.MapForInNoError(value)

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
	MapModel(original Model) (T, error)

	MapPrimitive(original Primitive) (T, error)

	MapList(original List) (T, error)

	MapMap(original Map) (T, error)
}

func MapType[T any](nodeType Type, mapper TypeMapper[T]) (T, error) {
	switch value := nodeType.(type) {

	case Model:
		return mapper.MapModel(value)

	case Primitive:
		return mapper.MapPrimitive(value)

	case List:
		return mapper.MapList(value)

	case Map:
		return mapper.MapMap(value)

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
	MapModelNoError(original Model) T

	MapPrimitiveNoError(original Primitive) T

	MapListNoError(original List) T

	MapMapNoError(original Map) T
}

func MapTypeNoError[T any](node Type, mapper TypeMapperNoError[T]) T {
	switch value := node.(type) {

	case Model:
		return mapper.MapModelNoError(value)

	case Primitive:
		return mapper.MapPrimitiveNoError(value)

	case List:
		return mapper.MapListNoError(value)

	case Map:
		return mapper.MapMapNoError(value)

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
	MapLiteralInt(original LiteralInt) (T, error)

	MapLiteralRune(original LiteralRune) (T, error)

	MapLiteralString(original LiteralString) (T, error)

	MapLiteralList(original LiteralList) (T, error)

	MapLiteralMap(original LiteralMap) (T, error)

	MapVariable(original Variable) (T, error)

	MapProperty(original Property) (T, error)

	MapUnaryOperation(original UnaryOperation) (T, error)

	MapBinaryOperation(original BinaryOperation) (T, error)

	MapCall(original Call) (T, error)

	MapLookup(original Lookup) (T, error)

	MapNew(original New) (T, error)

	MapLength(original Length) (T, error)
}

func MapValue[T any](nodeType Value, mapper ValueMapper[T]) (T, error) {
	switch value := nodeType.(type) {

	case LiteralInt:
		return mapper.MapLiteralInt(value)

	case LiteralRune:
		return mapper.MapLiteralRune(value)

	case LiteralString:
		return mapper.MapLiteralString(value)

	case LiteralList:
		return mapper.MapLiteralList(value)

	case LiteralMap:
		return mapper.MapLiteralMap(value)

	case Variable:
		return mapper.MapVariable(value)

	case Property:
		return mapper.MapProperty(value)

	case UnaryOperation:
		return mapper.MapUnaryOperation(value)

	case BinaryOperation:
		return mapper.MapBinaryOperation(value)

	case Call:
		return mapper.MapCall(value)

	case Lookup:
		return mapper.MapLookup(value)

	case New:
		return mapper.MapNew(value)

	case Length:
		return mapper.MapLength(value)

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
	MapLiteralIntNoError(original LiteralInt) T

	MapLiteralRuneNoError(original LiteralRune) T

	MapLiteralStringNoError(original LiteralString) T

	MapLiteralListNoError(original LiteralList) T

	MapLiteralMapNoError(original LiteralMap) T

	MapVariableNoError(original Variable) T

	MapPropertyNoError(original Property) T

	MapUnaryOperationNoError(original UnaryOperation) T

	MapBinaryOperationNoError(original BinaryOperation) T

	MapCallNoError(original Call) T

	MapLookupNoError(original Lookup) T

	MapNewNoError(original New) T

	MapLengthNoError(original Length) T
}

func MapValueNoError[T any](node Value, mapper ValueMapperNoError[T]) T {
	switch value := node.(type) {

	case LiteralInt:
		return mapper.MapLiteralIntNoError(value)

	case LiteralRune:
		return mapper.MapLiteralRuneNoError(value)

	case LiteralString:
		return mapper.MapLiteralStringNoError(value)

	case LiteralList:
		return mapper.MapLiteralListNoError(value)

	case LiteralMap:
		return mapper.MapLiteralMapNoError(value)

	case Variable:
		return mapper.MapVariableNoError(value)

	case Property:
		return mapper.MapPropertyNoError(value)

	case UnaryOperation:
		return mapper.MapUnaryOperationNoError(value)

	case BinaryOperation:
		return mapper.MapBinaryOperationNoError(value)

	case Call:
		return mapper.MapCallNoError(value)

	case Lookup:
		return mapper.MapLookupNoError(value)

	case New:
		return mapper.MapNewNoError(value)

	case Length:
		return mapper.MapLengthNoError(value)

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
