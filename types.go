package musgotest

import (
	"errors"
	"strconv"
	"time"

	"github.com/ymz-ncnk/musgotest/pkg"
)

// ErrInvalidInnerStruct is an ValidateInnerStruct validator error.
var ErrInvalidInnerStruct = errors.New("fields Number and Str have the same " +
	"value")

// MyStruct is a custom structure. We want to add a MUS format support for it.
// To do this, all custom types of this structure should support the MUS format
// too.
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

	// It's better to use raw encoding for int time representation.
	time  int64   `mus:"#raw"`
	money float32 `mus:"#raw"`
}

// MyInnerStruct is a custom structure. It's a part of the MyStruct.
type MyInnerStruct struct {
	// Note, pkg.MyInt is from an another package.
	Number pkg.MyInt
	Str    string
}

// NewMyStruct returns a new MyStruct.
func NewMyStruct(ist *MyInnerStruct, m pkg.MyMap, str string,
	number *int, money float32) MyStruct {
	return MyStruct{
		MyInnerStruct: ist,
		Map:           m,
		Str:           str,
		number:        number,
		time:          time.Now().UnixNano(),
		money:         money,
	}
}

// ValidateInnerStruct validates MyInnerStruct.
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
