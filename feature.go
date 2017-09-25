package confeito

import "fmt"

// Type of feature IDs is uint32.
type FeatureID uint32

// Feature ID for an illegal value.
const _FEATURE_ID_ILLEGAL = ^FeatureID(0)

// Feature is the interface for data points.
//
// TODO: Consider the way to handle non-float32 features.
type Feature interface {
	// Get returns the value of the feature.
	//
	// The function shoud return an error if id is illegal one.
	Get(id FeatureID) (value float32, err error)
	// Size returns the size of the feature.
	Size() int
}

// DenseFeature is a type for data points having dense features.
// This type implements interface Feature.
type DenseFeature []float32

// See Feature interface.
func (feature DenseFeature) Get(id FeatureID) (value float32, err error) {
	if id == _FEATURE_ID_ILLEGAL {
		err = fmt.Errorf("id must be legal one")
		return
	} else if int(id) >= len(feature) {
		err = fmt.Errorf("id %d is out of range [0:%d]", id, len(feature))
		return
	}
	return feature[id], nil
}

// See Feature interface.
func (feature DenseFeature) Size() int {
	return len(feature)
}

// SparseFeature is a type for data points having sparse features with id in [0:n].
// The default value is float32(0).
//
// This type implements interface Feature.
type SparseFeature struct {
	n int
	m map[FeatureID]float32
}

// NewSparseFeature returns an new empty SparseFeature with dimension n.
func NewSparseFeature(n int) *SparseFeature {
	return &SparseFeature{
		n: n,
		m: make(map[FeatureID]float32),
	}
}

// See Feature interface.
func (feature *SparseFeature) Get(id FeatureID) (value float32, err error) {
	if id == _FEATURE_ID_ILLEGAL {
		err = fmt.Errorf("id must be legal one")
		return
	} else if int(id) >= feature.n {
		err = fmt.Errorf("id %d is out of range [0:%d]", id, feature.n)
		return
	} else if value, ok := feature.m[id]; ok {
		return value, nil
	}
	return 0.0, nil
}

// Set sets the value at feature id.
//
// This function returns an error if id is not valid.
func (feature *SparseFeature) Set(id FeatureID, value float32) error {
	if id == _FEATURE_ID_ILLEGAL {
		return fmt.Errorf("id must be legal one")
	} else if int(id) >= feature.n {
		return fmt.Errorf("id %d is out of range [0:%d]", id, feature.n)
	}
	feature.m[id] = value
	return nil
}

// See Feature interface.
func (feature *SparseFeature) Size() int {
	return feature.n
}
