//go:generate go run ./make/musable.go
package musgotest

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/ymz-ncnk/musgo/errs"
	"github.com/ymz-ncnk/musgotest/pkg"
	mypkg "github.com/ymz-ncnk/musgotest/pkgpath"
)

func TestEquality(t *testing.T) {
	// Marshaling
	st := makeValidMyStruct()
	buf := make([]byte, st.SizeMUS())
	st.MarshalMUS(buf)

	// Unmarshaling
	var ast MyStruct
	_, err := ast.UnmarshalMUS(buf)
	if err != nil {
		t.Error(err)
	}

	// Are they equal?
	st.Str = ""
	if !reflect.DeepEqual(st, ast) {
		t.Error("something went wrong")
	}
}

func TestInnerStructValidation(t *testing.T) {
	// Marshaling
	st := makeMyStruct(true, false)
	buf := make([]byte, st.SizeMUS())
	st.MarshalMUS(buf)

	// Unmarshaling
	var ast MyStruct
	_, err := ast.UnmarshalMUS(buf)
	if err == nil {
		t.Error(err)
	}

	// Is error cause equal to ErrInvalidInnerStruct?
	fieldErr, ok := err.(errs.FieldError)
	if !ok {
		t.Error("wrong error")
	}
	if fieldErr.Cause() != ErrInvalidInnerStruct {
		t.Error("wrong cause")
	}
}

func TestMapFieldValidation(t *testing.T) {
	// Marshaling
	st := makeMyStruct(false, true)
	buf := make([]byte, st.SizeMUS())
	st.MarshalMUS(buf)

	// Unmarshaling
	var ast MyStruct
	_, err := ast.UnmarshalMUS(buf)
	if err == nil {
		t.Error(err)
	}

	// Is error cause equal to ErrElementsSumBiggerThenTen?
	fieldErr, ok := err.(errs.FieldError)
	if !ok {
		t.Error("wrong error")
	}
	if fieldErr.FieldName() != "Map" {
		t.Error("wrong error fieldName")
	}
	if fieldErr.Cause() != pkg.ErrElementsSumBiggerThenTen {
		t.Error("wrong cause")
	}
}

func TestSuffix(t *testing.T) {
	mb := mypkg.MyByte(3)
	buf := make([]byte, mb.SizeAMUS())
	mb.MarshalAMUS(buf)
	var amb mypkg.MyByte
	_, err := amb.UnmarshalAMUS(buf)
	if err != nil {
		t.Error(err)
	}
	if mb != amb {
		t.Error("something went wrong")
	}
}

func makeValidMyStruct() MyStruct {
	return makeMyStruct(false, false)
}

func makeMyStruct(ivalidInnerStruct bool, invalidMap bool) MyStruct {
	ist := MyInnerStruct{
		Number: 5,
		Str:    "world",
	}
	if ivalidInnerStruct {
		ist.Str = "5"
	}
	m := map[string]string{"one": "1", "two": "2"}
	if invalidMap {
		m["one"] = "8"
		m["two"] = "5"
	}
	number := 4
	money := rand.Float32()
	return NewMyStruct(&ist, m, "bird", &number, money)
}
