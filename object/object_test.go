package object

import "testing"

func TestStringHashKey(t *testing.T) {
	foo1 := &String{Value: "foo"}
	foo2 := &String{Value: "foo"}
	bar1 := &String{Value: "bar"}
	bar2 := &String{Value: "bar"}

	if foo1.HashKey() != foo2.HashKey() {
		t.Errorf("strings with same value have different hash keys")
	}
	if bar1.HashKey() != bar2.HashKey() {
		t.Errorf("strings with same value have different hash keys")
	}

	if foo1.HashKey() == bar1.HashKey() {
		t.Errorf("strings with different value have same hash keys")
	}
}
