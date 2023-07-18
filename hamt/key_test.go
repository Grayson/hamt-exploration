package hamt

import "testing"

type HashableInt int // For testing purposes

func (i HashableInt) Hash() uint8 {
	return uint8(i)
}

func TestKeyMaskedValues(t *testing.T) {
	tests := []struct {
		name  string
		key   HashableInt
		depth uint8
		want  uint
	}{
		{
			"Zero",
			0,
			0,
			0,
		},
		{
			"One",
			1,
			0,
			1,
		},
		{
			"Two",
			2,
			0,
			2,
		},
		{
			"Second level, first value",
			0b1_000,
			1,
			1,
		},
		{
			"Second level, all values set",
			0b111_000,
			1,
			7,
		},
		{
			"Third level, first and second values set",
			0b11_000_000,
			2,
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Key{
				item:  tt.key,
				hash:  tt.key.Hash(),
				depth: tt.depth,
			}
			if got := k.currentMaskedValue(); got != tt.want {
				t.Errorf("Key.currentMaskedValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
