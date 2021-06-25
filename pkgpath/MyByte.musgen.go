package mypkg

import "github.com/ymz-ncnk/musgo/errs"

// MarshalMUS fills buf with the MUS encoding of v.
func (v MyByte) MarshalMUS(buf []byte) int {
	i := 0
	{
		buf[i] = byte(v)
		i++
	}
	return i
}

// UnmarshalMUS parses the MUS-encoded buf, and sets the result to *v.
func (v *MyByte) UnmarshalMUS(buf []byte) (int, error) {
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

// SizeMUS returns the size of the MUS-encoded v.
func (v MyByte) SizeMUS() int {
	size := 0
	{
		_ = v
		size++
	}
	return size
}
