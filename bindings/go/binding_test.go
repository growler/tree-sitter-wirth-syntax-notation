package tree_sitter_wsn_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_wsn "github.com/tree-sitter/tree-sitter-wsn/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_wsn.Language())
	if language == nil {
		t.Errorf("Error loading Wsn grammar")
	}
}
