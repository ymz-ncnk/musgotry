package pkg3

func NotEmptyInnerStruct(st *InnerStruct) error {
	if st.Num == 0 && st.Str == "" {
		return ErrEmptyInnerStruct
	}
	return nil
}
