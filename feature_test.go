package confeito

import "testing"

func TestDenseFeature(t *testing.T) {
	x := DenseFeature{-2.0, -1.0, 0.0, 1.0, 2.0}
	NewAssert(t, float32(-2.0)).EqualWithoutError(x.Get(0))
	NewAssert(t, float32(-1.0)).EqualWithoutError(x.Get(1))
	NewAssert(t, float32(0.0)).EqualWithoutError(x.Get(2))
	NewAssert(t, float32(1.0)).EqualWithoutError(x.Get(3))
	NewAssert(t, float32(2.0)).EqualWithoutError(x.Get(4))
	NewAssert(t, "id 5 is out of range \\[0:5\\]").ExpectError(x.Get(5))
	NewAssert(t, "id must be legal one").ExpectError(x.Get(_FEATURE_ID_ILLEGAL))
	NewAssert(t, 5).Equal(x.Size())
}

func TestSparseFeature(t *testing.T) {
	x := NewSparseFeature(5)
	NewAssert(t).SucceedWithoutError(x.Set(0, -2.0))
	NewAssert(t).SucceedWithoutError(x.Set(1, -1.0))
	NewAssert(t).SucceedWithoutError(x.Set(2, 0.0))
	NewAssert(t).SucceedWithoutError(x.Set(4, 2.0))
	NewAssert(t, "id 5 is out of range \\[0:5\\]").ExpectError(x.Set(5, 3.0))
	NewAssert(t, "id must be legal one").ExpectError(x.Set(_FEATURE_ID_ILLEGAL, 3.0))
	NewAssert(t, float32(-2.0)).EqualWithoutError(x.Get(0))
	NewAssert(t, float32(-1.0)).EqualWithoutError(x.Get(1))
	NewAssert(t, float32(0.0)).EqualWithoutError(x.Get(2))
	NewAssert(t, float32(0.0)).EqualWithoutError(x.Get(3))
	NewAssert(t, float32(2.0)).EqualWithoutError(x.Get(4))
	NewAssert(t, "id 5 is out of range \\[0:5\\]").ExpectError(x.Get(5))
	NewAssert(t, "id must be legal one").ExpectError(x.Get(_FEATURE_ID_ILLEGAL))
	NewAssert(t, 5).Equal(x.Size())
}
