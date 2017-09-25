package confeito

import (
	"fmt"
	"testing"
)

func TestTerminalLeaf(t *testing.T) {
	x := DenseFeature{-2.0, -1.0, 0.0, 1.0, 2.0}

	terminal1 := NewAssert(t).SucceedNew(NewTerminalLeaf(float32(1.0))).(*Leaf)
	NewAssert(t, (*Leaf)(nil)).Equal(terminal1.GetLeft())
	NewAssert(t, (*Leaf)(nil)).Equal(terminal1.GetRight())
	NewAssert(t, "terminal leaf does not have threshold").ExpectError(terminal1.GetThreshold())
	NewAssert(t, float32(1.0)).EqualWithoutError(terminal1.GetValue())
	NewAssert(t, true).Equal(terminal1.IsTerminal())
	NewAssert(t, float32(1.0)).EqualWithoutError(terminal1.Predict(x))
	NewAssert(t, "terminal leaf cannot have left leaf").ExpectError(terminal1.SetLeft(terminal1))
	NewAssert(t, "terminal leaf cannot have right leaf").ExpectError(terminal1.SetRight(terminal1))
	NewAssert(t, "1").Equal(fmt.Sprintf("%s", terminal1))
}

func TestLeaf(t *testing.T) {
	x1 := DenseFeature{-2.0, -1.0, 0.0, 1.0, 2.0}
	x2 := DenseFeature{-2.0, -1.0, 1.0, 1.0, 2.0}
	x3 := DenseFeature{-2.0, 1.0, 0.0, 1.0, 2.0}
	x4 := DenseFeature{-2.0, 1.0, 0.0, 2.0, 2.0}
	x5 := DenseFeature{-2.0, 1.0, 0.0, 2.0, 3.0}
	x6 := DenseFeature{-2.0, 1.0, 0.0, 2.0}

	NewAssert(t, "featureID must be valid").ExpectError(NewLeaf(_FEATURE_ID_ILLEGAL, 0.0, nil, nil))

	leaf1 := NewAssert(t).SucceedNew(NewLeaf(1, -0.5, float32(1.0), float32(2.0))).(*Leaf)
	NewAssert(t, float32(1.0)).EqualWithoutError(leaf1.GetLeft().GetValue())
	NewAssert(t, float32(2.0)).EqualWithoutError(leaf1.GetRight().GetValue())
	NewAssert(t, FeatureID(1), float32(-0.5)).EqualWithoutError(leaf1.GetThreshold())
	NewAssert(t, "non-terminal leaf does not have value").ExpectError(leaf1.GetValue())
	NewAssert(t, false).Equal(leaf1.IsTerminal())
	NewAssert(t, float32(1.0)).EqualWithoutError(leaf1.Predict(x1))
	NewAssert(t, "left leaf of non-terminal leaf must not be nil").ExpectError(leaf1.SetLeft(nil))
	NewAssert(t, "right leaf of non-terminal leaf must not be nil").ExpectError(leaf1.SetRight(nil))

	leaf2 := NewAssert(t).SucceedNew(NewLeaf(2, 0.5, float32(1.0), float32(2.0))).(*Leaf)
	leaf3 := NewAssert(t).SucceedNew(NewLeaf(3, 1.5, float32(3.0), float32(4.0))).(*Leaf)
	leaf4 := NewAssert(t).SucceedNew(NewLeaf(4, 2.5, float32(4.0), float32(5.0))).(*Leaf)
	NewAssert(t).SucceedWithoutError(leaf1.SetLeft(leaf2))
	NewAssert(t).SucceedWithoutError(leaf1.SetRight(leaf3))
	NewAssert(t).SucceedWithoutError(leaf3.SetRight(leaf4))
	// (feature[1] <= -0.5
	//     ? (feature[2] <= 0.5 ? 1 : 2)
	//     : (feature[3] <= 1.5
	//           ? 3
	//           : (feature[4] <= 2.5 ? 4 : 5)
	//       )
	// )
	NewAssert(t, "(feature[1] <= -0.5 ? (feature[2] <= 0.5 ? 1 : 2) : (feature[3] <= 1.5 ? 3 : (feature[4] <= 2.5 ? 4 : 5)))").Equal(fmt.Sprintf("%s", leaf1))
	NewAssert(t, float32(1.0)).EqualWithoutError(leaf1.Predict(x1))
	NewAssert(t, float32(2.0)).EqualWithoutError(leaf1.Predict(x2))
	NewAssert(t, float32(3.0)).EqualWithoutError(leaf1.Predict(x3))
	NewAssert(t, float32(4.0)).EqualWithoutError(leaf1.Predict(x4))
	NewAssert(t, float32(5.0)).EqualWithoutError(leaf1.Predict(x5))
	NewAssert(t, "id 4 is out of range").ExpectError(leaf1.Predict(x6))
}
