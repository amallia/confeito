package confeito

import (
	"fmt"
	"math/bits"
	"sort"
)

// This type implements interface sort.Interface.
type forestFeature struct {
	thresholds []float32
	treeIDs    []int
	bvs        []uint64
}

// See sort.Interface.
func (ff *forestFeature) Len() int {
	return len(ff.thresholds)
}

// See sort.Interface.
func (ff *forestFeature) Less(i, j int) bool {
	return ff.thresholds[i] < ff.thresholds[j]
}

// See sort.Interface.
func (ff *forestFeature) Swap(i, j int) {
	ff.thresholds[i], ff.thresholds[j] = ff.thresholds[j], ff.thresholds[i]
	ff.treeIDs[i], ff.treeIDs[j] = ff.treeIDs[j], ff.treeIDs[i]
	ff.bvs[i], ff.bvs[j] = ff.bvs[j], ff.bvs[i]
}

type forestTree struct {
	values []interface{}
}

// Forest is a ensemble of tree (*Leaf).
// This is designed to compact and fast online prediction.
// Thus, there is no way to modify each tree, and users can enqueue/dequeue an tree, or get predicted values.
//
// Result of prediction is slice of the value predicted by each tree.
// This design enables users to use the predicted values for estimators weighted arbitrarily.
//
// NOTICE: Currently, Forest supports only trees having at most 64 terminal leaves.
type Forest struct {
	ntrees   int
	features map[FeatureID]*forestFeature
	trees    []*forestTree
}

// NewForest returns a new empty Forest.
func NewForest() *Forest {
	return &Forest{
		ntrees:   0,
		features: make(map[FeatureID]*forestFeature),
		trees:    []*forestTree{},
	}
}

// Dequeue dequeues the first enqueued tree from forest.
//
// This would be too slow because the implementation is not designed for frequent dequeues.
func (forest *Forest) Dequeue() {
	if len(forest.trees) == 0 {
		return
	}
	for _, feature := range forest.features {
		for p := 0; p < len(feature.treeIDs); {
			if feature.treeIDs[p] == 0 {
				feature.thresholds = append(feature.thresholds[:p], feature.thresholds[p+1:]...)
				feature.treeIDs = append(feature.treeIDs[:p], feature.treeIDs[p+1:]...)
				feature.bvs = append(feature.bvs[:p], feature.bvs[p+1:]...)
			} else {
				feature.treeIDs[p]--
				p++
			}
		}
	}
	forest.trees = forest.trees[1:]
}

func (forest *Forest) registerLeaf(leaf *Leaf, treeID int) (nleft, nright int, err error) {
	if leaf.IsTerminal() {
		value, _ := leaf.Value()
		tree := forest.trees[treeID]
		tree.values = append(tree.values, value)
		nleft = 1
		return
	}
	if rightLeaf := leaf.Right(); rightLeaf != nil {
		nleftAtRight, nrightAtRight, e := forest.registerLeaf(rightLeaf, treeID)
		if e != nil {
			err = e
			return
		}
		nright += nleftAtRight + nrightAtRight
	}
	if leftLeaf := leaf.Left(); leftLeaf != nil {
		nleftAtLeft, nrightAtLeft, e := forest.registerLeaf(leftLeaf, treeID)
		if e != nil {
			err = e
			return
		}
		nleft += nleftAtLeft + nrightAtLeft
	}
	nleaves := nleft + nright
	if nleaves > 64 {
		err = fmt.Errorf("the number of leaves in the tree must not be greater than 64")
		return
	}
	featureID, threshold, _ := leaf.Threshold()
	feature, ok := forest.features[featureID]
	if !ok {
		feature = &forestFeature{
			thresholds: []float32{},
			treeIDs:    []int{},
			bvs:        []uint64{},
		}
		forest.features[featureID] = feature
	}
	feature.thresholds = append(feature.thresholds, threshold)
	feature.treeIDs = append(feature.treeIDs, treeID)
	bv := ((^uint64(0)) & ^((1 << uint64(nleaves)) - 1)) | (((1 << uint64(nleaves)) | (1 << uint64(nright))) - 1)
	feature.bvs = append(feature.bvs, bv)
	return
}

func (forest *Forest) registerTree(treeRoot *Leaf) error {
	tree := &forestTree{
		values: []interface{}{},
	}
	treeID := len(forest.trees)
	forest.trees = append(forest.trees, tree)
	_, _, err := forest.registerLeaf(treeRoot, treeID)
	return err
}

// Enqueue enqueues the given trees to forest in order.
//
// This function returns an error if the number of leaves in tree is greater than 64.
func (forest *Forest) Enqueue(trees ...*Leaf) error {
	for _, tree := range trees {
		if err := forest.registerTree(tree); err != nil {
			return err
		}
	}
	for _, feature := range forest.features {
		sort.Sort(feature)
	}
	return nil
}

// Predict returns a slice of the value predicted by each tree of forest.
//
// This function returns an error at getting feature values of x.
func (forest *Forest) Predict(x FeatureVector) ([]interface{}, error) {
	bvs := make([]uint64, len(forest.trees))
	for t := 0; t < len(bvs); t++ {
		bvs[t] = (1 << uint64(len(forest.trees[t].values))) - 1
	}
	for i, feature := range forest.features {
		featureValue, _ := x.Get(FeatureID(i))
		left, right := 0, len(feature.thresholds)
		for left < right {
			middle := (left + right) / 2
			if feature.thresholds[middle] < featureValue {
				left = middle + 1
			} else {
				right = middle
			}
		}
		for p := 0; p < right; p++ {
			treeID := feature.treeIDs[p]
			bvs[treeID] &= feature.bvs[p]
		}
	}
	values := make([]interface{}, len(forest.trees))
	for t, tree := range forest.trees {
		leafID := bits.Len64(bvs[t]) - 1
		values[t] = tree.values[leafID]
	}
	return values, nil
}
