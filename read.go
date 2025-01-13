package bitio

import (
	"encoding/binary"
	"errors"
	"io"
)

type BitReader struct {
	r io.Reader
	c uint64
	p uint8 // pos START_POS -> 0
}

func NewReader(r io.Reader) *BitReader {
	return &BitReader{
		r: r,
		c: 0,
		p: R_START_POS,
	}
}

func (br *BitReader) fill() error {
	buf := make([]byte, READ_BUF_SIZE)
	_, err := br.r.Read(buf)
	if err != nil {
		return err
	}
	br.c = binary.BigEndian.Uint64(buf)
	br.p = R_START_POS
	return nil
}

func (br *BitReader) ReadBits(n uint8) (uint64, error) {
	if n > MAX_READ_SIZE {
		return 0, errors.New("invalid read size")
	}
	var res uint64 = 0
	for n > 0 {
		r, m, err := br.readamap(res, n)
		if err != nil {
			return 0, err
		}
		n -= m
		res = r
	}
	return res, nil
}

// read 'as much as possible'
func (br *BitReader) readamap(res uint64, n uint8) (uint64, uint8, error) {
	sz := min(n, br.p+1)
	mask := uint64(1<<br.p - 1)
	// make space for the new bits
	res <<= sz
	// mask the cache and cut off trailing bits
	res |= (br.c & mask) >> (br.p - sz + 1)
	// refill cache if necessary
	if br.p == sz-1 {
		return res, sz, br.fill()
	}
	br.p -= sz
	return res, sz, nil
}
