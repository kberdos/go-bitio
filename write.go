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
// FIXME: goes backwards for now lol
func (b *Bitio) WriteBits(r uint64, n uint8) error {
	for n > 0 {
		rr, m, err := b.writefull(r, n)
		if err != nil {
			return err
		}
		n -= m
		r = rr
	}
	return nil
}

// writes all possible bits
func (b *Bitio) writefull(r uint64, n uint8) (uint64, uint8, error) {
	n = min(n, b.p+1)
	mask := (1 << n) - 1
	rb := uint8(r) & uint8(mask)
	rb <<= b.p + 1 - n
	b.c |= rb
	r >>= n
	if b.p == n-1 {
		return r, n, b.cache()
	}
	b.p -= n
	return r, n, nil
}
