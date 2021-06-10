package validators

import "errors"

// General purpose validators.

// ErrBiggerThenTen is an validation error, happens when integer is bigger
// then ten.
var ErrBiggerThenTen = errors.New("bigger then ten")

// ErrHelloString is an validation error, happens when string is equal to
// "hello".
var ErrHelloString = errors.New("hello string")

func BiggerThenTen(number int) error {
	if number > 10 {
		return ErrBiggerThenTen
	}
	return nil
}

func NotHello(str string) error {
	if str == "hello" {
		return ErrHelloString
	}
	return nil
}
