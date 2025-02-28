package main

/** Returns a slice of length `dy`,
each element of which is a slice of `dx`
8-bit unsigned integers.*/
func Pic(dx, dy int) [][]uint8 {
	bigSlice := make([][]uint8, dy)

	for i := range bigSlice {
		subSlice := make([]uint8, dx)
		for j := range subSlice {
			subSlice[j] = uint8(imageFunc(i, j))
		}

		bigSlice[i] = subSlice

	}

	return bigSlice
}

func imageFunc(x, y int) int {
	return (x + y) / 2
}
