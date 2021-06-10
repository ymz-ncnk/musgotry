package musgotest

import (
	"github.com/ymz-ncnk/musgo/errs"
	"github.com/ymz-ncnk/musgotest/pkg"
)

func (v MyStruct) MarshalMUS(buf []byte) int {
	i := 0
	{
		si := (*v.MyInnerStruct).MarshalMUS(buf[i:])
		i += si
	}
	{
		si := v.Map.MarshalMUS(buf[i:])
		i += si
	}
	{
		uv := uint64((*v.number)<<1) ^ uint64((*v.number)>>63)
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
	return i
}

func (v *MyStruct) UnmarshalMUS(buf []byte) (int, error) {
	i := 0
	var err error
	v.MyInnerStruct = new(MyInnerStruct)
	{
		var sv MyInnerStruct
		si := 0
		si, err = sv.UnmarshalMUS(buf[i:])
		if err == nil {
			(*v.MyInnerStruct) = sv
			i += si
			err = ValidateInnerStruct(v.MyInnerStruct)
		}
	}
	if err != nil {
		return i, errs.NewFieldError("MyInnerStruct", err)
	}
	{
		var sv pkg.MyMap
		si := 0
		si, err = sv.UnmarshalMUS(buf[i:])
		if err == nil {
			v.Map = sv
			i += si
		}
	}
	if err != nil {
		return i, errs.NewFieldError("Map", err)
	}
	v.number = new(int)
	{
		var uv uint64
		{
			if i > len(buf)-1 {
				return i, errs.ErrSmallBuf
			} else {
				shift := 0
				done := false
				for l, b := range buf[i:] {
					if l > 9 || l == 9 && b > 1 {
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
		(*v.number) = int(uv)
	}
	if err != nil {
		return i, errs.NewFieldError("number", err)
	}
	return i, err
}

func (v MyStruct) SizeMUS() int {
	size := 0
	{
		ss := (*v.MyInnerStruct).SizeMUS()
		size += ss
	}
	{
		ss := v.Map.SizeMUS()
		size += ss
	}
	{
		uv := uint64((*v.number)<<1) ^ uint64((*v.number)>>63)
		{
			for uv >= 0x80 {
				uv >>= 7
				size++
			}
			size++
		}
	}
	return size
}
