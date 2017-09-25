package confeito

import "testing"

func TestForestEnqueueOnly(t *testing.T) {
	x := DenseFeature{-2.0, -1.0, 0.0, 1.0, 2.0, 3.0}

	// Each tree should be unbalanced for testing
	tree1 := NewAssert(t).SucceedNew(NewLeaf(0, -2.5, float32(0.0), float32(1.0))).(*Leaf)
	tree1.SetRight(NewAssert(t).SucceedNew(NewLeaf(1, 0.0, float32(1.0), float32(2.0))).(*Leaf))
	tree2 := NewAssert(t).SucceedNew(NewLeaf(0, -4.5, float32(0.0), float32(1.0))).(*Leaf)
	tree2R := NewAssert(t).SucceedNew(NewLeaf(2, 0.0, float32(1.0), float32(2.0))).(*Leaf)
	tree2RR := NewAssert(t).SucceedNew(NewLeaf(3, 0.0, float32(2.0), float32(3.0))).(*Leaf)
	tree2R.SetRight(tree2RR)
	tree2.SetRight(tree2R)
	tree3 := NewAssert(t).SucceedNew(NewLeaf(0, -3.5, float32(0.0), float32(1.0))).(*Leaf)
	tree3.SetRight(NewAssert(t).SucceedNew(NewLeaf(3, 0.0, float32(1.0), float32(2.0))).(*Leaf))
	tree4 := NewAssert(t).SucceedNew(NewLeaf(0, -1.5, float32(0.0), float32(1.0))).(*Leaf)
	tree4.SetRight(NewAssert(t).SucceedNew(NewLeaf(4, 0.0, float32(1.0), float32(2.0))).(*Leaf))
	forest := NewForest()
	NewAssert(t).SucceedWithoutError(forest.Enqueue(tree1))
	NewAssert(t).SucceedWithoutError(forest.Enqueue(tree2))
	NewAssert(t).SucceedWithoutError(forest.Enqueue(tree3))
	NewAssert(t).SucceedWithoutError(forest.Enqueue(tree4))
	NewAssert(t, []interface{}{float32(1.0), float32(1.0), float32(2.0), float32(0.0)}).EqualWithoutError(forest.Predict(x))
}

func TestForestEnqueueDequeue(t *testing.T) {
	x := DenseFeature{-2.0, -1.0, 0.0, 1.0, 2.0, 3.0}

	// Each tree should be unbalanced for testing
	tree1 := NewAssert(t).SucceedNew(NewLeaf(0, -2.5, float32(0.0), float32(1.0))).(*Leaf)
	tree1.SetRight(NewAssert(t).SucceedNew(NewLeaf(1, 0.0, float32(1.0), float32(2.0))).(*Leaf))
	tree2 := NewAssert(t).SucceedNew(NewLeaf(0, -4.5, float32(0.0), float32(1.0))).(*Leaf)
	tree2R := NewAssert(t).SucceedNew(NewLeaf(2, 0.0, float32(1.0), float32(2.0))).(*Leaf)
	tree2RR := NewAssert(t).SucceedNew(NewLeaf(3, 0.0, float32(2.0), float32(3.0))).(*Leaf)
	tree2R.SetRight(tree2RR)
	tree2.SetRight(tree2R)
	tree3 := NewAssert(t).SucceedNew(NewLeaf(0, -3.5, float32(0.0), float32(1.0))).(*Leaf)
	tree3.SetRight(NewAssert(t).SucceedNew(NewLeaf(3, 0.0, float32(1.0), float32(2.0))).(*Leaf))
	tree4 := NewAssert(t).SucceedNew(NewLeaf(0, -1.5, float32(0.0), float32(1.0))).(*Leaf)
	tree4.SetRight(NewAssert(t).SucceedNew(NewLeaf(4, 0.0, float32(1.0), float32(2.0))).(*Leaf))
	forest := NewForest()
	NewAssert(t).SucceedWithoutError(forest.Enqueue(tree1))
	NewAssert(t).SucceedWithoutError(forest.Enqueue(tree2))
	NewAssert(t).SucceedWithoutError(forest.Enqueue(tree3))
	NewAssert(t).SucceedWithoutError(forest.Enqueue(tree4))
	NewAssert(t, []interface{}{float32(1.0), float32(1.0), float32(2.0), float32(0.0)}).EqualWithoutError(forest.Predict(x))
	forest.Dequeue()
	NewAssert(t, []interface{}{float32(1.0), float32(2.0), float32(0.0)}).EqualWithoutError(forest.Predict(x))
	NewAssert(t).SucceedWithoutError(forest.Enqueue(tree1))
	NewAssert(t, []interface{}{float32(1.0), float32(2.0), float32(0.0), float32(1.0)}).EqualWithoutError(forest.Predict(x))
	forest.Dequeue()
	forest.Dequeue()
	forest.Dequeue()
	forest.Dequeue()
	NewAssert(t, []interface{}{}).EqualWithoutError(forest.Predict(x))
	forest.Dequeue()
	NewAssert(t, []interface{}{}).EqualWithoutError(forest.Predict(x))
}

