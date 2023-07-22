package hamt

import (
	"testing"
)

func TestCollisionsDuringPromotion(t *testing.T) {
	// n := valueNode[Hashable]{

	// }
}

func TestTerminalNodeSimpleInsertion(t *testing.T) {
	t.Run("Simple terminalNode insertion and retrieval", func(t *testing.T) {
		n := terminalNode[string]{}.insert(NewKey(foo{1}), "a")
		out, ok := n.retrieve(NewKey(foo{1}))
		if !ok {
			t.Errorf("Not okay!")
		}
		if out != "a" {
			t.Errorf("Unexpected result {%v instead of %v}", out, "a")
		}

		_, ok = n.retrieve(NewKey(foo{2}))
		if ok {
			t.Errorf("Unexpected ok!")
		}
	})
}

func TestTerminalNodeCollision(t *testing.T) {
	t.Run("Collision in terminalNode", func(t *testing.T) {
		n := terminalNode[string]{}.
			insert(NewKey(foo{1}), "a").
			insert(NewKey(foo{2}), "b")
		out, ok := n.retrieve(NewKey(foo{1}))
		if !ok || out != "a" {
			t.Errorf("Error retrieving %v (expected %v but got %v)", 1, "a", out)
		}

		out, ok = n.retrieve(NewKey(foo{2}))
		if !ok || out != "b" {
			t.Errorf("Error retrieving %v (expected %v but got %v)", 2, "b", out)
		}
	})
}

type foo struct {
	value int
}

func (f foo) Hash() uint8 {
	return 0
}

func (f foo) Equals(other Hashable) bool {
	switch o := other.(type) {
	case foo:
		return o.value == f.value
	default:
		return false
	}
}
