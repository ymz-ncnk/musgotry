//go:generate go run ./make/musable.go
package musgotest

import (
	"reflect"
	"testing"
	"time"

	"github.com/ymz-ncnk/musgo/errs"
	"github.com/ymz-ncnk/musgotry/pkg2"
	"github.com/ymz-ncnk/musgotry/pkg3"
	"github.com/ymz-ncnk/musgotry/validators"
)

func TestEquality(t *testing.T) {
	num := 5
	st := pkg3.NewStruct(&num, time.Now().Unix())
	st.InnerStruct = &pkg3.InnerStruct{Num: 1, Str: "hello"}
	st.Str = "yellow"
	st.Map = map[int]string{1: "hello"}

	buf := make([]byte, st.SizeMUS())
	st.MarshalMUS(buf)

	var ast pkg3.Struct
	_, err := ast.UnmarshalMUS(buf)
	if err != nil {
		t.Fatal(err)
	}
	ast.Str = "yellow"
	if !reflect.DeepEqual(st, ast) {
		t.Error("something went wrong")
	}
}

func TestInnerStructValidation(t *testing.T) {
	num := 8
	st := pkg3.NewStruct(&num, time.Now().Unix())
	st.InnerStruct = &pkg3.InnerStruct{}

	buf := make([]byte, st.SizeMUS())
	st.MarshalMUS(buf)

	var ast pkg3.Struct
	_, err := ast.UnmarshalMUS(buf)
	if err == nil {
		t.Error("unexpected nil error")
	}

	fieldErr, ok := err.(errs.FieldError)
	if !ok {
		t.Error("wrong error")
	}
	if fieldErr.Cause() != pkg3.ErrEmptyInnerStruct {
		t.Error("wrong error's cause")
	}
}

func TestMapFieldValidation(t *testing.T) {
	num := 10
	st := pkg3.NewStruct(&num, time.Now().Unix())
	st.InnerStruct = &pkg3.InnerStruct{Num: 1, Str: "hello"}
	st.Map = map[int]string{-1: "red"}

	buf := make([]byte, st.SizeMUS())
	st.MarshalMUS(buf)

	var ast pkg3.Struct
	_, err := ast.UnmarshalMUS(buf)
	if err == nil {
		t.Error("unexpected nil error")
	}

	if fieldErr, ok := err.(errs.FieldError); ok {
		if fieldErr.FieldName() != "Map" {
			t.Error("wrong error fieldName")
		}
		if fieldErr.Cause() != validators.ErrKeySumNegative {
			t.Error("wrong errors's cause")
		}
	} else {
		t.Error("wrong error")
	}
}

func TestCustomSuffix(t *testing.T) {
	b := pkg2.ByteAlias(3)
	buf := make([]byte, b.SizeCUSTOM())
	b.MarshalCUSTOM(buf)

	var ab pkg2.ByteAlias
	_, err := ab.UnmarshalCUSTOM(buf)
	if err != nil {
		t.Error(err)
	}
	if b != ab {
		t.Error("something went wrong")
	}
}
