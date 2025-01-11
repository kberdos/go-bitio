package bitio

import "io"

type BitWriter struct {
	w    io.Writer
	data []byte // all cached bytes
	c    uint8  // current byte
	p    uint8  // pos START_POS -> 0
}

func NewWriter(w io.Writer) *BitWriter {
	return &BitWriter{
		w:    w,
		data: make([]byte, 0),
		c:    0,
		p:    START_POS,
	}
}

func (bw *BitWriter) cache() error {
	bw.data = append(bw.data, bw.c)
	bw.c, bw.p = 0, START_POS
	return nil
}

func (bw *BitWriter) flush() error {
	if bw.p < START_POS {
		bw.cache()
	}
	_, err := bw.w.Write(bw.data)
	if err != nil {
		return err
	}
	return nil
}

func (bw *BitWriter) Close() error {
	return bw.flush()
}
func (bw *BitWriter) WriteOne()  { bw.writebit(true) }
func (bw *BitWriter) WriteZero() { bw.writebit(false) }
func (bw *BitWriter) writebit(one bool) error {
	if one {
		bw.c += 1 << bw.p
	}
	if bw.p == 0 {
		return bw.cache()
	}
	bw.p--
	return nil
}

// writes lowest n bits of r
func (bw *BitWriter) WriteBits(r uint64, n uint8) error {
	for n > 0 {
		m, err := bw.writeamap(r, n)
		if err != nil {
			return err
		}
		n -= m
	}
	return nil
}

// write 'as much as possible'
func (bw *BitWriter) writeamap(r uint64, n uint8) (uint8, error) {
	sz := min(n, bw.p+1)
	mask := uint64((1 << n) - 1)
	// relevant bytes
	r &= mask
	// cut off non-written bytes
	r >>= n - sz
	// add trailing 0s to place in correct pos in cache
	r <<= bw.p - sz + 1
	bw.c |= uint8(r)
	if bw.p == sz-1 {
		return sz, bw.cache()
	}
	bw.p -= sz
	return sz, nil
}
