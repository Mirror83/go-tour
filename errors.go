package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v\n", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := x / 2
	tolerance := 0.0000001
	for {
		oldZ := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-oldZ) < tolerance {
			break
		}
	}
	return z, nil
}
