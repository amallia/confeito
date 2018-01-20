package confeito

import (
	"fmt"
	"testing"

	"github.com/hiro4bbh/go-assert"
)

func TestTerminalLeaf(t *testing.T) {
	x := DenseFeatureVector{-2.0, -1.0, 0.0, 1.0, 2.0}

	terminal1 := goassert.New(t).SucceedNew(NewTerminalLeaf(float32(1.0))).(*Leaf)
	goassert.New(t, (*Leaf)(nil)).Equal(terminal1.Left())
	goassert.New(t, (*Leaf)(nil)).Equal(terminal1.Right())
	goassert.New(t, "terminal leaf does not have threshold").ExpectError(terminal1.Threshold())
	goassert.New(t, float32(1.0)).EqualWithoutError(terminal1.Value())
	goassert.New(t, true).Equal(terminal1.IsTerminal())
	goassert.New(t, float32(1.0)).EqualWithoutError(terminal1.Predict(x))
	goassert.New(t, "terminal leaf cannot have left leaf").ExpectError(terminal1.SetLeft(terminal1))
	goassert.New(t, "terminal leaf cannot have right leaf").ExpectError(terminal1.SetRight(terminal1))
	goassert.New(t, "1").Equal(fmt.Sprintf("%s", terminal1))
}

func TestLeaf(t *testing.T) {
	x1 := DenseFeatureVector{-2.0, -1.0, 0.0, 1.0, 2.0}
	x2 := DenseFeatureVector{-2.0, -1.0, 1.0, 1.0, 2.0}
	x3 := DenseFeatureVector{-2.0, 1.0, 0.0, 1.0, 2.0}
	x4 := DenseFeatureVector{-2.0, 1.0, 0.0, 2.0, 2.0}
	x5 := DenseFeatureVector{-2.0, 1.0, 0.0, 2.0, 3.0}
	x6 := DenseFeatureVector{-2.0, 1.0, 0.0, 2.0}

	goassert.New(t, "featureID must be valid").ExpectError(NewLeaf(_FEATURE_ID_ILLEGAL, 0.0, nil, nil))

	leaf1 := goassert.New(t).SucceedNew(NewLeaf(1, -0.5, float32(1.0), float32(2.0))).(*Leaf)
	goassert.New(t, float32(1.0)).EqualWithoutError(leaf1.Left().Value())
	goassert.New(t, float32(2.0)).EqualWithoutError(leaf1.Right().Value())
	goassert.New(t, FeatureID(1), float32(-0.5)).EqualWithoutError(leaf1.Threshold())
	goassert.New(t, "non-terminal leaf does not have value").ExpectError(leaf1.Value())
	goassert.New(t, false).Equal(leaf1.IsTerminal())
	goassert.New(t, float32(1.0)).EqualWithoutError(leaf1.Predict(x1))
	goassert.New(t, "left leaf of non-terminal leaf must not be nil").ExpectError(leaf1.SetLeft(nil))
	goassert.New(t, "right leaf of non-terminal leaf must not be nil").ExpectError(leaf1.SetRight(nil))

	leaf2 := goassert.New(t).SucceedNew(NewLeaf(2, 0.5, float32(1.0), float32(2.0))).(*Leaf)
	leaf3 := goassert.New(t).SucceedNew(NewLeaf(3, 1.5, float32(3.0), float32(4.0))).(*Leaf)
	leaf4 := goassert.New(t).SucceedNew(NewLeaf(4, 2.5, float32(4.0), float32(5.0))).(*Leaf)
	goassert.New(t).SucceedWithoutError(leaf1.SetLeft(leaf2))
	goassert.New(t).SucceedWithoutError(leaf1.SetRight(leaf3))
	goassert.New(t).SucceedWithoutError(leaf3.SetRight(leaf4))
	// (feature[1] <= -0.5
	//     ? (feature[2] <= 0.5 ? 1 : 2)
	//     : (feature[3] <= 1.5
	//           ? 3
	//           : (feature[4] <= 2.5 ? 4 : 5)
	//       )
	// )
	goassert.New(t, "(feature[1] <= -0.5 ? (feature[2] <= 0.5 ? 1 : 2) : (feature[3] <= 1.5 ? 3 : (feature[4] <= 2.5 ? 4 : 5)))").Equal(fmt.Sprintf("%s", leaf1))
	goassert.New(t, float32(1.0)).EqualWithoutError(leaf1.Predict(x1))
	goassert.New(t, float32(2.0)).EqualWithoutError(leaf1.Predict(x2))
	goassert.New(t, float32(3.0)).EqualWithoutError(leaf1.Predict(x3))
	goassert.New(t, float32(4.0)).EqualWithoutError(leaf1.Predict(x4))
	goassert.New(t, float32(5.0)).EqualWithoutError(leaf1.Predict(x5))
	goassert.New(t, float32(4.0)).EqualWithoutError(leaf1.Predict(x6))
}
