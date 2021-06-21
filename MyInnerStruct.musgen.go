package musgotest

import (
	"github.com/ymz-ncnk/musgo/errs"
	"github.com/ymz-ncnk/musgotest/pkg"
)

func (v MyInnerStruct) MarshalMUS(buf []byte) int {
	i := 0
	{
		si := v.Number.MarshalMUS(buf[i:])
		i += si
	}
	{
		length := len(v.Str)
		{
			uv := uint64(length<<1) ^ uint64(length>>63)
			{
				for uv >= 0x80 {
					buf[i] = byte(uv) | 0x80
					uv >>= 7
					i++
				}
				buf[i] = byte(uv)
				i++
			}
		}
		i += copy(buf[i:], v.Str)
	}
	return i
}

func (v *MyInnerStruct) UnmarshalMUS(buf []byte) (int, error) {
	i := 0
	var err error
	{
		var sv pkg.MyInt
		si := 0
		si, err = sv.UnmarshalMUS(buf[i:])
		if err == nil {
			v.Number = sv
			i += si
		}
	}
	if err != nil {
		return i, errs.NewFieldError("Number", err)
	}
	{
		var length int
		{
			var uv uint64
			{
				if i > len(buf)-1 {
					return i, errs.ErrSmallBuf
				} else {
					shift := 0
					done := false
					for l, b := range buf[i:] {
						if l == 9 && b > 1 {
							return i, errs.ErrOverflow
						}
						if b < 0x80 {
							uv = uv | uint64(b)<<shift
							done = true
							i += l + 1
							break
						}
						uv = uv | uint64(b&0x7F)<<shift
						shift += 7
					}
					if !done {
						return i, errs.ErrSmallBuf
					}
				}
			}
			uv = (uv >> 1) ^ uint64((int(uv&1)<<63)>>63)
			length = int(uv)
		}
		if length < 0 {
			return i, errs.ErrNegativeLength
		}
		if len(buf) < i+length {
			return i, errs.ErrSmallBuf
		}
		v.Str = string(buf[i : i+length])
		i += length
	}
	if err != nil {
		return i, errs.NewFieldError("Str", err)
	}
	return i, err
}

func (v MyInnerStruct) SizeMUS() int {
	size := 0
	{
		ss := v.Number.SizeMUS()
		size += ss
	}
	{
		length := len(v.Str)
		{
			uv := uint64(length<<1) ^ uint64(length>>63)
			{
				for uv >= 0x80 {
					uv >>= 7
					size++
				}
				size++
			}
		}
		size += len(v.Str)
	}
	return size
}
