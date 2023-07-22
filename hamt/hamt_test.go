package hamt

import (
	"testing"
)

func TestBasicInsertAndRetrieve(t *testing.T) {
	type kvp = struct {
		uint8
		string
	}

	tests := []struct {
		name  string
		pairs []kvp
	}{
		{
			"Insert one",
			[]kvp{
				{42, "test"},
			},
		},
		{
			"Insert two",
			[]kvp{
				{42, "test"},
				{1, "more"},
			},
		},
		{
			"Insert nine", // This will exceed the standard array size of 8
			[]kvp{
				{0, "one"},
				{1, "two"},
				{2, "three"},
				{3, "four"},
				{4, "five"},
				{5, "six"},
				{6, "seven"},
				{7, "eight"},
				{8, "nine"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := NewTrie[string]()

			// Inserts
			for _, x := range tt.pairs {
				trie = trie.Insert(HashableInt(x.uint8), x.string)
			}

			// Retrieval
			for _, x := range tt.pairs {
				got, ok := trie.Retrieve(HashableInt(x.uint8))
				if got != x.string || !ok {
					t.Errorf("Error retrieving %v got %v, want %v", x.uint8,
						got, x.string)
				}
			}
		})
	}
}
