package bash_test

import (
	"testing"

	sitter "github.com/colinwm/go-tree-sitter"
	"github.com/colinwm/go-tree-sitter/bash"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n := sitter.Parse([]byte("echo 1"), bash.GetLanguage())
	assert.Equal(
		"(program (command name: (command_name (word)) argument: (word)))",
		n.String(),
	)
}
