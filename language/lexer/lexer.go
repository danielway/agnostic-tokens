//go:generate go run ../../tool/consumer_caster_generator/main.go -nodeTypesFile ../../ast/node_types.go

package lexer

import (
	"errors"
	"github.com/JosephNaberhaus/agnostic/ast"
)

// TODO rename
func Test(rawText string) (ast.Module, error) {
	r := newRunes(rawText)

	var module ast.Module

	newR, _, err := inOrder(
		anyWhitespaceConsumer(),
		skip(stringConsumer("module")),
		allWhitespaceConsumer(),
		handleNoError(
			alphaConsumer(),
			func(name string) {
				module.Name = name
			},
		),
		emptyLineConsumer(),
		repeat(first(
			emptyLineConsumer(),
			handleNoError(
				constantDefConsumer(),
				func(constant ast.ConstantDef) {
					module.Constants = append(module.Constants, constant)
				},
			),
			handleNoError(
				modelDefConsumer(),
				func(model ast.ModelDef) {
					module.Models = append(module.Models, model)
				},
			),
			handleNoError(
				functionDefConsumer(),
				func(function ast.FunctionDef) {
					module.Functions = append(module.Functions, function)
				},
			),
		)),
	)(r)
	if err != nil {
		return ast.Module{}, contextualize(err, []rune(rawText))
	}

	if newR.isNotEmpty() {
		return ast.Module{}, contextualize(createError(newR, "expected end of module"), []rune(rawText))
	}

	return module, nil
}

func constantDefConsumer() consumer[ast.ConstantDef] {
	var result ast.ConstantDef
	return attempt(
		&result,
		inOrder(
			anyWhitespaceConsumer(),
			skip(stringConsumer("const")),
			allWhitespaceConsumer(),
			handleNoError(
				alphaConsumer(),
				func(name string) {
					result.Name = name
				},
			),
			anyWhitespaceConsumer(),
			skip(stringConsumer("=")),
			anyWhitespaceConsumer(),
			handle(
				valueConsumer(),
				func(value ast.Value) error {
					if constantValue, ok := value.(ast.ConstantValue); ok {
						result.Value = constantValue
						return nil
					}

					return errors.New("constant def can only take constant values")
				},
			),
			skip(stringConsumer(";")),
		),
	)
}
