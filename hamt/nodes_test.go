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

func TestValueNodeInsertionAtMaxLevel(t *testing.T) {
	t.Run("Insert value at max key level", func(t *testing.T) {
		key1 := Key{
			foo{42},
			0,
			MASK_BITS - 1,
		}
		key2 := Key{
			foo{1},
			0,
			MASK_BITS - 1,
		}
		n := valueNode[string]{}.insert(key1, "asdf").insert(key2, "test")
		switch x := n.(type) {
		case terminalNode[string]:
			// good!
		default:
			t.Errorf("Unexpected node type %T", x)
		}

		if v, ok := n.retrieve(NewKey(foo{42})); !ok || v != "asdf" {
			t.Errorf("Unexpected result (expected %v got %v; ok = %v)", "asdf", v, ok)
		}

		if v, ok := n.retrieve(NewKey(foo{1})); !ok || v != "test" {
			t.Errorf("Unexpected result (expected %v got %v; ok = %v)", "test", v, ok)
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
