package hamt

const ARRAY_NODE_SIZE = 8

type nodeType[TValue any] interface {
	canInsert() bool
	insert(uint8, TValue) nodeType[TValue]
	retrieve(uint8) TValue
}

/* Trampoline */

type trampolineNode[TValue any] struct {
}

func (trampolineNode[TValue]) canInsert() bool {
	return true
}

func (trampolineNode[TValue]) insert(uint8, TValue) nodeType[TValue] {
	panic("Unimplemented")
}

func (trampolineNode[TValue]) retrieve(uint8) TValue {
	panic("Unimplemented")
}

/* Value */

type valueNode[TValue any] struct {
	key   uint8
	value TValue
}

func (valueNode[TValue]) canInsert() bool {
	return false
}

func (valueNode[TValue]) insert(key uint8, value TValue) nodeType[TValue] {
	// TODO: Collisions
	// TODO: Expansions
	return valueNode[TValue]{
		key,
		value,
	}
}

func (n valueNode[TValue]) retrieve(key uint8) TValue {
	return n.value
}

/* Array */

type arrayNode[TValue any] struct {
	children [ARRAY_NODE_SIZE]nodeType[TValue]
}

func (n arrayNode[TValue]) canInsert() bool {
	return true
}

func (n arrayNode[TValue]) insert(key uint8, value TValue) nodeType[TValue] {
	index := key % ARRAY_NODE_SIZE
	children := n.children
	children[index] = valueNode[TValue]{key, value}
	return arrayNode[TValue]{
		children: children,
	}
}

func (n arrayNode[TValue]) retrieve(key uint8) TValue {
	index := key % ARRAY_NODE_SIZE
	return n.children[index].retrieve(key)
}
