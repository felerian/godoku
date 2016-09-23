/*
Package digitset provides a data type for a set of digits.
*/
package digitset

import (
	"errors"
)

// DigitSet is a set of digits.
type DigitSet int

// Add a digit to this set.
func (ds *DigitSet) Add(digit uint) {
	*ds = *ds | 1<<digit
}

// Remove a digit from this set.
func (ds *DigitSet) Remove(digit uint) {
	*ds = *ds &^ (1 << digit)
}

// Contains returns true if this set contains the given digit.
func (ds *DigitSet) Contains(digit uint) bool {
	return 1 == (*ds >> digit & 1)
}

// Count returns the number of digits contained in this set.
func (ds *DigitSet) Count() uint {
	var count uint
	for i := uint(0); i < 9; i++ {
		if ds.Contains(i) {
			count++
		}
	}
	return count
}

// All creates a new DigitSet containing the digits 1 through 9.
func All() DigitSet {
	return 1022
}

// Empty returns a new empty DigitSet.
func Empty() DigitSet {
	return 0
}

// Single returns a new DigitSet containing only the given digit.
func Single(value uint) DigitSet {
	set := Empty()
	set.Add(value)
	return set
}

// Value returns the single digit in this set. If this set contains more than one digit an error is returned.
func (ds *DigitSet) Value() (uint, error) {
	var result uint
	for i := uint(1); i <= uint(9); i++ {
		if ds.Contains(i) {
			if result != 0 {
				return result, errors.New("set contains multiple values")
			}
			result = i
		}
	}
	return result, nil
}
