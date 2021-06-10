package pkg

import "github.com/ymz-ncnk/musgo/errs"

func (v MyInt) MarshalMUS(buf []byte) int {
	i := 0
	{
		uv := uint64(v<<1) ^ uint64(v>>63)
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

func (v *MyInt) UnmarshalMUS(buf []byte) (int, error) {
	i := 0
	var err error
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
		(*v) = MyInt(uv)
		err = ValidateMyInt((*v))
	}
	return i, err
}

func (v MyInt) SizeMUS() int {
	size := 0
	{
		uv := uint64(v<<1) ^ uint64(v>>63)
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
