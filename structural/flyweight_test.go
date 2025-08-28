package structural

import "testing"

func TestFlyweight(t *testing.T) {
	forest := NewForest()
	forest.AddTree("tree1", "black", "regular")
	forest.AddTree("tree2", "black", "regular")

	if forest.trees[0].treeType != forest.trees[1].treeType {
		t.Error("Expected reuse of TreeType")
	}
}
