package main

import (
	"testing"
)

func TestShouldContainAllDigits(t *testing.T) {
	// when
	var ds DigitSet = All()
	// then
	for digit := uint(1); digit <= uint(9); digit++ {
		if !ds.Contains(digit) {
			t.Errorf("DigitSet should contain all digits (mising: %s).", digit)
		}
	}
}

func TestAddToDigitSet(t *testing.T) {
	// given
	var ds DigitSet = 0
	// when
	ds.Add(3)
	// then
	if !ds.Contains(3) {
		t.Error("DigitSet should contain 3.")
	}
}

func TestAddingToDigitSetShouldBeIdempotent(t *testing.T) {
	// given
	var ds DigitSet = 0
	// when
	ds.Add(3)
	ds.Add(3)
	ds.Add(3)
	// then
	if !ds.Contains(3) {
		t.Error("DigitSet should contain 3.")
	}
}

func TestRemoveFromDigitSet(t *testing.T) {
	// given
	var ds DigitSet = 0
	// when
	ds.Add(4)
	ds.Remove(4)
	// then
	if ds.Contains(4) {
		t.Error("DigitSet should not contain 4.")
	}
}

func TestRemovingFromDigitSetShouldBeIdempotent(t *testing.T) {
	// given
	var ds DigitSet = 0
	// when
	ds.Add(4)
	ds.Remove(4)
	ds.Remove(4)
	ds.Remove(4)
	// then
	if ds.Contains(4) {
		t.Error("DigitSet should not contain 4.")
	}
}

func TestDigitSetShouldHaveUniqueValue(t *testing.T) {
	// given
	var ds DigitSet = 0
	// when
	ds.Add(4)
	// then
	if value, err := ds.Value(); value != 4 || err != nil {
		t.Error("DigitSet should contain unique value 4.")
	}
}

func TestDigitSetShouldHaveMultipleValues(t *testing.T) {
	// given
	var ds DigitSet = 0
	// when
	ds.Add(4)
	ds.Add(5)
	// then
	if _, err := ds.Value(); err == nil {
		t.Error("DigitSet should contain multiple values 4 and 5.")
	}
}