// NOTICE: If Forest supports trees with more than 64 leaves, then this test should be modified.
func TestForestEnqueueTooDeepTree(t *testing.T) {
	treeLeft := NewAssert(t).SucceedNew(NewLeaf(0, 0.0, float32(0.0), float32(0.0))).(*Leaf)
	treeRight := NewAssert(t).SucceedNew(NewLeaf(0, 0.0, float32(0.0), float32(0.0))).(*Leaf)
	childLeft, childRight := treeLeft, treeRight
	for d := 0; d < 65; d++ {
		leafLeft := NewAssert(t).SucceedNew(NewLeaf(0, 0.0, float32(0.0), float32(0.0))).(*Leaf)
		leafRight := NewAssert(t).SucceedNew(NewLeaf(0, 0.0, float32(0.0), float32(0.0))).(*Leaf)
		childLeft.SetLeft(leafLeft)
		childRight.SetRight(leafRight)
		childLeft, childRight = leafLeft, leafRight
	}
	forest := NewForest()
	NewAssert(t, "the number of leaves in the tree must not be greater than 64").ExpectError(forest.Enqueue(treeLeft))
	NewAssert(t, "the number of leaves in the tree must not be greater than 64").ExpectError(forest.Enqueue(treeRight))
}

func TestForestGetIllegalFeature(t *testing.T) {
	x := DenseFeature{-2.0, -1.0, 0.0, 1.0, 2.0}
	tree := NewAssert(t).SucceedNew(NewLeaf(5, 0.0, float32(0.0), float32(0.0))).(*Leaf)
	forest := NewForest()
	NewAssert(t).SucceedWithoutError(forest.Enqueue(tree))
	NewAssert(t, "id 5 is out of range \\[0:5\\]").ExpectError(forest.Predict(x))
}

func BenchmarkBasicEnsembleTrees(b *testing.B) {
	dim, ntrees, depth := 65536, 65536, 12
	x := NewSparseFeature(dim)
	for i := 0; i < dim; i++ {
		x.Set(FeatureID(i), float32(i))
	}
	root := NewAssert(b).SucceedNew(NewLeaf(0, -1.0, float32(0.0), float32(1.0))).(*Leaf)
	leaf := root
	for d := 1; d < depth; d++ {
		child := NewAssert(b).SucceedNew(NewLeaf(FeatureID((dim/depth)*d), -1.0, float32(d), float32(d+1))).(*Leaf)
		leaf.SetRight(child)
		leaf = child
	}
	NewAssert(b, float32(depth)).EqualWithoutError(root.Predict(x))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for t := 0; t < ntrees; t++ {
			root.Predict(x)
		}
	}
}

func BenchmarkForest(b *testing.B) {
	dim, ntrees, depth := 65536, 65536, 12
	x := NewSparseFeature(dim)
	for i := 0; i < dim; i++ {
		x.Set(FeatureID(i), float32(i))
	}
	root := NewAssert(b).SucceedNew(NewLeaf(0, -1.0, float32(0.0), float32(1.0))).(*Leaf)
	leaf := root
	for d := 1; d < depth; d++ {
		child := NewAssert(b).SucceedNew(NewLeaf(FeatureID((dim/depth)*d), -1.0, float32(d), float32(d+1))).(*Leaf)
		leaf.SetRight(child)
		leaf = child
	}
	NewAssert(b, float32(depth)).EqualWithoutError(root.Predict(x))
	trees := make([]*Leaf, ntrees)
	forest := NewForest()
	for t := 0; t < ntrees; t++ {
		trees[t] = root
	}
	NewAssert(b).SucceedWithoutError(forest.Enqueue(trees...))
	y := make([]interface{}, ntrees)
	for t := 0; t < ntrees; t++ {
		y[t] = float32(depth)
	}
	NewAssert(b, y).EqualWithoutError(forest.Predict(x))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		forest.Predict(x)
	}
}
