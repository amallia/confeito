package confeito

import "fmt"

// Feature ID for a terminal leaf.
const _FEATURE_ID_TERMINAL_LEAF = _FEATURE_ID_ILLEGAL

// Leaf is an element in a tree.
// It is either of non-terminal or terminal.
// If it is non-terminal, then it has left and right leaf, otherwise it has a value which can be any object (interface{}).
//
// In predicting the value of the given feature, if feature[featureID] <= threshold, then the left leaf is taken, else the right one is taken.
// This process is repeated until the cursor points a terminal leaf, and returns the value of it.
//
// Leaf is slow, because it is designed to use manipulating tree structure in training-phase or testing its correctness.
//
// TODO: Consider the way to handle non-float32 features.
type Leaf struct {
	featureID   FeatureID
	threshold   float32
	value       interface{}
	left, right *Leaf
}

// NewLeaf returns a new leaf with feature ID and threshold.
// Also, the function sets the default value of the left and right leaf.
//
// This function returns an error if featureID is FEATURE_ID_TERMINAL_LEAF.
func NewLeaf(featureID FeatureID, threshold float32, leftValue, rightValue interface{}) (*Leaf, error) {
	if featureID == _FEATURE_ID_ILLEGAL {
		return nil, fmt.Errorf("featureID must be valid")
	}
	left, _ := NewTerminalLeaf(leftValue)
	right, _ := NewTerminalLeaf(rightValue)
	return &Leaf{
		featureID: featureID,
		threshold: threshold,
		left:      left,
		right:     right,
	}, nil
}

// NewTerminalLeaf returns a new terminal leaf with value.
//
// This function returns no error currently.
func NewTerminalLeaf(value interface{}) (*Leaf, error) {
	return &Leaf{
		featureID: _FEATURE_ID_TERMINAL_LEAF,
		value:     value,
	}, nil
}

// GetLeft returns the left leaf.
// If l is terminal, then the function returns nil.
func (l *Leaf) GetLeft() *Leaf {
	return l.left
}

// GetRight returns the right leaf.
// If l is terminal, then the function returns nil.
func (l *Leaf) GetRight() *Leaf {
	return l.right
}

// GetThreshold returns the threshold with feature ID of l.
//
// This function returns an error if l is terminal.
func (l *Leaf) GetThreshold() (featureID FeatureID, threshold float32, err error) {
	if l.IsTerminal() {
		err = fmt.Errorf("terminal leaf does not have threshold")
		return
	}
	return l.featureID, l.threshold, nil
}

// GetValue returns the value of l.
//
// This function returns an error if l is not terminal.
func (l *Leaf) GetValue() (value interface{}, err error) {
	if !l.IsTerminal() {
		err = fmt.Errorf("non-terminal leaf does not have value")
		return
	}
	return l.value, nil
}

// IsTerminal returns true if l is terminal, otherwise false.
func (l *Leaf) IsTerminal() bool {
	return l.featureID == _FEATURE_ID_TERMINAL_LEAF
}

// Predict returns the predicted value of the given feature.
//
// This function returns an errors at getting feature values of x.
func (l *Leaf) Predict(x FeatureVector) (value interface{}, err error) {
	if l.IsTerminal() {
		return l.value, nil
	}
	if fvalue, _ := x.Get(l.featureID); fvalue <= l.threshold {
		return l.left.Predict(x)
	} else {
		return l.right.Predict(x)
	}
}

// SetLeft sets left leaf.
//
// This function returns an error if l is terminal, or the new leaf is nil.
func (l *Leaf) SetLeft(left *Leaf) error {
	if l.IsTerminal() {
		return fmt.Errorf("terminal leaf cannot have left leaf")
	}
	if left == nil {
		return fmt.Errorf("left leaf of non-terminal leaf must not be nil")
	}
	l.left = left
	return nil
}

// SetRight sets right leaf.
//
// This function returns an error if l is terminal, or the new leaf is nil.
func (l *Leaf) SetRight(right *Leaf) error {
	if l.IsTerminal() {
		return fmt.Errorf("terminal leaf cannot have right leaf")
	}
	if right == nil {
		return fmt.Errorf("right leaf of non-terminal leaf must not be nil")
	}
	l.right = right
	return nil
}

// String returns the human-readable string representation of l.
func (l *Leaf) String() string {
	if l.IsTerminal() {
		return fmt.Sprintf("%g", l.value)
	}
	return fmt.Sprintf("(feature[%d] <= %g ? %s : %s)", l.featureID, l.threshold, l.left, l.right)
}
