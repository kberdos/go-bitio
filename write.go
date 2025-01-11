package bitio

func (b *Bitio) WriteOne()  { b.write(true) }
func (b *Bitio) WriteZero() { b.write(false) }

func (b *Bitio) write(add bool) error {
	if add {
		b.c += 1 << b.p
	}
	b.p--
	if b.p == -1 {
		return b.cache()
	}
	return nil
}
