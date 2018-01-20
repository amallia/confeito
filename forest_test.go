package confeito

import (
	"testing"

	"github.com/hiro4bbh/go-assert"
)

func TestForestEnqueueOnly(t *testing.T) {
	x := DenseFeatureVector{-2.0, -1.0, 0.0, 1.0, 2.0, 3.0}

	// Each tree should be unbalanced for testing
	tree1 := goassert.New(t).SucceedNew(NewLeaf(0, -2.5, float32(0.0), float32(1.0))).(*Leaf)
	tree1.SetRight(goassert.New(t).SucceedNew(NewLeaf(1, 0.0, float32(1.0), float32(2.0))).(*Leaf))
	tree2 := goassert.New(t).SucceedNew(NewLeaf(0, -4.5, float32(0.0), float32(1.0))).(*Leaf)
	tree2R := goassert.New(t).SucceedNew(NewLeaf(2, 0.0, float32(1.0), float32(2.0))).(*Leaf)
	tree2RR := goassert.New(t).SucceedNew(NewLeaf(3, 0.0, float32(2.0), float32(3.0))).(*Leaf)
	tree2R.SetRight(tree2RR)
	tree2.SetRight(tree2R)
	tree3 := goassert.New(t).SucceedNew(NewLeaf(0, -3.5, float32(0.0), float32(1.0))).(*Leaf)
	tree3.SetRight(goassert.New(t).SucceedNew(NewLeaf(3, 0.0, float32(1.0), float32(2.0))).(*Leaf))
	tree4 := goassert.New(t).SucceedNew(NewLeaf(0, -1.5, float32(0.0), float32(1.0))).(*Leaf)
	tree4.SetRight(goassert.New(t).SucceedNew(NewLeaf(4, 0.0, float32(1.0), float32(2.0))).(*Leaf))
	forest := NewForest()
	goassert.New(t).SucceedWithoutError(forest.Enqueue(tree1))
	goassert.New(t).SucceedWithoutError(forest.Enqueue(tree2))
	goassert.New(t).SucceedWithoutError(forest.Enqueue(tree3))
	goassert.New(t).SucceedWithoutError(forest.Enqueue(tree4))
	goassert.New(t, []interface{}{float32(1.0), float32(1.0), float32(2.0), float32(0.0)}).EqualWithoutError(forest.Predict(x))
}

func TestForestEnqueueDequeue(t *testing.T) {
	x := DenseFeatureVector{-2.0, -1.0, 0.0, 1.0, 2.0, 3.0}

	// Each tree should be unbalanced for testing
	tree1 := goassert.New(t).SucceedNew(NewLeaf(0, -2.5, float32(0.0), float32(1.0))).(*Leaf)
	tree1.SetRight(goassert.New(t).SucceedNew(NewLeaf(1, 0.0, float32(1.0), float32(2.0))).(*Leaf))
	tree2 := goassert.New(t).SucceedNew(NewLeaf(0, -4.5, float32(0.0), float32(1.0))).(*Leaf)
	tree2R := goassert.New(t).SucceedNew(NewLeaf(2, 0.0, float32(1.0), float32(2.0))).(*Leaf)
	tree2RR := goassert.New(t).SucceedNew(NewLeaf(3, 0.0, float32(2.0), float32(3.0))).(*Leaf)
	tree2R.SetRight(tree2RR)
	tree2.SetRight(tree2R)
	tree3 := goassert.New(t).SucceedNew(NewLeaf(0, -3.5, float32(0.0), float32(1.0))).(*Leaf)
	tree3.SetRight(goassert.New(t).SucceedNew(NewLeaf(3, 0.0, float32(1.0), float32(2.0))).(*Leaf))
	tree4 := goassert.New(t).SucceedNew(NewLeaf(0, -1.5, float32(0.0), float32(1.0))).(*Leaf)
	tree4.SetRight(goassert.New(t).SucceedNew(NewLeaf(4, 0.0, float32(1.0), float32(2.0))).(*Leaf))
	forest := NewForest()
	goassert.New(t).SucceedWithoutError(forest.Enqueue(tree1))
	goassert.New(t).SucceedWithoutError(forest.Enqueue(tree2))
	goassert.New(t).SucceedWithoutError(forest.Enqueue(tree3))
	goassert.New(t).SucceedWithoutError(forest.Enqueue(tree4))
	goassert.New(t, []interface{}{float32(1.0), float32(1.0), float32(2.0), float32(0.0)}).EqualWithoutError(forest.Predict(x))
	forest.Dequeue()
	goassert.New(t, []interface{}{float32(1.0), float32(2.0), float32(0.0)}).EqualWithoutError(forest.Predict(x))
	goassert.New(t).SucceedWithoutError(forest.Enqueue(tree1))
	goassert.New(t, []interface{}{float32(1.0), float32(2.0), float32(0.0), float32(1.0)}).EqualWithoutError(forest.Predict(x))
	forest.Dequeue()
	forest.Dequeue()
	forest.Dequeue()
	forest.Dequeue()
	goassert.New(t, []interface{}{}).EqualWithoutError(forest.Predict(x))
	forest.Dequeue()
	goassert.New(t, []interface{}{}).EqualWithoutError(forest.Predict(x))
}

