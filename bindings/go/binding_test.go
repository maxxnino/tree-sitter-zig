package tree_sitter_zig_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-zig"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_zig.Language())
	if language == nil {
		t.Errorf("Error loading Zig grammar")
	}
}
