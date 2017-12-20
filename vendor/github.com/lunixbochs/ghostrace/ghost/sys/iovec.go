package sys

import (
	"encoding/binary"
	"github.com/lunixbochs/struc"
	"io"
)

type Iovec32 struct {
	Base uint32
	Len  uint32
}

type Iovec64 struct {
	Base uint64
	Len  uint64
}

func iovecRead(r io.Reader, count uint64, bits int, endian binary.ByteOrder) []Iovec64 {
	ret := make([]Iovec64, 0, count)
	for i := uint64(0); i < count; i++ {
		if bits == 64 {
			var iovec Iovec64
			struc.UnpackWithOrder(r, &iovec, endian)
			ret = append(ret, iovec)
		} else {
			var iv32 Iovec32
			struc.UnpackWithOrder(r, &iv32, endian)
			ret = append(ret, Iovec64{uint64(iv32.Base), uint64(iv32.Len)})
		}
	}
	return ret
}