// NOTICE: If Forest supports trees with more than 64 leaves, then this test should be modified.
func TestForestEnqueueTooDeepTree(t *testing.T) {
	treeLeft := goassert.New(t).SucceedNew(NewLeaf(0, 0.0, float32(0.0), float32(0.0))).(*Leaf)
	treeRight := goassert.New(t).SucceedNew(NewLeaf(0, 0.0, float32(0.0), float32(0.0))).(*Leaf)
	childLeft, childRight := treeLeft, treeRight
	for d := 0; d < 65; d++ {
		leafLeft := goassert.New(t).SucceedNew(NewLeaf(0, 0.0, float32(0.0), float32(0.0))).(*Leaf)
		leafRight := goassert.New(t).SucceedNew(NewLeaf(0, 0.0, float32(0.0), float32(0.0))).(*Leaf)
		childLeft.SetLeft(leafLeft)
		childRight.SetRight(leafRight)
		childLeft, childRight = leafLeft, leafRight
	}
	forest := NewForest()
	goassert.New(t, "the number of leaves in the tree must not be greater than 64").ExpectError(forest.Enqueue(treeLeft))
	goassert.New(t, "the number of leaves in the tree must not be greater than 64").ExpectError(forest.Enqueue(treeRight))
}

func BenchmarkBasicEnsembleTrees(b *testing.B) {
	dim, ntrees, depth := 65536, 65536, 12
	x := make(DenseFeatureVector, dim)
	for i := 0; i < dim; i++ {
		x[i] = float32(i)
	}
	root := goassert.New(b).SucceedNew(NewLeaf(0, -1.0, float32(0.0), float32(1.0))).(*Leaf)
	leaf := root
	for d := 1; d < depth; d++ {
		child := goassert.New(b).SucceedNew(NewLeaf(FeatureID((dim/depth)*d), -1.0, float32(d), float32(d+1))).(*Leaf)
		leaf.SetRight(child)
		leaf = child
	}
	goassert.New(b, float32(depth)).EqualWithoutError(root.Predict(x))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for t := 0; t < ntrees; t++ {
			root.Predict(x)
		}
	}
}

func BenchmarkForest(b *testing.B) {
	dim, ntrees, depth := 65536, 65536, 12
	x := make(DenseFeatureVector, dim)
	for i := 0; i < dim; i++ {
		x[i] = float32(i)
	}
	root := goassert.New(b).SucceedNew(NewLeaf(0, -1.0, float32(0.0), float32(1.0))).(*Leaf)
	leaf := root
	for d := 1; d < depth; d++ {
		child := goassert.New(b).SucceedNew(NewLeaf(FeatureID((dim/depth)*d), -1.0, float32(d), float32(d+1))).(*Leaf)
		leaf.SetRight(child)
		leaf = child
	}
	goassert.New(b, float32(depth)).EqualWithoutError(root.Predict(x))
	trees := make([]*Leaf, ntrees)
	forest := NewForest()
	for t := 0; t < ntrees; t++ {
		trees[t] = root
	}
	goassert.New(b).SucceedWithoutError(forest.Enqueue(trees...))
	y := make([]interface{}, ntrees)
	for t := 0; t < ntrees; t++ {
		y[t] = float32(depth)
	}
	goassert.New(b, y).EqualWithoutError(forest.Predict(x))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		forest.Predict(x)
	}
}
