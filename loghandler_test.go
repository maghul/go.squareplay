package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeSelect(t *testing.T) {
	s := makeSelect("hej", "x", "a|bc|def")
	expected := `<select name="hej">
<option value="a">A</option>
<option value="bc">Bc</option>
<option value="def">Def</option>
</select>
`
	assert.Equal(t, expected, s)
}
