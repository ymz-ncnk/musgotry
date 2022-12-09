//go:build ignore

package main

import (
	"reflect"

	"github.com/ymz-ncnk/musgo"
	"github.com/ymz-ncnk/musgotry/pkg1"
	"github.com/ymz-ncnk/musgotry/pkg2"
	"github.com/ymz-ncnk/musgotry/pkg3"
)

func main() {
	musGo, err := musgo.New()
	if err != nil {
		panic(err)
	}
	{
		// Sets Validator, MaxLength and ElemValidator for pkg1.ValidMapAlias type.
		var v pkg1.ValidMapAlias
		conf := musgo.NewAliasConf()
		conf.T = reflect.TypeOf(v)
		conf.Validator = "KeySumPositive"
		conf.MaxLength = 3
		conf.ElemValidator = "validators.NotEmpty"
		conf.Path = "./pkg1"
		err := musGo.GenerateAliasAs(conf)
		if err != nil {
			panic(err)
		}
	}
	{
		// Sets Validator for pkg1.ValidIntAlias type.
		var v pkg1.ValidIntAlias
		conf := musgo.NewAliasConf()
		conf.T = reflect.TypeOf(v)
		conf.Validator = "Positive"
		conf.Path = "./pkg1"
		err := musGo.GenerateAliasAs(conf)
		if err != nil {
			panic(err)
		}
	}
	{
		// Generates Marshal, Unmarshal, Size methods for pkg2.ByteAlias type.
		var v pkg2.ByteAlias
		conf := musgo.NewConf()
		conf.T = reflect.TypeOf(v)
		conf.Path = "./pkg2"
		err := musGo.GenerateAs(conf)
		if err != nil {
			panic(err)
		}
	}
	{
		// Specifies custom suffix for Marshal, Unmarshal, Size methods of the
		// pkg2.ByteAlias type.
		var v pkg2.ByteAlias
		conf := musgo.NewConf()
		conf.T = reflect.TypeOf(v)
		conf.Suffix = "CUSTOM"
		conf.Path = "./pkg2"
		conf.Filename = "CustomSuffixByteAlias.musgen.go"
		err := musGo.GenerateAs(conf)
		if err != nil {
			panic(err)
		}
	}
	{
		// Generates Marshal, Unmarshal, Size methods for pkg3.Struct type.
		var v pkg3.Struct
		conf := musgo.NewConf()
		conf.T = reflect.TypeOf(v)
		conf.Path = "./pkg3"
		err := musGo.GenerateAs(conf)
		if err != nil {
			panic(err)
		}
	}
	{
		// Generates Marshal, Unmarshal, Size methods for pkg3.InnerStruct type.
		var v pkg3.InnerStruct
		conf := musgo.NewConf()
		conf.T = reflect.TypeOf(v)
		conf.Path = "./pkg3"
		err := musGo.GenerateAs(conf)
		if err != nil {
			panic(err)
		}
	}
}
