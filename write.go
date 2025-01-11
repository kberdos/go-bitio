package bitio

func (b *Bitio) WriteOne()  { b.writebit(true) }
func (b *Bitio) WriteZero() { b.writebit(false) }
func (b *Bitio) writebit(one bool) error {
	if one {
		b.c += 1 << b.p
	}
	if b.p == 0 {
		return b.cache()
	}
	b.p--
	return nil
}

// writes lowest n bits of r
func (b *Bitio) WriteBits(r uint64, n uint8) error {
	for n > 0 {
		m, err := b.writeamap(r, n)
		if err != nil {
			return err
		}
		n -= m
	}
	return nil
}

// write 'as much as possible'
func (b *Bitio) writeamap(r uint64, n uint8) (uint8, error) {
	sz := min(n, b.p+1)
	mask := uint64((1 << n) - 1)
	r &= mask
	r >>= n - sz
	r <<= b.p - sz + 1
	b.c |= uint8(r)
	if b.p == sz-1 {
		return sz, b.cache()
	}
	b.p -= sz
	return sz, nil
}
