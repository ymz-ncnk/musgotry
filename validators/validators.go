package validators

import "errors"

var ErrNegative = errors.New("negative")
var ErrEmpty = errors.New("empty")
var ErrKeySumNegative = errors.New("keys sum negative")

func Positive(n int) error {
	if n < 0 {
		return ErrNegative
	}
	return nil
}

func NotEmpty(s string) error {
	if s == "" {
		return ErrEmpty
	}
	return nil
}

func KeysSumPositive(m map[int]string) error {
	sum := 0
	for k := range m {
		sum = sum + k
	}
	if sum < 0 {
		return ErrKeySumNegative
	}
	return nil
}
