package pkg

import (
	"errors"
	"strconv"

	"github.com/ymz-ncnk/musgotest/validators"
)

var ErrElementsSumBiggerThenTen = errors.New("elements sum is bigger then ten")

type MyInt int
type MyMap map[string]string

// Validator of the MyInt type.
func ValidateMyInt(number *MyInt) error {
	return validators.BiggerThenTen(int(*number))
}

// Validator of the MyMap type.
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
