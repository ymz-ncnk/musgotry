// +build ignore

package main

import (
	"reflect"

	"github.com/ymz-ncnk/musgo"
	"github.com/ymz-ncnk/musgotest"
	"github.com/ymz-ncnk/musgotest/pkg"
	mypkg "github.com/ymz-ncnk/musgotest/pkgpath"
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
		conf := musgo.NewAliasConf()
		conf.T = reflect.TypeOf(v)
		conf.Validator = "ValidateMyMap"
		conf.MaxLength = 3
		conf.ElemValidator = "validators.NotHello"
		conf.KeyValidator = "validators.NotHello"
		conf.Path = "./pkg"
		err := musGo.GenerateAliasAs(conf)
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
		conf := musgo.NewAliasConf()
		conf.T = reflect.TypeOf(v)
		conf.Validator = "ValidateMyInt"
		conf.Path = "./pkg"
		err := musGo.GenerateAliasAs(conf)
		if err != nil {
			panic(err)
		}
	}
	{
		// In this case the pkg name differs from the last folder of the pkg path.
		var v mypkg.MyByte
		conf := musgo.NewConf()
		conf.T = reflect.TypeOf(v)
		conf.Path = "./pkgpath"
		err := musGo.GenerateAs(conf)
		if err != nil {
			panic(err)
		}
	}
	{
		// In this case the pkg name differs from the last folder of the pkg path.
		var v mypkg.MyByte
		conf := musgo.NewConf()
		conf.T = reflect.TypeOf(v)
		conf.Suffix = "AMUS"
		conf.Path = "./pkgpath"
		conf.Filename = "CustomSuffixMyByte.musgen.go"
		// err := musGo.GenerateAs(reflect.TypeOf(v), false, "./pkgpath", "")
		err := musGo.GenerateAs(conf)
		if err != nil {
			panic(err)
		}
	}
}
