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

func PolarToCartesian(matrix *dsputils.Matrix) (*dsputils.Matrix, error) {
	dims := matrix.Dimensions()
	if len(dims) > 2 {
		return nil, fmt.Errorf("CartesianToPolar expects 2 dimensions, found %d", len(dims))
	}

	// Clarify size of each dimension
	lenX, lenY, _ := func() (int, int, int) {
		// Reminder: matrices are matrix[y][x]
		return dims[1], dims[0], int(math.Hypot(float64(dims[0]), float64(dims[1])))
	}()

	cartesianMatrix := dsputils.MakeEmptyMatrix([]int{lenY, lenX})

	// Reminder: matrices are matrix[y][x]
	for cartesianY := 0; cartesianY < lenY; cartesianY++ {
		for cartesianX := 0; cartesianX < lenX; cartesianX++ {

			polarX := int(float64(cartesianX) * math.Cos(float64(cartesianY)*0.5*math.Pi/float64(lenY)))
			polarY := int(float64(cartesianX) * math.Sin(float64(cartesianY)*0.5*math.Pi/float64(lenY)))

			var value complex128
			if polarX >= lenX || polarY >= lenY {
				value = matrix.Value([]int{lenY - 1, lenX - 1})
			} else if polarX < 0 || polarY < 0 {
				value = matrix.Value([]int{0, 0})
			} else {
				value = matrix.Value([]int{polarY, polarX})
			}

			cartesianMatrix.SetValue(value, []int{cartesianY, cartesianX})
		}
	}

	return cartesianMatrix, nil
}

func CartesianToPolar(matrix *dsputils.Matrix) (*dsputils.Matrix, error) {
	dims := matrix.Dimensions()
	if len(dims) > 2 {
		return nil, fmt.Errorf("CartesianToPolar expects 2 dimensions, found %d", len(dims))
	}

	// Clarify size of each dimension
	lenX, lenY, _ := func() (int, int, int) {
		// Reminder: matrices are matrix[y][x]
		return dims[1], dims[0], int(math.Hypot(float64(dims[0]), float64(dims[1])))
	}()

	polarMatrix := dsputils.MakeEmptyMatrix([]int{lenY, lenX})

	// Reminder: matrices are matrix[y][x]
	for polarY := 0; polarY < lenY; polarY++ {
		for polarX := 0; polarX < lenX; polarX++ {

			cartesianX := int(math.Hypot(float64(polarX), float64(polarY)))
			cartesianY := int(math.Atan2(float64(polarY), float64(polarX)) * 2.0 / math.Pi * float64(lenY))

			var value complex128
			if cartesianX >= lenX || cartesianY >= lenY {
				value = matrix.Value([]int{lenY - 1, lenX - 1})
			} else if cartesianX < 0 || cartesianY < 0 {
				value = matrix.Value([]int{0, 0})
			} else {
				value = matrix.Value([]int{cartesianY, cartesianX})
			}

			polarMatrix.SetValue(value, []int{polarY, polarX})
		}
	}

	return polarMatrix, nil
}

func Translate(matrix *dsputils.Matrix) (*dsputils.Matrix, error) {
	dims := matrix.Dimensions()

	if len(dims) > 2 {
		return nil, fmt.Errorf("CartesianToPolar expects 2 dimensions, found %d", len(dims))
	}

	// Convert matrix to polar coordinates
	translatedMatrix := dsputils.MakeEmptyMatrix([]int{dims[0], dims[1]})

	for i := 0; i < dims[0]; i++ {
		for j := 0; j < dims[1]; j++ {
			oldX := (dims[0] + i + dims[0]/2) % dims[0]
			oldY := (dims[1] + j - dims[1]/2) % dims[1]

			translatedMatrix.SetValue(matrix.Value([]int{oldX, oldY}), []int{i, j})
		}
	}

	return translatedMatrix, nil
}
