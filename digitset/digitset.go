package digitset

import (
	"errors"
)

type DigitSet int

func (ds *DigitSet) Add(digit uint) {
	*ds = *ds | 1<<digit
}

func (ds *DigitSet) Remove(digit uint) {
	*ds = *ds &^ (1 << digit)
}

func (ds *DigitSet) Contains(digit uint) bool {
	return 1 == (*ds >> digit & 1)
}

func All() DigitSet {
	return 1022
}

func Empty() DigitSet {
	return 0
}

func Single(value uint) DigitSet {
	set := Empty()
	set.Add(value)
	return set
}

func (ds *DigitSet) Value() (uint, error) {
	var result uint
	for i := uint(1); i <= uint(9); i++ {
		if ds.Contains(i) {
			if result != 0 {
				return result, errors.New("Set contains multiple values.")
			}
			result = i
		}
	}
	return result, nil
}
