package musgotest

import (
	"errors"
	"strconv"

	"github.com/ymz-ncnk/musgotest/pkg"
)

// ErrInvalidInnerStruct is an ValidateInnerStruct validator error.
var ErrInvalidInnerStruct = errors.New("fields Number and Str have the same " +
	"value")

// For this structure, we want to add MUS format support. To do this, all
// custom types of this structure should support the MUS format too.
type MyStruct struct {
	// ValidateInnerStruct function will receive a pointer.
	*MyInnerStruct `mus:"ValidateInnerStruct"`

	// Map validators like `mus:,3,validators.NotHello` would not work
	// here. Instead, you should set them for the pkg.MyMap type with help
	// of a MusGo.GenerateAsAlias() method.
	Map pkg.MyMap

	// Str field will be skipped.
	Str string `mus:"-"`

	// Private fields are handled as well.
	number *int
}

type MyInnerStruct struct {
	// Note, pkg.MyInt is from an another package.
	Number pkg.MyInt
	Str    string
}

func NewMyStruct(ist *MyInnerStruct, m pkg.MyMap, str string,
	number *int) MyStruct {
	return MyStruct{
		MyInnerStruct: ist,
		Map:           m,
		Str:           str,
		number:        number,
	}
}

// Validator of the InnerStruct type.
func ValidateInnerStruct(st *MyInnerStruct) error {
	n, err := strconv.Atoi(st.Str)
	if err != nil {
		return nil
	}
	if int(st.Number) == n {
		return ErrInvalidInnerStruct
	}
	return nil
}
