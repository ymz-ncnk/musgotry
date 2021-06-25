package mypkg

import "github.com/ymz-ncnk/musgo/errs"

// MarshalAMUS fills buf with the MUS encoding of v.
func (v MyByte) MarshalAMUS(buf []byte) int {
	i := 0
	{
		buf[i] = byte(v)
		i++
	}
	return i
}

// UnmarshalAMUS parses the MUS-encoded buf, and sets the result to *v.
func (v *MyByte) UnmarshalAMUS(buf []byte) (int, error) {
	i := 0
	var err error
	{
		if i > len(buf)-1 {
			return i, errs.ErrSmallBuf
		}
		(*v) = MyByte(buf[i])
		i++
	}
	return i, err
}

// SizeAMUS returns the size of the MUS-encoded v.
func (v MyByte) SizeAMUS() int {
	size := 0
	{
		_ = v
		size++
	}
	return size
}
