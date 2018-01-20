package confeito

import (
	"fmt"
)

// Type of feature IDs is uint32.
type FeatureID uint32

// Feature ID for an illegal value.
const _FEATURE_ID_ILLEGAL = ^FeatureID(0)

// FeatureVector is the interface for a data point.
// FeatureVector interface can be used by any storage manager, so it should provide only read properties.
type FeatureVector interface {
	// Dim returns the vector dimension.
	Dim() int
	// Get returns the value of the feature.
	// The default value should be float32(0.0).
	// If id is larger than the vector dimension, then returns the default value.
	//
	// The function should return an error if id is illegal one.
	Get(id FeatureID) (value float32, err error)
}

// DenseFeatureVector is a type for a data point with dense feature values.
// This implements interface FeatureVector.
type DenseFeatureVector []float32

func (v DenseFeatureVector) Dim() int {
	return len(v)
}

func (v DenseFeatureVector) Get(id FeatureID) (value float32, err error) {
	if id == _FEATURE_ID_ILLEGAL {
		err = fmt.Errorf("id must be legal one")
		return
	} else if int(id) < len(v) {
		return v[id], nil
	}
	return 0.0, nil
}

// KeyValue is the pair of the FeatureID key and float32 value.
type KeyValue struct {
	Key   FeatureID
	Value float32
}

// SparseFeature is a type for a data point having sparse feature values.
// The features should be sorted in ascending order of its key.
//
// This implements interface FeatureVector and sort.Sort.
type SparseFeatureVector []KeyValue

func (v SparseFeatureVector) Dim() (d int) {
	for _, vpair := range v {
		if d < int(vpair.Key) {
			d = int(vpair.Key)
		}
	}
	return
}

func (v SparseFeatureVector) Get(id FeatureID) (value float32, err error) {
	if id == _FEATURE_ID_ILLEGAL {
		err = fmt.Errorf("id must be legal one")
		return
	}
	for _, vpair := range v {
		if vpair.Key == id {
			return vpair.Value, nil
		}
	}
	return 0.0, nil
}

func (v SparseFeatureVector) Len() int {
	return len(v)
}

func (v SparseFeatureVector) Less(i, j int) bool {
	return v[i].Key < v[j].Key
}

func (v SparseFeatureVector) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
