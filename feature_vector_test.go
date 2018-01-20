package confeito

import (
	"sort"
	"testing"

	"github.com/hiro4bbh/go-assert"
)

func TestDenseFeatureVector(t *testing.T) {
	x := DenseFeatureVector{-2.0, -1.0, 0.0, 1.0, 2.0}
	goassert.New(t, 5).Equal(x.Dim())
	goassert.New(t, float32(-2.0)).EqualWithoutError(x.Get(0))
	goassert.New(t, float32(-1.0)).EqualWithoutError(x.Get(1))
	goassert.New(t, float32(0.0)).EqualWithoutError(x.Get(2))
	goassert.New(t, float32(1.0)).EqualWithoutError(x.Get(3))
	goassert.New(t, float32(2.0)).EqualWithoutError(x.Get(4))
	goassert.New(t, float32(0.0)).EqualWithoutError(x.Get(5))
	goassert.New(t, "id must be legal one").ExpectError(x.Get(_FEATURE_ID_ILLEGAL))
}

func TestSparseFeatureVector(t *testing.T) {
	x := SparseFeatureVector{
		KeyValue{1, -1.0}, KeyValue{4, 2.0}, KeyValue{0, -2.0},
	}
	sort.Sort(x)
	goassert.New(t, SparseFeatureVector{
		KeyValue{0, -2.0}, KeyValue{1, -1.0}, KeyValue{4, 2.0},
	}).Equal(x)
	goassert.New(t, 4).Equal(x.Dim())
	goassert.New(t, float32(-2.0)).EqualWithoutError(x.Get(0))
	goassert.New(t, float32(-1.0)).EqualWithoutError(x.Get(1))
	goassert.New(t, float32(0.0)).EqualWithoutError(x.Get(2))
	goassert.New(t, float32(0.0)).EqualWithoutError(x.Get(3))
	goassert.New(t, float32(2.0)).EqualWithoutError(x.Get(4))
	goassert.New(t, float32(0.0)).EqualWithoutError(x.Get(5))
	goassert.New(t, "id must be legal one").ExpectError(x.Get(_FEATURE_ID_ILLEGAL))
}
