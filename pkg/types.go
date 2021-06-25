package pkg

import (
	"errors"
	"strconv"

	"github.com/ymz-ncnk/musgotest/validators"
)

// ErrElementsSumBiggerThenTen happens if sum of elements is bigger than ten.
var ErrElementsSumBiggerThenTen = errors.New("elements sum is bigger then ten")

// MyInt is an alias of int.
type MyInt int

// MyMap is an alias of map[string ]string
type MyMap map[string]string

// ValidateMyInt validates the MyInt type.
func ValidateMyInt(number *MyInt) error {
	return validators.BiggerThenTen(int(*number))
}

// ValidateMyMap valiadtes the MyMap type.
func ValidateMyMap(sl *MyMap) error {
	sum := 0
	for _, el := range *sl {
		n, err := strconv.Atoi(el)
		if err == nil {
			sum += n
		}
	}
	if sum > 10 {
		return ErrElementsSumBiggerThenTen
	}
	return nil
}
