package pkg3

import (
	"errors"

	"github.com/ymz-ncnk/musgotry/pkg1"
)

var ErrEmptyInnerStruct = errors.New("empty InnerStruct")

func NewStruct(num *int, time int64) Struct {
	return Struct{num: num, time: time}
}

type Struct struct {
	// ValidateInnerStruct function will receive a pointer.
	*InnerStruct `mus:"NotEmptyInnerStruct"`

	// Validators for map would not work here. Instead, you should set them for
	// the pkg1.MapAlias type with help of the musGo.GenerateAsAlias() method.
	Map pkg1.ValidMapAlias

	// Str field will be skipped.
	Str string `mus:"-"`

	// Private fields are handled as well.
	num *int

	// It's better to use raw encoding for int time representation.
	time int64 `mus:"#raw"`
}

type InnerStruct struct {
	// Note, pkg1.ValidIntAlias is from another package.
	Num pkg1.ValidIntAlias
	Str string
}
