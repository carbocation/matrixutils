/*
matrixutils provides utilities for working with dsputils and fft packages found at
github.com/mjibson/go-dsp/dsputils and github.com/mjibson/go-dsp/fft, respectively.
*/
package matrixutils

import (
	"fmt"
	"math"

	"github.com/mjibson/go-dsp/dsputils"
)

func CartesianToPolar(matrix *dsputils.Matrix) (*dsputils.Matrix, error) {
	dims := matrix.Dimensions()

	if len(dims) > 2 {
		return nil, fmt.Errorf("CartesianMatrixToPolar expects 2 dimensions, found %d", len(dims))
	}

	// Convert matrix to polar coordinates
	polarMatrix := dsputils.MakeEmptyMatrix([]int{dims[0], dims[1]})

	cartesian := []int{0, 0} // Slice providing cartesian coordinates to lookup value from original matrix
	polar := []int{0, 0}     // Slice providing polar coordinates to set that value
	for i := 0; i < dims[0]; i++ {
		for j := 0; j < dims[1]; j++ {
			radius := math.Hypot(float64(i), float64(j))
			angle := math.Atan(float64(i) / float64(j))

			// Mod to wraparound values that exceed the limits
			// Otherwise use conditions to set excessive values to something else
			x := int(radius) % dims[0]
			y := int(angle) % dims[1]

			cartesian[0], cartesian[1] = x, y
			polar[0], polar[1] = i, j

			polarMatrix.SetValue(matrix.Value(cartesian), polar)
		}
	}

	return polarMatrix, nil
}
