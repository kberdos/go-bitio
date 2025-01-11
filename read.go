package bitio

import "io"

type BitReader struct {
	r    io.Reader
	data []byte // all cached bytes
	c    uint8  // current byte
	p    uint8  // pos START_POS -> 0
}
