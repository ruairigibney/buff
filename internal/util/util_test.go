package util

import "testing"

func TestInt64ContainsTrue(t *testing.T) {
	ints := []int64{0, 1, 2}
	if found := Int64Contains(ints, 0); !found {
		t.Errorf("int found when it shouldn't have been")
	}
}

func TestInt64ContainsFalse(t *testing.T) {
	ints := []int64{0, 1, 2}
	if found := Int64Contains(ints, 4); found {
		t.Errorf("int wasn't found when it should have been")
	}
}
