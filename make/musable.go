// +build ignore

package main

import (
	"reflect"

	"github.com/ymz-ncnk/musgo"
	"github.com/ymz-ncnk/musgotest"
	"github.com/ymz-ncnk/musgotest/pkg"
)

// main adds MUS format support for MyStruct. Note, generated files should be
// placed in the right places. Files for musgotest.MyStruct - in 'musgotest'
// package, for pkg.MyStrSlice - in 'pkg', and so on...
func main() {
	musGo, err := musgo.New()
	if err != nil {
		panic(err)
	}
	{
		var v musgotest.MyStruct
		err := musGo.Generate(reflect.TypeOf(v), false)
		if err != nil {
			panic(err)
		}
	}
	{
		var v pkg.MyMap
		err := musGo.GenerateAliasAs(reflect.TypeOf(v), false, "ValidateMyMap",
			3, "validators.NotHello", "validators.NotHello", "./pkg", "")
		if err != nil {
			panic(err)
		}
	}
	{
		var v musgotest.MyInnerStruct
		err := musGo.Generate(reflect.TypeOf(v), false)
		if err != nil {
			panic(err)
		}
	}
	{
		var v pkg.MyInt
		// Specifies validator for the alias type.
		err := musGo.GenerateAliasAs(reflect.TypeOf(v), false, "ValidateMyInt",
			0, "", "", "./pkg", "")
		if err != nil {
			panic(err)
		}
	}
}
