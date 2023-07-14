package hamt

const MASK_BITS = 3 // 3 bits can index all 8 slots in array nodes

type Key struct {
	item  Hashable
	hash  uint8
	depth uint8
}

func NewKey(hashable Hashable) Key {
	return Key{hashable, hashable.Hash(), 0}
}

func (k *Key) moveToNextLevel() {
	k.depth++
}

func (k *Key) currentMaskedValue() uint {
	bitsToShift := MASK_BITS * k.depth
	shiftedBits := k.hash >> bitsToShift
	maskedValue := shiftedBits & 0b111
	return uint(maskedValue)
}
