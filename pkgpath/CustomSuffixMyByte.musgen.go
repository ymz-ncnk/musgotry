package mypkg

import "github.com/ymz-ncnk/musgo/errs"

func (v MyByte) MarshalAMUS(buf []byte) int {
	i := 0
	{
		buf[i] = byte(v)
		i++
	}
	return i
}

func (v *MyByte) UnmarshalAMUS(buf []byte) (int, error) {
	i := 0
	var err error
	{
		if i > len(buf)-1 {
			return i, errs.ErrSmallBuf
		} else {
			(*v) = MyByte(buf[i])
			i++
		}
	}
	return i, err
}

func (v MyByte) SizeAMUS() int {
	size := 0
	{
		_ = v
		size++
	}
	return size
}
