package elm_test

import (
	"testing"

	sitter "github.com/colinwm/go-tree-sitter"
	"github.com/colinwm/go-tree-sitter/elm"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n := sitter.Parse([]byte("import Html exposing (text)"), elm.GetLanguage())
	assert.Equal(
		"(file (import_clause (import) moduleName: (upper_case_qid (upper_case_identifier)) exposing: (exposing_list (exposing) (exposed_value (lower_case_identifier)))))",
		n.String(),
	)
}
