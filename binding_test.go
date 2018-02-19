package sitter_test

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/javascript"
	"github.com/stretchr/testify/assert"
)

func TestRootNode(t *testing.T) {
	assert := assert.New(t)

	n, close := sitter.Parse([]byte("let a = 1"), javascript.GetLanguage())
	defer close()

	assert.Equal(uint32(0), n.StartByte())
	assert.Equal(uint32(9), n.EndByte())
	assert.Equal("(program (lexical_declaration (variable_declarator (identifier) (number))))", n.String())
	assert.Equal("program", n.Type())

	assert.Equal(true, n.IsNamed())
	assert.Equal(false, n.IsMissing())
	assert.Equal(false, n.HasChanges())
	assert.Equal(false, n.HasError())

	assert.Equal(uint32(1), n.ChildCount())
	assert.Equal(uint32(1), n.NamedChildCount())
	assert.Equal(^uint32(0), n.ChildIndex())

	assert.Nil(n.Parent())
	assert.Nil(n.NextSibling())
	assert.Nil(n.NextNamedSibling())
	assert.Nil(n.PrevSibling())
	assert.Nil(n.PrevNamedSibling())

	assert.NotNil(n.Child(0))
	assert.NotNil(n.NamedChild(0))
}

func TestDocument(t *testing.T) {
	assert := assert.New(t)

	doc := sitter.NewDocument()
	defer doc.Close()

	doc.SetLanguage(javascript.GetLanguage())
	doc.SetInputBytes([]byte("let a = 1"))
	doc.Parse()
	n := doc.RootNode()

	assert.Equal(uint32(0), n.StartByte())
	assert.Equal(uint32(9), n.EndByte())
	assert.Equal("program", n.Type())
	assert.Equal("(program (lexical_declaration (variable_declarator (identifier) (number))))", n.String())

	doc.SetInputBytes([]byte("let a = 'a'"))
	doc.Parse()
	n = doc.RootNode()
	assert.Equal("(program (lexical_declaration (variable_declarator (identifier) (string))))", n.String())
}