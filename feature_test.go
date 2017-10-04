package confeito

import (
	"testing"

	"github.com/hiro4bbh/go-assert"
)

func TestDenseFeature(t *testing.T) {
	x := DenseFeature{-2.0, -1.0, 0.0, 1.0, 2.0}
	goassert.New(t, float32(-2.0)).EqualWithoutError(x.Get(0))
	goassert.New(t, float32(-1.0)).EqualWithoutError(x.Get(1))
	goassert.New(t, float32(0.0)).EqualWithoutError(x.Get(2))
	goassert.New(t, float32(1.0)).EqualWithoutError(x.Get(3))
	goassert.New(t, float32(2.0)).EqualWithoutError(x.Get(4))
	goassert.New(t, "id 5 is out of range \\[0:5\\]").ExpectError(x.Get(5))
	goassert.New(t, "id must be legal one").ExpectError(x.Get(_FEATURE_ID_ILLEGAL))
	goassert.New(t, 5).Equal(x.Size())
}

func TestSparseFeature(t *testing.T) {
	x := NewSparseFeature(5)
	goassert.New(t).SucceedWithoutError(x.Set(0, -2.0))
	goassert.New(t).SucceedWithoutError(x.Set(1, -1.0))
	goassert.New(t).SucceedWithoutError(x.Set(2, 0.0))
	goassert.New(t).SucceedWithoutError(x.Set(4, 2.0))
	goassert.New(t, "id 5 is out of range \\[0:5\\]").ExpectError(x.Set(5, 3.0))
	goassert.New(t, "id must be legal one").ExpectError(x.Set(_FEATURE_ID_ILLEGAL, 3.0))
	goassert.New(t, float32(-2.0)).EqualWithoutError(x.Get(0))
	goassert.New(t, float32(-1.0)).EqualWithoutError(x.Get(1))
	goassert.New(t, float32(0.0)).EqualWithoutError(x.Get(2))
	goassert.New(t, float32(0.0)).EqualWithoutError(x.Get(3))
	goassert.New(t, float32(2.0)).EqualWithoutError(x.Get(4))
	goassert.New(t, "id 5 is out of range \\[0:5\\]").ExpectError(x.Get(5))
	goassert.New(t, "id must be legal one").ExpectError(x.Get(_FEATURE_ID_ILLEGAL))
	goassert.New(t, 5).Equal(x.Size())
}
