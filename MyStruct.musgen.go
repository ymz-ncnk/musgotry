package musgotest

import (
	"math"
	"unsafe"

	"github.com/ymz-ncnk/musgo/errs"
	"github.com/ymz-ncnk/musgotest/pkg"
)

// MarshalMUS fills buf with the MUS encoding of v.
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
		uv := uint64((*v.number))
		if (*v.number) < 0 {
			uv = ^(uv << 1)
		} else {
			uv = uv << 1
		}
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
	{
		{
			*(*int64)(unsafe.Pointer(&buf[i])) = v.time
			i += 8
		}
	}
	{
		uv := math.Float32bits(float32(v.money))
		{
			*(*uint32)(unsafe.Pointer(&buf[i])) = uv
			i += 4
		}
	}
	return i
}

// UnmarshalMUS parses the MUS-encoded buf, and sets the result to *v.
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
			}
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
		if uv&1 == 1 {
			uv = ^(uv >> 1)
		} else {
			uv = uv >> 1
		}
		(*v.number) = int(uv)
	}
	if err != nil {
		return i, errs.NewFieldError("number", err)
	}
	{
		{
			if len(buf) < 8 {
				return i, errs.ErrSmallBuf
			}
			v.time = *(*int64)(unsafe.Pointer(&buf[i]))
			i += 8
		}
	}
	if err != nil {
		return i, errs.NewFieldError("time", err)
	}
	{
		var uv uint32
		{
			if len(buf) < 4 {
				return i, errs.ErrSmallBuf
			}
			uv = *(*uint32)(unsafe.Pointer(&buf[i]))
			i += 4
		}
		v.money = float32(math.Float32frombits(uv))
	}
	if err != nil {
		return i, errs.NewFieldError("money", err)
	}
	return i, err
}

// SizeMUS returns the size of the MUS-encoded v.
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
	{
		{
			_ = v.time
			size += 8
		}
	}
	{
		{
			_ = v.money
			size += 4
		}

	}
	return size
}
