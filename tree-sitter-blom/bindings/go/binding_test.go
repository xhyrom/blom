package tree_sitter_blom_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_blom "github.com/xhyrom/blom/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_blom.Language())
	if language == nil {
		t.Errorf("Error loading Blom grammar")
	}
}
