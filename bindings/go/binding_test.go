package zig_test

import (
	"context"
	"testing"

	zig "github.com/slimsag/tree-sitter-zig/bindings/go"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	code := `//! Looks good to me!
const std = @import("std");

/// Nice!
pub fn main() !void {
	const stdout = std.io.getStdOut().writer();
	try stdout.print("Hello, {s}!\n", .{"world"});
}`

	n, err := sitter.ParseCtx(context.Background(), []byte(code), zig.GetLanguage())
	assert.NoError(err)
	assert.Equal(
		"(source_file (container_doc_comment) (TopLevelDecl (VarDecl variable_type_function: (IDENTIFIER) (ErrorUnionExpr (SuffixExpr (BUILTINIDENTIFIER) (FnCallArguments (ErrorUnionExpr (SuffixExpr (STRINGLITERALSINGLE)))))))) (doc_comment) (TopLevelDecl (FnProto function: (IDENTIFIER) (ParamDeclList) (ErrorUnionExpr (SuffixExpr (BuildinTypeExpr)))) (Block (Statement (VarDecl variable_type_function: (IDENTIFIER) (ErrorUnionExpr (SuffixExpr variable_type_function: (IDENTIFIER) (FieldOrFnCall field_access: (IDENTIFIER)) (FieldOrFnCall function_call: (IDENTIFIER) (FnCallArguments)) (FieldOrFnCall function_call: (IDENTIFIER) (FnCallArguments)))))) (Statement (AssignExpr (UnaryExpr operator: (PrefixOp) left: (ErrorUnionExpr (SuffixExpr variable_type_function: (IDENTIFIER) (FieldOrFnCall function_call: (IDENTIFIER) (FnCallArguments (ErrorUnionExpr (SuffixExpr (STRINGLITERALSINGLE (FormatSequence) (EscapeSequence)))) (ErrorUnionExpr (SuffixExpr (InitList (ErrorUnionExpr (SuffixExpr (STRINGLITERALSINGLE))))))))))))))))",
		n.String(),
	)
}
