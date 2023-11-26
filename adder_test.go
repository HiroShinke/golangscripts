package adder

import (
	"adder"
	"testing"
)

func TestBasic(t *testing.T) {
	ret := Adder(1, 2)
	if ret != 3 {
		t.Fatalf(`adder(1,2) = %d != 3`, ret)
	}
}
