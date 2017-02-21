package core

import (
	"testing"
	"strings"
)

func TestPermute(t *testing.T) {
	in := "1 2 3"
	result := Permute(strings.Fields(in))
	t.Log(result)
}

func TestPermuteOperator(t *testing.T) {
	result := PermuteOperator(3)
	t.Log(result)
}
