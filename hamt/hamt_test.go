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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := NewTrie[string]()

			// Inserts
			for _, x := range tt.pairs {
				trie = trie.Insert(x.uint8, x.string)
			}

			// Retrieval
			for _, x := range tt.pairs {
				got := trie.Retrieve(x.uint8)
				if got != x.string {
					t.Errorf("Error retrieving %v got %v, want %v", x.uint8,
						got, x.string)
				}
			}
		})
	}
}
