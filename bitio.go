package bitio

import (
	"io"
)

type Bitio struct {
	rw   io.ReadWriter
	data []byte // all cached bytes
	c    uint8  // current byte
	p    int8   // pos START_POS -> 0
}

func New(rw io.ReadWriter) *Bitio {
	return &Bitio{
		rw:   rw,
		data: make([]byte, 0),
		c:    0,
		p:    START_POS,
	}
}

func (b *Bitio) cache() error {
	b.data = append(b.data, b.c)
	b.c, b.p = 0, START_POS
	return nil
}

func (b *Bitio) flush() error {
	b.cache()
	_, err := b.rw.Write(b.data)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bitio) Close() error {
	return b.flush()
}
