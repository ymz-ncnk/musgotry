package pkg1

import "github.com/ymz-ncnk/musgotry/validators"

func Positive(n *ValidIntAlias) error {
	return validators.Positive(int(*n))
}

func KeySumPositive(m *ValidMapAlias) error {
	return validators.KeysSumPositive(map[int]string(*m))
}
