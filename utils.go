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

func positiveMod(n, k int) int {
	if n < 0 {
		return k - ((-1 * n) % k)
	}

	return n % k
}

func RasterToCartesian(matrix *dsputils.Matrix) (*dsputils.Matrix, error) {
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
	for y := 0; y < lenY; y++ {
		for x := 0; x < lenX; x++ {

			cartesianX := x - lenX/2
			cartesianY := lenY/2 - y

			value := matrix.Value([]int{positiveMod(cartesianY, lenY), positiveMod(cartesianX, lenX)})

			cartesianMatrix.SetValue(value, []int{y, x})
		}
	}

	return cartesianMatrix, nil
}

// Using degrees, not radians
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

			value := matrix.Value([]int{positiveMod(polarY, lenY), positiveMod(polarX, lenX)})

			cartesianMatrix.SetValue(value, []int{cartesianY, cartesianX})
		}
	}

	return cartesianMatrix, nil
}

// Using degrees, not radians
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

			value := matrix.Value([]int{positiveMod(cartesianY, lenY), positiveMod(cartesianX, lenX)})

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
