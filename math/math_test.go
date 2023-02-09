package math

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	fmt.Println("Abs: ", Abs(-45.6))

	fmt.Println("Round: ", Round(3.1), "  ", Round(3.6))

	fmt.Println("Floor: ", Floor(3.1), "  ", Floor(3.6))

	fmt.Println("Ceil: ", Ceil(3.1), "  ", Ceil(3.6))

	fmt.Println("Max: ", Max(1, 2, 3.5, 4))

	fmt.Println("Min: ", Min(1, 2, 3.5, 4))

	fmt.Println("DecBin: ", DecBin(10))

	fmt.Println("DecHex: ", DecHex(10))
}
